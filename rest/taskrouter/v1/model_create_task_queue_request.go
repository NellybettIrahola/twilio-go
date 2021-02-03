/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateTaskQueueRequest struct for CreateTaskQueueRequest
type CreateTaskQueueRequest struct {
	// The SID of the Activity to assign Workers when a task is assigned to them.
	AssignmentActivitySid string `json:"AssignmentActivitySid,omitempty"`
	// A descriptive string that you create to describe the TaskQueue. For example `Support-Tier 1`, `Sales`, or `Escalation`.
	FriendlyName string `json:"FriendlyName"`
	// The maximum number of Workers to reserve for the assignment of a Task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1.
	MaxReservedWorkers int32 `json:"MaxReservedWorkers,omitempty"`
	// The SID of the Activity to assign Workers when a task is reserved for them.
	ReservationActivitySid string `json:"ReservationActivitySid,omitempty"`
	// A string that describes the Worker selection criteria for any Tasks that enter the TaskQueue. For example, `'\"language\" == \"spanish\"'`. The default value is `1==1`. If this value is empty, Tasks will wait in the TaskQueue until they are deleted or moved to another TaskQueue. For more information about Worker selection, see [Describing Worker selection criteria](https://www.twilio.com/docs/taskrouter/api/taskqueues#target-workers).
	TargetWorkers string `json:"TargetWorkers,omitempty"`
	// How Tasks will be assigned to Workers. Set this parameter to `LIFO` to assign most recently created Task first or FIFO to assign the oldest Task first. Default is `FIFO`. [Click here](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo) to learn more.
	TaskOrder string `json:"TaskOrder,omitempty"`
}