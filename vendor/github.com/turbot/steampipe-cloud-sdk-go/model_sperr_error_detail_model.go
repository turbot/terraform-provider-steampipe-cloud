/*
Steampipe Cloud API

Steampipe Cloud API

API version: 1.0
Contact: support@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SperrErrorDetailModel struct for SperrErrorDetailModel
type SperrErrorDetailModel struct {
	Location *string `json:"location,omitempty"`
	Message *string `json:"message,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewSperrErrorDetailModel instantiates a new SperrErrorDetailModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSperrErrorDetailModel() *SperrErrorDetailModel {
	this := SperrErrorDetailModel{}
	return &this
}

// NewSperrErrorDetailModelWithDefaults instantiates a new SperrErrorDetailModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSperrErrorDetailModelWithDefaults() *SperrErrorDetailModel {
	this := SperrErrorDetailModel{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SperrErrorDetailModel) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SperrErrorDetailModel) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SperrErrorDetailModel) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SperrErrorDetailModel) SetLocation(v string) {
	o.Location = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SperrErrorDetailModel) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SperrErrorDetailModel) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SperrErrorDetailModel) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SperrErrorDetailModel) SetMessage(v string) {
	o.Message = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SperrErrorDetailModel) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SperrErrorDetailModel) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SperrErrorDetailModel) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SperrErrorDetailModel) SetValue(v string) {
	o.Value = &v
}

func (o SperrErrorDetailModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSperrErrorDetailModel struct {
	value *SperrErrorDetailModel
	isSet bool
}

func (v NullableSperrErrorDetailModel) Get() *SperrErrorDetailModel {
	return v.value
}

func (v *NullableSperrErrorDetailModel) Set(val *SperrErrorDetailModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSperrErrorDetailModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSperrErrorDetailModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSperrErrorDetailModel(val *SperrErrorDetailModel) *NullableSperrErrorDetailModel {
	return &NullableSperrErrorDetailModel{value: val, isSet: true}
}

func (v NullableSperrErrorDetailModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSperrErrorDetailModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


