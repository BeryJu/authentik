---
title: Grafana
---

## What is Grafana

From https://en.wikipedia.org/wiki/Grafana

:::note
Grafana is a multi-platform open source analytics and interactive visualization web application. It provides charts, graphs, and alerts for the web when connected to supported data sources, Grafana Enterprise version with additional capabilities is also available. It is expandable through a plug-in system.
:::

## Preparation

The following placeholders will be used:

-   `grafana.company` is the FQDN of the Grafana install.
-   `authentik.company` is the FQDN of the authentik install.

Create an application in authentik. Create an OAuth2/OpenID provider with the following parameters:

-   Client Type: `Confidential`
-   JWT Algorithm: `RS256`
-   Scopes: OpenID, Email and Profile
-   RSA Key: Select any available key
-   Redirect URIs: `https://grafana.company/login/generic_oauth`

Note the Client ID and Client Secret values. Create an application, using the provider you've created above. Note the slug of the application you've created.

## Grafana

If your Grafana is running in docker, set the following environment variables:

```yaml
environment:
    GF_AUTH_GENERIC_OAUTH_ENABLED: "true"
    GF_AUTH_GENERIC_OAUTH_NAME: "authentik"
    GF_AUTH_GENERIC_OAUTH_CLIENT_ID: "<Client ID from above>"
    GF_AUTH_GENERIC_OAUTH_CLIENT_SECRET: "<Client Secret from above>"
    GF_AUTH_GENERIC_OAUTH_SCOPES: "openid profile email"
    GF_AUTH_GENERIC_OAUTH_AUTH_URL: "https://authentik.company/application/o/authorize/"
    GF_AUTH_GENERIC_OAUTH_TOKEN_URL: "https://authentik.company/application/o/token/"
    GF_AUTH_GENERIC_OAUTH_API_URL: "https://authentik.company/application/o/userinfo/"
    GF_AUTH_SIGNOUT_REDIRECT_URL: "https://authentik.company/application/o/<Slug of the application from above>/end-session/"
    # Optionally enable auto-login
    GF_AUTH_OAUTH_AUTO_LOGIN: "true"
```

If you are using a config-file instead, you have to set these options:

```ini
[auth]
signout_redirect_url = https://authentik.company/application/o/<Slug of the application from above>/end-session/
# Optionally enable auto-login
oauth_auto_login = true

[auth.generic_oauth]
name = authentik
enabled = true
client_id = <Client ID from above>
client_secret = <Client Secret from above>
scopes = openid email profile
auth_url = https://authentik.company/application/o/authorize/
token_url = https://authentik.company/application/o/token/
api_url = https://authentik.company/application/o/userinfo/
```
