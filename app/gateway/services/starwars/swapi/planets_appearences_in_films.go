package swapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
)

type totalAppearencesResponses struct {
	Results []struct {
		FilmsUrls []string `json:"films"`
	} `json:"results"`
}

func (c Client) PlanetsAppearancesInFilms(ctx context.Context, planetName string) (int, error) {
	fromCache, err := c.cache.Get(planetName)
	if err != nil {
		return 0, err
	}

	if fromCache != nil {
		total, errCache := strconv.Atoi(string(fromCache.([]byte)))
		if errCache != nil {
			return 0, errCache
		}

		return total, nil
	}

	path := fmt.Sprintf("/api/planets/?search=%s", planetName)
	req, err := c.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	q.Set("limit", "1")

	res, err := c.doRequest(req)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != http.StatusOK {
		dumpResp, _ := httputil.DumpResponse(res, true)

		return 0, fmt.Errorf(`call to %s returned status %d: %s`, path, res.StatusCode, dumpResp)
	}

	defer res.Body.Close()

	var data totalAppearencesResponses
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return 0, fmt.Errorf(`decoding body response: %w`, err)
	}

	if len(data.Results) == 0 {
		return 0, nil
	}

	total := len(data.Results[0].FilmsUrls)
	err = c.cache.Save(planetName, []byte(fmt.Sprint(total)))
	if err != nil {
		return 0, err
	}

	return total, nil
}
