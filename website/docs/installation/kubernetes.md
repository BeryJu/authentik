---
title: Kubernetes installation
---

For a mid to high-load installation, Kubernetes is recommended. authentik is installed using a helm-chart.

To install authentik using the helm chart, run these commands:

```
helm repo add authentik https://docker.beryju.org/chartrepo/authentik
helm repo update
helm install authentik/authentik --devel -f values.yaml
```

This installation automatically applies database migrations on startup. After the installation is done, navigate to the `https://<ingress you've specified>/if/flow/initial-setup/`, to set a password for the akadmin user.

It is also recommended to configure global email credentials. These are used by authentik to notify you about alerts, configuration issues. They can also be used by [Email stages](flow/stages/email/index.md) to send verification/recovery emails.

```yaml
###################################
# Values directly affecting authentik
###################################
image:
  name: beryju/authentik
  name_static: beryju/authentik-static
  name_outposts: beryju/authentik # Prefix used for Outpost deployments, Outpost type and version is appended
  tag: 2021.3.4

serverReplicas: 1
workerReplicas: 1

# Enable the Kubernetes integration which lets authentik deploy outposts into kubernetes
kubernetesIntegration: true

config:
  # Optionally specify fixed secret_key, otherwise generated automatically
  # secretKey: _k*@6h2u2@q-dku57hhgzb7tnx*ba9wodcb^s9g0j59@=y(@_o
  # Enable error reporting
  errorReporting:
    enabled: false
    environment: customer
    sendPii: false
  # Log level used by web and worker
  # Can be either debug, info, warning, error
  logLevel: warning
  # Global Email settings
  email:
    # SMTP Host Emails are sent to
    host: localhost
    port: 25
    # Optionally authenticate
    username: ""
    password: ""
    # Use StartTLS
    useTls: false
    # Use SSL
    useSsl: false
    timeout: 10
    # Email address authentik will send from, should have a correct @domain
    from: authentik@localhost

# Enable MaxMind GeoIP
# geoip:
#   enabled: false
#   accountId: ""
#   licenseKey: ""
#   image: maxmindinc/geoipupdate:latest

# Enable Database Backups to S3
# backup:
#   accessKey: access-key
#   secretKey: secret-key
#   bucket: s3-bucket
#   region: eu-central-1
#   host: s3-host

ingress:
  annotations:
    {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  hosts:
    - authentik.k8s.local
  tls: []
  #  - secretName: chart-example-tls
  #    hosts:
  #      - authentik.k8s.local

###################################
# Values controlling dependencies
###################################

install:
  postgresql: true
  redis: true
```
