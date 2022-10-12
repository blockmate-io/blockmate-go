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

// BalanceResponseAccountsInnerState struct for BalanceResponseAccountsInnerState
type BalanceResponseAccountsInnerState struct {
	// null if initial sync was not yet done
	LastSync NullableString `json:"last_sync,omitempty"`
}

// NewBalanceResponseAccountsInnerState instantiates a new BalanceResponseAccountsInnerState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceResponseAccountsInnerState() *BalanceResponseAccountsInnerState {
	this := BalanceResponseAccountsInnerState{}
	return &this
}

// NewBalanceResponseAccountsInnerStateWithDefaults instantiates a new BalanceResponseAccountsInnerState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceResponseAccountsInnerStateWithDefaults() *BalanceResponseAccountsInnerState {
	this := BalanceResponseAccountsInnerState{}
	return &this
}

// GetLastSync returns the LastSync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceResponseAccountsInnerState) GetLastSync() string {
	if o == nil || o.LastSync.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastSync.Get()
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceResponseAccountsInnerState) GetLastSyncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSync.Get(), o.LastSync.IsSet()
}

// HasLastSync returns a boolean if a field has been set.
func (o *BalanceResponseAccountsInnerState) HasLastSync() bool {
	if o != nil && o.LastSync.IsSet() {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given NullableString and assigns it to the LastSync field.
func (o *BalanceResponseAccountsInnerState) SetLastSync(v string) {
	o.LastSync.Set(&v)
}
// SetLastSyncNil sets the value for LastSync to be an explicit nil
func (o *BalanceResponseAccountsInnerState) SetLastSyncNil() {
	o.LastSync.Set(nil)
}

// UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
func (o *BalanceResponseAccountsInnerState) UnsetLastSync() {
	o.LastSync.Unset()
}

func (o BalanceResponseAccountsInnerState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastSync.IsSet() {
		toSerialize["last_sync"] = o.LastSync.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBalanceResponseAccountsInnerState struct {
	value *BalanceResponseAccountsInnerState
	isSet bool
}

func (v NullableBalanceResponseAccountsInnerState) Get() *BalanceResponseAccountsInnerState {
	return v.value
}

func (v *NullableBalanceResponseAccountsInnerState) Set(val *BalanceResponseAccountsInnerState) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceResponseAccountsInnerState) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceResponseAccountsInnerState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceResponseAccountsInnerState(val *BalanceResponseAccountsInnerState) *NullableBalanceResponseAccountsInnerState {
	return &NullableBalanceResponseAccountsInnerState{value: val, isSet: true}
}

func (v NullableBalanceResponseAccountsInnerState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceResponseAccountsInnerState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

