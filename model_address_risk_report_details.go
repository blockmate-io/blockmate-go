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

// AddressRiskReportDetails struct for AddressRiskReportDetails
type AddressRiskReportDetails struct {
	OwnCategories []RiskReportCategory `json:"own_categories,omitempty"`
	SourceOfFundsCategories []RiskReportCategory `json:"source_of_funds_categories,omitempty"`
}

// NewAddressRiskReportDetails instantiates a new AddressRiskReportDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressRiskReportDetails() *AddressRiskReportDetails {
	this := AddressRiskReportDetails{}
	return &this
}

// NewAddressRiskReportDetailsWithDefaults instantiates a new AddressRiskReportDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressRiskReportDetailsWithDefaults() *AddressRiskReportDetails {
	this := AddressRiskReportDetails{}
	return &this
}

// GetOwnCategories returns the OwnCategories field value if set, zero value otherwise.
func (o *AddressRiskReportDetails) GetOwnCategories() []RiskReportCategory {
	if o == nil || o.OwnCategories == nil {
		var ret []RiskReportCategory
		return ret
	}
	return o.OwnCategories
}

// GetOwnCategoriesOk returns a tuple with the OwnCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRiskReportDetails) GetOwnCategoriesOk() ([]RiskReportCategory, bool) {
	if o == nil || o.OwnCategories == nil {
		return nil, false
	}
	return o.OwnCategories, true
}

// HasOwnCategories returns a boolean if a field has been set.
func (o *AddressRiskReportDetails) HasOwnCategories() bool {
	if o != nil && o.OwnCategories != nil {
		return true
	}

	return false
}

// SetOwnCategories gets a reference to the given []RiskReportCategory and assigns it to the OwnCategories field.
func (o *AddressRiskReportDetails) SetOwnCategories(v []RiskReportCategory) {
	o.OwnCategories = v
}

// GetSourceOfFundsCategories returns the SourceOfFundsCategories field value if set, zero value otherwise.
func (o *AddressRiskReportDetails) GetSourceOfFundsCategories() []RiskReportCategory {
	if o == nil || o.SourceOfFundsCategories == nil {
		var ret []RiskReportCategory
		return ret
	}
	return o.SourceOfFundsCategories
}

// GetSourceOfFundsCategoriesOk returns a tuple with the SourceOfFundsCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRiskReportDetails) GetSourceOfFundsCategoriesOk() ([]RiskReportCategory, bool) {
	if o == nil || o.SourceOfFundsCategories == nil {
		return nil, false
	}
	return o.SourceOfFundsCategories, true
}

// HasSourceOfFundsCategories returns a boolean if a field has been set.
func (o *AddressRiskReportDetails) HasSourceOfFundsCategories() bool {
	if o != nil && o.SourceOfFundsCategories != nil {
		return true
	}

	return false
}

// SetSourceOfFundsCategories gets a reference to the given []RiskReportCategory and assigns it to the SourceOfFundsCategories field.
func (o *AddressRiskReportDetails) SetSourceOfFundsCategories(v []RiskReportCategory) {
	o.SourceOfFundsCategories = v
}

func (o AddressRiskReportDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnCategories != nil {
		toSerialize["own_categories"] = o.OwnCategories
	}
	if o.SourceOfFundsCategories != nil {
		toSerialize["source_of_funds_categories"] = o.SourceOfFundsCategories
	}
	return json.Marshal(toSerialize)
}

type NullableAddressRiskReportDetails struct {
	value *AddressRiskReportDetails
	isSet bool
}

func (v NullableAddressRiskReportDetails) Get() *AddressRiskReportDetails {
	return v.value
}

func (v *NullableAddressRiskReportDetails) Set(val *AddressRiskReportDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressRiskReportDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressRiskReportDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressRiskReportDetails(val *AddressRiskReportDetails) *NullableAddressRiskReportDetails {
	return &NullableAddressRiskReportDetails{value: val, isSet: true}
}

func (v NullableAddressRiskReportDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressRiskReportDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


