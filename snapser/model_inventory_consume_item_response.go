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

// InventoryConsumeItemResponse struct for InventoryConsumeItemResponse
type InventoryConsumeItemResponse struct {
	// Id of the item present in user inventory
	InstanceId *string `json:"instance_id,omitempty"`
	// Timestamp of the last time the item was consumed
	LastConsumedAt *int64 `json:"last_consumed_at,omitempty"`
	// Number of remaining uses for the user inventory item
	RemainingUses *int32 `json:"remaining_uses,omitempty"`
}

// NewInventoryConsumeItemResponse instantiates a new InventoryConsumeItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryConsumeItemResponse() *InventoryConsumeItemResponse {
	this := InventoryConsumeItemResponse{}
	return &this
}

// NewInventoryConsumeItemResponseWithDefaults instantiates a new InventoryConsumeItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryConsumeItemResponseWithDefaults() *InventoryConsumeItemResponse {
	this := InventoryConsumeItemResponse{}
	return &this
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *InventoryConsumeItemResponse) GetInstanceId() string {
	if o == nil || isNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryConsumeItemResponse) GetInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *InventoryConsumeItemResponse) HasInstanceId() bool {
	if o != nil && !isNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *InventoryConsumeItemResponse) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetLastConsumedAt returns the LastConsumedAt field value if set, zero value otherwise.
func (o *InventoryConsumeItemResponse) GetLastConsumedAt() int64 {
	if o == nil || isNil(o.LastConsumedAt) {
		var ret int64
		return ret
	}
	return *o.LastConsumedAt
}

// GetLastConsumedAtOk returns a tuple with the LastConsumedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryConsumeItemResponse) GetLastConsumedAtOk() (*int64, bool) {
	if o == nil || isNil(o.LastConsumedAt) {
		return nil, false
	}
	return o.LastConsumedAt, true
}

// HasLastConsumedAt returns a boolean if a field has been set.
func (o *InventoryConsumeItemResponse) HasLastConsumedAt() bool {
	if o != nil && !isNil(o.LastConsumedAt) {
		return true
	}

	return false
}

// SetLastConsumedAt gets a reference to the given int64 and assigns it to the LastConsumedAt field.
func (o *InventoryConsumeItemResponse) SetLastConsumedAt(v int64) {
	o.LastConsumedAt = &v
}

// GetRemainingUses returns the RemainingUses field value if set, zero value otherwise.
func (o *InventoryConsumeItemResponse) GetRemainingUses() int32 {
	if o == nil || isNil(o.RemainingUses) {
		var ret int32
		return ret
	}
	return *o.RemainingUses
}

// GetRemainingUsesOk returns a tuple with the RemainingUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryConsumeItemResponse) GetRemainingUsesOk() (*int32, bool) {
	if o == nil || isNil(o.RemainingUses) {
		return nil, false
	}
	return o.RemainingUses, true
}

// HasRemainingUses returns a boolean if a field has been set.
func (o *InventoryConsumeItemResponse) HasRemainingUses() bool {
	if o != nil && !isNil(o.RemainingUses) {
		return true
	}

	return false
}

// SetRemainingUses gets a reference to the given int32 and assigns it to the RemainingUses field.
func (o *InventoryConsumeItemResponse) SetRemainingUses(v int32) {
	o.RemainingUses = &v
}

func (o InventoryConsumeItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InstanceId) {
		toSerialize["instance_id"] = o.InstanceId
	}
	if !isNil(o.LastConsumedAt) {
		toSerialize["last_consumed_at"] = o.LastConsumedAt
	}
	if !isNil(o.RemainingUses) {
		toSerialize["remaining_uses"] = o.RemainingUses
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryConsumeItemResponse struct {
	value *InventoryConsumeItemResponse
	isSet bool
}

func (v NullableInventoryConsumeItemResponse) Get() *InventoryConsumeItemResponse {
	return v.value
}

func (v *NullableInventoryConsumeItemResponse) Set(val *InventoryConsumeItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryConsumeItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryConsumeItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryConsumeItemResponse(val *InventoryConsumeItemResponse) *NullableInventoryConsumeItemResponse {
	return &NullableInventoryConsumeItemResponse{value: val, isSet: true}
}

func (v NullableInventoryConsumeItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryConsumeItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
