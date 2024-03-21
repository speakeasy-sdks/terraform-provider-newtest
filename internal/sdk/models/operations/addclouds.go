// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/testing/terraform-provider-newtest/internal/sdk/models/shared"
	"net/http"
)

type AddCloudsRequestBody struct {
	Zone shared.ZoneCreate `json:"zone"`
}

func (o *AddCloudsRequestBody) GetZone() shared.ZoneCreate {
	if o == nil {
		return shared.ZoneCreate{}
	}
	return o.Zone
}

// AddCloudsResponseBody - Successful Request
type AddCloudsResponseBody struct {
	Success *bool        `json:"success,omitempty"`
	Zone    *shared.Zone `json:"zone,omitempty"`
}

func (o *AddCloudsResponseBody) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}

func (o *AddCloudsResponseBody) GetZone() *shared.Zone {
	if o == nil {
		return nil
	}
	return o.Zone
}

type AddCloudsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error Codes
	DefaultError *shared.DefaultError
	// Successful Request
	Object *AddCloudsResponseBody
}

func (o *AddCloudsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddCloudsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddCloudsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddCloudsResponse) GetDefaultError() *shared.DefaultError {
	if o == nil {
		return nil
	}
	return o.DefaultError
}

func (o *AddCloudsResponse) GetObject() *AddCloudsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}