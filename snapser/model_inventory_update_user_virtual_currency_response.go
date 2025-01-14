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

// InventoryUpdateUserVirtualCurrencyResponse struct for InventoryUpdateUserVirtualCurrencyResponse
type InventoryUpdateUserVirtualCurrencyResponse struct {
	// Current Balance of given currency with user after adding
	CurrentBalance *int32 `json:"current_balance,omitempty"`
	// Balance of given currency with user before adding
	PreviousBalance *int32 `json:"previous_balance,omitempty"`
}

// NewInventoryUpdateUserVirtualCurrencyResponse instantiates a new InventoryUpdateUserVirtualCurrencyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryUpdateUserVirtualCurrencyResponse() *InventoryUpdateUserVirtualCurrencyResponse {
	this := InventoryUpdateUserVirtualCurrencyResponse{}
	return &this
}

// NewInventoryUpdateUserVirtualCurrencyResponseWithDefaults instantiates a new InventoryUpdateUserVirtualCurrencyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryUpdateUserVirtualCurrencyResponseWithDefaults() *InventoryUpdateUserVirtualCurrencyResponse {
	this := InventoryUpdateUserVirtualCurrencyResponse{}
	return &this
}

// GetCurrentBalance returns the CurrentBalance field value if set, zero value otherwise.
func (o *InventoryUpdateUserVirtualCurrencyResponse) GetCurrentBalance() int32 {
	if o == nil || isNil(o.CurrentBalance) {
		var ret int32
		return ret
	}
	return *o.CurrentBalance
}

// GetCurrentBalanceOk returns a tuple with the CurrentBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdateUserVirtualCurrencyResponse) GetCurrentBalanceOk() (*int32, bool) {
	if o == nil || isNil(o.CurrentBalance) {
		return nil, false
	}
	return o.CurrentBalance, true
}

// HasCurrentBalance returns a boolean if a field has been set.
func (o *InventoryUpdateUserVirtualCurrencyResponse) HasCurrentBalance() bool {
	if o != nil && !isNil(o.CurrentBalance) {
		return true
	}

	return false
}

// SetCurrentBalance gets a reference to the given int32 and assigns it to the CurrentBalance field.
func (o *InventoryUpdateUserVirtualCurrencyResponse) SetCurrentBalance(v int32) {
	o.CurrentBalance = &v
}

// GetPreviousBalance returns the PreviousBalance field value if set, zero value otherwise.
func (o *InventoryUpdateUserVirtualCurrencyResponse) GetPreviousBalance() int32 {
	if o == nil || isNil(o.PreviousBalance) {
		var ret int32
		return ret
	}
	return *o.PreviousBalance
}

// GetPreviousBalanceOk returns a tuple with the PreviousBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdateUserVirtualCurrencyResponse) GetPreviousBalanceOk() (*int32, bool) {
	if o == nil || isNil(o.PreviousBalance) {
		return nil, false
	}
	return o.PreviousBalance, true
}

// HasPreviousBalance returns a boolean if a field has been set.
func (o *InventoryUpdateUserVirtualCurrencyResponse) HasPreviousBalance() bool {
	if o != nil && !isNil(o.PreviousBalance) {
		return true
	}

	return false
}

// SetPreviousBalance gets a reference to the given int32 and assigns it to the PreviousBalance field.
func (o *InventoryUpdateUserVirtualCurrencyResponse) SetPreviousBalance(v int32) {
	o.PreviousBalance = &v
}

func (o InventoryUpdateUserVirtualCurrencyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CurrentBalance) {
		toSerialize["current_balance"] = o.CurrentBalance
	}
	if !isNil(o.PreviousBalance) {
		toSerialize["previous_balance"] = o.PreviousBalance
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryUpdateUserVirtualCurrencyResponse struct {
	value *InventoryUpdateUserVirtualCurrencyResponse
	isSet bool
}

func (v NullableInventoryUpdateUserVirtualCurrencyResponse) Get() *InventoryUpdateUserVirtualCurrencyResponse {
	return v.value
}

func (v *NullableInventoryUpdateUserVirtualCurrencyResponse) Set(val *InventoryUpdateUserVirtualCurrencyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryUpdateUserVirtualCurrencyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryUpdateUserVirtualCurrencyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryUpdateUserVirtualCurrencyResponse(val *InventoryUpdateUserVirtualCurrencyResponse) *NullableInventoryUpdateUserVirtualCurrencyResponse {
	return &NullableInventoryUpdateUserVirtualCurrencyResponse{value: val, isSet: true}
}

func (v NullableInventoryUpdateUserVirtualCurrencyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryUpdateUserVirtualCurrencyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
