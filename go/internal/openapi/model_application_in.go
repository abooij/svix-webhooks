/*
 * Svix
 *
 * The Svix server API documentation
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplicationIn struct for ApplicationIn
type ApplicationIn struct {
	Uid *string `json:"uid,omitempty"`
	Name string `json:"name"`
}

// NewApplicationIn instantiates a new ApplicationIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationIn(name string, ) *ApplicationIn {
	this := ApplicationIn{}
	this.Name = name
	return &this
}

// NewApplicationInWithDefaults instantiates a new ApplicationIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationInWithDefaults() *ApplicationIn {
	this := ApplicationIn{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ApplicationIn) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationIn) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ApplicationIn) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *ApplicationIn) SetUid(v string) {
	o.Uid = &v
}

// GetName returns the Name field value
func (o *ApplicationIn) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationIn) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationIn) SetName(v string) {
	o.Name = v
}

func (o ApplicationIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationIn struct {
	value *ApplicationIn
	isSet bool
}

func (v NullableApplicationIn) Get() *ApplicationIn {
	return v.value
}

func (v *NullableApplicationIn) Set(val *ApplicationIn) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationIn) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationIn(val *ApplicationIn) *NullableApplicationIn {
	return &NullableApplicationIn{value: val, isSet: true}
}

func (v NullableApplicationIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

