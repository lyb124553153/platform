// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type LicenseRecord struct {
	Id       string `json:"id"`
	CreateAt int64  `json:"create_at"`
	Bytes    string `json:"-"`
}

type License struct {
	Id        string    `json:"id"`
	IssuedAt  int64     `json:"issued_at"`
	StartsAt  int64     `json:"starts_at"`
	ExpiresAt int64     `json:"expires_at"`
	Customer  *Customer `json:"customer"`
	Features  *Features `json:"features"`
}

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	PhoneNumber string `json:"phone_number"`
}

type Features struct {
	Users      *int  `json:"users"`
	LDAP       *bool `json:"ldap"`
	MFA        *bool `json:"mfa"`
	GoogleSSO  *bool `json:"google_sso"`
	Compliance *bool `json:"compliance"`
}

func (f *Features) SetDefaults() {
	if f.Users == nil {
		f.Users = new(int)
		*f.Users = 0
	}

	if f.LDAP == nil {
		f.LDAP = new(bool)
		*f.LDAP = true
	}

	if f.MFA == nil {
		f.MFA = new(bool)
		*f.MFA = true
	}

	if f.GoogleSSO == nil {
		f.GoogleSSO = new(bool)
		*f.GoogleSSO = true
	}

	if f.Compliance == nil {
		f.Compliance = new(bool)
		*f.Compliance = true
	}
}

func (l *License) IsExpired() bool {
	now := GetMillis()
	if l.ExpiresAt < now {
		return true
	}
	return false
}

func (l *License) IsStarted() bool {
	now := GetMillis()
	if l.StartsAt < now {
		return true
	}
	return false
}

func (l *License) ToJson() string {
	b, err := json.Marshal(l)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func LicenseFromJson(data io.Reader) *License {
	decoder := json.NewDecoder(data)
	var o License
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (lr *LicenseRecord) IsValid() *AppError {
	if len(lr.Id) != 26 {
		return NewLocAppError("LicenseRecord.IsValid", "model.license_record.is_valid.id.app_error", nil, "")
	}

	if lr.CreateAt == 0 {
		return NewLocAppError("LicenseRecord.IsValid", "model.license_record.is_valid.create_at.app_error", nil, "")
	}

	if len(lr.Bytes) == 0 || len(lr.Bytes) > 10000 {
		return NewLocAppError("LicenseRecord.IsValid", "model.license_record.is_valid.create_at.app_error", nil, "")
	}

	return nil
}

func (lr *LicenseRecord) PreSave() {
	lr.CreateAt = GetMillis()
}
