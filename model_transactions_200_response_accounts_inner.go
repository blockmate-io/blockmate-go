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

// Transactions200ResponseAccountsInner struct for Transactions200ResponseAccountsInner
type Transactions200ResponseAccountsInner struct {
	Id string `json:"id"`
	Description string `json:"description"`
	Verified bool `json:"verified"`
	Type string `json:"type"`
	State BalanceResponseAccountsInnerState `json:"state"`
}

// NewTransactions200ResponseAccountsInner instantiates a new Transactions200ResponseAccountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactions200ResponseAccountsInner(id string, description string, verified bool, type_ string, state BalanceResponseAccountsInnerState) *Transactions200ResponseAccountsInner {
	this := Transactions200ResponseAccountsInner{}
	this.Id = id
	this.Description = description
	this.Verified = verified
	this.Type = type_
	this.State = state
	return &this
}

// NewTransactions200ResponseAccountsInnerWithDefaults instantiates a new Transactions200ResponseAccountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactions200ResponseAccountsInnerWithDefaults() *Transactions200ResponseAccountsInner {
	this := Transactions200ResponseAccountsInner{}
	return &this
}

// GetId returns the Id field value
func (o *Transactions200ResponseAccountsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transactions200ResponseAccountsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transactions200ResponseAccountsInner) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *Transactions200ResponseAccountsInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Transactions200ResponseAccountsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Transactions200ResponseAccountsInner) SetDescription(v string) {
	o.Description = v
}

// GetVerified returns the Verified field value
func (o *Transactions200ResponseAccountsInner) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *Transactions200ResponseAccountsInner) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *Transactions200ResponseAccountsInner) SetVerified(v bool) {
	o.Verified = v
}

// GetType returns the Type field value
func (o *Transactions200ResponseAccountsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Transactions200ResponseAccountsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Transactions200ResponseAccountsInner) SetType(v string) {
	o.Type = v
}

// GetState returns the State field value
func (o *Transactions200ResponseAccountsInner) GetState() BalanceResponseAccountsInnerState {
	if o == nil {
		var ret BalanceResponseAccountsInnerState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Transactions200ResponseAccountsInner) GetStateOk() (*BalanceResponseAccountsInnerState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Transactions200ResponseAccountsInner) SetState(v BalanceResponseAccountsInnerState) {
	o.State = v
}

func (o Transactions200ResponseAccountsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["verified"] = o.Verified
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableTransactions200ResponseAccountsInner struct {
	value *Transactions200ResponseAccountsInner
	isSet bool
}

func (v NullableTransactions200ResponseAccountsInner) Get() *Transactions200ResponseAccountsInner {
	return v.value
}

func (v *NullableTransactions200ResponseAccountsInner) Set(val *Transactions200ResponseAccountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactions200ResponseAccountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactions200ResponseAccountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactions200ResponseAccountsInner(val *Transactions200ResponseAccountsInner) *NullableTransactions200ResponseAccountsInner {
	return &NullableTransactions200ResponseAccountsInner{value: val, isSet: true}
}

func (v NullableTransactions200ResponseAccountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactions200ResponseAccountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


