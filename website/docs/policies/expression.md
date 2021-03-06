---
title: Expression Policies
---

:::note
These variables are available in addition to the common variables/functions defined in [**Expressions**](../expressions/index.md)
:::

The passing of the policy is determined by the return value of the code. Use `return True` to pass a policy and `return False` to fail it.

### Available Functions

#### `ak_message(message: str)`

Add a message, visible by the end user. This can be used to show the reason why they were denied.

Example:

```python
ak_message("Access denied")
return False
```

### Context variables

- `request`: A PolicyRequest object, which has the following properties:
    - `request.user`: ([ref](../expressions/reference/user-object.md)) The current user, against which the policy is applied.
    - `request.http_request`: ([ref](https://docs.djangoproject.com/en/3.0/ref/request-response/#httprequest-objects)) The Django HTTP Request.
    - `request.obj`: A Django Model instance. This is only set if the policy is ran against an object.
    - `request.context`: A dictionary with dynamic data. This depends on the origin of the execution.
- `geoip`: ([ref](https://geoip2.readthedocs.io/en/latest/#geoip2.models.City)) GeoIP object, which is added when GeoIP is enabled.
- `ak_is_sso_flow`: Boolean which is true if request was initiated by authenticating through an external provider.
- `ak_client_ip`: ([ref](https://docs.python.org/3/library/ipaddress.html#ipaddress.ip_address))  Client's IP Address or 255.255.255.255 if no IP Address could be extracted. Can be [compared](../expressions/index.md#comparing-ip-addresses), for example

    ```python
    return ak_client_ip in ip_network('10.0.0.0/24')
    # or
    return ak_client_ip.is_private
    ```

Additionally, when the policy is executed from a flow, every variable from the flow's current context is accessible under the `context` object.

This includes the following:

- `prompt_data`: Data which has been saved from a prompt stage or an external source.
- `application`: The application the user is in the process of authorizing.
- `pending_user`: The currently pending user
