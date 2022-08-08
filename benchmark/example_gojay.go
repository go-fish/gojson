package benchmark

import (
	"github.com/francoispqt/gojay"
)

func (m *DSUser) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "username":
		return dec.AddString(&m.Username)
	}
	return nil
}
func (m *DSUser) NKeys() int {
	return 1
}
func (m *DSUser) IsNil() bool {
	return m == nil
}
func (m *DSUser) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddStringKey("username", m.Username)
}

func (m *DSTopic) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "id":
		return dec.AddInt(&m.Id)
	case "slug":
		return dec.AddString(&m.Slug)
	}
	return nil
}
func (m *DSTopic) NKeys() int {
	return 2
}
func (m *DSTopic) IsNil() bool {
	return m == nil
}
func (m *DSTopic) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddIntKey("id", m.Id)
	enc.AddStringKey("slug", m.Slug)
}

func (t *DSTopics) UnmarshalJSONArray(dec *gojay.Decoder) error {
	dsTopic := &DSTopic{}
	*t = append(*t, dsTopic)
	return dec.AddObject(dsTopic)
}

func (m *DSTopics) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *m {
		enc.AddObject(e)
	}
}
func (m *DSTopics) IsNil() bool {
	return m == nil
}

func (m *DSTopicsList) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "topics":
		m.Topics = DSTopics{}
		return dec.AddArray(&m.Topics)
	case "more_topics_url":
		return dec.AddString(&m.MoreTopicsUrl)
	}
	return nil
}
func (m *DSTopicsList) NKeys() int {
	return 2
}

func (m *DSTopicsList) IsNil() bool {
	return m == nil
}

func (m *DSTopicsList) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey("users", &m.Topics)
	enc.AddStringKey("more_topics_url", m.MoreTopicsUrl)
}

func (t *DSUsers) UnmarshalJSONArray(dec *gojay.Decoder) error {
	dsUser := DSUser{}
	*t = append(*t, &dsUser)
	return dec.AddObject(&dsUser)
}

func (m *DSUsers) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *m {
		enc.AddObject(e)
	}
}
func (m *DSUsers) IsNil() bool {
	return m == nil
}

func (m *LargePayload) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "users":
		return dec.AddArray(&m.Users)
	case "topics":
		m.Topics = &DSTopicsList{}
		return dec.AddObject(m.Topics)
	}
	return nil
}

func (m *LargePayload) NKeys() int {
	return 2
}

//easyjson:json
func (m *LargePayload) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey("users", &m.Users)
	enc.AddObjectKey("topics", m.Topics)
}

func (m *LargePayload) IsNil() bool {
	return m == nil
}

func (m *CBAvatar) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "avatars":
		return dec.AddString(&m.Url)
	}
	return nil
}
func (m *CBAvatar) NKeys() int {
	return 1
}

func (m *CBAvatar) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddStringKey("url", m.Url)
}

func (m *CBAvatar) IsNil() bool {
	return m == nil
}

func (t *Avatars) UnmarshalJSONArray(dec *gojay.Decoder) error {
	avatar := CBAvatar{}
	*t = append(*t, &avatar)
	return dec.AddObject(&avatar)
}

func (m *Avatars) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *m {
		enc.AddObject(e)
	}
}
func (m *Avatars) IsNil() bool {
	return m == nil
}

func (m *CBGravatar) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "avatars":
		return dec.AddArray(&m.Avatars)
	}
	return nil
}
func (m *CBGravatar) NKeys() int {
	return 1
}

func (m *CBGravatar) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey("avatars", &m.Avatars)
}

func (m *CBGravatar) IsNil() bool {
	return m == nil
}

func (m *CBGithub) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "followers":
		return dec.AddInt(&m.Followers)
	}
	return nil
}

func (m *CBGithub) NKeys() int {
	return 1
}

func (m *CBGithub) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddIntKey("followers", m.Followers)
}

func (m *CBGithub) IsNil() bool {
	return m == nil
}

func (m *CBName) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "fullName":
		return dec.AddString(&m.FullName)
	}
	return nil
}

func (m *CBName) NKeys() int {
	return 1
}

func (m *CBName) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddStringKey("fullName", m.FullName)
}

func (m *CBName) IsNil() bool {
	return m == nil
}

func (m *CBPerson) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "name":
		m.Name = &CBName{}
		return dec.AddObject(m.Name)
	case "github":
		m.Github = &CBGithub{}
		return dec.AddObject(m.Github)
	case "gravatar":
		m.Gravatar = &CBGravatar{}
		return dec.AddObject(m.Gravatar)
	}
	return nil
}

func (m *CBPerson) NKeys() int {
	return 3
}

func (m *CBPerson) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddObjectKey("name", m.Name)
	enc.AddObjectKey("github", m.Github)
	enc.AddObjectKey("gravatar", m.Gravatar)
}

func (m *CBPerson) IsNil() bool {
	return m == nil
}

func (m *MediumPayload) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "person":
		m.Person = &CBPerson{}
		return dec.AddObject(m.Person)
	case "company":
		dec.AddString(&m.Company)
	}
	return nil
}

func (m *MediumPayload) NKeys() int {
	return 2
}

func (m *MediumPayload) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddObjectKey("person", m.Person)
	enc.AddStringKey("company", m.Company)
}

func (m *MediumPayload) IsNil() bool {
	return m == nil
}

func (t *SmallPayload) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddIntKey("st", t.St)
	enc.AddIntKey("sid", t.Sid)
	enc.AddStringKey("tt", t.Tt)
	enc.AddIntKey("gr", t.Gr)
	enc.AddStringKey("uuid", t.Uuid)
	enc.AddStringKey("ip", t.Ip)
	enc.AddStringKey("ua", t.Ua)
	enc.AddIntKey("tz", t.Tz)
	enc.AddIntKey("v", t.V)
}

func (t *SmallPayload) IsNil() bool {
	return t == nil
}

func (t *SmallPayload) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "st":
		return dec.AddInt(&t.St)
	case "sid":
		return dec.AddInt(&t.Sid)
	case "gr":
		return dec.AddInt(&t.Gr)
	case "tz":
		return dec.AddInt(&t.Tz)
	case "v":
		return dec.AddInt(&t.V)
	case "tt":
		return dec.AddString(&t.Tt)
	case "uuid":
		return dec.AddString(&t.Uuid)
	case "ip":
		return dec.AddString(&t.Ip)
	case "ua":
		return dec.AddString(&t.Ua)
	}
	return nil
}

func (t *SmallPayload) NKeys() int {
	return 9
}
