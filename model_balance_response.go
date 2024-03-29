/*
Blockmate

Blockmate API OpenAPI documentation

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blockmate

import (
	"encoding/json"
)

// BalanceResponse Represents a response from aggregated balance request.
type BalanceResponse struct {
	BalanceSum []Amount `json:"balance_sum"`
	Accounts []BalanceResponseAccountsInner `json:"accounts"`
}

// NewBalanceResponse instantiates a new BalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceResponse(balanceSum []Amount, accounts []BalanceResponseAccountsInner) *BalanceResponse {
	this := BalanceResponse{}
	this.BalanceSum = balanceSum
	this.Accounts = accounts
	return &this
}

// NewBalanceResponseWithDefaults instantiates a new BalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceResponseWithDefaults() *BalanceResponse {
	this := BalanceResponse{}
	return &this
}

// GetBalanceSum returns the BalanceSum field value
func (o *BalanceResponse) GetBalanceSum() []Amount {
	if o == nil {
		var ret []Amount
		return ret
	}

	return o.BalanceSum
}

// GetBalanceSumOk returns a tuple with the BalanceSum field value
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetBalanceSumOk() ([]Amount, bool) {
	if o == nil {
		return nil, false
	}
	return o.BalanceSum, true
}

// SetBalanceSum sets field value
func (o *BalanceResponse) SetBalanceSum(v []Amount) {
	o.BalanceSum = v
}

// GetAccounts returns the Accounts field value
func (o *BalanceResponse) GetAccounts() []BalanceResponseAccountsInner {
	if o == nil {
		var ret []BalanceResponseAccountsInner
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *BalanceResponse) GetAccountsOk() ([]BalanceResponseAccountsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *BalanceResponse) SetAccounts(v []BalanceResponseAccountsInner) {
	o.Accounts = v
}

func (o BalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["balance_sum"] = o.BalanceSum
	}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullableBalanceResponse struct {
	value *BalanceResponse
	isSet bool
}

func (v NullableBalanceResponse) Get() *BalanceResponse {
	return v.value
}

func (v *NullableBalanceResponse) Set(val *BalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceResponse(val *BalanceResponse) *NullableBalanceResponse {
	return &NullableBalanceResponse{value: val, isSet: true}
}

func (v NullableBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


