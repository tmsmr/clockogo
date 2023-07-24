package clockogo

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	BaseURL = "https://my.clockodo.com"
)

type Client struct {
	hc   *http.Client
	auth Auth

	Entries      *EntriesAPI
	EntriesTexts *EntriesTextsAPI
}

type HTTPClientOption func(c *http.Client)

func WithTimeout(timout time.Duration) HTTPClientOption {
	return func(c *http.Client) {
		c.Timeout = timout
	}
}

func NewClient(auth Auth, opts ...HTTPClientOption) *Client {
	hc := &http.Client{}
	for _, opt := range opts {
		opt(hc)
	}
	c := &Client{
		hc:   hc,
		auth: auth,
	}
	c.Entries = &EntriesAPI{client: c}
	c.EntriesTexts = &EntriesTextsAPI{client: c}
	return c
}

func (c *Client) Do(req *http.Request, v any) error {
	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return NewAPIError(resp.StatusCode, body)
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}
	return nil
}
