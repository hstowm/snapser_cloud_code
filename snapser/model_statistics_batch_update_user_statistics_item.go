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

// StatisticsBatchUpdateUserStatisticsItem struct for StatisticsBatchUpdateUserStatisticsItem
type StatisticsBatchUpdateUserStatisticsItem struct {
	// Index used to determine the order in which operations are executed
	Idx int64 `json:"idx"`
	// Name of the user statistic
	Key string `json:"key"`
	// Operation to be performed (increment / set)
	Operation string `json:"operation"`
	// Value that the statistic should be set to / incremented by
	Value string `json:"value"`
}

// NewStatisticsBatchUpdateUserStatisticsItem instantiates a new StatisticsBatchUpdateUserStatisticsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticsBatchUpdateUserStatisticsItem(idx int64, key string, operation string, value string) *StatisticsBatchUpdateUserStatisticsItem {
	this := StatisticsBatchUpdateUserStatisticsItem{}
	this.Idx = idx
	this.Key = key
	this.Operation = operation
	this.Value = value
	return &this
}

// NewStatisticsBatchUpdateUserStatisticsItemWithDefaults instantiates a new StatisticsBatchUpdateUserStatisticsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticsBatchUpdateUserStatisticsItemWithDefaults() *StatisticsBatchUpdateUserStatisticsItem {
	this := StatisticsBatchUpdateUserStatisticsItem{}
	return &this
}

// GetIdx returns the Idx field value
func (o *StatisticsBatchUpdateUserStatisticsItem) GetIdx() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Idx
}

// GetIdxOk returns a tuple with the Idx field value
// and a boolean to check if the value has been set.
func (o *StatisticsBatchUpdateUserStatisticsItem) GetIdxOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idx, true
}

// SetIdx sets field value
func (o *StatisticsBatchUpdateUserStatisticsItem) SetIdx(v int64) {
	o.Idx = v
}

// GetKey returns the Key field value
func (o *StatisticsBatchUpdateUserStatisticsItem) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *StatisticsBatchUpdateUserStatisticsItem) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *StatisticsBatchUpdateUserStatisticsItem) SetKey(v string) {
	o.Key = v
}

// GetOperation returns the Operation field value
func (o *StatisticsBatchUpdateUserStatisticsItem) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *StatisticsBatchUpdateUserStatisticsItem) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *StatisticsBatchUpdateUserStatisticsItem) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value
func (o *StatisticsBatchUpdateUserStatisticsItem) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *StatisticsBatchUpdateUserStatisticsItem) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *StatisticsBatchUpdateUserStatisticsItem) SetValue(v string) {
	o.Value = v
}

func (o StatisticsBatchUpdateUserStatisticsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["idx"] = o.Idx
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticsBatchUpdateUserStatisticsItem struct {
	value *StatisticsBatchUpdateUserStatisticsItem
	isSet bool
}

func (v NullableStatisticsBatchUpdateUserStatisticsItem) Get() *StatisticsBatchUpdateUserStatisticsItem {
	return v.value
}

func (v *NullableStatisticsBatchUpdateUserStatisticsItem) Set(val *StatisticsBatchUpdateUserStatisticsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticsBatchUpdateUserStatisticsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticsBatchUpdateUserStatisticsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticsBatchUpdateUserStatisticsItem(val *StatisticsBatchUpdateUserStatisticsItem) *NullableStatisticsBatchUpdateUserStatisticsItem {
	return &NullableStatisticsBatchUpdateUserStatisticsItem{value: val, isSet: true}
}

func (v NullableStatisticsBatchUpdateUserStatisticsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticsBatchUpdateUserStatisticsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
