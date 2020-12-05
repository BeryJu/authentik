---
title: Upgrading to 0.13 (passbook -> authentik)
---

After a long back and forth, we've finally switched to a more permanent name. Whilst the upgrade is pretty much seamless, there are some things you have to change before upgrading.

## Upgrading

### docker-compose

Docker-compose users should download the latest docker-compose file from [here](https://raw.githubusercontent.com/BeryJu/authentik/master/docker-compose.yml).

:::caution
If you decided to rename the folder you're running the docker-compose file from, be aware that this will also change the name, that docker-compose will give the database volume. To mitigate this, either
- Keep the original directory name
- Move the directory and set `COMPOSE_PROJECT_NAME` to the name of the old directory (see [here](https://docs.docker.com/compose/reference/envvars/#compose_project_name))
- Create a backup, rename the directory and restore from backup.
:::

The only manual change you have to do is replace the `PASSBOOK_` prefix in your `.env` file, so `PASSBOOK_SECRET_KEY` gets changed to `AUTHENTIK_SECRET_KEY`.

Afterwards, you can simply run `docker-compose up -d` and then the normal upgrade command of `docker-compose run --rm server migrate`.

### Kubernetes

The helm repository changes from passbook to authentik. To update your repository, execute these commands:

```
helm repo remove passbook
helm repo add authentik https://docker.beryju.org/chartrepo/authentik
```

:::notice
If you've set any custom image names in your values file, make sure to change them to authentik before upgrading.
:::

Additionally, you need to change the database name that authentik uses, as the database name doesn't change. Add this snippet to your `values.yaml` file:

```yaml
postgresql:
    postgresqlDatabase: passbook
```

Afterwards you can upgrade as usual from the new repository:

```
helm upgrade passbook authentik/authentik --devel -f values.yaml
```

## Post-upgrade notes

- Some default values change, for example the SAML Provider's default issuer.

    This only makes a difference for newly created objects.

- Expression Policies variables change

    Anything prefixed with `pb_` changes to `ak_`, this change is done **automatically**
