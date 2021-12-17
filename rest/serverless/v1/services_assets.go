/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/NellybettIrahola/twilio-go/client"
)

// Optional parameters for the method 'CreateAsset'
type CreateAssetParams struct {
	// A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateAssetParams) SetFriendlyName(FriendlyName string) *CreateAssetParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new Asset resource.
func (c *ApiService) CreateAsset(ServiceSid string, params *CreateAssetParams) (*ServerlessV1Asset, error) {
	path := "/v1/Services/{ServiceSid}/Assets"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Asset{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete an Asset resource.
func (c *ApiService) DeleteAsset(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Assets/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Retrieve a specific Asset resource.
func (c *ApiService) FetchAsset(ServiceSid string, Sid string) (*ServerlessV1Asset, error) {
	path := "/v1/Services/{ServiceSid}/Assets/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Asset{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAsset'
type ListAssetParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAssetParams) SetPageSize(PageSize int) *ListAssetParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAssetParams) SetLimit(Limit int) *ListAssetParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Asset records from the API. Request is executed immediately.
func (c *ApiService) PageAsset(ServiceSid string, params *ListAssetParams, pageToken, pageNumber string) (*ListAssetResponse, error) {
	path := "/v1/Services/{ServiceSid}/Assets"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAssetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Asset records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAsset(ServiceSid string, params *ListAssetParams) ([]ServerlessV1Asset, error) {
	if params == nil {
		params = &ListAssetParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAsset(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ServerlessV1Asset

	for response != nil {
		records = append(records, response.Assets...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAssetResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListAssetResponse)
	}

	return records, err
}

// Streams Asset records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAsset(ServiceSid string, params *ListAssetParams) (chan ServerlessV1Asset, error) {
	if params == nil {
		params = &ListAssetParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAsset(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ServerlessV1Asset, 1)

	go func() {
		for response != nil {
			for item := range response.Assets {
				channel <- response.Assets[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAssetResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListAssetResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListAssetResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAssetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateAsset'
type UpdateAssetParams struct {
	// A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateAssetParams) SetFriendlyName(FriendlyName string) *UpdateAssetParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Update a specific Asset resource.
func (c *ApiService) UpdateAsset(ServiceSid string, Sid string, params *UpdateAssetParams) (*ServerlessV1Asset, error) {
	path := "/v1/Services/{ServiceSid}/Assets/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Asset{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
