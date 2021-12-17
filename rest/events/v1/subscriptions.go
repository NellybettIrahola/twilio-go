/*
 * Twilio - Events
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

// Optional parameters for the method 'CreateSubscription'
type CreateSubscriptionParams struct {
	// A human readable description for the Subscription **This value should not contain PII.**
	Description *string `json:"Description,omitempty"`
	// The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
	SinkSid *string `json:"SinkSid,omitempty"`
	// An array of objects containing the subscribed Event Types
	Types *[]map[string]interface{} `json:"Types,omitempty"`
}

func (params *CreateSubscriptionParams) SetDescription(Description string) *CreateSubscriptionParams {
	params.Description = &Description
	return params
}
func (params *CreateSubscriptionParams) SetSinkSid(SinkSid string) *CreateSubscriptionParams {
	params.SinkSid = &SinkSid
	return params
}
func (params *CreateSubscriptionParams) SetTypes(Types []map[string]interface{}) *CreateSubscriptionParams {
	params.Types = &Types
	return params
}

// Create a new Subscription.
func (c *ApiService) CreateSubscription(params *CreateSubscriptionParams) (*EventsV1Subscription, error) {
	path := "/v1/Subscriptions"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.SinkSid != nil {
		data.Set("SinkSid", *params.SinkSid)
	}
	if params != nil && params.Types != nil {
		for _, item := range *params.Types {
			v, err := json.Marshal(item)

			if err != nil {
				return nil, err
			}

			data.Add("Types", string(v))
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Subscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Subscription.
func (c *ApiService) DeleteSubscription(Sid string) error {
	path := "/v1/Subscriptions/{Sid}"
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

// Fetch a specific Subscription.
func (c *ApiService) FetchSubscription(Sid string) (*EventsV1Subscription, error) {
	path := "/v1/Subscriptions/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Subscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSubscription'
type ListSubscriptionParams struct {
	// The SID of the sink that the list of Subscriptions should be filtered by.
	SinkSid *string `json:"SinkSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSubscriptionParams) SetSinkSid(SinkSid string) *ListSubscriptionParams {
	params.SinkSid = &SinkSid
	return params
}
func (params *ListSubscriptionParams) SetPageSize(PageSize int) *ListSubscriptionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSubscriptionParams) SetLimit(Limit int) *ListSubscriptionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Subscription records from the API. Request is executed immediately.
func (c *ApiService) PageSubscription(params *ListSubscriptionParams, pageToken, pageNumber string) (*ListSubscriptionResponse, error) {
	path := "/v1/Subscriptions"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.SinkSid != nil {
		data.Set("SinkSid", *params.SinkSid)
	}
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

	ps := &ListSubscriptionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Subscription records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSubscription(params *ListSubscriptionParams) ([]EventsV1Subscription, error) {
	if params == nil {
		params = &ListSubscriptionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSubscription(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []EventsV1Subscription

	for response != nil {
		records = append(records, response.Subscriptions...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSubscriptionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSubscriptionResponse)
	}

	return records, err
}

// Streams Subscription records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSubscription(params *ListSubscriptionParams) (chan EventsV1Subscription, error) {
	if params == nil {
		params = &ListSubscriptionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSubscription(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan EventsV1Subscription, 1)

	go func() {
		for response != nil {
			for item := range response.Subscriptions {
				channel <- response.Subscriptions[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSubscriptionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSubscriptionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSubscriptionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSubscriptionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSubscription'
type UpdateSubscriptionParams struct {
	// A human readable description for the Subscription.
	Description *string `json:"Description,omitempty"`
	// The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
	SinkSid *string `json:"SinkSid,omitempty"`
}

func (params *UpdateSubscriptionParams) SetDescription(Description string) *UpdateSubscriptionParams {
	params.Description = &Description
	return params
}
func (params *UpdateSubscriptionParams) SetSinkSid(SinkSid string) *UpdateSubscriptionParams {
	params.SinkSid = &SinkSid
	return params
}

// Update a Subscription.
func (c *ApiService) UpdateSubscription(Sid string, params *UpdateSubscriptionParams) (*EventsV1Subscription, error) {
	path := "/v1/Subscriptions/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.SinkSid != nil {
		data.Set("SinkSid", *params.SinkSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Subscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
