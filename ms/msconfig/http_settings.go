package msconfig

import (
	"fmt"
	"os"
)

type HTTPSettings struct {
	Requests    HTTPRequests     `yaml:"requests"`
	Downstreams []HTTPDownstream `yaml:"downstreams"`
}

type HTTPRequests struct {
	Headers HTTPRequestHeaders `yaml:"headers"`
}

type HTTPRequestHeaders struct {
	XClientHost          *HTTPRequestHeaderValue `yaml:"X-Client-Host" validate:"omitnil"`
	XClientCloudService  *HTTPRequestHeaderValue `yaml:"X-Client-CloudService" validate:"omitnil"`
	XClientCloudInstance *HTTPRequestHeaderValue `yaml:"X-Client-CloudInstance" validate:"omitnil"`
}

type HTTPRequestHeaderValue struct {
	Const *string `yaml:"const"`
	Env   *string `yaml:"env"`

	cachedValue *string
}

type HTTPDownstream struct {
	Id      string              `yaml:"id"`
	Match   HTTPDownstreamMatch `yaml:"match"`
	OpenAPI *OpenAPI            `yaml:"openapi" validate:"omitnil"`
}

type HTTPDownstreamMatch struct {
	Host HTTPDownstreamHostMatch `yaml:"host"`
}

type HTTPDownstreamHostMatch struct {
	Regexp string `yaml:"regexp"`
}

type OpenAPI struct {
	FilePath string `yaml:"file_path" validate:"file"`
}

func (hv *HTTPRequestHeaderValue) Value() string {
	if hv.cachedValue == nil {
		hv.invalidateCache()
	}
	return *hv.cachedValue
}

func (hv *HTTPRequestHeaderValue) invalidateCache() {
	hv.cachedValue = nil
	if hv.Const != nil {
		hv.cachedValue = hv.Const
	} else if hv.Env != nil {
		v := os.Getenv(*hv.Env)
		hv.cachedValue = &v
	}

	if hv.cachedValue == nil {
		panic("HTTPRequestHeaderValue is invalid - unable to determine header value")
	}
}

func (hv *HTTPRequestHeaderValue) ValidateSelf() error {
	vsOpts := []*string{hv.Const, hv.Env}
	set := 0
	for _, v := range vsOpts {
		if v != nil {
			set++
		}
	}
	if set != 1 {
		return fmt.Errorf("exactly one option should be specified, but %d given (%v)", set, vsOpts)
	}
	return nil
}
