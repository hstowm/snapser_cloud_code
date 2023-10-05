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

// AuthRefreshRequest struct for AuthRefreshRequest
type AuthRefreshRequest struct {
	// Session token to refresh
	SessionToken string `json:"session_token"`
}

// NewAuthRefreshRequest instantiates a new AuthRefreshRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRefreshRequest(sessionToken string) *AuthRefreshRequest {
	this := AuthRefreshRequest{}
	this.SessionToken = sessionToken
	return &this
}

// NewAuthRefreshRequestWithDefaults instantiates a new AuthRefreshRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRefreshRequestWithDefaults() *AuthRefreshRequest {
	this := AuthRefreshRequest{}
	return &this
}

// GetSessionToken returns the SessionToken field value
func (o *AuthRefreshRequest) GetSessionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value
// and a boolean to check if the value has been set.
func (o *AuthRefreshRequest) GetSessionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionToken, true
}

// SetSessionToken sets field value
func (o *AuthRefreshRequest) SetSessionToken(v string) {
	o.SessionToken = v
}

func (o AuthRefreshRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["session_token"] = o.SessionToken
	}
	return json.Marshal(toSerialize)
}

type NullableAuthRefreshRequest struct {
	value *AuthRefreshRequest
	isSet bool
}

func (v NullableAuthRefreshRequest) Get() *AuthRefreshRequest {
	return v.value
}

func (v *NullableAuthRefreshRequest) Set(val *AuthRefreshRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRefreshRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRefreshRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRefreshRequest(val *AuthRefreshRequest) *NullableAuthRefreshRequest {
	return &NullableAuthRefreshRequest{value: val, isSet: true}
}

func (v NullableAuthRefreshRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRefreshRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
