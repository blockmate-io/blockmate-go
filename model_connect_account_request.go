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

// ConnectAccountRequest struct for ConnectAccountRequest
type ConnectAccountRequest struct {
	Wallet string `json:"wallet"`
	Description string `json:"description"`
}

// NewConnectAccountRequest instantiates a new ConnectAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectAccountRequest(wallet string, description string) *ConnectAccountRequest {
	this := ConnectAccountRequest{}
	this.Wallet = wallet
	this.Description = description
	return &this
}

// NewConnectAccountRequestWithDefaults instantiates a new ConnectAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectAccountRequestWithDefaults() *ConnectAccountRequest {
	this := ConnectAccountRequest{}
	return &this
}

// GetWallet returns the Wallet field value
func (o *ConnectAccountRequest) GetWallet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value
// and a boolean to check if the value has been set.
func (o *ConnectAccountRequest) GetWalletOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wallet, true
}

// SetWallet sets field value
func (o *ConnectAccountRequest) SetWallet(v string) {
	o.Wallet = v
}

// GetDescription returns the Description field value
func (o *ConnectAccountRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ConnectAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ConnectAccountRequest) SetDescription(v string) {
	o.Description = v
}

func (o ConnectAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wallet"] = o.Wallet
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableConnectAccountRequest struct {
	value *ConnectAccountRequest
	isSet bool
}

func (v NullableConnectAccountRequest) Get() *ConnectAccountRequest {
	return v.value
}

func (v *NullableConnectAccountRequest) Set(val *ConnectAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectAccountRequest(val *ConnectAccountRequest) *NullableConnectAccountRequest {
	return &NullableConnectAccountRequest{value: val, isSet: true}
}

func (v NullableConnectAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


