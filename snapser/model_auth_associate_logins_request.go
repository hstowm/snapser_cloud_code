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

// AuthAssociateLoginsRequest struct for AuthAssociateLoginsRequest
type AuthAssociateLoginsRequest struct {
	// Token of the user to discard
	DiscardUserToken string `json:"discard_user_token"`
	// Token of the user to keep
	KeepUserToken string `json:"keep_user_token"`
}

// NewAuthAssociateLoginsRequest instantiates a new AuthAssociateLoginsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAssociateLoginsRequest(discardUserToken string, keepUserToken string) *AuthAssociateLoginsRequest {
	this := AuthAssociateLoginsRequest{}
	this.DiscardUserToken = discardUserToken
	this.KeepUserToken = keepUserToken
	return &this
}

// NewAuthAssociateLoginsRequestWithDefaults instantiates a new AuthAssociateLoginsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAssociateLoginsRequestWithDefaults() *AuthAssociateLoginsRequest {
	this := AuthAssociateLoginsRequest{}
	return &this
}

// GetDiscardUserToken returns the DiscardUserToken field value
func (o *AuthAssociateLoginsRequest) GetDiscardUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscardUserToken
}

// GetDiscardUserTokenOk returns a tuple with the DiscardUserToken field value
// and a boolean to check if the value has been set.
func (o *AuthAssociateLoginsRequest) GetDiscardUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscardUserToken, true
}

// SetDiscardUserToken sets field value
func (o *AuthAssociateLoginsRequest) SetDiscardUserToken(v string) {
	o.DiscardUserToken = v
}

// GetKeepUserToken returns the KeepUserToken field value
func (o *AuthAssociateLoginsRequest) GetKeepUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeepUserToken
}

// GetKeepUserTokenOk returns a tuple with the KeepUserToken field value
// and a boolean to check if the value has been set.
func (o *AuthAssociateLoginsRequest) GetKeepUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeepUserToken, true
}

// SetKeepUserToken sets field value
func (o *AuthAssociateLoginsRequest) SetKeepUserToken(v string) {
	o.KeepUserToken = v
}

func (o AuthAssociateLoginsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["discard_user_token"] = o.DiscardUserToken
	}
	if true {
		toSerialize["keep_user_token"] = o.KeepUserToken
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAssociateLoginsRequest struct {
	value *AuthAssociateLoginsRequest
	isSet bool
}

func (v NullableAuthAssociateLoginsRequest) Get() *AuthAssociateLoginsRequest {
	return v.value
}

func (v *NullableAuthAssociateLoginsRequest) Set(val *AuthAssociateLoginsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAssociateLoginsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAssociateLoginsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAssociateLoginsRequest(val *AuthAssociateLoginsRequest) *NullableAuthAssociateLoginsRequest {
	return &NullableAuthAssociateLoginsRequest{value: val, isSet: true}
}

func (v NullableAuthAssociateLoginsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAssociateLoginsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}