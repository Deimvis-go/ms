package ms

import (
	"net/http"
)

// PathNormalizer returns normalized path with low cardinality.
// For example it may be able to identify that request path '/users/sDaf34Fb9'
// is a variation of normalized path '/users/:id'.
// It returns nil if no normalized path found.
type PathNormalizer func(*http.Request) *string
