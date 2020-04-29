package application

import (
	"fmt"
	"github.com/alexey-zayats/claim-handler/internal/form"
	"github.com/alexey-zayats/claim-handler/internal/util"
	"html"
	"regexp"
	"strings"
)

// Vehicle ..
type Vehicle struct {
	Dirty        bool
	DistrictID   int64
	PassType     int
	Title        string
	Address      string
	Inn          string
	Ogrn         string
	CeoName      string
	CeoPhone     string
	CeoEmail     string
	ActivityKind int64
	Agreement    int
	Reliability  int
	Passes       []Pass
}

// NewVehicle ...
func NewVehicle(form *form.Vehicle) *Vehicle {
	app := &Vehicle{}

	app.DistrictID = parseInt64(form.DistrictID)
	app.PassType = int(parseInt64(form.PassType))
	app.Title = html.EscapeString(form.Title)
	app.Address = html.EscapeString(form.Address)
	app.Inn = strings.TrimSpace(form.Inn)
	app.Ogrn = strings.TrimSpace(form.Ogrn)
	app.CeoName = html.EscapeString(form.CeoName)
	app.CeoPhone = html.EscapeString(form.CeoPhone)
	app.CeoEmail = html.EscapeString(form.CeoEmail)
	app.ActivityKind = parseInt64(form.ActivityKind)
	app.Agreement = int(parseInt64(form.Personal))
	app.Reliability = int(parseInt64(form.Authenticity))

	for _, p := range form.Passes {
		app.Passes = append(app.Passes, Pass{
			Car: html.EscapeString(p.Car),
			FIO: FIO{
				Lastname:   html.EscapeString(p.Lastname),
				Firstname:  html.EscapeString(p.Firstname),
				Middlename: html.EscapeString(p.Middlename),
			},
		})
	}

	return app
}

var re = []*regexp.Regexp{
	// Е 100 EE 123 RUS
	regexp.MustCompile(`((?:\p{L}{1})(?:\s+)?(?:\d{3})(?:\s+)?(?:\p{L}{2})(?:\s+)?(?:\d{2,3})(?:\s+)?(?i:rus?)?)`),
}

// Validate ...
func (a *Vehicle) Validate() ValidationErrors {
	ve := make(ValidationErrors)

	for i, p := range a.Passes {

		ok := false
		for _, r := range re {
			if r.MatchString(p.Car) == true {
				ok = true
				break
			}
		}

		a.Dirty = !ok
		a.Passes[i].Car = util.TrimNumber(util.NormalizeCarNumber(p.Car))
	}

	if err := util.CheckINN(a.Inn); err != nil {
		ve["inn"] = append(ve["inn"], fmt.Sprintf("Некорректный ИНН(%s): %s", a.Inn, err))
	}

	if err := util.CheckOGRN(a.Ogrn); err != nil {
		ve["ogrn"] = append(ve["org"], fmt.Sprintf("Некорректный ОРГН(%s): %s", a.Ogrn, err))
	}

	return ve
}
