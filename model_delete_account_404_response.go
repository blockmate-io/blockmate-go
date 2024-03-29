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

// DeleteAccount404Response struct for DeleteAccount404Response
type DeleteAccount404Response struct {
	Detail *string `json:"detail,omitempty"`
}

// NewDeleteAccount404Response instantiates a new DeleteAccount404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAccount404Response() *DeleteAccount404Response {
	this := DeleteAccount404Response{}
	return &this
}

// NewDeleteAccount404ResponseWithDefaults instantiates a new DeleteAccount404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAccount404ResponseWithDefaults() *DeleteAccount404Response {
	this := DeleteAccount404Response{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *DeleteAccount404Response) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAccount404Response) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *DeleteAccount404Response) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *DeleteAccount404Response) SetDetail(v string) {
	o.Detail = &v
}

func (o DeleteAccount404Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteAccount404Response struct {
	value *DeleteAccount404Response
	isSet bool
}

func (v NullableDeleteAccount404Response) Get() *DeleteAccount404Response {
	return v.value
}

func (v *NullableDeleteAccount404Response) Set(val *DeleteAccount404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAccount404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAccount404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAccount404Response(val *DeleteAccount404Response) *NullableDeleteAccount404Response {
	return &NullableDeleteAccount404Response{value: val, isSet: true}
}

func (v NullableDeleteAccount404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAccount404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


