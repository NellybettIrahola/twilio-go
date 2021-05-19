/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSyncMapResponse struct for ListSyncMapResponse
type ListSyncMapResponse struct {
	Maps []SyncV1ServiceSyncMap  `json:"maps,omitempty"`
	Meta ListServiceResponseMeta `json:"meta,omitempty"`
}