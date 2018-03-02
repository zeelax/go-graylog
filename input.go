package graylog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// InputAttributes represents Input's attributes.
type InputAttributes struct {
	// OverrideSource string `json:"override_source,omitempty"`
	RecvBufferSize      int    `json:"recv_buffer_size,omitempty"`
	BindAddress         string `json:"bind_address,omitempty"`
	Port                int    `json:"port,omitempty"`
	DecompressSizeLimit int    `json:"decompress_size_limit,omitempty"`
}

// InputConfiguration represents Input's configuration.
type InputConfiguration struct {
	BindAddress    string `json:"bind_address,omitempty"`
	Port           int    `json:"port,omitempty"`
	RecvBufferSize int    `json:"recv_buffer_size,omitempty"`
}

// Input represents Graylog Input.
type Input struct {
	Id            string              `json:"id,omitempty"`
	Title         string              `json:"title,omitempty"`
	Type          string              `json:"type,omitempty"`
	Global        bool                `json:"global,omitempty"`
	Node          string              `json:"node,omitempty"`
	CreatedAt     string              `json:"created_at,omitempty"`
	CreatorUserId string              `json:"creator_user_id,omitempty"`
	Attributes    *InputAttributes    `json:"attributes,omitempty"`
	Configuration *InputConfiguration `json:"configuration,omitempty"`
	// ContextPack `json:"context_pack,omitempty"`
	// StaticFields `json:"static_fields,omitempty"`
}

// CreateInput creates an input.
func (client *Client) CreateInput(input *Input) (
	id string, ei *ErrorInfo, err error,
) {
	return client.CreateInputContext(context.Background(), input)
}

// CreateInputContext creates an input with a context.
func (client *Client) CreateInputContext(
	ctx context.Context, input *Input,
) (id string, ei *ErrorInfo, err error) {
	b, err := json.Marshal(input)
	if err != nil {
		return "", nil, errors.Wrap(err, "Failed to json.Marshal(input)")
	}

	ei, err = client.callReq(
		ctx, http.MethodPost, client.endpoints.Inputs, b, true)
	if err != nil {
		return "", ei, err
	}

	ret := &Input{}

	if err := json.Unmarshal(ei.ResponseBody, ret); err != nil {
		return "", ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Input: %s",
				string(ei.ResponseBody)))
	}
	return ret.Id, ei, nil
}

type inputsBody struct {
	Inputs []Input `json:"inputs"`
	Total  int     `json:"total"`
}

// GetInputs returns all inputs.
func (client *Client) GetInputs() ([]Input, *ErrorInfo, error) {
	return client.GetInputsContext(context.Background())
}

// GetInputsContext returns all inputs with a context.
func (client *Client) GetInputsContext(ctx context.Context) (
	[]Input, *ErrorInfo, error,
) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.endpoints.Inputs, nil, true)
	if err != nil {
		return nil, ei, err
	}

	inputs := &inputsBody{}
	err = json.Unmarshal(ei.ResponseBody, inputs)
	if err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Inputs: %s",
				string(ei.ResponseBody)))
	}
	return inputs.Inputs, ei, nil
}

// GetInput returns a given input.
func (client *Client) GetInput(id string) (*Input, *ErrorInfo, error) {
	return client.GetInputContext(context.Background(), id)
}

// GetInputContext returns a given input with a context.
func (client *Client) GetInputContext(
	ctx context.Context, id string,
) (*Input, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet, fmt.Sprintf("%s/%s", client.endpoints.Inputs, id),
		nil, true)
	if err != nil {
		return nil, ei, err
	}

	input := &Input{}
	err = json.Unmarshal(ei.ResponseBody, input)
	if err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Input: %s", string(ei.ResponseBody)))
	}
	return input, ei, nil
}

// UpdateInput updates an given input.
func (client *Client) UpdateInput(id string, input *Input) (
	*Input, *ErrorInfo, error,
) {
	return client.UpdateInputContext(context.Background(), id, input)
}

// UpdateInputContext updates an given input with a context.
func (client *Client) UpdateInputContext(
	ctx context.Context, id string, input *Input,
) (*Input, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	b, err := json.Marshal(input)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to json.Marshal(input)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut, fmt.Sprintf("%s/%s", client.endpoints.Inputs, id),
		b, true)
	if err != nil {
		return nil, ei, err
	}

	ret := &Input{}
	if err := json.Unmarshal(ei.ResponseBody, ret); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Input: %s", string(ei.ResponseBody)))
	}
	return ret, ei, nil
}

// DeleteInput deletes an given input.
func (client *Client) DeleteInput(id string) (*ErrorInfo, error) {
	return client.DeleteInputContext(context.Background(), id)
}

// DeleteInputContext deletes an given input with a context.
func (client *Client) DeleteInputContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	return client.callReq(
		ctx, http.MethodDelete, fmt.Sprintf("%s/%s", client.endpoints.Inputs, id),
		nil, false)
}
