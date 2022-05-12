package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/systemizer/ArendtArchives/backend/locapi"
)

const BASEURL = "https://www.loc.gov/collections/hannah-arendt-papers?fo=json&c=150"

type APIClient struct {
	c http.Client
}

func NewAPIClient() locapi.Client {
	return &APIClient{
		c: http.Client{Timeout: 15 * time.Second},
	}
}

func (c *APIClient) ListCollections() ([]locapi.Collection, error) {
	cols := []locapi.Collection{}
	curUrl := BASEURL

	for curUrl != "" {
		fmt.Printf("Making request to: %v\n", curUrl)
		res, err := c.c.Get(curUrl)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()

		apiData := locapi.APICollectionResponse{}
		json.NewDecoder(res.Body).Decode(&apiData)
		cols = append(cols, apiData.Content.Results...)

		curUrl = apiData.Content.Pagination.Next

		res.Body.Close()

		time.Sleep(time.Second * 2)
	}
	return cols, nil
}

func (c *APIClient) ListCollectionItems(col locapi.Collection) ([]locapi.CollectionItem, error) {
	cols := []locapi.CollectionItem{}
	curUrl := col.Resources[0].Search

	for curUrl != "" {
		fmt.Printf("Making a request: %s\n", curUrl)
		res, err := c.c.Get(curUrl)
		if err != nil {
			return nil, err
		}
		apiData := locapi.APICollectionItemResponse{}
		json.NewDecoder(res.Body).Decode(&apiData)
		cols = append(cols, apiData.Content.Results...)

		curUrl = apiData.Content.Pagination.Next

		res.Body.Close()

		time.Sleep(2 * time.Second)
	}

	return cols, nil
}
