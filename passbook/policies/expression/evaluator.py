"""passbook expression policy evaluator"""
import re
from typing import TYPE_CHECKING, Any, Dict, Optional

from django.core.exceptions import ValidationError
from jinja2 import Undefined
from jinja2.exceptions import TemplateSyntaxError, UndefinedError
from jinja2.nativetypes import NativeEnvironment
from requests import Session
from structlog import get_logger

from passbook.flows.planner import PLAN_CONTEXT_SSO
from passbook.flows.views import SESSION_KEY_PLAN
from passbook.lib.utils.http import get_client_ip
from passbook.policies.types import PolicyRequest, PolicyResult

if TYPE_CHECKING:
    from passbook.core.models import User

LOGGER = get_logger()


class Evaluator:
    """Validate and evaluate jinja2-based expressions"""

    _env: NativeEnvironment

    def __init__(self):
        self._env = NativeEnvironment()
        # update passbook/policies/expression/templates/policy/expression/form.html
        # update docs/policies/expression/index.md
        self._env.filters["regex_match"] = Evaluator.jinja2_filter_regex_match
        self._env.filters["regex_replace"] = Evaluator.jinja2_filter_regex_replace

    @staticmethod
    def jinja2_filter_regex_match(value: Any, regex: str) -> bool:
        """Jinja2 Filter to run re.search"""
        return re.search(regex, value) is None

    @staticmethod
    def jinja2_filter_regex_replace(value: Any, regex: str, repl: str) -> str:
        """Jinja2 Filter to run re.sub"""
        return re.sub(regex, repl, value)

    @staticmethod
    def jinja2_func_is_group_member(user: "User", group_name: str) -> bool:
        """Check if `user` is member of group with name `group_name`"""
        return user.groups.filter(name=group_name).exists()

    def _get_expression_context(
        self, request: PolicyRequest, **kwargs
    ) -> Dict[str, Any]:
        """Return dictionary with additional global variables passed to expression"""
        # update passbook/policies/expression/templates/policy/expression/form.html
        # update docs/policies/expression/index.md
        kwargs["pb_is_group_member"] = Evaluator.jinja2_func_is_group_member
        kwargs["pb_logger"] = get_logger()
        kwargs["requests"] = Session()
        kwargs["pb_is_sso_flow"] = request.context.get(PLAN_CONTEXT_SSO, False)
        if request.http_request:
            kwargs["pb_client_ip"] = (
                get_client_ip(request.http_request) or "255.255.255.255"
            )
            if SESSION_KEY_PLAN in request.http_request.session:
                kwargs["pb_flow_plan"] = request.http_request.session[SESSION_KEY_PLAN]
        return kwargs

    def evaluate(self, expression_source: str, request: PolicyRequest) -> PolicyResult:
        """Parse and evaluate expression.
        If the Expression evaluates to a list with 2 items, the first is used as passing bool and
        the second as messages.
        If the Expression evaluates to a truthy-object, it is used as passing bool."""
        try:
            expression = self._env.from_string(expression_source)
        except TemplateSyntaxError as exc:
            return PolicyResult(False, str(exc))
        try:
            context = self._get_expression_context(request)
            LOGGER.debug("expression context", **context)
            result: Optional[Any] = expression.render(request=request, **context)
            if isinstance(result, Undefined):
                LOGGER.warning(
                    "Expression policy returned undefined",
                    src=expression_source,
                    req=request,
                )
                return PolicyResult(False)
            if isinstance(result, (list, tuple)) and len(result) == 2:
                return PolicyResult(*result)
            if result:
                return PolicyResult(bool(result))
            return PolicyResult(False)
        except UndefinedError as exc:
            return PolicyResult(False, str(exc))

    def validate(self, expression: str):
        """Validate expression's syntax, raise ValidationError if Syntax is invalid"""
        try:
            self._env.from_string(expression)
            return True
        except TemplateSyntaxError as exc:
            raise ValidationError("Expression Syntax Error") from exc
