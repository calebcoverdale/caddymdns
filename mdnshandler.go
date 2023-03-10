package github.com/calebcoverdale/caddymdns

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/hashicorp/mdns"
)

func init() {
	caddy.RegisterModule(MDNSHandler{})
	httpcaddyfile.RegisterHandlerDirective("mdns", parseCaddyfile)
}

// MDNSHandler is a Caddy module that resolves the .local hostname using mDNS.
type MDNSHandler struct {
	Name      string  `json:"name,omitempty"`
	Service   string  `json:"service,omitempty"`
	Resolvers *string `json:"resolvers,omitempty"`
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

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var handler MDNSHandler

	for h.Next() {
		if !h.Args(&handler.Name, &handler.Service) {
			return nil, h.ArgErr()
		}
		if h.NextBlock(0) {
			for {
				switch h.Val() {
				case "resolvers":
					var resolvers string
					if !h.Args(&resolvers) {
						return nil, h.ArgErr()
					}
					handler.Resolvers = &resolvers // create a new pointer to the string
				default:
					if h.Val() != "}" {
						return nil, h.Errf("unrecognized subdirective: %s", h.Val())
					}
					return &handler, nil
				}
			}
		}

	}

	return nil, h.Err("missing closing brace for mdns directive")
}

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*MDNSHandler)(nil)
	_ caddy.Module                = (*MDNSHandler)(nil)
)
