/*
Blockmate

Blockmate API OpenAPI documentation

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blockmate

import (
	"encoding/json"
	"fmt"
)

// GetDomainFromAddress200ResponseMetadataAttributesInnerValue - struct for GetDomainFromAddress200ResponseMetadataAttributesInnerValue
type GetDomainFromAddress200ResponseMetadataAttributesInnerValue struct {
	Int32 *int32
	String *string
}

// int32AsGetDomainFromAddress200ResponseMetadataAttributesInnerValue is a convenience function that returns int32 wrapped in GetDomainFromAddress200ResponseMetadataAttributesInnerValue
func Int32AsGetDomainFromAddress200ResponseMetadataAttributesInnerValue(v *int32) GetDomainFromAddress200ResponseMetadataAttributesInnerValue {
	return GetDomainFromAddress200ResponseMetadataAttributesInnerValue{
		Int32: v,
	}
}

// stringAsGetDomainFromAddress200ResponseMetadataAttributesInnerValue is a convenience function that returns string wrapped in GetDomainFromAddress200ResponseMetadataAttributesInnerValue
func StringAsGetDomainFromAddress200ResponseMetadataAttributesInnerValue(v *string) GetDomainFromAddress200ResponseMetadataAttributesInnerValue {
	return GetDomainFromAddress200ResponseMetadataAttributesInnerValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetDomainFromAddress200ResponseMetadataAttributesInnerValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GetDomainFromAddress200ResponseMetadataAttributesInnerValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetDomainFromAddress200ResponseMetadataAttributesInnerValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetDomainFromAddress200ResponseMetadataAttributesInnerValue) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetDomainFromAddress200ResponseMetadataAttributesInnerValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableGetDomainFromAddress200ResponseMetadataAttributesInnerValue struct {
	value *GetDomainFromAddress200ResponseMetadataAttributesInnerValue
	isSet bool
}

func (v NullableGetDomainFromAddress200ResponseMetadataAttributesInnerValue) Get() *GetDomainFromAddress200ResponseMetadataAttributesInnerValue {
	return v.value
}

func (v *NullableGetDomainFromAddress200ResponseMetadataAttributesInnerValue) Set(val *GetDomainFromAddress200ResponseMetadataAttributesInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDomainFromAddress200ResponseMetadataAttributesInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDomainFromAddress200ResponseMetadataAttributesInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDomainFromAddress200ResponseMetadataAttributesInnerValue(val *GetDomainFromAddress200ResponseMetadataAttributesInnerValue) *NullableGetDomainFromAddress200ResponseMetadataAttributesInnerValue {
	return &NullableGetDomainFromAddress200ResponseMetadataAttributesInnerValue{value: val, isSet: true}
}

func (v NullableGetDomainFromAddress200ResponseMetadataAttributesInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDomainFromAddress200ResponseMetadataAttributesInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


