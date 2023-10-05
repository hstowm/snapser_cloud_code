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

// InventoryGetUserItemsResponse struct for InventoryGetUserItemsResponse
type InventoryGetUserItemsResponse struct {
	// Items in user inventory
	Items []InventoryUserInventoryItem `json:"items,omitempty"`
}

// NewInventoryGetUserItemsResponse instantiates a new InventoryGetUserItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryGetUserItemsResponse() *InventoryGetUserItemsResponse {
	this := InventoryGetUserItemsResponse{}
	return &this
}

// NewInventoryGetUserItemsResponseWithDefaults instantiates a new InventoryGetUserItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryGetUserItemsResponseWithDefaults() *InventoryGetUserItemsResponse {
	this := InventoryGetUserItemsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InventoryGetUserItemsResponse) GetItems() []InventoryUserInventoryItem {
	if o == nil || isNil(o.Items) {
		var ret []InventoryUserInventoryItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGetUserItemsResponse) GetItemsOk() ([]InventoryUserInventoryItem, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InventoryGetUserItemsResponse) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InventoryUserInventoryItem and assigns it to the Items field.
func (o *InventoryGetUserItemsResponse) SetItems(v []InventoryUserInventoryItem) {
	o.Items = v
}

func (o InventoryGetUserItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryGetUserItemsResponse struct {
	value *InventoryGetUserItemsResponse
	isSet bool
}

func (v NullableInventoryGetUserItemsResponse) Get() *InventoryGetUserItemsResponse {
	return v.value
}

func (v *NullableInventoryGetUserItemsResponse) Set(val *InventoryGetUserItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryGetUserItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryGetUserItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryGetUserItemsResponse(val *InventoryGetUserItemsResponse) *NullableInventoryGetUserItemsResponse {
	return &NullableInventoryGetUserItemsResponse{value: val, isSet: true}
}

func (v NullableInventoryGetUserItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryGetUserItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
