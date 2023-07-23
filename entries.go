package clockogo

import (
	"github.com/google/go-querystring/query"
	"net/http"
	"strconv"
)

type EntriesAPI struct {
	client *Client
}

type EntryType uint8

const (
	TimeEntry EntryType = iota + 1
	LumpsumValue
	LumpsumService
)

type BillableType uint8

const (
	NotBillable BillableType = iota
	Billable
	AlreadyBilled
)

type BudgetType string

const (
	Strict           BudgetType = "strict"
	StrictCompleted  BudgetType = "strict-completed"
	StrictIncomplete BudgetType = "strict-incomplete"
	Soft             BudgetType = "soft"
	SoftCompleted    BudgetType = "soft-completed"
	SoftIncomplete   BudgetType = "soft-incomplete"
	Without          BudgetType = "without"
	WithoutStrict    BudgetType = "without-strict"
)

type Entry struct {
	Id                     *int          `json:"id" url:"-"`
	CustomersId            int           `json:"customers_id" url:"customers_id"`
	ProjectsId             *int          `json:"projects_id" url:"projects_id,omitempty"`
	UsersId                *int          `json:"users_id" url:"users_id,omitempty"`
	Billable               *BillableType `json:"billable" url:"billable"`
	TextsId                *int          `json:"texts_id" url:"-"`
	TimeSince              ISO8601UTC    `json:"time_since" url:"time_since"`
	TimeUntil              *ISO8601UTC   `json:"time_until" url:"time_until,omitempty"`
	TimeInsert             *ISO8601UTC   `json:"time_insert" url:"-"`
	TimeLastChange         *ISO8601UTC   `json:"time_last_change" url:"-"`
	CustomersName          *string       `json:"customers_name" url:"-"`
	ProjectsName           *string       `json:"projects_name" url:"-"`
	UsersName              *string       `json:"users_name" url:"-"`
	Text                   *string       `json:"text" url:"text,omitempty"`
	Revenue                *float64      `json:"revenue" url:"-"`
	Type                   EntryType     `json:"type" url:"type"`
	ServicesId             *int          `json:"services_id" url:"services_id,omitempty"`
	Duration               *int          `json:"duration" url:"duration,omitempty"`
	Offset                 *int          `json:"offset" url:"-"`
	Clocked                *bool         `json:"clocked" url:"-"`
	ClockedOffline         *bool         `json:"clocked_offline" url:"-"`
	TimeClockedSince       *ISO8601UTC   `json:"time_clocked_since" url:"-"`
	TimeLastChangeWorktime *ISO8601UTC   `json:"time_last_change_worktime" url:"-"`
	HourlyRate             *float64      `json:"hourly_rate" url:"hourly_rate,omitempty"`
	ServicesName           *string       `json:"services_name" url:"-"`
	Lumpsum                *float64      `json:"lumpsum" url:"-"`
	LumpsumServicesId      *int          `json:"lumpsum_services_id" url:"-"`
	LumpsumServicesAmount  *float64      `json:"lumpsum_services_amount" url:"-"`
	LumpsumServicesPrice   *float64      `json:"lumpsum_services_price" url:"-"`
}

func NewTimeEntry(customersId int, servicesId int, billable BillableType, timeSince ISO8601UTC, timeUntil ISO8601UTC) Entry {
	return Entry{
		Type:        TimeEntry,
		CustomersId: customersId,
		ServicesId:  &servicesId,
		Billable:    &billable,
		TimeSince:   timeSince,
		TimeUntil:   &timeUntil,
	}
}

type Entries struct {
	Paging  `json:"paging"`
	Filter  interface{} `json:"filter"`
	Entries []Entry     `json:"entries"`
}

type EntriesListParams struct {
	TimeSince                                 ISO8601UTC   `url:"time_since"`
	TimeUntil                                 ISO8601UTC   `url:"time_until"`
	FilterUsersId                             int          `url:"filter[users_id],omitempty"`
	FilterCustomersId                         int          `url:"filter[customers_id],omitempty"`
	FilterProjectsId                          int          `url:"filter[projects_id],omitempty"`
	FilterServicesId                          int          `url:"filter[services_id],omitempty"`
	FilterLumpsumServicesId                   int          `url:"filter[lumpsum_services_id],omitempty"`
	FilterBillable                            BillableType `url:"filter[billable],omitempty"`
	FilterText                                string       `url:"filter[text],omitempty"`
	FilterTextsId                             int          `url:"filter[texts_id],omitempty"`
	FilterBudgetType                          BudgetType   `url:"filter[budget_type],omitempty"`
	CalcAlsoRevenuesForProjectsWithHardBudget bool         `url:"calc_also_revenues_for_projects_with_hard_budget,omitempty"`
	EnhancedList                              bool         `url:"enhanced_list,omitempty"`
	Page                                      int          `url:"page,omitempty"`
}

func (api EntriesAPI) List(q EntriesListParams) (*Entries, error) {
	params, err := query.Values(&q)
	if err != nil {
		return nil, err
	}
	req, err := api.client.auth.NewRequest(http.MethodGet, BaseURL+"/api/v2/entries?"+params.Encode(), nil)
	var data Entries
	err = api.client.Do(req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type SingleEntry struct {
	Entry `json:"entry"`
}

func (api EntriesAPI) Get(id int) (*Entry, error) {
	req, err := api.client.auth.NewRequest(http.MethodGet, BaseURL+"/api/v2/entries/"+strconv.Itoa(id), nil)
	var data SingleEntry
	err = api.client.Do(req, &data)
	if err != nil {
		return nil, err
	}
	return &data.Entry, nil
}

func (api EntriesAPI) Post(entry Entry) (*Entry, error) {
	params, err := query.Values(&entry)
	if err != nil {
		return nil, err
	}
	req, err := api.client.auth.NewRequest(http.MethodPost, BaseURL+"/api/v2/entries?"+params.Encode(), nil)
	var data SingleEntry
	err = api.client.Do(req, &data)
	if err != nil {
		return nil, err
	}
	return &data.Entry, nil
}
