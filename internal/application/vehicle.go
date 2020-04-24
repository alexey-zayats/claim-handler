package application

import (
	"github.com/alexey-zayats/claim-handler/internal/form"
	"github.com/alexey-zayats/claim-handler/internal/util"
	"html"
	"regexp"
)

// Vehicle ..
type Vehicle struct {
	Dirty        bool
	DistrictID   int64
	PassType     int
	Title        string
	Address      string
	Inn          int64
	Ogrn         int64
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
	app.Inn = parseInt64(form.Inn)
	app.Ogrn = parseInt64(form.Ogrn)
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
		a.Passes[i].Car = util.TrimNumber(p.Car)
	}

	if util.CheckINN(a.Inn) == false {
		ve["inn"] = append(ve["inn"], "Некорректный ИНН")
	}

	if util.CheckOGRN(a.Ogrn) == false {
		ve["ogrn"] = append(ve["org"], "Некорректный ОРГН")
	}

	return ve
}
