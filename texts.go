package clockogo

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type EntriesTextsMode string

const (
	ExactMatch EntriesTextsMode = "exact_match"
	StartsWith EntriesTextsMode = "starts_with"
	EndsWith   EntriesTextsMode = "ends_with"
	Contains   EntriesTextsMode = "contains"
)

type EntriesTextsSort string

const (
	TextAsc  EntriesTextsMode = "text_asc"
	TextDesc EntriesTextsMode = "text_desc"
	TimeAsc  EntriesTextsMode = "time_asc"
	TimeDesc EntriesTextsMode = "time_desc"
)

type EntriesTextsAPI struct {
	client *Client
}

type EntriesTextsListParams struct {
	Text                    string           `url:"text"`
	Mode                    EntriesTextsMode `url:"mode,omitempty"`
	Sort                    EntriesTextsSort `url:"sort,omitempty"`
	FilterTimeSince         *ISO8601UTC      `url:"filter[time_since],omitempty"`
	FilterTimeUntil         *ISO8601UTC      `url:"filter[time_until],omitempty"`
	FilterUsersId           int              `url:"filter[users_id],omitempty"`
	FilterCustomersId       int              `url:"filter[customers_id],omitempty"`
	FilterProjectsId        int              `url:"filter[projects_id],omitempty"`
	FilterServicesId        int              `url:"filter[services_id],omitempty"`
	FilterLumpsumServicesId int              `url:"filter[lumpsum_services_id],omitempty"`
	FilterBillable          BillableType     `url:"filter[billable],omitempty"`
	Page                    int              `url:"page,omitempty"`
}

type EntriesTexts struct {
	Paging `json:"paging"`
	Filter interface{}      `json:"filter"`
	Mode   EntriesTextsMode `json:"mode"`
	Sort   EntriesTextsSort `json:"sort"`
	Texts  interface{}      `json:"texts"`
}

func (api EntriesTextsAPI) List(q EntriesTextsListParams) (*EntriesTexts, error) {
	params, err := query.Values(&q)
	if err != nil {
		return nil, err
	}
	req, err := api.client.auth.NewRequest(http.MethodGet, BaseURL+"/api/v2/entriesTexts?"+params.Encode(), nil)
	var data EntriesTexts
	err = api.client.Do(req, &data)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
