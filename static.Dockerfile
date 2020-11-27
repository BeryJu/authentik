FROM python:3.9-slim-buster as locker

COPY ./Pipfile /app/
COPY ./Pipfile.lock /app/

WORKDIR /app/

RUN pip install pipenv && \
    pipenv lock -r > requirements.txt && \
    pipenv lock -rd > requirements-dev.txt

FROM python:3.9-slim-buster as static-build

COPY --from=locker /app/requirements.txt /app/
COPY --from=locker /app/requirements-dev.txt /app/

WORKDIR /app/

RUN apt-get update && \
    apt-get install -y --no-install-recommends postgresql-client-11 build-essential libxmlsec1-dev pkg-config && \
    rm -rf /var/lib/apt/ && \
    pip install -r requirements.txt  --no-cache-dir && \
    apt-get remove --purge -y build-essential && \
    apt-get autoremove --purge && \
    adduser --system --no-create-home --uid 1000 --group --home /app passbook

COPY ./passbook/ /app/passbook
COPY ./manage.py /app/

WORKDIR /app/

ENV PASSBOOK_POSTGRESQL__HOST=postgres
ENV PASSBOOK_REDIS__HOST=redis
ENV PASSBOOK_POSTGRESQL__USER=passbook
# CI Password, same as in .github/workflows/ci.yml
ENV PASSBOOK_POSTGRESQL__PASSWORD="EK-5jnKfjrGRm<77"
RUN ./manage.py collectstatic --no-input

FROM node as npm-builder

COPY ./web /static/

RUN cd /static && npm i && npm run build

FROM nginx

COPY --from=static-build /app/static /usr/share/nginx/html/static
COPY --from=static-build /app/static/robots.txt /usr/share/nginx/html/robots.txt
COPY --from=npm-builder /static/node_modules /usr/share/nginx/html/static/node_modules
COPY --from=npm-builder /static/dist/* /usr/share/nginx/html/static/dist/
