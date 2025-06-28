package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"stockfetcher/models"
)

const baseURL = "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
const token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6OCwiZW1haWwiOiJhYW5nZWwxOGNAZ21haWwuY29tIiwiZXhwIjoxNzUwMDM2MDc3LCJpZCI6IjAiLCJwYXNzd29yZCI6IicgT1IgJzEnPScxIn0.jB-fTZVi0NHwW031HpZqRqAnmsHHeTYWAsJCyZPTD_s"

func FetchAllStocks() ([]models.Stock, error) {
	var allStocks []models.Stock
	page := ""

	for {
		url := baseURL
		if page != "" {
			url += "?next_page=" + page
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("error status: %d, body: %s", resp.StatusCode, body)
		}

		var apiResp models.APIResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
			return nil, err
		}

		allStocks = append(allStocks, apiResp.Items...)

		if apiResp.NextPage == "" {
			break
		}
		page = apiResp.NextPage
	}

	return allStocks, nil
}
