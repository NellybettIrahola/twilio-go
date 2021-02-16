/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Status the model 'Status'
type Status string

// List of status
const (
	STATUS_ENQUEUED Status = "enqueued"
	STATUS_BUILDING Status = "building"
	STATUS_COMPLETED Status = "completed"
	STATUS_FAILED Status = "failed"
	STATUS_CANCELED Status = "canceled"
)