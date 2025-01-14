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

// SetUserStatisticRequest struct for SetUserStatisticRequest
type SetUserStatisticRequest struct {
	// Value to set the statistic to
	Value string `json:"value"`
}

// NewSetUserStatisticRequest instantiates a new SetUserStatisticRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetUserStatisticRequest(value string) *SetUserStatisticRequest {
	this := SetUserStatisticRequest{}
	this.Value = value
	return &this
}

// NewSetUserStatisticRequestWithDefaults instantiates a new SetUserStatisticRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetUserStatisticRequestWithDefaults() *SetUserStatisticRequest {
	this := SetUserStatisticRequest{}
	return &this
}

// GetValue returns the Value field value
func (o *SetUserStatisticRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SetUserStatisticRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SetUserStatisticRequest) SetValue(v string) {
	o.Value = v
}

func (o SetUserStatisticRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSetUserStatisticRequest struct {
	value *SetUserStatisticRequest
	isSet bool
}

func (v NullableSetUserStatisticRequest) Get() *SetUserStatisticRequest {
	return v.value
}

func (v *NullableSetUserStatisticRequest) Set(val *SetUserStatisticRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetUserStatisticRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetUserStatisticRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetUserStatisticRequest(val *SetUserStatisticRequest) *NullableSetUserStatisticRequest {
	return &NullableSetUserStatisticRequest{value: val, isSet: true}
}

func (v NullableSetUserStatisticRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetUserStatisticRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
