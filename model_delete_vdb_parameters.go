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

// DeleteVDBParameters Parameters to delete a VDB.
type DeleteVDBParameters struct {
	// Whether to continue the operation upon failures.
	Force *bool `json:"force,omitempty"`
}

// NewDeleteVDBParameters instantiates a new DeleteVDBParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVDBParameters() *DeleteVDBParameters {
	this := DeleteVDBParameters{}
	var force bool = false
	this.Force = &force
	return &this
}

// NewDeleteVDBParametersWithDefaults instantiates a new DeleteVDBParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVDBParametersWithDefaults() *DeleteVDBParameters {
	this := DeleteVDBParameters{}
	var force bool = false
	this.Force = &force
	return &this
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *DeleteVDBParameters) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVDBParameters) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *DeleteVDBParameters) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *DeleteVDBParameters) SetForce(v bool) {
	o.Force = &v
}

func (o DeleteVDBParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteVDBParameters struct {
	value *DeleteVDBParameters
	isSet bool
}

func (v NullableDeleteVDBParameters) Get() *DeleteVDBParameters {
	return v.value
}

func (v *NullableDeleteVDBParameters) Set(val *DeleteVDBParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVDBParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVDBParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVDBParameters(val *DeleteVDBParameters) *NullableDeleteVDBParameters {
	return &NullableDeleteVDBParameters{value: val, isSet: true}
}

func (v NullableDeleteVDBParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVDBParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


