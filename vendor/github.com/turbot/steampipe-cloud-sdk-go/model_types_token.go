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

// TypesToken struct for TypesToken
type TypesToken struct {
	CreatedAt string `json:"created_at"`
	Id string `json:"id"`
	Last4 *string `json:"last4,omitempty"`
	Status string `json:"status"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UserId string `json:"user_id"`
	VersionId int32 `json:"version_id"`
}

// NewTypesToken instantiates a new TypesToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesToken(createdAt string, id string, status string, userId string, versionId int32) *TypesToken {
	this := TypesToken{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Status = status
	this.UserId = userId
	this.VersionId = versionId
	return &this
}

// NewTypesTokenWithDefaults instantiates a new TypesToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesTokenWithDefaults() *TypesToken {
	this := TypesToken{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *TypesToken) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TypesToken) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TypesToken) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *TypesToken) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TypesToken) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TypesToken) SetId(v string) {
	o.Id = v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *TypesToken) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesToken) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *TypesToken) HasLast4() bool {
	if o != nil && o.Last4 != nil {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *TypesToken) SetLast4(v string) {
	o.Last4 = &v
}

// GetStatus returns the Status field value
func (o *TypesToken) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TypesToken) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TypesToken) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TypesToken) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesToken) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TypesToken) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TypesToken) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value
func (o *TypesToken) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *TypesToken) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *TypesToken) SetUserId(v string) {
	o.UserId = v
}

// GetVersionId returns the VersionId field value
func (o *TypesToken) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *TypesToken) GetVersionIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *TypesToken) SetVersionId(v int32) {
	o.VersionId = v
}

func (o TypesToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["version_id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableTypesToken struct {
	value *TypesToken
	isSet bool
}

func (v NullableTypesToken) Get() *TypesToken {
	return v.value
}

func (v *NullableTypesToken) Set(val *TypesToken) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesToken) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesToken(val *TypesToken) *NullableTypesToken {
	return &NullableTypesToken{value: val, isSet: true}
}

func (v NullableTypesToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


