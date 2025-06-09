package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location *string) (RespLocationDetail, error) {
	url := baseURL + "/location-area/" + *location

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocationDetail{}, err
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocationDetail{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return RespLocationDetail{}, err
		}

		c.cache.Add(url, data)
	}

	locationsResp := RespLocationDetail{}
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return RespLocationDetail{}, err
	}

	return locationsResp, nil
}
