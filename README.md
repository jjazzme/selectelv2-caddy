# Selectel DNS v2 module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [Selectel DNS v2 API](https://developers.selectel.ru/docs/cloud-services/dns_api/dns_api_actual/).

## Caddy module name

```
dns.providers.selectelv2
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "selectelv2",
        "user": "SELECTEL_USER", // see link "Selectel Service user" below
        "password": "SELECTEL_PASSWORD",
        "account_id": "SELECTEL_ACCOUNT_ID", // your main Selectel account id, like "123456"
        "project_name": "SELECTEL_PROJECT_NAME" // a project in which the service user is an administrator
      }
    }
  }
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns selectelv2 {
		user <SELECTEL_USER>
		password <SELECTEL_PASSWORD>
		account_id <SELECTEL_ACCOUNT_ID>
		project_name <SELECTEL_PROJECT_NAME>
	}
}
```

```
# one site
tls {
	dns selectelv2 {
		user <SELECTEL_USER>
		password <SELECTEL_PASSWORD>
		account_id <SELECTEL_ACCOUNT_ID>
		project_name <SELECTEL_PROJECT_NAME>
	}
}
```

Selectel [Service user](https://my.selectel.ru/iam/users_management/users?type=service) management
