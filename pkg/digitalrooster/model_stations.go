/*
 * DigitalRooster
 *
 * Open API for Digital Rooster
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package digitalrooster

import (
	"encoding/json"
)

// Stations struct for Stations
type Stations struct {
	Items []Station
}

// NewStations instantiates a new Stations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStations() *Stations {
	this := Stations{}
	return &this
}

// NewStationsWithDefaults instantiates a new Stations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStationsWithDefaults() *Stations {
	this := Stations{}
	return &this
}

func (o Stations) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return json.Marshal(toSerialize)
}

func (o *Stations) UnmarshalJSON(bytes []byte) (err error) {
	return json.Unmarshal(bytes, &o.Items)
}

type NullableStations struct {
	value *Stations
	isSet bool
}

func (v NullableStations) Get() *Stations {
	return v.value
}

func (v *NullableStations) Set(val *Stations) {
	v.value = val
	v.isSet = true
}

func (v NullableStations) IsSet() bool {
	return v.isSet
}

func (v *NullableStations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStations(val *Stations) *NullableStations {
	return &NullableStations{value: val, isSet: true}
}

func (v NullableStations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

