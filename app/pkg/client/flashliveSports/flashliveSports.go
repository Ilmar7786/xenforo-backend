package flashliveSports

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"xenforo/app/pkg/logging"
)

const (
	prefix             = "/v1"
	QueryParamSport    = "sport_id"
	QueryTimeParamZone = "timezone"
	QueryParamLocale   = "locale"
)

type Client struct {
	token   string
	host    string
	ctx     context.Context
	baseUrl *url.URL
}

func NewFlashlightSportClient(ctx context.Context, token string) *Client {
	baseUrl, _ := url.Parse("https://flashlive-sports.p.rapidapi.com/v1")

	return &Client{
		token:   token,
		host:    "flashlive-sport.p.rapidapi.com",
		ctx:     ctx,
		baseUrl: baseUrl,
	}
}

func (c *Client) DefaultQueryParams() {
	query := c.baseUrl.Query()

	if !query.Has(QueryParamLocale) {
		query.Set(QueryParamLocale, DefaultLocale)
	}

	if !query.Has(QueryTimeParamZone) {
		query.Set(QueryTimeParamZone, DefaultTimeZone)
	}

	c.baseUrl.RawQuery = query.Encode()
}

func (c *Client) AddQuery(key, value string) *Client {
	query := c.baseUrl.Query()
	query.Add(key, value)
	c.baseUrl.RawQuery = query.Encode()

	return c
}

func (c *Client) Request(method, query string, entity interface{}) error {
	c.DefaultQueryParams()

	c.baseUrl.Path = prefix + query
	baseUrlString := c.baseUrl.String()
	req, _ := http.NewRequest(method, baseUrlString, nil)

	req.Header.Add("X-RapidAPI-Key", c.token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logging.Error(c.ctx, err)
		return err
	}
	if res.StatusCode < 200 || res.StatusCode > 300 {
		return errors.New(res.Status)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, &entity)
	if err != nil {
		return err
	}

	c.baseUrl.RawQuery = ""
	return nil
}
