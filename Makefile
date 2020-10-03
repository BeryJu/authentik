all: lint-fix lint coverage gen

coverage:
	coverage run --concurrency=multiprocessing manage.py test --failfast -v 3
	coverage combine
	coverage html
	coverage report

lint-fix:
	isort -rc .
	black passbook e2e lifecycle

lint:
	pyright passbook e2e lifecycle
	bandit -r passbook e2e lifecycle
	pylint passbook e2e lifecycle
	prospector

gen: coverage
	./manage.py generate_swagger -o swagger.yaml -f yaml

local-stack:
	export PASSBOOK_TAG=testing
	docker build -t beryju/passbook:testng .
	docker-compose up -d
	docker-compose run --rm server migrate
