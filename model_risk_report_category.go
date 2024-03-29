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

// RiskReportCategory RiskReportCategory
type RiskReportCategory struct {
	Adress *string `json:"adress,omitempty"`
	Name string `json:"name"`
	CategoryName string `json:"category_name"`
	Risk int32 `json:"risk"`
}

// NewRiskReportCategory instantiates a new RiskReportCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskReportCategory(name string, categoryName string, risk int32) *RiskReportCategory {
	this := RiskReportCategory{}
	this.Name = name
	this.CategoryName = categoryName
	this.Risk = risk
	return &this
}

// NewRiskReportCategoryWithDefaults instantiates a new RiskReportCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskReportCategoryWithDefaults() *RiskReportCategory {
	this := RiskReportCategory{}
	return &this
}

// GetAdress returns the Adress field value if set, zero value otherwise.
func (o *RiskReportCategory) GetAdress() string {
	if o == nil || o.Adress == nil {
		var ret string
		return ret
	}
	return *o.Adress
}

// GetAdressOk returns a tuple with the Adress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskReportCategory) GetAdressOk() (*string, bool) {
	if o == nil || o.Adress == nil {
		return nil, false
	}
	return o.Adress, true
}

// HasAdress returns a boolean if a field has been set.
func (o *RiskReportCategory) HasAdress() bool {
	if o != nil && o.Adress != nil {
		return true
	}

	return false
}

// SetAdress gets a reference to the given string and assigns it to the Adress field.
func (o *RiskReportCategory) SetAdress(v string) {
	o.Adress = &v
}

// GetName returns the Name field value
func (o *RiskReportCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskReportCategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskReportCategory) SetName(v string) {
	o.Name = v
}

// GetCategoryName returns the CategoryName field value
func (o *RiskReportCategory) GetCategoryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value
// and a boolean to check if the value has been set.
func (o *RiskReportCategory) GetCategoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryName, true
}

// SetCategoryName sets field value
func (o *RiskReportCategory) SetCategoryName(v string) {
	o.CategoryName = v
}

// GetRisk returns the Risk field value
func (o *RiskReportCategory) GetRisk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Risk
}

// GetRiskOk returns a tuple with the Risk field value
// and a boolean to check if the value has been set.
func (o *RiskReportCategory) GetRiskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Risk, true
}

// SetRisk sets field value
func (o *RiskReportCategory) SetRisk(v int32) {
	o.Risk = v
}

func (o RiskReportCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Adress != nil {
		toSerialize["adress"] = o.Adress
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["category_name"] = o.CategoryName
	}
	if true {
		toSerialize["risk"] = o.Risk
	}
	return json.Marshal(toSerialize)
}

type NullableRiskReportCategory struct {
	value *RiskReportCategory
	isSet bool
}

func (v NullableRiskReportCategory) Get() *RiskReportCategory {
	return v.value
}

func (v *NullableRiskReportCategory) Set(val *RiskReportCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskReportCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskReportCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskReportCategory(val *RiskReportCategory) *NullableRiskReportCategory {
	return &NullableRiskReportCategory{value: val, isSet: true}
}

func (v NullableRiskReportCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskReportCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


