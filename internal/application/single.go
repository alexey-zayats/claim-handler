package application

import (
	"fmt"
	"github.com/alexey-zayats/claim-handler/internal/form"
	"github.com/alexey-zayats/claim-handler/internal/util"
	"html"
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
	app.Title = html.EscapeString(form.Title)
	app.Address = html.EscapeString(form.Address)
	app.Inn = parseInt64(form.Inn)
	app.Ogrn = 3333333333333
	app.CeoName = html.EscapeString(form.CeoName)
	app.CeoPhone = html.EscapeString(form.CeoPhone)
	app.CeoEmail = html.EscapeString(form.CeoEmail)
	app.ActivityKind = parseInt64(form.ActivityKind)
	app.Agreement = int(parseInt64(form.Personal))
	app.Reliability = int(parseInt64(form.Authenticity))

	app.CityFrom = html.EscapeString(form.CityFrom)
	app.CityTo = html.EscapeString(form.CityTo)
	app.AddressDest = html.EscapeString(form.AddressDest)
	app.DateFrom = form.DateFrom.Time
	app.DateTo = form.DateTo.Time
	app.OtherReason = html.EscapeString(form.OtherReason)
	app.WhoNeedsHelp = html.EscapeString(form.WhoNeedsHelp)
	app.WhoNeedsHelpPhone = html.EscapeString(form.WhoNeedsHelpPhone)
	app.DocLinks = html.EscapeString(form.DocLinks)

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
		a.Passes[i].Car = util.TrimNumber(util.NormalizeCarNumber(p.Car))
	}

	var err error
	err = util.CheckINN(a.Inn)
	if err != nil {
		ve["inn"] = append(ve["inn"], fmt.Sprintf("Некорректный ИНН(%d): %s", a.Inn, err))
	}

	return ve
}
