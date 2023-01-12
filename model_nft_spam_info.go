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

// NftSpamInfo Information about whether and why a contract was marked as spam.
type NftSpamInfo struct {
	// \"true\" if contract is spam, else \"false\"
	IsSpam *string `json:"isSpam,omitempty"`
	// List of reasons why a contract was classified as spam.
	Classifications []string `json:"classifications,omitempty"`
}

// NewNftSpamInfo instantiates a new NftSpamInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftSpamInfo() *NftSpamInfo {
	this := NftSpamInfo{}
	return &this
}

// NewNftSpamInfoWithDefaults instantiates a new NftSpamInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftSpamInfoWithDefaults() *NftSpamInfo {
	this := NftSpamInfo{}
	return &this
}

// GetIsSpam returns the IsSpam field value if set, zero value otherwise.
func (o *NftSpamInfo) GetIsSpam() string {
	if o == nil || o.IsSpam == nil {
		var ret string
		return ret
	}
	return *o.IsSpam
}

// GetIsSpamOk returns a tuple with the IsSpam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftSpamInfo) GetIsSpamOk() (*string, bool) {
	if o == nil || o.IsSpam == nil {
		return nil, false
	}
	return o.IsSpam, true
}

// HasIsSpam returns a boolean if a field has been set.
func (o *NftSpamInfo) HasIsSpam() bool {
	if o != nil && o.IsSpam != nil {
		return true
	}

	return false
}

// SetIsSpam gets a reference to the given string and assigns it to the IsSpam field.
func (o *NftSpamInfo) SetIsSpam(v string) {
	o.IsSpam = &v
}

// GetClassifications returns the Classifications field value if set, zero value otherwise.
func (o *NftSpamInfo) GetClassifications() []string {
	if o == nil || o.Classifications == nil {
		var ret []string
		return ret
	}
	return o.Classifications
}

// GetClassificationsOk returns a tuple with the Classifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftSpamInfo) GetClassificationsOk() ([]string, bool) {
	if o == nil || o.Classifications == nil {
		return nil, false
	}
	return o.Classifications, true
}

// HasClassifications returns a boolean if a field has been set.
func (o *NftSpamInfo) HasClassifications() bool {
	if o != nil && o.Classifications != nil {
		return true
	}

	return false
}

// SetClassifications gets a reference to the given []string and assigns it to the Classifications field.
func (o *NftSpamInfo) SetClassifications(v []string) {
	o.Classifications = v
}

func (o NftSpamInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsSpam != nil {
		toSerialize["isSpam"] = o.IsSpam
	}
	if o.Classifications != nil {
		toSerialize["classifications"] = o.Classifications
	}
	return json.Marshal(toSerialize)
}

type NullableNftSpamInfo struct {
	value *NftSpamInfo
	isSet bool
}

func (v NullableNftSpamInfo) Get() *NftSpamInfo {
	return v.value
}

func (v *NullableNftSpamInfo) Set(val *NftSpamInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNftSpamInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNftSpamInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftSpamInfo(val *NftSpamInfo) *NullableNftSpamInfo {
	return &NullableNftSpamInfo{value: val, isSet: true}
}

func (v NullableNftSpamInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftSpamInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


