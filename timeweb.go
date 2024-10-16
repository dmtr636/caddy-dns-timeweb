package timeweb

import (
	// "fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	timeweb "github.com/dmtr636/libdns-timeweb"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *timeweb.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.timeweb",
		New: func() caddy.Module { return &Provider{new(timeweb.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.ApiURL = caddy.NewReplacer().ReplaceAll(p.Provider.ApiURL, "")
	p.Provider.ApiToken = caddy.NewReplacer().ReplaceAll(p.Provider.ApiToken, "")
	return nil
}

func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "ApiURL":
				if p.Provider.ApiURL != "" {
					return d.Err("ApiURL already set")
				}
				if d.NextArg() {
					p.Provider.ApiURL = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "ApiToken":
				if p.Provider.ApiToken != "" {
					return d.Err("ApiToken already set")
				}
				if d.NextArg() {
					p.Provider.ApiToken = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	if p.Provider.ApiURL == "" {
		return d.Err("missing User")
	}
	if p.Provider.ApiToken == "" {
		return d.Err("missing Password")
	}

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
