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

// ListEnginesResponse struct for ListEnginesResponse
type ListEnginesResponse struct {
	Items []Engine `json:"items,omitempty"`
	// Sadly, sometimes requests to the API are not successful. Failures can occur for a wide range of reasons. The Errors object contains information about full or partial failures which might have occurred during the request.
	Errors []Error `json:"errors,omitempty"`
}

// NewListEnginesResponse instantiates a new ListEnginesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEnginesResponse() *ListEnginesResponse {
	this := ListEnginesResponse{}
	return &this
}

// NewListEnginesResponseWithDefaults instantiates a new ListEnginesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEnginesResponseWithDefaults() *ListEnginesResponse {
	this := ListEnginesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListEnginesResponse) GetItems() []Engine {
	if o == nil || o.Items == nil {
		var ret []Engine
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEnginesResponse) GetItemsOk() ([]Engine, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListEnginesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Engine and assigns it to the Items field.
func (o *ListEnginesResponse) SetItems(v []Engine) {
	o.Items = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ListEnginesResponse) GetErrors() []Error {
	if o == nil || o.Errors == nil {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEnginesResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ListEnginesResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ListEnginesResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o ListEnginesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableListEnginesResponse struct {
	value *ListEnginesResponse
	isSet bool
}

func (v NullableListEnginesResponse) Get() *ListEnginesResponse {
	return v.value
}

func (v *NullableListEnginesResponse) Set(val *ListEnginesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListEnginesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListEnginesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEnginesResponse(val *ListEnginesResponse) *NullableListEnginesResponse {
	return &NullableListEnginesResponse{value: val, isSet: true}
}

func (v NullableListEnginesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEnginesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


