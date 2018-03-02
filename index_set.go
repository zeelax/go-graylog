package graylog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// IndexSet represents a Graylog's Index Set.
type IndexSet struct {
	Id                              string             `json:"id,omitempty"`
	Title                           string             `json:"title,omitempty"`
	Description                     string             `json:"description,omitempty"`
	IndexPrefix                     string             `json:"index_prefix,omitempty"`
	Shards                          int                `json:"shards,omitempty"`
	Replicas                        int                `json:"replicas,omitempty"`
	RotationStrategyClass           string             `json:"rotation_strategy_class,omitempty"`
	RotationStrategy                *RotationStrategy  `json:"rotation_strategy,omitempty"`
	RetentionStrategyClass          string             `json:"retention_strategy_class,omitempty"`
	RetentionStrategy               *RetentionStrategy `json:"retention_strategy,omitempty"`
	CreationDate                    string             `json:"creation_date,omitempty"`
	IndexAnalyzer                   string             `json:"index_analyzer,omitempty"`
	IndexOptimizationMaxNumSegments int                `json:"index_optimization_max_num_segments,omitempty"`
	IndexOptimizationDisabled       bool               `json:"index_optimization_disabled,omitempty"`
	Writable                        bool               `json:"writable,omitempty"`
	Default                         bool               `json:"default,omitempty"`
}

// IndexSetStats represents a Graylog's Index Set Stats.
type IndexSetStats struct {
	Indices   int `json:"indices"`
	Documents int `json:"documents"`
	Size      int `json:"size"`
}

// RotationStrategy represents a Graylog's Index Set Rotation Strategy.
type RotationStrategy struct {
	Type            string `json:"type,omitempty"`
	MaxDocsPerIndex int    `json:"max_docs_per_index,omitempty"`
}

// RetentionStrategy represents a Graylog's Index Set Retention Strategy.
type RetentionStrategy struct {
	Type               string `json:"type,omitempty"`
	MaxNumberOfIndices int    `json:"max_number_of_indices,omitempty"`
}

type indexSetsBody struct {
	IndexSets []IndexSet     `json:"index_sets"`
	Stats     *IndexSetStats `json:"stats"`
	Total     int            `json:"total"`
}

// GetIndexSets returns a list of all index sets.
func (client *Client) GetIndexSets(
	skip, limit int,
) ([]IndexSet, *IndexSetStats, *ErrorInfo, error) {
	return client.GetIndexSetsContext(context.Background(), skip, limit)
}

// GetIndexSetStatsContext returns a list of all index sets with a context.
func (client *Client) GetIndexSetsContext(
	ctx context.Context, skip, limit int,
) ([]IndexSet, *IndexSetStats, *ErrorInfo, error) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.endpoints.IndexSets, nil, true)
	if err != nil {
		return nil, nil, ei, err
	}
	indexSets := &indexSetsBody{}
	err = json.Unmarshal(ei.ResponseBody, indexSets)
	if err != nil {
		return nil, nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as indexSetsBody: %s",
				string(ei.ResponseBody)))
	}
	return indexSets.IndexSets, indexSets.Stats, ei, nil
}

// GetIndexSet returns a given index set.
func (client *Client) GetIndexSet(id string) (*IndexSet, *ErrorInfo, error) {
	return client.GetIndexSetContext(context.Background(), id)
}

// GetIndexSetContext returns a given index set with a context.
func (client *Client) GetIndexSetContext(
	ctx context.Context, id string,
) (*IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	ei, err := client.callReq(
		ctx, http.MethodGet,
		fmt.Sprintf("%s/%s", client.endpoints.IndexSets, id), nil, true)
	if err != nil {
		return nil, ei, err
	}
	indexSet := &IndexSet{}
	err = json.Unmarshal(ei.ResponseBody, indexSet)
	if err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return indexSet, ei, nil
}

// CreateIndexSet creates a Index Set.
func (client *Client) CreateIndexSet(indexSet *IndexSet) (
	*IndexSet, *ErrorInfo, error,
) {
	return client.CreateIndexSetContext(context.Background(), indexSet)
}

// CreateIndexSetContext creates a Index Set with a context.
func (client *Client) CreateIndexSetContext(
	ctx context.Context, indexSet *IndexSet,
) (*IndexSet, *ErrorInfo, error) {
	b, err := json.Marshal(indexSet)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to json.Marshal(indexSet)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPost, client.endpoints.IndexSets, b, true)
	if err != nil {
		return nil, ei, err
	}

	is := &IndexSet{}
	if err := json.Unmarshal(ei.ResponseBody, is); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return is, ei, nil
}

// UpdateIndexSet updates a given Index Set.
func (client *Client) UpdateIndexSet(
	id string, indexSet *IndexSet,
) (*IndexSet, *ErrorInfo, error) {
	return client.UpdateIndexSetContext(context.Background(), id, indexSet)
}

// UpdateIndexSetContext updates a given Index Set with a context.
func (client *Client) UpdateIndexSetContext(
	ctx context.Context, id string, indexSet *IndexSet,
) (*IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	b, err := json.Marshal(indexSet)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to json.Marshal(indexSet)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut,
		fmt.Sprintf("%s/%s", client.endpoints.IndexSets, id), b, true)
	if err != nil {
		return nil, ei, err
	}

	is := &IndexSet{}
	err = json.Unmarshal(ei.ResponseBody, is)
	if err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return is, ei, nil
}

// DeleteIndexSet deletes a given Index Set.
func (client *Client) DeleteIndexSet(id string) (*ErrorInfo, error) {
	return client.DeleteIndexSetContext(context.Background(), id)
}

// DeleteIndexSet deletes a given Index Set with a context.
func (client *Client) DeleteIndexSetContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	return client.callReq(
		ctx, http.MethodDelete,
		fmt.Sprintf("%s/%s", client.endpoints.IndexSets, id), nil, false)
}

// SetDefaultIndexSet sets default Index Set.
func (client *Client) SetDefaultIndexSet(id string) (
	*IndexSet, *ErrorInfo, error,
) {
	return client.SetDefaultIndexSetContext(context.Background(), id)
}

// SetDefaultIndexSet sets default Index Set with a context.
func (client *Client) SetDefaultIndexSetContext(
	ctx context.Context, id string,
) (*IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut,
		fmt.Sprintf("%s/%s/default", client.endpoints.IndexSets, id),
		nil, true)
	if err != nil {
		return nil, ei, err
	}

	is := &IndexSet{}
	if err := json.Unmarshal(ei.ResponseBody, is); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return is, ei, nil
}

// GetIndexSetStats returns a given Index Set statistics.
func (client *Client) GetIndexSetStats(id string) (
	*IndexSetStats, *ErrorInfo, error,
) {
	return client.GetIndexSetStatsContext(context.Background(), id)
}

// GetIndexSetStatsContext returns a given Index Set statistics with a context.
func (client *Client) GetIndexSetStatsContext(
	ctx context.Context, id string,
) (*IndexSetStats, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet,
		fmt.Sprintf("%s/%s/stats", client.endpoints.IndexSets, id),
		nil, true)
	if err != nil {
		return nil, ei, err
	}

	indexSetStats := &IndexSetStats{}
	err = json.Unmarshal(ei.ResponseBody, indexSetStats)
	if err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSetStats: %s",
				string(ei.ResponseBody)))
	}
	return indexSetStats, ei, nil
}

// GetAllIndexSetsStats returns stats of all Index Sets.
func (client *Client) GetAllIndexSetsStats() (
	*IndexSetStats, *ErrorInfo, error,
) {
	return client.GetAllIndexSetsStatsContext(context.Background())
}

// GetAllIndexSetsStats returns stats of all Index Sets with a context.
func (client *Client) GetAllIndexSetsStatsContext(
	ctx context.Context,
) (*IndexSetStats, *ErrorInfo, error) {

	ei, err := client.callReq(
		ctx, http.MethodGet,
		fmt.Sprintf("%s/stats", client.endpoints.IndexSets),
		nil, true)
	if err != nil {
		return nil, ei, err
	}

	indexSetStats := &IndexSetStats{}
	err = json.Unmarshal(ei.ResponseBody, indexSetStats)
	if err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as IndexSetStats: %s",
				string(ei.ResponseBody)))
	}
	return indexSetStats, ei, nil
}
