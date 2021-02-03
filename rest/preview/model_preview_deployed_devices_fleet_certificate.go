/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// PreviewDeployedDevicesFleetCertificate struct for PreviewDeployedDevicesFleetCertificate
type PreviewDeployedDevicesFleetCertificate struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	DeviceSid string `json:"DeviceSid,omitempty"`
	FleetSid string `json:"FleetSid,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Thumbprint string `json:"Thumbprint,omitempty"`
	Url string `json:"Url,omitempty"`
}