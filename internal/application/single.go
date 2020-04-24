package application

import (
	"github.com/alexey-zayats/claim-handler/internal/form"
	"github.com/alexey-zayats/claim-handler/internal/util"
	"time"
)

// Single ..
type Single struct {
	Dirty             bool
	DistrictID        int64
	PassType          int
	Title             string
	Address           string
	Inn               int64
	Ogrn              int64
	CeoName           string
	CeoPhone          string
	CeoEmail          string
	ActivityKind      int64
	Agreement         int
	Reliability       int
	Passes            []Pass
	CityFrom          string
	CityTo            string
	AddressDest       string
	DateFrom          time.Time
	DateTo            time.Time
	OtherReason       string
	WhoNeedsHelp      string
	WhoNeedsHelpPhone string
	DocLinks          string
}

// NewSingle ...
func NewSingle(form *form.Single) *Single {

	app := &Single{}
	app.DistrictID = parseInt64(form.DistrictID)
	app.PassType = int(parseInt64(form.PassType))
	app.Title = form.Title
	app.Address = form.Address
	app.Inn = parseInt64(form.Inn)
	app.Ogrn = 3333333333333
	app.CeoName = form.CeoName
	app.CeoPhone = form.CeoPhone
	app.CeoEmail = form.CeoEmail
	app.ActivityKind = parseInt64(form.ActivityKind)
	app.Agreement = int(parseInt64(form.Personal))
	app.Reliability = int(parseInt64(form.Authenticity))

	app.CityFrom = form.CityFrom
	app.CityTo = form.CityTo
	app.AddressDest = form.AddressDest
	app.DateFrom = form.DateFrom.Time
	app.DateTo = form.DateTo.Time
	app.OtherReason = form.OtherReason
	app.WhoNeedsHelp = form.WhoNeedsHelp
	app.WhoNeedsHelpPhone = form.WhoNeedsHelpPhone
	app.DocLinks = form.DocLinks

	for _, p := range form.Passes {
		app.Passes = append(app.Passes, Pass{
			Car: p.Car,
			FIO: FIO{
				Lastname:   p.Lastname,
				Firstname:  p.Firstname,
				Middlename: p.Middlename,
			},
		})
	}

	return app
}

// Validate ...
func (a *Single) Validate() ValidationErrors {
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

	return ve
}
