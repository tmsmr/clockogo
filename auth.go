package clockogo

import (
	"io"
	"net/http"
)

const (
	extAppHeader      = "X-Clockodo-External-Application"
	apiUserHeader     = "X-ClockodoApiUser"
	apiKeyHeader      = "X-ClockodoApiKey"
	defaultClientName = "clockogo"
)

type ClientId struct {
	name  string
	email string
}

type Auth interface {
	NewRequest(method string, url string, body io.Reader) (*http.Request, error)
}

type APIKeyAuth struct {
	ClientId
	login string
	key   string
}

func (a APIKeyAuth) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if a.name == "" {
		a.name = defaultClientName
	}
	req.Header.Add(extAppHeader, a.name+":"+a.email)
	req.Header.Add(apiUserHeader, a.login)
	req.Header.Add(apiKeyHeader, a.key)
	return req, nil
}

type BasicAuth struct {
	ClientId
	login    string
	password string
}

func (a BasicAuth) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if a.name == "" {
		a.name = defaultClientName
	}
	req.Header.Add(extAppHeader, a.name+":"+a.email)
	req.SetBasicAuth(a.login, a.password)
	return req, nil
}
