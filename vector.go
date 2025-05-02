// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"github.com/goperigon/perigon-go-sdk/option"
)

// VectorService contains methods and other services that help with interacting
// with the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVectorService] method instead.
type VectorService struct {
	Options []option.RequestOption
	News    VectorNewsService
}

// NewVectorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVectorService(opts ...option.RequestOption) (r VectorService) {
	r = VectorService{}
	r.Options = opts
	r.News = NewVectorNewsService(opts...)
	return
}
