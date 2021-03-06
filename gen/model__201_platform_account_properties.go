/*
CyberArkIAG

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 18a45ad8-77e8-4ecc-873e-787c6de10a60
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// 201PlatformAccountProperties struct for 201PlatformAccountProperties
type 201PlatformAccountProperties struct {
	Index *string `json:"Index,omitempty"`
	DualAccountStatus *string `json:"DualAccountStatus,omitempty"`
	VirtualUsername *string `json:"VirtualUsername,omitempty"`
}

// New201PlatformAccountProperties instantiates a new 201PlatformAccountProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func New201PlatformAccountProperties() *201PlatformAccountProperties {
	this := 201PlatformAccountProperties{}
	var index string = "1"
	this.Index = &index
	var dualAccountStatus string = "Active"
	this.DualAccountStatus = &dualAccountStatus
	var virtualUsername string = "cluster02sqluser"
	this.VirtualUsername = &virtualUsername
	return &this
}

// New201PlatformAccountPropertiesWithDefaults instantiates a new 201PlatformAccountProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func New201PlatformAccountPropertiesWithDefaults() *201PlatformAccountProperties {
	this := 201PlatformAccountProperties{}
	var index string = "1"
	this.Index = &index
	var dualAccountStatus string = "Active"
	this.DualAccountStatus = &dualAccountStatus
	var virtualUsername string = "cluster02sqluser"
	this.VirtualUsername = &virtualUsername
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *201PlatformAccountProperties) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *201PlatformAccountProperties) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *201PlatformAccountProperties) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *201PlatformAccountProperties) SetIndex(v string) {
	o.Index = &v
}

// GetDualAccountStatus returns the DualAccountStatus field value if set, zero value otherwise.
func (o *201PlatformAccountProperties) GetDualAccountStatus() string {
	if o == nil || o.DualAccountStatus == nil {
		var ret string
		return ret
	}
	return *o.DualAccountStatus
}

// GetDualAccountStatusOk returns a tuple with the DualAccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *201PlatformAccountProperties) GetDualAccountStatusOk() (*string, bool) {
	if o == nil || o.DualAccountStatus == nil {
		return nil, false
	}
	return o.DualAccountStatus, true
}

// HasDualAccountStatus returns a boolean if a field has been set.
func (o *201PlatformAccountProperties) HasDualAccountStatus() bool {
	if o != nil && o.DualAccountStatus != nil {
		return true
	}

	return false
}

// SetDualAccountStatus gets a reference to the given string and assigns it to the DualAccountStatus field.
func (o *201PlatformAccountProperties) SetDualAccountStatus(v string) {
	o.DualAccountStatus = &v
}

// GetVirtualUsername returns the VirtualUsername field value if set, zero value otherwise.
func (o *201PlatformAccountProperties) GetVirtualUsername() string {
	if o == nil || o.VirtualUsername == nil {
		var ret string
		return ret
	}
	return *o.VirtualUsername
}

// GetVirtualUsernameOk returns a tuple with the VirtualUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *201PlatformAccountProperties) GetVirtualUsernameOk() (*string, bool) {
	if o == nil || o.VirtualUsername == nil {
		return nil, false
	}
	return o.VirtualUsername, true
}

// HasVirtualUsername returns a boolean if a field has been set.
func (o *201PlatformAccountProperties) HasVirtualUsername() bool {
	if o != nil && o.VirtualUsername != nil {
		return true
	}

	return false
}

// SetVirtualUsername gets a reference to the given string and assigns it to the VirtualUsername field.
func (o *201PlatformAccountProperties) SetVirtualUsername(v string) {
	o.VirtualUsername = &v
}

func (o 201PlatformAccountProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["Index"] = o.Index
	}
	if o.DualAccountStatus != nil {
		toSerialize["DualAccountStatus"] = o.DualAccountStatus
	}
	if o.VirtualUsername != nil {
		toSerialize["VirtualUsername"] = o.VirtualUsername
	}
	return json.Marshal(toSerialize)
}

type Nullable201PlatformAccountProperties struct {
	value *201PlatformAccountProperties
	isSet bool
}

func (v Nullable201PlatformAccountProperties) Get() *201PlatformAccountProperties {
	return v.value
}

func (v *Nullable201PlatformAccountProperties) Set(val *201PlatformAccountProperties) {
	v.value = val
	v.isSet = true
}

func (v Nullable201PlatformAccountProperties) IsSet() bool {
	return v.isSet
}

func (v *Nullable201PlatformAccountProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullable201PlatformAccountProperties(val *201PlatformAccountProperties) *Nullable201PlatformAccountProperties {
	return &Nullable201PlatformAccountProperties{value: val, isSet: true}
}

func (v Nullable201PlatformAccountProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *Nullable201PlatformAccountProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


