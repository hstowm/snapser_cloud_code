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

// SetScoreRequest struct for SetScoreRequest
type SetScoreRequest struct {
	// New score to be inserted
	Score float64 `json:"score"`
	// Name of the tier
	TierName *string `json:"tier_name,omitempty"`
}

// NewSetScoreRequest instantiates a new SetScoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetScoreRequest(score float64) *SetScoreRequest {
	this := SetScoreRequest{}
	this.Score = score
	return &this
}

// NewSetScoreRequestWithDefaults instantiates a new SetScoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetScoreRequestWithDefaults() *SetScoreRequest {
	this := SetScoreRequest{}
	return &this
}

// GetScore returns the Score field value
func (o *SetScoreRequest) GetScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *SetScoreRequest) GetScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *SetScoreRequest) SetScore(v float64) {
	o.Score = v
}

// GetTierName returns the TierName field value if set, zero value otherwise.
func (o *SetScoreRequest) GetTierName() string {
	if o == nil || isNil(o.TierName) {
		var ret string
		return ret
	}
	return *o.TierName
}

// GetTierNameOk returns a tuple with the TierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetScoreRequest) GetTierNameOk() (*string, bool) {
	if o == nil || isNil(o.TierName) {
		return nil, false
	}
	return o.TierName, true
}

// HasTierName returns a boolean if a field has been set.
func (o *SetScoreRequest) HasTierName() bool {
	if o != nil && !isNil(o.TierName) {
		return true
	}

	return false
}

// SetTierName gets a reference to the given string and assigns it to the TierName field.
func (o *SetScoreRequest) SetTierName(v string) {
	o.TierName = &v
}

func (o SetScoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["score"] = o.Score
	}
	if !isNil(o.TierName) {
		toSerialize["tier_name"] = o.TierName
	}
	return json.Marshal(toSerialize)
}

type NullableSetScoreRequest struct {
	value *SetScoreRequest
	isSet bool
}

func (v NullableSetScoreRequest) Get() *SetScoreRequest {
	return v.value
}

func (v *NullableSetScoreRequest) Set(val *SetScoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetScoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetScoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetScoreRequest(val *SetScoreRequest) *NullableSetScoreRequest {
	return &NullableSetScoreRequest{value: val, isSet: true}
}

func (v NullableSetScoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetScoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
