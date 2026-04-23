package msconfig

import (
	"net/http"
)

type RoundTripFn func(*http.Request) (*http.Response, error)
type RoundTripWrapFn func(RoundTripFn) RoundTripFn

func NewRoundTripWrap(cfg *HTTPRequests) RoundTripWrapFn {
	h := cfg.Headers
	return func(fn RoundTripFn) RoundTripFn {
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
