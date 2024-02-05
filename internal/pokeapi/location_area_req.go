package pokeapi

import (
	"fmt"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
  endpoint := "/location-area?offset=0&limit=20"
  fullURL := baseURL + endpoint


  req, err := http.NewRequest("GET", fullURL, nil)
  if err != nil {
    return LocationAreasResp{}, err
  }

  resp, err := c.httpClient.Do(req)
  if err != nil {
    return LocationAreasResp{}, err
  }

  if resp.StatusCode > 399 {
    return LocationAreasResp{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
  }


  return LocationAreasResp{}, nil
}
