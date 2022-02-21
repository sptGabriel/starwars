package swapi

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"net/url"

	"github.com/sptGabriel/starwars/app/gateway/services/cache"
)

const (
	urlScheme = "https"
	urlHost   = "swapi.dev"
	userAgent = "swapi.go"
)

type Client struct {
	host   *url.URL
	client *http.Client
	cache  cache.Cache
}

func New(cache cache.Cache) Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return Client{
		host: &url.URL{
			Scheme: urlScheme,
			Host:   urlHost,
		},
		client: client,
		cache:  cache,
	}
}

func (c Client) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	URL, err := c.host.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, URL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Add("User-Agent", userAgent)

	return req, nil
}

func (c Client) doRequest(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
