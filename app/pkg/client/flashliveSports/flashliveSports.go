package flashliveSports

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"xenforo/app/pkg/logging"
)

type (
	Client struct {
		token string
		host  string
		ctx   context.Context
	}
	Request struct {
		Data any `json:"Data"`
	}
)

func NewFlashlightSportClient(ctx context.Context, token string) *Client {
	return &Client{
		token: token,
		host:  "flashlive-sport.p.rapidapi.com",
		ctx:   ctx,
	}
}

func (c *Client) Request(method, query string, filter *Filter) ([]byte, error) {
	baseURL := fmt.Sprintf(
		"https://flashlive-sports.p.rapidapi.com/v1/%s?locale=%s&timezone=%s",
		query, filter.Locale, filter.TimeZone,
	)

	req, _ := http.NewRequest(method, baseURL, nil)

	req.Header.Add("X-RapidAPI-Key", c.token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logging.Error(c.ctx, err)
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		logging.Error(c.ctx, err)
		return nil, err
	}

	return resBody, nil
}
