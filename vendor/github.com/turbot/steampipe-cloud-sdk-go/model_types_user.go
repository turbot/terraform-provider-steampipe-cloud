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

// TypesUser struct for TypesUser
type TypesUser struct {
	AvatarUrl *string `json:"avatar_url,omitempty"`
	CreatedAt string `json:"created_at"`
	DisplayName *string `json:"display_name,omitempty"`
	Email string `json:"email"`
	Handle string `json:"handle"`
	Id string `json:"id"`
	PreviewAccessMode *string `json:"preview_access_mode,omitempty"`
	Status string `json:"status"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
	VersionId int32 `json:"version_id"`
}

// NewTypesUser instantiates a new TypesUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesUser(createdAt string, email string, handle string, id string, status string, versionId int32) *TypesUser {
	this := TypesUser{}
	this.CreatedAt = createdAt
	this.Email = email
	this.Handle = handle
	this.Id = id
	this.Status = status
	this.VersionId = versionId
	return &this
}

// NewTypesUserWithDefaults instantiates a new TypesUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesUserWithDefaults() *TypesUser {
	this := TypesUser{}
	return &this
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *TypesUser) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUser) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *TypesUser) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *TypesUser) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TypesUser) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TypesUser) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TypesUser) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TypesUser) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUser) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TypesUser) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TypesUser) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value
func (o *TypesUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *TypesUser) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *TypesUser) SetEmail(v string) {
	o.Email = v
}

// GetHandle returns the Handle field value
func (o *TypesUser) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *TypesUser) GetHandleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *TypesUser) SetHandle(v string) {
	o.Handle = v
}

// GetId returns the Id field value
func (o *TypesUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TypesUser) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TypesUser) SetId(v string) {
	o.Id = v
}

// GetPreviewAccessMode returns the PreviewAccessMode field value if set, zero value otherwise.
func (o *TypesUser) GetPreviewAccessMode() string {
	if o == nil || o.PreviewAccessMode == nil {
		var ret string
		return ret
	}
	return *o.PreviewAccessMode
}

// GetPreviewAccessModeOk returns a tuple with the PreviewAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUser) GetPreviewAccessModeOk() (*string, bool) {
	if o == nil || o.PreviewAccessMode == nil {
		return nil, false
	}
	return o.PreviewAccessMode, true
}

// HasPreviewAccessMode returns a boolean if a field has been set.
func (o *TypesUser) HasPreviewAccessMode() bool {
	if o != nil && o.PreviewAccessMode != nil {
		return true
	}

	return false
}

// SetPreviewAccessMode gets a reference to the given string and assigns it to the PreviewAccessMode field.
func (o *TypesUser) SetPreviewAccessMode(v string) {
	o.PreviewAccessMode = &v
}

// GetStatus returns the Status field value
func (o *TypesUser) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TypesUser) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TypesUser) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TypesUser) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUser) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TypesUser) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TypesUser) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TypesUser) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUser) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TypesUser) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TypesUser) SetUrl(v string) {
	o.Url = &v
}

// GetVersionId returns the VersionId field value
func (o *TypesUser) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *TypesUser) GetVersionIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *TypesUser) SetVersionId(v int32) {
	o.VersionId = v
}

func (o TypesUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvatarUrl != nil {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["handle"] = o.Handle
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.PreviewAccessMode != nil {
		toSerialize["preview_access_mode"] = o.PreviewAccessMode
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["version_id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableTypesUser struct {
	value *TypesUser
	isSet bool
}

func (v NullableTypesUser) Get() *TypesUser {
	return v.value
}

func (v *NullableTypesUser) Set(val *TypesUser) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesUser) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesUser(val *TypesUser) *NullableTypesUser {
	return &NullableTypesUser{value: val, isSet: true}
}

func (v NullableTypesUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


