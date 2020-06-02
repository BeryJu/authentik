# passbook

![](https://img.shields.io/github/workflow/status/beryju/passbook/passbook-ci?style=flat-square)
![](https://img.shields.io/docker/pulls/beryju/passbook.svg?style=flat-square)
![](https://img.shields.io/docker/pulls/beryju/passbook-gatekeeper.svg?style=flat-square)
![](https://img.shields.io/docker/pulls/beryju/passbook-static.svg?style=flat-square)
![](https://img.shields.io/docker/v/beryju/passbook?sort=semver&style=flat-square)
![](https://img.shields.io/codecov/c/gh/beryju/passbook?style=flat-square)

## What is passbook?

passbook is an open-source Identity Provider focused on flexibility and versatility. You can use passbook in an existing environment to add support for new protocols. passbook is also a great solution for implementing signup/recovery/etc in your application, so you don't have to deal with it.

## Installation

For small/test setups it is recommended to use docker-compose.

```
wget https://raw.githubusercontent.com/BeryJu/passbook/master/docker-compose.yml
# Optionally enable Error-reporting
# export PASSBOOK_ERROR_REPORTING=true
# Optionally deploy a different version
# export PASSBOOK_TAG=0.8.15-beta
# If this is a productive installation, set a different PostgreSQL Password
# export PG_PASS=$(pwgen 40 1)
docker-compose pull
docker-compose up -d
docker-compose exec server ./manage.py migrate
```

For bigger setups, there is a Helm Chart in the `helm/` directory. This is documented [here](https://beryju.github.io/passbook/installation/kubernetes/)

## Screenshots

![](.github/screen_apps.png)
![](.github/screen_admin.png)

## Development

To develop on passbook, you need a system with Python 3.7+ (3.8 is recommended). passbook uses [pipenv](https://pipenv.pypa.io/en/latest/) for managing dependencies.

To get started, run

```
python3 -m pip install pipenv
git clone https://github.com/BeryJu/passbook.git
cd passbook
pipenv shell
pipenv sync -d
```

Since passbook uses PostgreSQL-specific fields, you also need a local PostgreSQL instance to develop. passbook also uses redis for caching and message queueing.
For these databases you can use [Postgres.app](https://postgresapp.com/) and [Redis.app](https://jpadilla.github.io/redisapp/) on macOS or use it via docker-comppose:

```yaml
version: '3.7'

services:
  postgresql:
    container_name: postgres
    image: postgres:11
    volumes:
    - db-data:/var/lib/postgresql/data
    ports:
    - 127.0.0.1:5432:5432
    restart: always
  redis:
    container_name: redis
    image: redis
    ports:
    - 127.0.0.1:6379:6379
    restart: always

volumes:
  db-data:
    driver: local
```

To tell passbook about these databases, create a file in the project root called `local.env.yml` with the following contents:

```yaml
debug: true
postgresql:
  user: postgres

log_level: debug
error_reporting: false
```

## Security

See [SECURITY.md](SECURITY.md)
