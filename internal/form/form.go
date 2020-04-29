package form

import (
	"strings"
)

// Form ...
type Form struct {
	DistrictID   string `json:"district_id" validate:"required,numeric"`
	PassType     string `json:"pass_type" validate:"required,numeric"`
	Title        string `json:"title" validate:"required"`
	Address      string `json:"address" validate:"required"`
	Inn          string `json:"inn" validate:"required,min=10,max=15"`
	Ogrn         string `json:"ogrn" validate:"required,min=13,max=15"`
	CeoName      string `json:"ceo_name" validate:"required"`
	CeoPhone     string `json:"ceo_phone" validate:"required"`
	CeoEmail     string `json:"ceo_email" validate:"required,email"`
	ActivityKind string `json:"activity_kind" validate:"required"`
	Personal     string `json:"personal" validate:"required"`
	Authenticity string `json:"authenticity" validate:"required"`
}

// Trim ...
func (f *Form) Trim() {
	f.DistrictID = strings.TrimSpace(f.DistrictID)
	f.PassType = strings.TrimSpace(f.PassType)
	f.Title = strings.TrimSpace(f.Title)
	f.Address = strings.TrimSpace(f.Address)
	f.Inn = strings.TrimSpace(f.Inn)
	f.Ogrn = strings.TrimSpace(f.Ogrn)
	f.CeoName = strings.TrimSpace(f.CeoName)
	f.CeoPhone = strings.TrimSpace(f.CeoPhone)
	f.CeoEmail = strings.TrimSpace(f.CeoEmail)
	f.ActivityKind = strings.TrimSpace(f.ActivityKind)
	f.Personal = strings.TrimSpace(f.Personal)
	f.Authenticity = strings.TrimSpace(f.Authenticity)
}

// FIO ...
type FIO struct {
	Lastname   string `json:"lastname" validate:"required"`
	Firstname  string `json:"firstname" validate:"required"`
	Middlename string `json:"middlename" validate:"required"`
}

// Trim ...
func (f FIO) Trim() {
	f.Lastname = strings.TrimSpace(f.Lastname)
	f.Firstname = strings.TrimSpace(f.Firstname)
	f.Middlename = strings.TrimSpace(f.Middlename)
}

// Car ...
type Car struct {
	FIO
	Car string `json:"car" validate:"required,max=15"`
}

// Trim ...
func (c Car) Trim() {
	c.FIO.Trim()
	c.Car = strings.TrimSpace(c.Car)
}

// Vehicle ...
type Vehicle struct {
	Form
	Passes []Car `json:"people" validate:"required,dive,required"`
}

// Trim ...
func (v *Vehicle) Trim() {
	v.Form.Trim()
	for i := range v.Passes {
		v.Passes[i].Trim()
	}
}

// People ...
type People struct {
	Form
	Passes []FIO `json:"people" validate:"required,dive,required"`
}

// Trim ...
func (v *People) Trim() {
	v.Form.Trim()
	for i := range v.Passes {
		v.Passes[i].Trim()
	}
}
