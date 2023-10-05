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

// IncrementCounterRequest struct for IncrementCounterRequest
type IncrementCounterRequest struct {
	// Value to increment the counter by. Use negative to decrement
	Count int32 `json:"count"`
}

// NewIncrementCounterRequest instantiates a new IncrementCounterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncrementCounterRequest(count int32) *IncrementCounterRequest {
	this := IncrementCounterRequest{}
	this.Count = count
	return &this
}

// NewIncrementCounterRequestWithDefaults instantiates a new IncrementCounterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncrementCounterRequestWithDefaults() *IncrementCounterRequest {
	this := IncrementCounterRequest{}
	return &this
}

// GetCount returns the Count field value
func (o *IncrementCounterRequest) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *IncrementCounterRequest) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *IncrementCounterRequest) SetCount(v int32) {
	o.Count = v
}

func (o IncrementCounterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableIncrementCounterRequest struct {
	value *IncrementCounterRequest
	isSet bool
}

func (v NullableIncrementCounterRequest) Get() *IncrementCounterRequest {
	return v.value
}

func (v *NullableIncrementCounterRequest) Set(val *IncrementCounterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncrementCounterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncrementCounterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncrementCounterRequest(val *IncrementCounterRequest) *NullableIncrementCounterRequest {
	return &NullableIncrementCounterRequest{value: val, isSet: true}
}

func (v NullableIncrementCounterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncrementCounterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}