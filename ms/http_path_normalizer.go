package ms

import (
	"net/http"

	"github.com/Deimvis-go/ms/ms/msconfig"
)

func NewHTTPPathNormalizer(downstreamsCfg []msconfig.HTTPDownstream) PathNormalizer {
	return func(req *http.Request) *string {
		// TODO: implement
		return nil
	}
}
