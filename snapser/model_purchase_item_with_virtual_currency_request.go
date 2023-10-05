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

// PurchaseItemWithVirtualCurrencyRequest struct for PurchaseItemWithVirtualCurrencyRequest
type PurchaseItemWithVirtualCurrencyRequest struct {
	// Name of catalog item to purchase
	CatalogItemName *string `json:"catalog_item_name,omitempty"`
	// Name of virtual currency to purchase item with
	CurrencyName *string `json:"currency_name,omitempty"`
}

// NewPurchaseItemWithVirtualCurrencyRequest instantiates a new PurchaseItemWithVirtualCurrencyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseItemWithVirtualCurrencyRequest() *PurchaseItemWithVirtualCurrencyRequest {
	this := PurchaseItemWithVirtualCurrencyRequest{}
	return &this
}

// NewPurchaseItemWithVirtualCurrencyRequestWithDefaults instantiates a new PurchaseItemWithVirtualCurrencyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseItemWithVirtualCurrencyRequestWithDefaults() *PurchaseItemWithVirtualCurrencyRequest {
	this := PurchaseItemWithVirtualCurrencyRequest{}
	return &this
}

// GetCatalogItemName returns the CatalogItemName field value if set, zero value otherwise.
func (o *PurchaseItemWithVirtualCurrencyRequest) GetCatalogItemName() string {
	if o == nil || isNil(o.CatalogItemName) {
		var ret string
		return ret
	}
	return *o.CatalogItemName
}

// GetCatalogItemNameOk returns a tuple with the CatalogItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseItemWithVirtualCurrencyRequest) GetCatalogItemNameOk() (*string, bool) {
	if o == nil || isNil(o.CatalogItemName) {
		return nil, false
	}
	return o.CatalogItemName, true
}

// HasCatalogItemName returns a boolean if a field has been set.
func (o *PurchaseItemWithVirtualCurrencyRequest) HasCatalogItemName() bool {
	if o != nil && !isNil(o.CatalogItemName) {
		return true
	}

	return false
}

// SetCatalogItemName gets a reference to the given string and assigns it to the CatalogItemName field.
func (o *PurchaseItemWithVirtualCurrencyRequest) SetCatalogItemName(v string) {
	o.CatalogItemName = &v
}

// GetCurrencyName returns the CurrencyName field value if set, zero value otherwise.
func (o *PurchaseItemWithVirtualCurrencyRequest) GetCurrencyName() string {
	if o == nil || isNil(o.CurrencyName) {
		var ret string
		return ret
	}
	return *o.CurrencyName
}

// GetCurrencyNameOk returns a tuple with the CurrencyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseItemWithVirtualCurrencyRequest) GetCurrencyNameOk() (*string, bool) {
	if o == nil || isNil(o.CurrencyName) {
		return nil, false
	}
	return o.CurrencyName, true
}

// HasCurrencyName returns a boolean if a field has been set.
func (o *PurchaseItemWithVirtualCurrencyRequest) HasCurrencyName() bool {
	if o != nil && !isNil(o.CurrencyName) {
		return true
	}

	return false
}

// SetCurrencyName gets a reference to the given string and assigns it to the CurrencyName field.
func (o *PurchaseItemWithVirtualCurrencyRequest) SetCurrencyName(v string) {
	o.CurrencyName = &v
}

func (o PurchaseItemWithVirtualCurrencyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CatalogItemName) {
		toSerialize["catalog_item_name"] = o.CatalogItemName
	}
	if !isNil(o.CurrencyName) {
		toSerialize["currency_name"] = o.CurrencyName
	}
	return json.Marshal(toSerialize)
}

type NullablePurchaseItemWithVirtualCurrencyRequest struct {
	value *PurchaseItemWithVirtualCurrencyRequest
	isSet bool
}

func (v NullablePurchaseItemWithVirtualCurrencyRequest) Get() *PurchaseItemWithVirtualCurrencyRequest {
	return v.value
}

func (v *NullablePurchaseItemWithVirtualCurrencyRequest) Set(val *PurchaseItemWithVirtualCurrencyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseItemWithVirtualCurrencyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseItemWithVirtualCurrencyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseItemWithVirtualCurrencyRequest(val *PurchaseItemWithVirtualCurrencyRequest) *NullablePurchaseItemWithVirtualCurrencyRequest {
	return &NullablePurchaseItemWithVirtualCurrencyRequest{value: val, isSet: true}
}

func (v NullablePurchaseItemWithVirtualCurrencyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseItemWithVirtualCurrencyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}