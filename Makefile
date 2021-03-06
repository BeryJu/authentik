all: lint-fix lint test gen

test-integration:
	k3d cluster create || exit 0
	k3d kubeconfig write -o ~/.kube/config --overwrite
	coverage run manage.py test -v 3 tests/integration

test-e2e:
	coverage run manage.py test --failfast -v 3 tests/e2e

test:
	coverage run manage.py test -v 3 authentik
	coverage html
	coverage report

lint-fix:
	isort authentik tests lifecycle
	black authentik tests lifecycle

lint:
	pyright authentik tests lifecycle
	bandit -r authentik tests lifecycle -x node_modules
	pylint authentik tests lifecycle

gen:
	./manage.py generate_swagger -o swagger.yaml -f yaml

local-stack:
	export AUTHENTIK_TAG=testing
	docker build -t beryju/authentik:testng .
	docker-compose up -d
	docker-compose run --rm server migrate

run:
	go run -v cmd/server/main.go
