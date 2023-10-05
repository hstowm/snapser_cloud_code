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

// StorageDeleteBlobResponse struct for StorageDeleteBlobResponse
type StorageDeleteBlobResponse struct {
	// CAS value used for future operations
	Cas *string `json:"cas,omitempty"`
}

// NewStorageDeleteBlobResponse instantiates a new StorageDeleteBlobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDeleteBlobResponse() *StorageDeleteBlobResponse {
	this := StorageDeleteBlobResponse{}
	return &this
}

// NewStorageDeleteBlobResponseWithDefaults instantiates a new StorageDeleteBlobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDeleteBlobResponseWithDefaults() *StorageDeleteBlobResponse {
	this := StorageDeleteBlobResponse{}
	return &this
}

// GetCas returns the Cas field value if set, zero value otherwise.
func (o *StorageDeleteBlobResponse) GetCas() string {
	if o == nil || isNil(o.Cas) {
		var ret string
		return ret
	}
	return *o.Cas
}

// GetCasOk returns a tuple with the Cas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDeleteBlobResponse) GetCasOk() (*string, bool) {
	if o == nil || isNil(o.Cas) {
		return nil, false
	}
	return o.Cas, true
}

// HasCas returns a boolean if a field has been set.
func (o *StorageDeleteBlobResponse) HasCas() bool {
	if o != nil && !isNil(o.Cas) {
		return true
	}

	return false
}

// SetCas gets a reference to the given string and assigns it to the Cas field.
func (o *StorageDeleteBlobResponse) SetCas(v string) {
	o.Cas = &v
}

func (o StorageDeleteBlobResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Cas) {
		toSerialize["cas"] = o.Cas
	}
	return json.Marshal(toSerialize)
}

type NullableStorageDeleteBlobResponse struct {
	value *StorageDeleteBlobResponse
	isSet bool
}

func (v NullableStorageDeleteBlobResponse) Get() *StorageDeleteBlobResponse {
	return v.value
}

func (v *NullableStorageDeleteBlobResponse) Set(val *StorageDeleteBlobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDeleteBlobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDeleteBlobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDeleteBlobResponse(val *StorageDeleteBlobResponse) *NullableStorageDeleteBlobResponse {
	return &NullableStorageDeleteBlobResponse{value: val, isSet: true}
}

func (v NullableStorageDeleteBlobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDeleteBlobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}