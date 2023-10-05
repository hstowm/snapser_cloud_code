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

// AuthRefreshResponse struct for AuthRefreshResponse
type AuthRefreshResponse struct {
	User *AuthUser `json:"user,omitempty"`
}

// NewAuthRefreshResponse instantiates a new AuthRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRefreshResponse() *AuthRefreshResponse {
	this := AuthRefreshResponse{}
	return &this
}

// NewAuthRefreshResponseWithDefaults instantiates a new AuthRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRefreshResponseWithDefaults() *AuthRefreshResponse {
	this := AuthRefreshResponse{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AuthRefreshResponse) GetUser() AuthUser {
	if o == nil || isNil(o.User) {
		var ret AuthUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRefreshResponse) GetUserOk() (*AuthUser, bool) {
	if o == nil || isNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AuthRefreshResponse) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given AuthUser and assigns it to the User field.
func (o *AuthRefreshResponse) SetUser(v AuthUser) {
	o.User = &v
}

func (o AuthRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAuthRefreshResponse struct {
	value *AuthRefreshResponse
	isSet bool
}

func (v NullableAuthRefreshResponse) Get() *AuthRefreshResponse {
	return v.value
}

func (v *NullableAuthRefreshResponse) Set(val *AuthRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRefreshResponse(val *AuthRefreshResponse) *NullableAuthRefreshResponse {
	return &NullableAuthRefreshResponse{value: val, isSet: true}
}

func (v NullableAuthRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}