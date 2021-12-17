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

// Optional parameters for the method 'CreateBuild'
type CreateBuildParams struct {
	// The list of Asset Version resource SIDs to include in the Build.
	AssetVersions *[]string `json:"AssetVersions,omitempty"`
	// A list of objects that describe the Dependencies included in the Build. Each object contains the `name` and `version` of the dependency.
	Dependencies *string `json:"Dependencies,omitempty"`
	// The list of the Function Version resource SIDs to include in the Build.
	FunctionVersions *[]string `json:"FunctionVersions,omitempty"`
	// The Runtime version that will be used to run the Build resource when it is deployed.
	Runtime *string `json:"Runtime,omitempty"`
}

func (params *CreateBuildParams) SetAssetVersions(AssetVersions []string) *CreateBuildParams {
	params.AssetVersions = &AssetVersions
	return params
}
func (params *CreateBuildParams) SetDependencies(Dependencies string) *CreateBuildParams {
	params.Dependencies = &Dependencies
	return params
}
func (params *CreateBuildParams) SetFunctionVersions(FunctionVersions []string) *CreateBuildParams {
	params.FunctionVersions = &FunctionVersions
	return params
}
func (params *CreateBuildParams) SetRuntime(Runtime string) *CreateBuildParams {
	params.Runtime = &Runtime
	return params
}

// Create a new Build resource. At least one function version or asset version is required.
func (c *ApiService) CreateBuild(ServiceSid string, params *CreateBuildParams) (*ServerlessV1Build, error) {
	path := "/v1/Services/{ServiceSid}/Builds"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AssetVersions != nil {
		for _, item := range *params.AssetVersions {
			data.Add("AssetVersions", item)
		}
	}
	if params != nil && params.Dependencies != nil {
		data.Set("Dependencies", *params.Dependencies)
	}
	if params != nil && params.FunctionVersions != nil {
		for _, item := range *params.FunctionVersions {
			data.Add("FunctionVersions", item)
		}
	}
	if params != nil && params.Runtime != nil {
		data.Set("Runtime", *params.Runtime)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Build{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Build resource.
func (c *ApiService) DeleteBuild(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Builds/{Sid}"
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

// Retrieve a specific Build resource.
func (c *ApiService) FetchBuild(ServiceSid string, Sid string) (*ServerlessV1Build, error) {
	path := "/v1/Services/{ServiceSid}/Builds/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Build{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListBuild'
type ListBuildParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListBuildParams) SetPageSize(PageSize int) *ListBuildParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBuildParams) SetLimit(Limit int) *ListBuildParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Build records from the API. Request is executed immediately.
func (c *ApiService) PageBuild(ServiceSid string, params *ListBuildParams, pageToken, pageNumber string) (*ListBuildResponse, error) {
	path := "/v1/Services/{ServiceSid}/Builds"

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

	ps := &ListBuildResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Build records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBuild(ServiceSid string, params *ListBuildParams) ([]ServerlessV1Build, error) {
	if params == nil {
		params = &ListBuildParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageBuild(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ServerlessV1Build

	for response != nil {
		records = append(records, response.Builds...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListBuildResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListBuildResponse)
	}

	return records, err
}

// Streams Build records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBuild(ServiceSid string, params *ListBuildParams) (chan ServerlessV1Build, error) {
	if params == nil {
		params = &ListBuildParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageBuild(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ServerlessV1Build, 1)

	go func() {
		for response != nil {
			for item := range response.Builds {
				channel <- response.Builds[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListBuildResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListBuildResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListBuildResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBuildResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
