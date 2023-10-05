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

// StatisticsUserStatistic struct for StatisticsUserStatistic
type StatisticsUserStatistic struct {
	Key    *string `json:"key,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	Value  *string `json:"value,omitempty"`
}

// NewStatisticsUserStatistic instantiates a new StatisticsUserStatistic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticsUserStatistic() *StatisticsUserStatistic {
	this := StatisticsUserStatistic{}
	return &this
}

// NewStatisticsUserStatisticWithDefaults instantiates a new StatisticsUserStatistic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticsUserStatisticWithDefaults() *StatisticsUserStatistic {
	this := StatisticsUserStatistic{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StatisticsUserStatistic) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticsUserStatistic) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StatisticsUserStatistic) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StatisticsUserStatistic) SetKey(v string) {
	o.Key = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *StatisticsUserStatistic) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticsUserStatistic) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *StatisticsUserStatistic) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *StatisticsUserStatistic) SetUserId(v string) {
	o.UserId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StatisticsUserStatistic) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticsUserStatistic) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StatisticsUserStatistic) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *StatisticsUserStatistic) SetValue(v string) {
	o.Value = &v
}

func (o StatisticsUserStatistic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticsUserStatistic struct {
	value *StatisticsUserStatistic
	isSet bool
}

func (v NullableStatisticsUserStatistic) Get() *StatisticsUserStatistic {
	return v.value
}

func (v *NullableStatisticsUserStatistic) Set(val *StatisticsUserStatistic) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticsUserStatistic) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticsUserStatistic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticsUserStatistic(val *StatisticsUserStatistic) *NullableStatisticsUserStatistic {
	return &NullableStatisticsUserStatistic{value: val, isSet: true}
}

func (v NullableStatisticsUserStatistic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticsUserStatistic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
