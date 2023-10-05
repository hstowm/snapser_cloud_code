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

// StorageInsertBlobResponse struct for StorageInsertBlobResponse
type StorageInsertBlobResponse struct {
	// CAS value used for future operations
	Cas *string `json:"cas,omitempty"`
}

// NewStorageInsertBlobResponse instantiates a new StorageInsertBlobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageInsertBlobResponse() *StorageInsertBlobResponse {
	this := StorageInsertBlobResponse{}
	return &this
}

// NewStorageInsertBlobResponseWithDefaults instantiates a new StorageInsertBlobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageInsertBlobResponseWithDefaults() *StorageInsertBlobResponse {
	this := StorageInsertBlobResponse{}
	return &this
}

// GetCas returns the Cas field value if set, zero value otherwise.
func (o *StorageInsertBlobResponse) GetCas() string {
	if o == nil || isNil(o.Cas) {
		var ret string
		return ret
	}
	return *o.Cas
}

// GetCasOk returns a tuple with the Cas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageInsertBlobResponse) GetCasOk() (*string, bool) {
	if o == nil || isNil(o.Cas) {
		return nil, false
	}
	return o.Cas, true
}

// HasCas returns a boolean if a field has been set.
func (o *StorageInsertBlobResponse) HasCas() bool {
	if o != nil && !isNil(o.Cas) {
		return true
	}

	return false
}

// SetCas gets a reference to the given string and assigns it to the Cas field.
func (o *StorageInsertBlobResponse) SetCas(v string) {
	o.Cas = &v
}

func (o StorageInsertBlobResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Cas) {
		toSerialize["cas"] = o.Cas
	}
	return json.Marshal(toSerialize)
}

type NullableStorageInsertBlobResponse struct {
	value *StorageInsertBlobResponse
	isSet bool
}

func (v NullableStorageInsertBlobResponse) Get() *StorageInsertBlobResponse {
	return v.value
}

func (v *NullableStorageInsertBlobResponse) Set(val *StorageInsertBlobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageInsertBlobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageInsertBlobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageInsertBlobResponse(val *StorageInsertBlobResponse) *NullableStorageInsertBlobResponse {
	return &NullableStorageInsertBlobResponse{value: val, isSet: true}
}

func (v NullableStorageInsertBlobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageInsertBlobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
