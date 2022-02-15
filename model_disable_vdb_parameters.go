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

// DisableVDBParameters Parameters to disable a VDB.
type DisableVDBParameters struct {
	// Whether to attempt a cleanup of the VDB before the disable.
	AttemptCleanup *bool `json:"attempt_cleanup,omitempty"`
}

// NewDisableVDBParameters instantiates a new DisableVDBParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableVDBParameters() *DisableVDBParameters {
	this := DisableVDBParameters{}
	var attemptCleanup bool = true
	this.AttemptCleanup = &attemptCleanup
	return &this
}

// NewDisableVDBParametersWithDefaults instantiates a new DisableVDBParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableVDBParametersWithDefaults() *DisableVDBParameters {
	this := DisableVDBParameters{}
	var attemptCleanup bool = true
	this.AttemptCleanup = &attemptCleanup
	return &this
}

// GetAttemptCleanup returns the AttemptCleanup field value if set, zero value otherwise.
func (o *DisableVDBParameters) GetAttemptCleanup() bool {
	if o == nil || o.AttemptCleanup == nil {
		var ret bool
		return ret
	}
	return *o.AttemptCleanup
}

// GetAttemptCleanupOk returns a tuple with the AttemptCleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableVDBParameters) GetAttemptCleanupOk() (*bool, bool) {
	if o == nil || o.AttemptCleanup == nil {
		return nil, false
	}
	return o.AttemptCleanup, true
}

// HasAttemptCleanup returns a boolean if a field has been set.
func (o *DisableVDBParameters) HasAttemptCleanup() bool {
	if o != nil && o.AttemptCleanup != nil {
		return true
	}

	return false
}

// SetAttemptCleanup gets a reference to the given bool and assigns it to the AttemptCleanup field.
func (o *DisableVDBParameters) SetAttemptCleanup(v bool) {
	o.AttemptCleanup = &v
}

func (o DisableVDBParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttemptCleanup != nil {
		toSerialize["attempt_cleanup"] = o.AttemptCleanup
	}
	return json.Marshal(toSerialize)
}

type NullableDisableVDBParameters struct {
	value *DisableVDBParameters
	isSet bool
}

func (v NullableDisableVDBParameters) Get() *DisableVDBParameters {
	return v.value
}

func (v *NullableDisableVDBParameters) Set(val *DisableVDBParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableVDBParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableVDBParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableVDBParameters(val *DisableVDBParameters) *NullableDisableVDBParameters {
	return &NullableDisableVDBParameters{value: val, isSet: true}
}

func (v NullableDisableVDBParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableVDBParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

