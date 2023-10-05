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

// UpsertProfileRequest struct for UpsertProfileRequest
type UpsertProfileRequest struct {
	// JSON representation of the profile being upserted
	Profile map[string]interface{} `json:"profile"`
}

// NewUpsertProfileRequest instantiates a new UpsertProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertProfileRequest(profile map[string]interface{}) *UpsertProfileRequest {
	this := UpsertProfileRequest{}
	this.Profile = profile
	return &this
}

// NewUpsertProfileRequestWithDefaults instantiates a new UpsertProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertProfileRequestWithDefaults() *UpsertProfileRequest {
	this := UpsertProfileRequest{}
	return &this
}

// GetProfile returns the Profile field value
func (o *UpsertProfileRequest) GetProfile() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *UpsertProfileRequest) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Profile, true
}

// SetProfile sets field value
func (o *UpsertProfileRequest) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

func (o UpsertProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profile"] = o.Profile
	}
	return json.Marshal(toSerialize)
}

type NullableUpsertProfileRequest struct {
	value *UpsertProfileRequest
	isSet bool
}

func (v NullableUpsertProfileRequest) Get() *UpsertProfileRequest {
	return v.value
}

func (v *NullableUpsertProfileRequest) Set(val *UpsertProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertProfileRequest(val *UpsertProfileRequest) *NullableUpsertProfileRequest {
	return &NullableUpsertProfileRequest{value: val, isSet: true}
}

func (v NullableUpsertProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
