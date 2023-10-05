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

// AuthAppleLoginRequest struct for AuthAppleLoginRequest
type AuthAppleLoginRequest struct {
	// Whether to create a user, if it does not exist
	CreateUser *bool `json:"create_user,omitempty"`
	// Apple token generated on the client
	Token string `json:"token"`
}

// NewAuthAppleLoginRequest instantiates a new AuthAppleLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAppleLoginRequest(token string) *AuthAppleLoginRequest {
	this := AuthAppleLoginRequest{}
	this.Token = token
	return &this
}

// NewAuthAppleLoginRequestWithDefaults instantiates a new AuthAppleLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAppleLoginRequestWithDefaults() *AuthAppleLoginRequest {
	this := AuthAppleLoginRequest{}
	return &this
}

// GetCreateUser returns the CreateUser field value if set, zero value otherwise.
func (o *AuthAppleLoginRequest) GetCreateUser() bool {
	if o == nil || isNil(o.CreateUser) {
		var ret bool
		return ret
	}
	return *o.CreateUser
}

// GetCreateUserOk returns a tuple with the CreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAppleLoginRequest) GetCreateUserOk() (*bool, bool) {
	if o == nil || isNil(o.CreateUser) {
		return nil, false
	}
	return o.CreateUser, true
}

// HasCreateUser returns a boolean if a field has been set.
func (o *AuthAppleLoginRequest) HasCreateUser() bool {
	if o != nil && !isNil(o.CreateUser) {
		return true
	}

	return false
}

// SetCreateUser gets a reference to the given bool and assigns it to the CreateUser field.
func (o *AuthAppleLoginRequest) SetCreateUser(v bool) {
	o.CreateUser = &v
}

// GetToken returns the Token field value
func (o *AuthAppleLoginRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AuthAppleLoginRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AuthAppleLoginRequest) SetToken(v string) {
	o.Token = v
}

func (o AuthAppleLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreateUser) {
		toSerialize["create_user"] = o.CreateUser
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAppleLoginRequest struct {
	value *AuthAppleLoginRequest
	isSet bool
}

func (v NullableAuthAppleLoginRequest) Get() *AuthAppleLoginRequest {
	return v.value
}

func (v *NullableAuthAppleLoginRequest) Set(val *AuthAppleLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAppleLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAppleLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAppleLoginRequest(val *AuthAppleLoginRequest) *NullableAuthAppleLoginRequest {
	return &NullableAuthAppleLoginRequest{value: val, isSet: true}
}

func (v NullableAuthAppleLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAppleLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
