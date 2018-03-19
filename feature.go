package client

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FeaturesResponse represent the http response for listAll
type FeaturesResponse struct {
	Results  []FeatureModel `json:"results"`
	Metadata Metadata       `json:"metadata"`
}

// Feature represents a feature in izanami point of view
type FeatureModel struct {
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

func (c *FeatureClient) List(page int, pageSize int) (FeaturesResponse, error) {
	var features FeaturesResponse

	httpParams := make(map[string]string)
	httpParams[httpParamPage] = strconv.Itoa(page)
	httpParams[httpParamPageSize] = strconv.Itoa(pageSize)

	res, errListAll := c.client.get("/features", httpParams)
	if errListAll != nil {
		return features, errListAll
	}

	if err := json.Unmarshal(res, &features); err != nil {
		return features, err
	}
	return features, nil
}

func (c *FeatureClient) ListAll() ([]FeatureModel, error) {
	features := []FeatureModel{}

	currentPage := 0
	pageSize := 20

	for {
		res, err := c.List(currentPage, pageSize)
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

func (c *FeatureClient) Create(feat FeatureModel) error {
	_, errPost := c.client.post("/features", feat)
	if errPost != nil {
		return errPost
	}
	return nil
}

func (c *FeatureClient) Get(id string) (FeatureModel, error) {
	var feature FeatureModel
	body, errGet := c.client.get(fmt.Sprintf("/features/%s", id), nil)
	if errGet != nil {
		return feature, errGet
	}
	if err := json.Unmarshal(body, &feature); err != nil {
		return feature, err
	}
	return feature, nil
}

func (c *FeatureClient) Update(feat FeatureModel) error {
	_, errPut := c.client.put(fmt.Sprintf("/features/%s", feat.ID), feat)
	if errPut != nil {
		return errPut
	}
	return nil
}

func (c *FeatureClient) Delete(id string) error {
	return c.client.delete(fmt.Sprintf("/features/%s", id))
}
