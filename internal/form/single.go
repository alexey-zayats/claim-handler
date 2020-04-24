package form

import (
	"github.com/pkg/errors"
	"strings"
	"time"
)

// Single ...
type Single struct {
	DistrictID        string `json:"district_id" validate:"required,numeric"`
	PassType          string `json:"pass_type" validate:"required,numeric"`
	Title             string `json:"title,omitempty"`
	Address           string `json:"address" validate:"required"`
	Inn               string `json:"inn" validate:"required,numeric,min=10"`
	CeoName           string `json:"ceo_name" validate:"required"`
	CeoPhone          string `json:"ceo_phone" validate:"required"`
	CeoEmail          string `json:"ceo_email" validate:"required,email"`
	ActivityKind      string `json:"reason" validate:"required"`
	Personal          string `json:"personal" validate:"required"`
	Authenticity      string `json:"authenticity" validate:"required"`
	Passes            []Car  `json:"people" validate:"required,dive,required"`
	CityFrom          string `json:"city_from" validate:"required"`
	CityTo            string `json:"city_to" validate:"required"`
	AddressDest       string `json:"address_dest" validate:"required"`
	DateFrom          Date   `json:"date_from" validate:"required"`
	DateTo            Date   `json:"date_to" validate:"required"`
	OtherReason       string `json:"other_reason,omitempty"`
	WhoNeedsHelp      string `json:"who_needs_help" validate:"required"`
	WhoNeedsHelpPhone string `json:"who_needs_help_phone" validate:"required"`
	DocLinks          string `json:"doc_links" validate:"required"`
}

// Date ...
type Date struct {
	time.Time
}

// UnmarshalJSON ...
func (d *Date) UnmarshalJSON(input []byte) error {
	s := string(input)
	s = strings.Trim(s, `"`)
	newTime, err := time.Parse("02.01.2006", s)
	if err != nil {
		return errors.Wrap(err, "указана неверная дата")
	}

	d.Time = newTime
	return nil
}
