package multivers

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewCountryInfoListGetRequest() CountryInfoListGetRequest {
	r := CountryInfoListGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewCountryInfoListGetQueryParams()
	r.pathParams = r.NewCountryInfoListGetPathParams()
	r.requestBody = r.NewCountryInfoListGetRequestBody()
	return r
}

type CountryInfoListGetRequest struct {
	client      *Client
	queryParams *CountryInfoListGetQueryParams
	pathParams  *CountryInfoListGetPathParams
	method      string
	headers     http.Header
	requestBody CountryInfoListGetRequestBody
}

func (r CountryInfoListGetRequest) NewCountryInfoListGetQueryParams() *CountryInfoListGetQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &CountryInfoListGetQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type CountryInfoListGetQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p CountryInfoListGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CountryInfoListGetRequest) QueryParams() *CountryInfoListGetQueryParams {
	return r.queryParams
}

func (r CountryInfoListGetRequest) NewCountryInfoListGetPathParams() *CountryInfoListGetPathParams {
	return &CountryInfoListGetPathParams{}
}

type CountryInfoListGetPathParams struct {
}

func (p *CountryInfoListGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CountryInfoListGetRequest) PathParams() *CountryInfoListGetPathParams {
	return r.pathParams
}

func (r *CountryInfoListGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CountryInfoListGetRequest) Method() string {
	return r.method
}

func (r CountryInfoListGetRequest) NewCountryInfoListGetRequestBody() CountryInfoListGetRequestBody {
	return CountryInfoListGetRequestBody{}
}

type CountryInfoListGetRequestBody struct{}

func (r *CountryInfoListGetRequest) RequestBody() *CountryInfoListGetRequestBody {
	return &r.requestBody
}

func (r *CountryInfoListGetRequest) SetRequestBody(body CountryInfoListGetRequestBody) {
	r.requestBody = body
}

func (r *CountryInfoListGetRequest) NewResponseBody() *CountryInfoListGetResponseBody {
	return &CountryInfoListGetResponseBody{}
}

type CountryInfoListGetResponseBody []struct {
}

func (r *CountryInfoListGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/CountryInfoList", r.PathParams())
}

func (r *CountryInfoListGetRequest) Do() (CountryInfoListGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
