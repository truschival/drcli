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

// Alarms struct for Alarms
type Alarms struct {
	Items []Alarm
}

// NewAlarms instantiates a new Alarms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarms() *Alarms {
	this := Alarms{}
	return &this
}

// NewAlarmsWithDefaults instantiates a new Alarms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmsWithDefaults() *Alarms {
	this := Alarms{}
	return &this
}

func (o Alarms) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return json.Marshal(toSerialize)
}

func (o *Alarms) UnmarshalJSON(bytes []byte) (err error) {
	return json.Unmarshal(bytes, &o.Items)
}

type NullableAlarms struct {
	value *Alarms
	isSet bool
}

func (v NullableAlarms) Get() *Alarms {
	return v.value
}

func (v *NullableAlarms) Set(val *Alarms) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarms) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarms(val *Alarms) *NullableAlarms {
	return &NullableAlarms{value: val, isSet: true}
}

func (v NullableAlarms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

