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

	"github.com/NellybettIrahola/twilio-go/client"
)

// Optional parameters for the method 'ListParticipantConversation'
type ListParticipantConversationParams struct {
	// A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
	Identity *string `json:"Identity,omitempty"`
	// A unique string identifier for the conversation participant who's not a Conversation User. This parameter could be found in messaging_binding.address field of Participant resource. It should be url-encoded.
	Address *string `json:"Address,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListParticipantConversationParams) SetIdentity(Identity string) *ListParticipantConversationParams {
	params.Identity = &Identity
	return params
}
func (params *ListParticipantConversationParams) SetAddress(Address string) *ListParticipantConversationParams {
	params.Address = &Address
	return params
}
func (params *ListParticipantConversationParams) SetPageSize(PageSize int) *ListParticipantConversationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListParticipantConversationParams) SetLimit(Limit int) *ListParticipantConversationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ParticipantConversation records from the API. Request is executed immediately.
func (c *ApiService) PageParticipantConversation(params *ListParticipantConversationParams, pageToken, pageNumber string) (*ListParticipantConversationResponse, error) {
	path := "/v1/ParticipantConversations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.Address != nil {
		data.Set("Address", *params.Address)
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

	ps := &ListParticipantConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ParticipantConversation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListParticipantConversation(params *ListParticipantConversationParams) ([]ConversationsV1ParticipantConversation, error) {
	if params == nil {
		params = &ListParticipantConversationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageParticipantConversation(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ConversationsV1ParticipantConversation

	for response != nil {
		records = append(records, response.Conversations...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListParticipantConversationResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListParticipantConversationResponse)
	}

	return records, err
}

// Streams ParticipantConversation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamParticipantConversation(params *ListParticipantConversationParams) (chan ConversationsV1ParticipantConversation, error) {
	if params == nil {
		params = &ListParticipantConversationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageParticipantConversation(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ParticipantConversation, 1)

	go func() {
		for response != nil {
			for item := range response.Conversations {
				channel <- response.Conversations[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListParticipantConversationResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListParticipantConversationResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListParticipantConversationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListParticipantConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
