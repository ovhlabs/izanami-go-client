package izanami_client

import (
	"encoding/json"
	"strconv"
)

// FeaturesResponse represent the http response for listAll
type FeaturesResponse struct {
	Results  []Feature `json:"results"`
	Metadata Metadata  `json:"metadata"`
}

// Feature represents a feature in izanami point of view
type Feature struct {
	ID         string             `json:"id"`
	Enabled    bool               `json:"enabled"`
	Parameters map[string]string  `json:"parameters"`
	Strategy   ActivationStrategy `json:"activationStrategy"`
}

// ActivationStrategy represents the different way to activate a feature
type ActivationStrategy string

const (
	NoStrategy   ActivationStrategy = "NO_STRATEGY"
	ReleaseDate  ActivationStrategy = "RELEASE_DATE"
	Script       ActivationStrategy = "SCRIPT"
	GlobalScript ActivationStrategy = "GLOBAL_SCRIPT"
)

func (c *client) list(page int, pageSize int) (FeaturesResponse, error) {
	var features FeaturesResponse

	httpParams := make(map[string]string)
	httpParams[httpParamPage] = strconv.Itoa(page)
	httpParams[httpParamPageSize] = strconv.Itoa(pageSize)

	res, errListAll := c.get("/features", httpParams)
	if errListAll != nil {
		return features, errListAll
	}

	if err := json.Unmarshal(res, &features); err != nil {
		return features, err
	}
	return features, nil
}

func (c *client) listAll() ([]Feature, error) {
	features := []Feature{}

	currentPage := 0
	pageSize := 20

	for {
		res, err := c.list(currentPage, pageSize)
		if err != nil {
			return features, err
		}
		features = append(features, res.Results...)
		if res.Metadata.Page == res.Metadata.PageSize {
			break
		}

	}
	return features, nil
}
