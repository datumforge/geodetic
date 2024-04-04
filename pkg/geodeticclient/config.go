package geodeticclient

import (
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

type Config struct {
	// Enabled is a flag to enable the geodetic client
	Enabled bool `json:"enabled" koanf:"enabled" jsonschema:"description=Enable the geodetic client" default:"true"`
	// BaseURL is the base url for the geodetic service
	BaseURL string `json:"baseUrl" koanf:"baseUrl" jsonschema:"description=Base URL for the geodetic service" default:"http://localhost:1337"`
	// Endpoint is the endpoint for the graphql api
	Endpoint string `json:"endpoint" koanf:"endpoint" jsonschema:"description=Endpoint for the graphql api" default:"query"`
	// Debug is a flag to enable debug mode
	Debug bool `json:"debug" koanf:"debug" jsonschema:"description=Enable debug mode" default:"false"`
}

// NewDefaultClient creates a new default geodetic client based on the config
func (c Config) NewDefaultClient() GeodeticClient {
	i := WithEmptyInterceptor()
	interceptors := []clientv2.RequestInterceptor{i}

	if c.Debug {
		interceptors = append(interceptors, WithLoggingInterceptor())
	}

	return c.NewClientWithInterceptors(interceptors)
}

// NewClientWithInterceptors creates a new default geodetic client with the provided interceptors
func (c Config) NewClientWithInterceptors(i []clientv2.RequestInterceptor) GeodeticClient {
	h := http.DefaultClient

	// set options
	opts := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	gc := NewClient(h, c.BaseURL+"/"+c.Endpoint, opts, i...)

	return gc
}
