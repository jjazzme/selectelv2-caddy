package selectelv2

import (
	// "fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	selectelv2 "github.com/jjazzme/selectelv2-libdns"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *selectelv2.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.selectelv2",
		New: func() caddy.Module { return &Provider{new(selectelv2.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.User = caddy.NewReplacer().ReplaceAll(p.Provider.User, "")
	p.Provider.Password = caddy.NewReplacer().ReplaceAll(p.Provider.Password, "")
	p.Provider.AccountId = caddy.NewReplacer().ReplaceAll(p.Provider.AccountId, "")
	p.Provider.ProjectName = caddy.NewReplacer().ReplaceAll(p.Provider.ProjectName, "")
	return nil
}

func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "user":
				if p.Provider.User != "" {
					return d.Err("User already set")
				}
				if d.NextArg() {
					p.Provider.User = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "password":
				if p.Provider.Password != "" {
					return d.Err("Password already set")
				}
				if d.NextArg() {
					p.Provider.Password = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "account_id":
				if p.Provider.AccountId != "" {
					return d.Err("AccountId already set")
				}
				if d.NextArg() {
					p.Provider.AccountId = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "project_name":
				if p.Provider.ProjectName != "" {
					return d.Err("ProjectName already set")
				}
				if d.NextArg() {
					p.Provider.ProjectName = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}


	if p.Provider.User == "" {
		return d.Err("missing User")
	}
	if p.Provider.Password == "" {
		return d.Err("missing Password")
	}
	if p.Provider.AccountId == "" {
		return d.Err("missing AccountId")
	}
	if p.Provider.ProjectName == "" {
		return d.Err("missing ProjectName")
	}

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
