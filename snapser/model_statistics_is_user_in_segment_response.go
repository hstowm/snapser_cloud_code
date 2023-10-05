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

// StatisticsIsUserInSegmentResponse struct for StatisticsIsUserInSegmentResponse
type StatisticsIsUserInSegmentResponse struct {
	UserInSegment *bool `json:"userInSegment,omitempty"`
}

// NewStatisticsIsUserInSegmentResponse instantiates a new StatisticsIsUserInSegmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticsIsUserInSegmentResponse() *StatisticsIsUserInSegmentResponse {
	this := StatisticsIsUserInSegmentResponse{}
	return &this
}

// NewStatisticsIsUserInSegmentResponseWithDefaults instantiates a new StatisticsIsUserInSegmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticsIsUserInSegmentResponseWithDefaults() *StatisticsIsUserInSegmentResponse {
	this := StatisticsIsUserInSegmentResponse{}
	return &this
}

// GetUserInSegment returns the UserInSegment field value if set, zero value otherwise.
func (o *StatisticsIsUserInSegmentResponse) GetUserInSegment() bool {
	if o == nil || isNil(o.UserInSegment) {
		var ret bool
		return ret
	}
	return *o.UserInSegment
}

// GetUserInSegmentOk returns a tuple with the UserInSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticsIsUserInSegmentResponse) GetUserInSegmentOk() (*bool, bool) {
	if o == nil || isNil(o.UserInSegment) {
		return nil, false
	}
	return o.UserInSegment, true
}

// HasUserInSegment returns a boolean if a field has been set.
func (o *StatisticsIsUserInSegmentResponse) HasUserInSegment() bool {
	if o != nil && !isNil(o.UserInSegment) {
		return true
	}

	return false
}

// SetUserInSegment gets a reference to the given bool and assigns it to the UserInSegment field.
func (o *StatisticsIsUserInSegmentResponse) SetUserInSegment(v bool) {
	o.UserInSegment = &v
}

func (o StatisticsIsUserInSegmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UserInSegment) {
		toSerialize["userInSegment"] = o.UserInSegment
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticsIsUserInSegmentResponse struct {
	value *StatisticsIsUserInSegmentResponse
	isSet bool
}

func (v NullableStatisticsIsUserInSegmentResponse) Get() *StatisticsIsUserInSegmentResponse {
	return v.value
}

func (v *NullableStatisticsIsUserInSegmentResponse) Set(val *StatisticsIsUserInSegmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticsIsUserInSegmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticsIsUserInSegmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticsIsUserInSegmentResponse(val *StatisticsIsUserInSegmentResponse) *NullableStatisticsIsUserInSegmentResponse {
	return &NullableStatisticsIsUserInSegmentResponse{value: val, isSet: true}
}

func (v NullableStatisticsIsUserInSegmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticsIsUserInSegmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
