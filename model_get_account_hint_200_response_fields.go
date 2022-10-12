/*
Blockmate

Blockmate API OpenAPI documentation

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blockmate

import (
	"encoding/json"
)

// GetAccountHint200ResponseFields struct for GetAccountHint200ResponseFields
type GetAccountHint200ResponseFields struct {
	Description *string `json:"description,omitempty"`
	Wallet *string `json:"wallet,omitempty"`
}

// NewGetAccountHint200ResponseFields instantiates a new GetAccountHint200ResponseFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountHint200ResponseFields() *GetAccountHint200ResponseFields {
	this := GetAccountHint200ResponseFields{}
	return &this
}

// NewGetAccountHint200ResponseFieldsWithDefaults instantiates a new GetAccountHint200ResponseFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountHint200ResponseFieldsWithDefaults() *GetAccountHint200ResponseFields {
	this := GetAccountHint200ResponseFields{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetAccountHint200ResponseFields) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountHint200ResponseFields) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetAccountHint200ResponseFields) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetAccountHint200ResponseFields) SetDescription(v string) {
	o.Description = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *GetAccountHint200ResponseFields) GetWallet() string {
	if o == nil || o.Wallet == nil {
		var ret string
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountHint200ResponseFields) GetWalletOk() (*string, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *GetAccountHint200ResponseFields) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given string and assigns it to the Wallet field.
func (o *GetAccountHint200ResponseFields) SetWallet(v string) {
	o.Wallet = &v
}

func (o GetAccountHint200ResponseFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Wallet != nil {
		toSerialize["wallet"] = o.Wallet
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccountHint200ResponseFields struct {
	value *GetAccountHint200ResponseFields
	isSet bool
}

func (v NullableGetAccountHint200ResponseFields) Get() *GetAccountHint200ResponseFields {
	return v.value
}

func (v *NullableGetAccountHint200ResponseFields) Set(val *GetAccountHint200ResponseFields) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountHint200ResponseFields) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountHint200ResponseFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountHint200ResponseFields(val *GetAccountHint200ResponseFields) *NullableGetAccountHint200ResponseFields {
	return &NullableGetAccountHint200ResponseFields{value: val, isSet: true}
}

func (v NullableGetAccountHint200ResponseFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountHint200ResponseFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

