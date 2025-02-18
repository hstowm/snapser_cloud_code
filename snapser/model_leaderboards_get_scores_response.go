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

// LeaderboardsGetScoresResponse struct for LeaderboardsGetScoresResponse
type LeaderboardsGetScoresResponse struct {
	Tier *LeaderboardsTier `json:"tier,omitempty"`
	// List of users in the order defined by the leaderboard
	UserScores []LeaderboardsUserScore `json:"user_scores,omitempty"`
}

// NewLeaderboardsGetScoresResponse instantiates a new LeaderboardsGetScoresResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLeaderboardsGetScoresResponse() *LeaderboardsGetScoresResponse {
	this := LeaderboardsGetScoresResponse{}
	return &this
}

// NewLeaderboardsGetScoresResponseWithDefaults instantiates a new LeaderboardsGetScoresResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeaderboardsGetScoresResponseWithDefaults() *LeaderboardsGetScoresResponse {
	this := LeaderboardsGetScoresResponse{}
	return &this
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *LeaderboardsGetScoresResponse) GetTier() LeaderboardsTier {
	if o == nil || isNil(o.Tier) {
		var ret LeaderboardsTier
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeaderboardsGetScoresResponse) GetTierOk() (*LeaderboardsTier, bool) {
	if o == nil || isNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *LeaderboardsGetScoresResponse) HasTier() bool {
	if o != nil && !isNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given LeaderboardsTier and assigns it to the Tier field.
func (o *LeaderboardsGetScoresResponse) SetTier(v LeaderboardsTier) {
	o.Tier = &v
}

// GetUserScores returns the UserScores field value if set, zero value otherwise.
func (o *LeaderboardsGetScoresResponse) GetUserScores() []LeaderboardsUserScore {
	if o == nil || isNil(o.UserScores) {
		var ret []LeaderboardsUserScore
		return ret
	}
	return o.UserScores
}

// GetUserScoresOk returns a tuple with the UserScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeaderboardsGetScoresResponse) GetUserScoresOk() ([]LeaderboardsUserScore, bool) {
	if o == nil || isNil(o.UserScores) {
		return nil, false
	}
	return o.UserScores, true
}

// HasUserScores returns a boolean if a field has been set.
func (o *LeaderboardsGetScoresResponse) HasUserScores() bool {
	if o != nil && !isNil(o.UserScores) {
		return true
	}

	return false
}

// SetUserScores gets a reference to the given []LeaderboardsUserScore and assigns it to the UserScores field.
func (o *LeaderboardsGetScoresResponse) SetUserScores(v []LeaderboardsUserScore) {
	o.UserScores = v
}

func (o LeaderboardsGetScoresResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !isNil(o.UserScores) {
		toSerialize["user_scores"] = o.UserScores
	}
	return json.Marshal(toSerialize)
}

type NullableLeaderboardsGetScoresResponse struct {
	value *LeaderboardsGetScoresResponse
	isSet bool
}

func (v NullableLeaderboardsGetScoresResponse) Get() *LeaderboardsGetScoresResponse {
	return v.value
}

func (v *NullableLeaderboardsGetScoresResponse) Set(val *LeaderboardsGetScoresResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLeaderboardsGetScoresResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLeaderboardsGetScoresResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeaderboardsGetScoresResponse(val *LeaderboardsGetScoresResponse) *NullableLeaderboardsGetScoresResponse {
	return &NullableLeaderboardsGetScoresResponse{value: val, isSet: true}
}

func (v NullableLeaderboardsGetScoresResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeaderboardsGetScoresResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
