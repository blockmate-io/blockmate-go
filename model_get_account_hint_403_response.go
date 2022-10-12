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

// GetAccountHint403Response struct for GetAccountHint403Response
type GetAccountHint403Response struct {
	Detail *string `json:"detail,omitempty"`
}

// NewGetAccountHint403Response instantiates a new GetAccountHint403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountHint403Response() *GetAccountHint403Response {
	this := GetAccountHint403Response{}
	return &this
}

// NewGetAccountHint403ResponseWithDefaults instantiates a new GetAccountHint403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountHint403ResponseWithDefaults() *GetAccountHint403Response {
	this := GetAccountHint403Response{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *GetAccountHint403Response) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountHint403Response) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *GetAccountHint403Response) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *GetAccountHint403Response) SetDetail(v string) {
	o.Detail = &v
}

func (o GetAccountHint403Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccountHint403Response struct {
	value *GetAccountHint403Response
	isSet bool
}

func (v NullableGetAccountHint403Response) Get() *GetAccountHint403Response {
	return v.value
}

func (v *NullableGetAccountHint403Response) Set(val *GetAccountHint403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountHint403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountHint403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountHint403Response(val *GetAccountHint403Response) *NullableGetAccountHint403Response {
	return &NullableGetAccountHint403Response{value: val, isSet: true}
}

func (v NullableGetAccountHint403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountHint403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

