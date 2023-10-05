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

// AuthValidateRequest struct for AuthValidateRequest
type AuthValidateRequest struct {
	// Session token to validate
	SessionToken string `json:"session_token"`
}

// NewAuthValidateRequest instantiates a new AuthValidateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthValidateRequest(sessionToken string) *AuthValidateRequest {
	this := AuthValidateRequest{}
	this.SessionToken = sessionToken
	return &this
}

// NewAuthValidateRequestWithDefaults instantiates a new AuthValidateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthValidateRequestWithDefaults() *AuthValidateRequest {
	this := AuthValidateRequest{}
	return &this
}

// GetSessionToken returns the SessionToken field value
func (o *AuthValidateRequest) GetSessionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value
// and a boolean to check if the value has been set.
func (o *AuthValidateRequest) GetSessionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionToken, true
}

// SetSessionToken sets field value
func (o *AuthValidateRequest) SetSessionToken(v string) {
	o.SessionToken = v
}

func (o AuthValidateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["session_token"] = o.SessionToken
	}
	return json.Marshal(toSerialize)
}

type NullableAuthValidateRequest struct {
	value *AuthValidateRequest
	isSet bool
}

func (v NullableAuthValidateRequest) Get() *AuthValidateRequest {
	return v.value
}

func (v *NullableAuthValidateRequest) Set(val *AuthValidateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthValidateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthValidateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthValidateRequest(val *AuthValidateRequest) *NullableAuthValidateRequest {
	return &NullableAuthValidateRequest{value: val, isSet: true}
}

func (v NullableAuthValidateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthValidateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
