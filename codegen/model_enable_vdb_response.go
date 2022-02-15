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

// EnableVDBResponse struct for EnableVDBResponse
type EnableVDBResponse struct {
	// The initiated job id.
	JobId *string `json:"job_id,omitempty"`
}

// NewEnableVDBResponse instantiates a new EnableVDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableVDBResponse() *EnableVDBResponse {
	this := EnableVDBResponse{}
	return &this
}

// NewEnableVDBResponseWithDefaults instantiates a new EnableVDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableVDBResponseWithDefaults() *EnableVDBResponse {
	this := EnableVDBResponse{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *EnableVDBResponse) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableVDBResponse) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *EnableVDBResponse) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *EnableVDBResponse) SetJobId(v string) {
	o.JobId = &v
}

func (o EnableVDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobId != nil {
		toSerialize["job_id"] = o.JobId
	}
	return json.Marshal(toSerialize)
}

type NullableEnableVDBResponse struct {
	value *EnableVDBResponse
	isSet bool
}

func (v NullableEnableVDBResponse) Get() *EnableVDBResponse {
	return v.value
}

func (v *NullableEnableVDBResponse) Set(val *EnableVDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableVDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableVDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableVDBResponse(val *EnableVDBResponse) *NullableEnableVDBResponse {
	return &NullableEnableVDBResponse{value: val, isSet: true}
}

func (v NullableEnableVDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableVDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


