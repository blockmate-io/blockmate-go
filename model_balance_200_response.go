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

// Balance200Response struct for Balance200Response
type Balance200Response struct {
	Balance BalanceResponse `json:"balance"`
}

// NewBalance200Response instantiates a new Balance200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalance200Response(balance BalanceResponse) *Balance200Response {
	this := Balance200Response{}
	this.Balance = balance
	return &this
}

// NewBalance200ResponseWithDefaults instantiates a new Balance200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalance200ResponseWithDefaults() *Balance200Response {
	this := Balance200Response{}
	return &this
}

// GetBalance returns the Balance field value
func (o *Balance200Response) GetBalance() BalanceResponse {
	if o == nil {
		var ret BalanceResponse
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *Balance200Response) GetBalanceOk() (*BalanceResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *Balance200Response) SetBalance(v BalanceResponse) {
	o.Balance = v
}

func (o Balance200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["balance"] = o.Balance
	}
	return json.Marshal(toSerialize)
}

type NullableBalance200Response struct {
	value *Balance200Response
	isSet bool
}

func (v NullableBalance200Response) Get() *Balance200Response {
	return v.value
}

func (v *NullableBalance200Response) Set(val *Balance200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBalance200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBalance200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalance200Response(val *Balance200Response) *NullableBalance200Response {
	return &NullableBalance200Response{value: val, isSet: true}
}

func (v NullableBalance200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalance200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


