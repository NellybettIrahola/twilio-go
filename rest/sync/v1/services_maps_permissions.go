/*
 * Twilio - Sync
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

// Delete a specific Sync Map Permission.
func (c *ApiService) DeleteSyncMapPermission(ServiceSid string, MapSid string, Identity string) error {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific Sync Map Permission.
func (c *ApiService) FetchSyncMapPermission(ServiceSid string, MapSid string, Identity string) (*SyncV1SyncMapPermission, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMapPermission{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSyncMapPermission'
type ListSyncMapPermissionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSyncMapPermissionParams) SetPageSize(PageSize int) *ListSyncMapPermissionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSyncMapPermissionParams) SetLimit(Limit int) *ListSyncMapPermissionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SyncMapPermission records from the API. Request is executed immediately.
func (c *ApiService) PageSyncMapPermission(ServiceSid string, MapSid string, params *ListSyncMapPermissionParams, pageToken, pageNumber string) (*ListSyncMapPermissionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)

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

	ps := &ListSyncMapPermissionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SyncMapPermission records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncMapPermission(ServiceSid string, MapSid string, params *ListSyncMapPermissionParams) ([]SyncV1SyncMapPermission, error) {
	if params == nil {
		params = &ListSyncMapPermissionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncMapPermission(ServiceSid, MapSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []SyncV1SyncMapPermission

	for response != nil {
		records = append(records, response.Permissions...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncMapPermissionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSyncMapPermissionResponse)
	}

	return records, err
}

// Streams SyncMapPermission records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncMapPermission(ServiceSid string, MapSid string, params *ListSyncMapPermissionParams) (chan SyncV1SyncMapPermission, error) {
	if params == nil {
		params = &ListSyncMapPermissionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncMapPermission(ServiceSid, MapSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan SyncV1SyncMapPermission, 1)

	go func() {
		for response != nil {
			for item := range response.Permissions {
				channel <- response.Permissions[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncMapPermissionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSyncMapPermissionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSyncMapPermissionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncMapPermissionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSyncMapPermission'
type UpdateSyncMapPermissionParams struct {
	// Whether the identity can delete the Sync Map. Default value is `false`.
	Manage *bool `json:"Manage,omitempty"`
	// Whether the identity can read the Sync Map and its Items. Default value is `false`.
	Read *bool `json:"Read,omitempty"`
	// Whether the identity can create, update, and delete Items in the Sync Map. Default value is `false`.
	Write *bool `json:"Write,omitempty"`
}

func (params *UpdateSyncMapPermissionParams) SetManage(Manage bool) *UpdateSyncMapPermissionParams {
	params.Manage = &Manage
	return params
}
func (params *UpdateSyncMapPermissionParams) SetRead(Read bool) *UpdateSyncMapPermissionParams {
	params.Read = &Read
	return params
}
func (params *UpdateSyncMapPermissionParams) SetWrite(Write bool) *UpdateSyncMapPermissionParams {
	params.Write = &Write
	return params
}

// Update an identity&#39;s access to a specific Sync Map.
func (c *ApiService) UpdateSyncMapPermission(ServiceSid string, MapSid string, Identity string, params *UpdateSyncMapPermissionParams) (*SyncV1SyncMapPermission, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Manage != nil {
		data.Set("Manage", fmt.Sprint(*params.Manage))
	}
	if params != nil && params.Read != nil {
		data.Set("Read", fmt.Sprint(*params.Read))
	}
	if params != nil && params.Write != nil {
		data.Set("Write", fmt.Sprint(*params.Write))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMapPermission{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
