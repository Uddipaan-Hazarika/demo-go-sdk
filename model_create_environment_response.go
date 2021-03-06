/*
Delphix API Gateway

Delphix API Gateway API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// CreateEnvironmentResponse struct for CreateEnvironmentResponse
type CreateEnvironmentResponse struct {
	// The initiated job id.
	JobId *string `json:"job_id,omitempty"`
	// The id of environment created.
	EnvironmentId *string `json:"environment_id,omitempty"`
}

// NewCreateEnvironmentResponse instantiates a new CreateEnvironmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEnvironmentResponse() *CreateEnvironmentResponse {
	this := CreateEnvironmentResponse{}
	return &this
}

// NewCreateEnvironmentResponseWithDefaults instantiates a new CreateEnvironmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEnvironmentResponseWithDefaults() *CreateEnvironmentResponse {
	this := CreateEnvironmentResponse{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *CreateEnvironmentResponse) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentResponse) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *CreateEnvironmentResponse) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *CreateEnvironmentResponse) SetJobId(v string) {
	o.JobId = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *CreateEnvironmentResponse) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentResponse) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *CreateEnvironmentResponse) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId != nil {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *CreateEnvironmentResponse) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

func (o CreateEnvironmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobId != nil {
		toSerialize["job_id"] = o.JobId
	}
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEnvironmentResponse struct {
	value *CreateEnvironmentResponse
	isSet bool
}

func (v NullableCreateEnvironmentResponse) Get() *CreateEnvironmentResponse {
	return v.value
}

func (v *NullableCreateEnvironmentResponse) Set(val *CreateEnvironmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEnvironmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEnvironmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEnvironmentResponse(val *CreateEnvironmentResponse) *NullableCreateEnvironmentResponse {
	return &NullableCreateEnvironmentResponse{value: val, isSet: true}
}

func (v NullableCreateEnvironmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEnvironmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


