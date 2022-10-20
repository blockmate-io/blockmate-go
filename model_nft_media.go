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

// NftMedia struct for NftMedia
type NftMedia struct {
	// String - Uri representing the location of the NFT's original metadata blob. This is a backup for you to parse when the metadata field is not automatically populated.
	Raw *string `json:"raw,omitempty"`
	// String - Public gateway uri for the raw uri above.
	Gateway *string `json:"gateway,omitempty"`
	// URL for a resized thumbnail of the NFT media asset.
	Thumbnail *string `json:"thumbnail,omitempty"`
	// The media format (jpg, gif, png, etc.) of the gateway and thumbnail assets.
	Format *string `json:"format,omitempty"`
	// The size of the media asset in bytes.
	Bytes *int32 `json:"bytes,omitempty"`
}

// NewNftMedia instantiates a new NftMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftMedia() *NftMedia {
	this := NftMedia{}
	return &this
}

// NewNftMediaWithDefaults instantiates a new NftMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftMediaWithDefaults() *NftMedia {
	this := NftMedia{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *NftMedia) GetRaw() string {
	if o == nil || o.Raw == nil {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftMedia) GetRawOk() (*string, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *NftMedia) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *NftMedia) SetRaw(v string) {
	o.Raw = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *NftMedia) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftMedia) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *NftMedia) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *NftMedia) SetGateway(v string) {
	o.Gateway = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *NftMedia) GetThumbnail() string {
	if o == nil || o.Thumbnail == nil {
		var ret string
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftMedia) GetThumbnailOk() (*string, bool) {
	if o == nil || o.Thumbnail == nil {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *NftMedia) HasThumbnail() bool {
	if o != nil && o.Thumbnail != nil {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given string and assigns it to the Thumbnail field.
func (o *NftMedia) SetThumbnail(v string) {
	o.Thumbnail = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *NftMedia) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftMedia) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *NftMedia) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *NftMedia) SetFormat(v string) {
	o.Format = &v
}

// GetBytes returns the Bytes field value if set, zero value otherwise.
func (o *NftMedia) GetBytes() int32 {
	if o == nil || o.Bytes == nil {
		var ret int32
		return ret
	}
	return *o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftMedia) GetBytesOk() (*int32, bool) {
	if o == nil || o.Bytes == nil {
		return nil, false
	}
	return o.Bytes, true
}

// HasBytes returns a boolean if a field has been set.
func (o *NftMedia) HasBytes() bool {
	if o != nil && o.Bytes != nil {
		return true
	}

	return false
}

// SetBytes gets a reference to the given int32 and assigns it to the Bytes field.
func (o *NftMedia) SetBytes(v int32) {
	o.Bytes = &v
}

func (o NftMedia) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Thumbnail != nil {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Bytes != nil {
		toSerialize["bytes"] = o.Bytes
	}
	return json.Marshal(toSerialize)
}

type NullableNftMedia struct {
	value *NftMedia
	isSet bool
}

func (v NullableNftMedia) Get() *NftMedia {
	return v.value
}

func (v *NullableNftMedia) Set(val *NftMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableNftMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableNftMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftMedia(val *NftMedia) *NullableNftMedia {
	return &NullableNftMedia{value: val, isSet: true}
}

func (v NullableNftMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

