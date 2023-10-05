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

// AuthValidateResponse struct for AuthValidateResponse
type AuthValidateResponse struct {
	User *AuthUser `json:"user,omitempty"`
}

// NewAuthValidateResponse instantiates a new AuthValidateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthValidateResponse() *AuthValidateResponse {
	this := AuthValidateResponse{}
	return &this
}

// NewAuthValidateResponseWithDefaults instantiates a new AuthValidateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthValidateResponseWithDefaults() *AuthValidateResponse {
	this := AuthValidateResponse{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AuthValidateResponse) GetUser() AuthUser {
	if o == nil || isNil(o.User) {
		var ret AuthUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthValidateResponse) GetUserOk() (*AuthUser, bool) {
	if o == nil || isNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AuthValidateResponse) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given AuthUser and assigns it to the User field.
func (o *AuthValidateResponse) SetUser(v AuthUser) {
	o.User = &v
}

func (o AuthValidateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAuthValidateResponse struct {
	value *AuthValidateResponse
	isSet bool
}

func (v NullableAuthValidateResponse) Get() *AuthValidateResponse {
	return v.value
}

func (v *NullableAuthValidateResponse) Set(val *AuthValidateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthValidateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthValidateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthValidateResponse(val *AuthValidateResponse) *NullableAuthValidateResponse {
	return &NullableAuthValidateResponse{value: val, isSet: true}
}

func (v NullableAuthValidateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthValidateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}