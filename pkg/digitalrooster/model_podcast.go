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

// Podcast struct for Podcast
type Podcast struct {
	Id *string `json:"id,omitempty"`
	Title string `json:"title"`
	Url string `json:"url"`
	UpdateInterval *int32 `json:"updateInterval,omitempty"`
}

// NewPodcast instantiates a new Podcast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodcast(title string, url string, ) *Podcast {
	this := Podcast{}
	this.Title = title
	this.Url = url
	return &this
}

// NewPodcastWithDefaults instantiates a new Podcast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodcastWithDefaults() *Podcast {
	this := Podcast{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Podcast) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Podcast) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Podcast) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Podcast) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value
func (o *Podcast) GetTitle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Podcast) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Podcast) SetTitle(v string) {
	o.Title = v
}

// GetUrl returns the Url field value
func (o *Podcast) GetUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Podcast) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Podcast) SetUrl(v string) {
	o.Url = v
}

// GetUpdateInterval returns the UpdateInterval field value if set, zero value otherwise.
func (o *Podcast) GetUpdateInterval() int32 {
	if o == nil || o.UpdateInterval == nil {
		var ret int32
		return ret
	}
	return *o.UpdateInterval
}

// GetUpdateIntervalOk returns a tuple with the UpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Podcast) GetUpdateIntervalOk() (*int32, bool) {
	if o == nil || o.UpdateInterval == nil {
		return nil, false
	}
	return o.UpdateInterval, true
}

// HasUpdateInterval returns a boolean if a field has been set.
func (o *Podcast) HasUpdateInterval() bool {
	if o != nil && o.UpdateInterval != nil {
		return true
	}

	return false
}

// SetUpdateInterval gets a reference to the given int32 and assigns it to the UpdateInterval field.
func (o *Podcast) SetUpdateInterval(v int32) {
	o.UpdateInterval = &v
}

func (o Podcast) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.UpdateInterval != nil {
		toSerialize["updateInterval"] = o.UpdateInterval
	}
	return json.Marshal(toSerialize)
}

type NullablePodcast struct {
	value *Podcast
	isSet bool
}

func (v NullablePodcast) Get() *Podcast {
	return v.value
}

func (v *NullablePodcast) Set(val *Podcast) {
	v.value = val
	v.isSet = true
}

func (v NullablePodcast) IsSet() bool {
	return v.isSet
}

func (v *NullablePodcast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodcast(val *Podcast) *NullablePodcast {
	return &NullablePodcast{value: val, isSet: true}
}

func (v NullablePodcast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodcast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


