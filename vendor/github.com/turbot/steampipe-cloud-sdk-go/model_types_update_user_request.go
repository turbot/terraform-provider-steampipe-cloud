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

// TypesUpdateUserRequest struct for TypesUpdateUserRequest
type TypesUpdateUserRequest struct {
	AvatarUrl *string `json:"avatar_url,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Handle *string `json:"handle,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewTypesUpdateUserRequest instantiates a new TypesUpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesUpdateUserRequest() *TypesUpdateUserRequest {
	this := TypesUpdateUserRequest{}
	return &this
}

// NewTypesUpdateUserRequestWithDefaults instantiates a new TypesUpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesUpdateUserRequestWithDefaults() *TypesUpdateUserRequest {
	this := TypesUpdateUserRequest{}
	return &this
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *TypesUpdateUserRequest) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUpdateUserRequest) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *TypesUpdateUserRequest) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *TypesUpdateUserRequest) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TypesUpdateUserRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUpdateUserRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TypesUpdateUserRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TypesUpdateUserRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *TypesUpdateUserRequest) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUpdateUserRequest) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *TypesUpdateUserRequest) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *TypesUpdateUserRequest) SetHandle(v string) {
	o.Handle = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TypesUpdateUserRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUpdateUserRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TypesUpdateUserRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TypesUpdateUserRequest) SetUrl(v string) {
	o.Url = &v
}

func (o TypesUpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvatarUrl != nil {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableTypesUpdateUserRequest struct {
	value *TypesUpdateUserRequest
	isSet bool
}

func (v NullableTypesUpdateUserRequest) Get() *TypesUpdateUserRequest {
	return v.value
}

func (v *NullableTypesUpdateUserRequest) Set(val *TypesUpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesUpdateUserRequest(val *TypesUpdateUserRequest) *NullableTypesUpdateUserRequest {
	return &NullableTypesUpdateUserRequest{value: val, isSet: true}
}

func (v NullableTypesUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


