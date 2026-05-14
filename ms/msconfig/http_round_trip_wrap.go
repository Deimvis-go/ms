package msconfig

import (
	"net/http"

	"github.com/Deimvis/go-ext/go1.25/xhttp"
)

func NewRoundTripWrap(cfg *HTTPRequests) xhttp.RoundTripWrapFn {
	h := cfg.Headers
	return func(fn xhttp.RoundTripFn) xhttp.RoundTripFn {
		return func(req *http.Request) (*http.Response, error) {
			if h.XClientHost != nil {
				req.Header.Set(XClientHost, h.XClientCloudService.Value())
			}
			if h.XClientCloudService != nil {
				req.Header.Set(XClientCloudService, h.XClientCloudService.Value())
			}
			if h.XClientCloudInstance != nil {
				req.Header.Set(XClientCloudInstance, h.XClientCloudInstance.Value())
			}
			return fn(req)
		}
	}
}
