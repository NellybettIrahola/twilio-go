/*
 * Twilio - Conversations
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

// Fetch the delivery and read receipts of the conversation message
func (c *ApiService) FetchConversationMessageReceipt(ConversationSid string, MessageSid string, Sid string) (*ConversationsV1ConversationMessageReceipt, error) {
	path := "/v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationMessageReceipt{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConversationMessageReceipt'
type ListConversationMessageReceiptParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConversationMessageReceiptParams) SetPageSize(PageSize int) *ListConversationMessageReceiptParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConversationMessageReceiptParams) SetLimit(Limit int) *ListConversationMessageReceiptParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConversationMessageReceipt records from the API. Request is executed immediately.
func (c *ApiService) PageConversationMessageReceipt(ConversationSid string, MessageSid string, params *ListConversationMessageReceiptParams, pageToken, pageNumber string) (*ListConversationMessageReceiptResponse, error) {
	path := "/v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts"

	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)

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

	ps := &ListConversationMessageReceiptResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConversationMessageReceipt records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConversationMessageReceipt(ConversationSid string, MessageSid string, params *ListConversationMessageReceiptParams) ([]ConversationsV1ConversationMessageReceipt, error) {
	if params == nil {
		params = &ListConversationMessageReceiptParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConversationMessageReceipt(ConversationSid, MessageSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ConversationsV1ConversationMessageReceipt

	for response != nil {
		records = append(records, response.DeliveryReceipts...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConversationMessageReceiptResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListConversationMessageReceiptResponse)
	}

	return records, err
}

// Streams ConversationMessageReceipt records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConversationMessageReceipt(ConversationSid string, MessageSid string, params *ListConversationMessageReceiptParams) (chan ConversationsV1ConversationMessageReceipt, error) {
	if params == nil {
		params = &ListConversationMessageReceiptParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConversationMessageReceipt(ConversationSid, MessageSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ConversationMessageReceipt, 1)

	go func() {
		for response != nil {
			for item := range response.DeliveryReceipts {
				channel <- response.DeliveryReceipts[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConversationMessageReceiptResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListConversationMessageReceiptResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListConversationMessageReceiptResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConversationMessageReceiptResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
