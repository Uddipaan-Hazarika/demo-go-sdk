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

// RefreshVDBByTimestampParametersAllOf struct for RefreshVDBByTimestampParametersAllOf
type RefreshVDBByTimestampParametersAllOf struct {
	// ID of the dataset to refresh to
	DatasetId *string `json:"dataset_id,omitempty"`
}

// NewRefreshVDBByTimestampParametersAllOf instantiates a new RefreshVDBByTimestampParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBByTimestampParametersAllOf() *RefreshVDBByTimestampParametersAllOf {
	this := RefreshVDBByTimestampParametersAllOf{}
	return &this
}

// NewRefreshVDBByTimestampParametersAllOfWithDefaults instantiates a new RefreshVDBByTimestampParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBByTimestampParametersAllOfWithDefaults() *RefreshVDBByTimestampParametersAllOf {
	this := RefreshVDBByTimestampParametersAllOf{}
	return &this
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *RefreshVDBByTimestampParametersAllOf) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBByTimestampParametersAllOf) GetDatasetIdOk() (*string, bool) {
	if o == nil || o.DatasetId == nil {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *RefreshVDBByTimestampParametersAllOf) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *RefreshVDBByTimestampParametersAllOf) SetDatasetId(v string) {
	o.DatasetId = &v
}

func (o RefreshVDBByTimestampParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatasetId != nil {
		toSerialize["dataset_id"] = o.DatasetId
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshVDBByTimestampParametersAllOf struct {
	value *RefreshVDBByTimestampParametersAllOf
	isSet bool
}

func (v NullableRefreshVDBByTimestampParametersAllOf) Get() *RefreshVDBByTimestampParametersAllOf {
	return v.value
}

func (v *NullableRefreshVDBByTimestampParametersAllOf) Set(val *RefreshVDBByTimestampParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBByTimestampParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBByTimestampParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBByTimestampParametersAllOf(val *RefreshVDBByTimestampParametersAllOf) *NullableRefreshVDBByTimestampParametersAllOf {
	return &NullableRefreshVDBByTimestampParametersAllOf{value: val, isSet: true}
}

func (v NullableRefreshVDBByTimestampParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBByTimestampParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

