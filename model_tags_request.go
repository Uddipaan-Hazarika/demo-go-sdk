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

// TagsRequest struct for TagsRequest
type TagsRequest struct {
	// Array of tags with key value pairs
	Tags []Tag `json:"tags,omitempty"`
}

// NewTagsRequest instantiates a new TagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagsRequest() *TagsRequest {
	this := TagsRequest{}
	return &this
}

// NewTagsRequestWithDefaults instantiates a new TagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagsRequestWithDefaults() *TagsRequest {
	this := TagsRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TagsRequest) GetTags() []Tag {
	if o == nil || o.Tags == nil {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsRequest) GetTagsOk() ([]Tag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TagsRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *TagsRequest) SetTags(v []Tag) {
	o.Tags = v
}

func (o TagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableTagsRequest struct {
	value *TagsRequest
	isSet bool
}

func (v NullableTagsRequest) Get() *TagsRequest {
	return v.value
}

func (v *NullableTagsRequest) Set(val *TagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagsRequest(val *TagsRequest) *NullableTagsRequest {
	return &NullableTagsRequest{value: val, isSet: true}
}

func (v NullableTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


