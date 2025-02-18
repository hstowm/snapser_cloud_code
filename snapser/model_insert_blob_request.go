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

// InsertBlobRequest struct for InsertBlobRequest
type InsertBlobRequest struct {
	// Optional TTL for the blob
	Ttl *int64 `json:"ttl,omitempty"`
	// Blob value to be inserted
	Value string `json:"value"`
}

// NewInsertBlobRequest instantiates a new InsertBlobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsertBlobRequest(value string) *InsertBlobRequest {
	this := InsertBlobRequest{}
	this.Value = value
	return &this
}

// NewInsertBlobRequestWithDefaults instantiates a new InsertBlobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsertBlobRequestWithDefaults() *InsertBlobRequest {
	this := InsertBlobRequest{}
	return &this
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *InsertBlobRequest) GetTtl() int64 {
	if o == nil || isNil(o.Ttl) {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertBlobRequest) GetTtlOk() (*int64, bool) {
	if o == nil || isNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *InsertBlobRequest) HasTtl() bool {
	if o != nil && !isNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *InsertBlobRequest) SetTtl(v int64) {
	o.Ttl = &v
}

// GetValue returns the Value field value
func (o *InsertBlobRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InsertBlobRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InsertBlobRequest) SetValue(v string) {
	o.Value = v
}

func (o InsertBlobRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInsertBlobRequest struct {
	value *InsertBlobRequest
	isSet bool
}

func (v NullableInsertBlobRequest) Get() *InsertBlobRequest {
	return v.value
}

func (v *NullableInsertBlobRequest) Set(val *InsertBlobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInsertBlobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInsertBlobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsertBlobRequest(val *InsertBlobRequest) *NullableInsertBlobRequest {
	return &NullableInsertBlobRequest{value: val, isSet: true}
}

func (v NullableInsertBlobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsertBlobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
