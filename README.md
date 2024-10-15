# Timeweb module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Timeweb API.

## Caddy module name

```
dns.providers.timeweb
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "timeweb",
        "ApiURL": "",
        "ApiToken": ""
      }
    }
  }
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns timeweb {
		ApiURL <API_URL>
		ApiToken <API_TOKEN>
	}
}
```

```
# one site
tls {
	dns timeweb {
		ApiURL <API_URL>
		ApiToken <API_TOKEN>
	}
}
```