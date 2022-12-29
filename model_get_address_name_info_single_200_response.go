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

// GetAddressNameInfoSingle200Response struct for GetAddressNameInfoSingle200Response
type GetAddressNameInfoSingle200Response struct {
	Name *string `json:"name,omitempty"`
	Category *string `json:"category,omitempty"`
}

// NewGetAddressNameInfoSingle200Response instantiates a new GetAddressNameInfoSingle200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAddressNameInfoSingle200Response() *GetAddressNameInfoSingle200Response {
	this := GetAddressNameInfoSingle200Response{}
	return &this
}

// NewGetAddressNameInfoSingle200ResponseWithDefaults instantiates a new GetAddressNameInfoSingle200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAddressNameInfoSingle200ResponseWithDefaults() *GetAddressNameInfoSingle200Response {
	this := GetAddressNameInfoSingle200Response{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAddressNameInfoSingle200Response) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAddressNameInfoSingle200Response) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAddressNameInfoSingle200Response) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAddressNameInfoSingle200Response) SetName(v string) {
	o.Name = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *GetAddressNameInfoSingle200Response) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAddressNameInfoSingle200Response) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GetAddressNameInfoSingle200Response) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *GetAddressNameInfoSingle200Response) SetCategory(v string) {
	o.Category = &v
}

func (o GetAddressNameInfoSingle200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	return json.Marshal(toSerialize)
}

type NullableGetAddressNameInfoSingle200Response struct {
	value *GetAddressNameInfoSingle200Response
	isSet bool
}

func (v NullableGetAddressNameInfoSingle200Response) Get() *GetAddressNameInfoSingle200Response {
	return v.value
}

func (v *NullableGetAddressNameInfoSingle200Response) Set(val *GetAddressNameInfoSingle200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAddressNameInfoSingle200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAddressNameInfoSingle200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAddressNameInfoSingle200Response(val *GetAddressNameInfoSingle200Response) *NullableGetAddressNameInfoSingle200Response {
	return &NullableGetAddressNameInfoSingle200Response{value: val, isSet: true}
}

func (v NullableGetAddressNameInfoSingle200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAddressNameInfoSingle200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


