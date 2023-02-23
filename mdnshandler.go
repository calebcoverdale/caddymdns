package mdnshandler

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/hashicorp/mdns"
)

func init() {
	caddy.RegisterModule(MDNSHandler{})
}

type MDNSHandler struct {
	Name    string
	Service string
}

func (h MDNSHandler) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.mdns",
		New: func() caddy.Module { return new(MDNSHandler) },
	}
}

func (h *MDNSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	entriesCh := make(chan *mdns.ServiceEntry)
	mdns.Lookup(h.Service, entriesCh)

	for entry := range entriesCh {
		if entry.Host == h.Name {
			r.URL.Scheme = "http"
			r.URL.Host = entry.AddrV4.String()
			break
		}
	}

	return next.ServeHTTP(w, r)
}

func (h *MDNSHandler) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	if !d.Next() {
		return d.Err("missing service name")
	}
	h.Service = d.Val()

	if !d.NextArg() {
		return d.Err("missing host name")
	}
	h.Name = d.Val()

	if d.NextBlock() {
		return d.Err("unexpected block")
	}

	return nil
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var hnd MDNSHandler
	err := hnd.UnmarshalCaddyfile(h.Dispenser)
	return hnd, err
}

func (h MDNSHandler) Validate() error {
	if h.Service == "" {
		return caddyhttp.Error("service name is required")
	}
	if h.Name == "" {
		return caddyhttp.Error("host name is required")
	}
	return nil
}

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*MDNSHandler)(nil)
	_ caddyfile.Unmarshaler       = (*MDNSHandler)(nil)
)
