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

// TransactionRiskReport TransactionRiskReport
type TransactionRiskReport struct {
	CaseId *string `json:"case_id,omitempty"`
	RequestDatetime *string `json:"request_datetime,omitempty"`
	ResponseDatetime *string `json:"response_datetime,omitempty"`
	Transaction *string `json:"transaction,omitempty"`
	Chain string `json:"chain"`
	Risk int32 `json:"risk"`
	// Keys are addresses from transaction inputs or outputs
	Details map[string]AddressRiskReportDetails `json:"details"`
}

// NewTransactionRiskReport instantiates a new TransactionRiskReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRiskReport(chain string, risk int32, details map[string]AddressRiskReportDetails) *TransactionRiskReport {
	this := TransactionRiskReport{}
	this.Chain = chain
	this.Risk = risk
	this.Details = details
	return &this
}

// NewTransactionRiskReportWithDefaults instantiates a new TransactionRiskReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRiskReportWithDefaults() *TransactionRiskReport {
	this := TransactionRiskReport{}
	return &this
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *TransactionRiskReport) GetCaseId() string {
	if o == nil || o.CaseId == nil {
		var ret string
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRiskReport) GetCaseIdOk() (*string, bool) {
	if o == nil || o.CaseId == nil {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *TransactionRiskReport) HasCaseId() bool {
	if o != nil && o.CaseId != nil {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given string and assigns it to the CaseId field.
func (o *TransactionRiskReport) SetCaseId(v string) {
	o.CaseId = &v
}

// GetRequestDatetime returns the RequestDatetime field value if set, zero value otherwise.
func (o *TransactionRiskReport) GetRequestDatetime() string {
	if o == nil || o.RequestDatetime == nil {
		var ret string
		return ret
	}
	return *o.RequestDatetime
}

// GetRequestDatetimeOk returns a tuple with the RequestDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRiskReport) GetRequestDatetimeOk() (*string, bool) {
	if o == nil || o.RequestDatetime == nil {
		return nil, false
	}
	return o.RequestDatetime, true
}

// HasRequestDatetime returns a boolean if a field has been set.
func (o *TransactionRiskReport) HasRequestDatetime() bool {
	if o != nil && o.RequestDatetime != nil {
		return true
	}

	return false
}

// SetRequestDatetime gets a reference to the given string and assigns it to the RequestDatetime field.
func (o *TransactionRiskReport) SetRequestDatetime(v string) {
	o.RequestDatetime = &v
}

// GetResponseDatetime returns the ResponseDatetime field value if set, zero value otherwise.
func (o *TransactionRiskReport) GetResponseDatetime() string {
	if o == nil || o.ResponseDatetime == nil {
		var ret string
		return ret
	}
	return *o.ResponseDatetime
}

// GetResponseDatetimeOk returns a tuple with the ResponseDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRiskReport) GetResponseDatetimeOk() (*string, bool) {
	if o == nil || o.ResponseDatetime == nil {
		return nil, false
	}
	return o.ResponseDatetime, true
}

// HasResponseDatetime returns a boolean if a field has been set.
func (o *TransactionRiskReport) HasResponseDatetime() bool {
	if o != nil && o.ResponseDatetime != nil {
		return true
	}

	return false
}

// SetResponseDatetime gets a reference to the given string and assigns it to the ResponseDatetime field.
func (o *TransactionRiskReport) SetResponseDatetime(v string) {
	o.ResponseDatetime = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *TransactionRiskReport) GetTransaction() string {
	if o == nil || o.Transaction == nil {
		var ret string
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRiskReport) GetTransactionOk() (*string, bool) {
	if o == nil || o.Transaction == nil {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *TransactionRiskReport) HasTransaction() bool {
	if o != nil && o.Transaction != nil {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given string and assigns it to the Transaction field.
func (o *TransactionRiskReport) SetTransaction(v string) {
	o.Transaction = &v
}

// GetChain returns the Chain field value
func (o *TransactionRiskReport) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *TransactionRiskReport) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *TransactionRiskReport) SetChain(v string) {
	o.Chain = v
}

// GetRisk returns the Risk field value
func (o *TransactionRiskReport) GetRisk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Risk
}

// GetRiskOk returns a tuple with the Risk field value
// and a boolean to check if the value has been set.
func (o *TransactionRiskReport) GetRiskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Risk, true
}

// SetRisk sets field value
func (o *TransactionRiskReport) SetRisk(v int32) {
	o.Risk = v
}

// GetDetails returns the Details field value
func (o *TransactionRiskReport) GetDetails() map[string]AddressRiskReportDetails {
	if o == nil {
		var ret map[string]AddressRiskReportDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *TransactionRiskReport) GetDetailsOk() (*map[string]AddressRiskReportDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *TransactionRiskReport) SetDetails(v map[string]AddressRiskReportDetails) {
	o.Details = v
}

func (o TransactionRiskReport) MarshalJSON() ([]byte, error) {
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
	if o.Transaction != nil {
		toSerialize["transaction"] = o.Transaction
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

type NullableTransactionRiskReport struct {
	value *TransactionRiskReport
	isSet bool
}

func (v NullableTransactionRiskReport) Get() *TransactionRiskReport {
	return v.value
}

func (v *NullableTransactionRiskReport) Set(val *TransactionRiskReport) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRiskReport) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRiskReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRiskReport(val *TransactionRiskReport) *NullableTransactionRiskReport {
	return &NullableTransactionRiskReport{value: val, isSet: true}
}

func (v NullableTransactionRiskReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRiskReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

