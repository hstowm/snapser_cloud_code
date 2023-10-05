/*
Roi_dev

Your custom SDK

API version: Roi_dev: v9 SDK
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package snapser

import (
	"encoding/json"
)

// UpdateAppendBlobRequest struct for UpdateAppendBlobRequest
type UpdateAppendBlobRequest struct {
	// Value to be appended
	Value string `json:"value"`
}

// NewUpdateAppendBlobRequest instantiates a new UpdateAppendBlobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAppendBlobRequest(value string) *UpdateAppendBlobRequest {
	this := UpdateAppendBlobRequest{}
	this.Value = value
	return &this
}

// NewUpdateAppendBlobRequestWithDefaults instantiates a new UpdateAppendBlobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAppendBlobRequestWithDefaults() *UpdateAppendBlobRequest {
	this := UpdateAppendBlobRequest{}
	return &this
}

// GetValue returns the Value field value
func (o *UpdateAppendBlobRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateAppendBlobRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdateAppendBlobRequest) SetValue(v string) {
	o.Value = v
}

func (o UpdateAppendBlobRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAppendBlobRequest struct {
	value *UpdateAppendBlobRequest
	isSet bool
}

func (v NullableUpdateAppendBlobRequest) Get() *UpdateAppendBlobRequest {
	return v.value
}

func (v *NullableUpdateAppendBlobRequest) Set(val *UpdateAppendBlobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAppendBlobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAppendBlobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAppendBlobRequest(val *UpdateAppendBlobRequest) *NullableUpdateAppendBlobRequest {
	return &NullableUpdateAppendBlobRequest{value: val, isSet: true}
}

func (v NullableUpdateAppendBlobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAppendBlobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}