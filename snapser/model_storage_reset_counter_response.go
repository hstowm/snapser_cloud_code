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

// StorageResetCounterResponse struct for StorageResetCounterResponse
type StorageResetCounterResponse struct {
	// Value of the counter
	Count int32 `json:"count"`
}

// NewStorageResetCounterResponse instantiates a new StorageResetCounterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageResetCounterResponse(count int32) *StorageResetCounterResponse {
	this := StorageResetCounterResponse{}
	this.Count = count
	return &this
}

// NewStorageResetCounterResponseWithDefaults instantiates a new StorageResetCounterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageResetCounterResponseWithDefaults() *StorageResetCounterResponse {
	this := StorageResetCounterResponse{}
	return &this
}

// GetCount returns the Count field value
func (o *StorageResetCounterResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *StorageResetCounterResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *StorageResetCounterResponse) SetCount(v int32) {
	o.Count = v
}

func (o StorageResetCounterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableStorageResetCounterResponse struct {
	value *StorageResetCounterResponse
	isSet bool
}

func (v NullableStorageResetCounterResponse) Get() *StorageResetCounterResponse {
	return v.value
}

func (v *NullableStorageResetCounterResponse) Set(val *StorageResetCounterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageResetCounterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageResetCounterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageResetCounterResponse(val *StorageResetCounterResponse) *NullableStorageResetCounterResponse {
	return &NullableStorageResetCounterResponse{value: val, isSet: true}
}

func (v NullableStorageResetCounterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageResetCounterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
