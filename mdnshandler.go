package mdnshandler

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/hashicorp/mdns"
)

func init() {
	caddy.RegisterModule(MDNSHandler{})
}

// MDNSHandler is a Caddy module that resolves the .local hostname using mDNS.
type MDNSHandler struct {
	Name    string `json:"name,omitempty"`
	Service string `json:"service,omitempty"`
}

func (h MDNSHandler) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.mdnshandler",
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

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*MDNSHandler)(nil)
	_ caddy.Module                = (*MDNSHandler)(nil)
)
