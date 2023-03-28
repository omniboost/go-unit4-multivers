package multivers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-unit4-multivers/utils"
)

func (c *Client) NewAccountInfoListByYearGetRequest() AccountInfoListByYearGetRequest {
	r := AccountInfoListByYearGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewAccountInfoListByYearGetQueryParams()
	r.pathParams = r.NewAccountInfoListByYearGetPathParams()
	r.requestBody = r.NewAccountInfoListByYearGetRequestBody()
	return r
}

type AccountInfoListByYearGetRequest struct {
	client      *Client
	queryParams *AccountInfoListByYearGetQueryParams
	pathParams  *AccountInfoListByYearGetPathParams
	method      string
	headers     http.Header
	requestBody AccountInfoListByYearGetRequestBody
}

func (r AccountInfoListByYearGetRequest) NewAccountInfoListByYearGetQueryParams() *AccountInfoListByYearGetQueryParams {
	// selectFields, _ := utils.Fields(&Customer{})
	return &AccountInfoListByYearGetQueryParams{
		// Select: odata.NewSelect(selectFields),
		// Filter: odata.NewFilter(),
		// Top:    odata.NewTop(),
		// Skip:   odata.NewSkip(),
	}
}

type AccountInfoListByYearGetQueryParams struct {
	// Select *odata.Select `schema:"$select,omitempty"`
	// Filter *odata.Filter `schema:"$filter,omitempty"`
	// Top    *odata.Top    `schema:"$top,omitempty"`
	// Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p AccountInfoListByYearGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountInfoListByYearGetRequest) QueryParams() *AccountInfoListByYearGetQueryParams {
	return r.queryParams
}

func (r AccountInfoListByYearGetRequest) NewAccountInfoListByYearGetPathParams() *AccountInfoListByYearGetPathParams {
	return &AccountInfoListByYearGetPathParams{}
}

type AccountInfoListByYearGetPathParams struct {
	Year int `schema:"year"`
}

func (p *AccountInfoListByYearGetPathParams) Params() map[string]string {
	return map[string]string{
		"year": strconv.Itoa(p.Year),
	}
}

func (r *AccountInfoListByYearGetRequest) PathParams() *AccountInfoListByYearGetPathParams {
	return r.pathParams
}

func (r *AccountInfoListByYearGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountInfoListByYearGetRequest) Method() string {
	return r.method
}

func (r AccountInfoListByYearGetRequest) NewAccountInfoListByYearGetRequestBody() AccountInfoListByYearGetRequestBody {
	return AccountInfoListByYearGetRequestBody{}
}

type AccountInfoListByYearGetRequestBody struct{}

func (r *AccountInfoListByYearGetRequest) RequestBody() *AccountInfoListByYearGetRequestBody {
	return &r.requestBody
}

func (r *AccountInfoListByYearGetRequest) SetRequestBody(body AccountInfoListByYearGetRequestBody) {
	r.requestBody = body
}

func (r *AccountInfoListByYearGetRequest) NewResponseBody() *AccountInfoListByYearGetResponseBody {
	return &AccountInfoListByYearGetResponseBody{}
}

type AccountInfoListByYearGetResponseBody []struct {
	AdministrationID    string   `json:"administrationId"`
	BackupDate          string   `json:"backupDate"`
	Code                string   `json:"code"`
	Description         string   `json:"description"`
	GroupID             int      `json:"groupId"`
	OnlineState         int      `json:"onlineState"`
	PreviousOnlineState int      `json:"previousOnlineState"`
	ReportPath          string   `json:"reportPath"`
	ShortName           string   `json:"shortName"`
	Users               []string `json:"users"`
	Version             int      `json:"version"`
}

func (r *AccountInfoListByYearGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/{{.administration_id}}/AccountInfoList/{{.year}}", r.PathParams())
}

func (r *AccountInfoListByYearGetRequest) Do() (AccountInfoListByYearGetResponseBody, error) {
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
