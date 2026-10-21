/*
ce script a pour but de fournir un client pour interagir avec une API, dans le projet on utilisera Groq ou Jev pour que les agents puissent prendre des décisions.
Cela nous permettra d'éviter les erreurs 429 si on monte en nombre d'agents, et de gérer les erreurs de manière plus robuste.

j'ai simplifié ce bout de code que j'ai trouvé:
https://github.com/hashicorp/go-retryablehttp/blob/main/client.go

c'est globalement assez basique, y'a juste un retry avec un backoff expo 
*/
package env

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)


type Client struct {
	HTTPClient   *http.Client
	Header       http.Header 
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration
	RetryMax     int
}

func NewClient() *Client {
	return &Client{
		HTTPClient:   &http.Client{},
		Header:       http.Header{},
		RetryWaitMin: time.Second,
		RetryWaitMax: 30 * time.Second,
		RetryMax:     4,
	}
}


func (c *Client) Do(ctx context.Context, method, url string, body []byte) (*http.Response, error) {
	for attempt := 0; ; attempt++ {
		req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
		if err != nil {
			return nil, err
		}

		for k, v := range c.Header {
			req.Header[k] = v
		}

		resp, err := c.HTTPClient.Do(req)
		if !retryable(resp, err) {
			return resp, err
		}

		if attempt == c.RetryMax {
			if err == nil {
				err = fmt.Errorf("unexpected HTTP status %s", resp.Status)
				resp.Body.Close()
			}
			return nil, fmt.Errorf("%s %s giving up after %d attempt(s): %w", method, url, attempt+1, err)
		}

		if resp != nil {
			io.CopyN(io.Discard, resp.Body, 4096) 
			resp.Body.Close()
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(c.backoff(attempt)):
		}
	}
}


func retryable(resp *http.Response, err error) bool {
	if err != nil {
		return true
	}
	code := resp.StatusCode
	return code == http.StatusTooManyRequests || code == 0 ||
		(code >= 500 && code != http.StatusNotImplemented)
}


func (c *Client) backoff(attempt int) time.Duration {
	d := float64(c.RetryWaitMin) * math.Pow(2, float64(attempt))
	if d > float64(c.RetryWaitMax) {
		return c.RetryWaitMax
	}
	return time.Duration(d)
}
