package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// LocationDetail retrieves detailed information about a specific location area.
func (c *Client) LocationDetail(name string) (RespDetailLocation, error) {
	url := baseURL + "/location-area/" + name
	// if pageURL != nil {
	// 	url = *pageURL
	// }

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespDetailLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespDetailLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailLocation{}, err
	}

	locationsResp := RespDetailLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespDetailLocation{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
