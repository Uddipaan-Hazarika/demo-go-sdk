/*
Delphix API Gateway

Delphix API Gateway API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Job An asynchronous task.
type Job struct {
	// The Job entity ID.
	Id *string `json:"id,omitempty"`
	// The status of the job.
	Status *string `json:"status,omitempty"`
	// The type of job being done.
	Type *string `json:"type,omitempty"`
	// Details about the failure for FAILED jobs.
	ErrorDetails *string `json:"error_details,omitempty"`
}

// NewJob instantiates a new Job object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJob() *Job {
	this := Job{}
	return &this
}

// NewJobWithDefaults instantiates a new Job object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobWithDefaults() *Job {
	this := Job{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Job) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Job) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Job) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Job) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Job) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Job) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Job) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Job) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Job) SetType(v string) {
	o.Type = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *Job) GetErrorDetails() string {
	if o == nil || o.ErrorDetails == nil {
		var ret string
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetErrorDetailsOk() (*string, bool) {
	if o == nil || o.ErrorDetails == nil {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *Job) HasErrorDetails() bool {
	if o != nil && o.ErrorDetails != nil {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given string and assigns it to the ErrorDetails field.
func (o *Job) SetErrorDetails(v string) {
	o.ErrorDetails = &v
}

func (o Job) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ErrorDetails != nil {
		toSerialize["error_details"] = o.ErrorDetails
	}
	return json.Marshal(toSerialize)
}

type NullableJob struct {
	value *Job
	isSet bool
}

func (v NullableJob) Get() *Job {
	return v.value
}

func (v *NullableJob) Set(val *Job) {
	v.value = val
	v.isSet = true
}

func (v NullableJob) IsSet() bool {
	return v.isSet
}

func (v *NullableJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJob(val *Job) *NullableJob {
	return &NullableJob{value: val, isSet: true}
}

func (v NullableJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

