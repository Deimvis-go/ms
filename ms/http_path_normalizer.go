package ms

import (
	"net/http"

	"github.com/Deimvis-go/ms/ms/msconfig"
	"github.com/Deimvis/go-ext/go1.25/xhttp"
)

func NewHTTPPathNormalizer(downstreamsCfg []msconfig.HTTPDownstream) xhttp.PathNormalizer {
	return func(req *http.Request) *string {
		// TODO: implement
		return nil
	}
}
