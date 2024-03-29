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

// GetAddressNameInfoMulti200ResponseValue struct for GetAddressNameInfoMulti200ResponseValue
type GetAddressNameInfoMulti200ResponseValue struct {
	Name *string `json:"name,omitempty"`
	Category *string `json:"category,omitempty"`
}

// NewGetAddressNameInfoMulti200ResponseValue instantiates a new GetAddressNameInfoMulti200ResponseValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAddressNameInfoMulti200ResponseValue() *GetAddressNameInfoMulti200ResponseValue {
	this := GetAddressNameInfoMulti200ResponseValue{}
	return &this
}

// NewGetAddressNameInfoMulti200ResponseValueWithDefaults instantiates a new GetAddressNameInfoMulti200ResponseValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAddressNameInfoMulti200ResponseValueWithDefaults() *GetAddressNameInfoMulti200ResponseValue {
	this := GetAddressNameInfoMulti200ResponseValue{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAddressNameInfoMulti200ResponseValue) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAddressNameInfoMulti200ResponseValue) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAddressNameInfoMulti200ResponseValue) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAddressNameInfoMulti200ResponseValue) SetName(v string) {
	o.Name = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *GetAddressNameInfoMulti200ResponseValue) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAddressNameInfoMulti200ResponseValue) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GetAddressNameInfoMulti200ResponseValue) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *GetAddressNameInfoMulti200ResponseValue) SetCategory(v string) {
	o.Category = &v
}

func (o GetAddressNameInfoMulti200ResponseValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	return json.Marshal(toSerialize)
}

type NullableGetAddressNameInfoMulti200ResponseValue struct {
	value *GetAddressNameInfoMulti200ResponseValue
	isSet bool
}

func (v NullableGetAddressNameInfoMulti200ResponseValue) Get() *GetAddressNameInfoMulti200ResponseValue {
	return v.value
}

func (v *NullableGetAddressNameInfoMulti200ResponseValue) Set(val *GetAddressNameInfoMulti200ResponseValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAddressNameInfoMulti200ResponseValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAddressNameInfoMulti200ResponseValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAddressNameInfoMulti200ResponseValue(val *GetAddressNameInfoMulti200ResponseValue) *NullableGetAddressNameInfoMulti200ResponseValue {
	return &NullableGetAddressNameInfoMulti200ResponseValue{value: val, isSet: true}
}

func (v NullableGetAddressNameInfoMulti200ResponseValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAddressNameInfoMulti200ResponseValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


