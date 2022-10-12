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

// AddressRiskReport AddressRiskReport
type AddressRiskReport struct {
	CaseId *string `json:"case_id,omitempty"`
	RequestDatetime *string `json:"request_datetime,omitempty"`
	ResponseDatetime *string `json:"response_datetime,omitempty"`
	Address string `json:"address"`
	Chain string `json:"chain"`
	Risk int32 `json:"risk"`
	Details AddressRiskReportDetails `json:"details"`
}

// NewAddressRiskReport instantiates a new AddressRiskReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressRiskReport(address string, chain string, risk int32, details AddressRiskReportDetails) *AddressRiskReport {
	this := AddressRiskReport{}
	this.Address = address
	this.Chain = chain
	this.Risk = risk
	this.Details = details
	return &this
}

// NewAddressRiskReportWithDefaults instantiates a new AddressRiskReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressRiskReportWithDefaults() *AddressRiskReport {
	this := AddressRiskReport{}
	return &this
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *AddressRiskReport) GetCaseId() string {
	if o == nil || o.CaseId == nil {
		var ret string
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRiskReport) GetCaseIdOk() (*string, bool) {
	if o == nil || o.CaseId == nil {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *AddressRiskReport) HasCaseId() bool {
	if o != nil && o.CaseId != nil {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given string and assigns it to the CaseId field.
func (o *AddressRiskReport) SetCaseId(v string) {
	o.CaseId = &v
}

// GetRequestDatetime returns the RequestDatetime field value if set, zero value otherwise.
func (o *AddressRiskReport) GetRequestDatetime() string {
	if o == nil || o.RequestDatetime == nil {
		var ret string
		return ret
	}
	return *o.RequestDatetime
}

// GetRequestDatetimeOk returns a tuple with the RequestDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRiskReport) GetRequestDatetimeOk() (*string, bool) {
	if o == nil || o.RequestDatetime == nil {
		return nil, false
	}
	return o.RequestDatetime, true
}

// HasRequestDatetime returns a boolean if a field has been set.
func (o *AddressRiskReport) HasRequestDatetime() bool {
	if o != nil && o.RequestDatetime != nil {
		return true
	}

	return false
}

// SetRequestDatetime gets a reference to the given string and assigns it to the RequestDatetime field.
func (o *AddressRiskReport) SetRequestDatetime(v string) {
	o.RequestDatetime = &v
}

// GetResponseDatetime returns the ResponseDatetime field value if set, zero value otherwise.
func (o *AddressRiskReport) GetResponseDatetime() string {
	if o == nil || o.ResponseDatetime == nil {
		var ret string
		return ret
	}
	return *o.ResponseDatetime
}

// GetResponseDatetimeOk returns a tuple with the ResponseDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRiskReport) GetResponseDatetimeOk() (*string, bool) {
	if o == nil || o.ResponseDatetime == nil {
		return nil, false
	}
	return o.ResponseDatetime, true
}

// HasResponseDatetime returns a boolean if a field has been set.
func (o *AddressRiskReport) HasResponseDatetime() bool {
	if o != nil && o.ResponseDatetime != nil {
		return true
	}

	return false
}

// SetResponseDatetime gets a reference to the given string and assigns it to the ResponseDatetime field.
func (o *AddressRiskReport) SetResponseDatetime(v string) {
	o.ResponseDatetime = &v
}

// GetAddress returns the Address field value
func (o *AddressRiskReport) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddressRiskReport) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddressRiskReport) SetAddress(v string) {
	o.Address = v
}

// GetChain returns the Chain field value
func (o *AddressRiskReport) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *AddressRiskReport) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *AddressRiskReport) SetChain(v string) {
	o.Chain = v
}

// GetRisk returns the Risk field value
func (o *AddressRiskReport) GetRisk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Risk
}

// GetRiskOk returns a tuple with the Risk field value
// and a boolean to check if the value has been set.
func (o *AddressRiskReport) GetRiskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Risk, true
}

// SetRisk sets field value
func (o *AddressRiskReport) SetRisk(v int32) {
	o.Risk = v
}

// GetDetails returns the Details field value
func (o *AddressRiskReport) GetDetails() AddressRiskReportDetails {
	if o == nil {
		var ret AddressRiskReportDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *AddressRiskReport) GetDetailsOk() (*AddressRiskReportDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *AddressRiskReport) SetDetails(v AddressRiskReportDetails) {
	o.Details = v
}

func (o AddressRiskReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CaseId != nil {
		toSerialize["case_id"] = o.CaseId
	}
	if o.RequestDatetime != nil {
		toSerialize["request_datetime"] = o.RequestDatetime
	}
	if o.ResponseDatetime != nil {
		toSerialize["response_datetime"] = o.ResponseDatetime
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["chain"] = o.Chain
	}
	if true {
		toSerialize["risk"] = o.Risk
	}
	if true {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableAddressRiskReport struct {
	value *AddressRiskReport
	isSet bool
}

func (v NullableAddressRiskReport) Get() *AddressRiskReport {
	return v.value
}

func (v *NullableAddressRiskReport) Set(val *AddressRiskReport) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressRiskReport) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressRiskReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressRiskReport(val *AddressRiskReport) *NullableAddressRiskReport {
	return &NullableAddressRiskReport{value: val, isSet: true}
}

func (v NullableAddressRiskReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressRiskReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


