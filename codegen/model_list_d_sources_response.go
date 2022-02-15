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

// ListDSourcesResponse struct for ListDSourcesResponse
type ListDSourcesResponse struct {
	Items []DSource `json:"items,omitempty"`
	// Sadly, sometimes requests to the API are not successful. Failures can occur for a wide range of reasons. The Errors object contains information about full or partial failures which might have occurred during the request.
	Errors []Error `json:"errors,omitempty"`
}

// NewListDSourcesResponse instantiates a new ListDSourcesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDSourcesResponse() *ListDSourcesResponse {
	this := ListDSourcesResponse{}
	return &this
}

// NewListDSourcesResponseWithDefaults instantiates a new ListDSourcesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDSourcesResponseWithDefaults() *ListDSourcesResponse {
	this := ListDSourcesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListDSourcesResponse) GetItems() []DSource {
	if o == nil || o.Items == nil {
		var ret []DSource
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDSourcesResponse) GetItemsOk() ([]DSource, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListDSourcesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DSource and assigns it to the Items field.
func (o *ListDSourcesResponse) SetItems(v []DSource) {
	o.Items = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ListDSourcesResponse) GetErrors() []Error {
	if o == nil || o.Errors == nil {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDSourcesResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ListDSourcesResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ListDSourcesResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o ListDSourcesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableListDSourcesResponse struct {
	value *ListDSourcesResponse
	isSet bool
}

func (v NullableListDSourcesResponse) Get() *ListDSourcesResponse {
	return v.value
}

func (v *NullableListDSourcesResponse) Set(val *ListDSourcesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDSourcesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDSourcesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDSourcesResponse(val *ListDSourcesResponse) *NullableListDSourcesResponse {
	return &NullableListDSourcesResponse{value: val, isSet: true}
}

func (v NullableListDSourcesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDSourcesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

