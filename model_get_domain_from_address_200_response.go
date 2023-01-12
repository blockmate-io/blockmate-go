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

// GetDomainFromAddress200Response struct for GetDomainFromAddress200Response
type GetDomainFromAddress200Response struct {
	Domain *string `json:"domain,omitempty"`
	Metadata *GetDomainFromAddress200ResponseMetadata `json:"metadata,omitempty"`
}

// NewGetDomainFromAddress200Response instantiates a new GetDomainFromAddress200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDomainFromAddress200Response() *GetDomainFromAddress200Response {
	this := GetDomainFromAddress200Response{}
	return &this
}

// NewGetDomainFromAddress200ResponseWithDefaults instantiates a new GetDomainFromAddress200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDomainFromAddress200ResponseWithDefaults() *GetDomainFromAddress200Response {
	this := GetDomainFromAddress200Response{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *GetDomainFromAddress200Response) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDomainFromAddress200Response) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *GetDomainFromAddress200Response) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *GetDomainFromAddress200Response) SetDomain(v string) {
	o.Domain = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetDomainFromAddress200Response) GetMetadata() GetDomainFromAddress200ResponseMetadata {
	if o == nil || o.Metadata == nil {
		var ret GetDomainFromAddress200ResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDomainFromAddress200Response) GetMetadataOk() (*GetDomainFromAddress200ResponseMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetDomainFromAddress200Response) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given GetDomainFromAddress200ResponseMetadata and assigns it to the Metadata field.
func (o *GetDomainFromAddress200Response) SetMetadata(v GetDomainFromAddress200ResponseMetadata) {
	o.Metadata = &v
}

func (o GetDomainFromAddress200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableGetDomainFromAddress200Response struct {
	value *GetDomainFromAddress200Response
	isSet bool
}

func (v NullableGetDomainFromAddress200Response) Get() *GetDomainFromAddress200Response {
	return v.value
}

func (v *NullableGetDomainFromAddress200Response) Set(val *GetDomainFromAddress200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDomainFromAddress200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDomainFromAddress200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDomainFromAddress200Response(val *GetDomainFromAddress200Response) *NullableGetDomainFromAddress200Response {
	return &NullableGetDomainFromAddress200Response{value: val, isSet: true}
}

func (v NullableGetDomainFromAddress200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDomainFromAddress200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


