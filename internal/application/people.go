package application

import (
	"fmt"
	"github.com/alexey-zayats/claim-handler/internal/form"
	"github.com/alexey-zayats/claim-handler/internal/util"
	"html"
	"strings"
)

// People ..
type People struct {
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

// NewPeople ...
func NewPeople(form *form.People) *People {
	app := &People{}

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
func (a *People) Validate() ValidationErrors {
	ve := make(ValidationErrors)

	if err := util.CheckINN(a.Inn); err != nil {
		ve["inn"] = append(ve["inn"], fmt.Sprintf("Некорректный ИНН(%d): %s", a.Inn, err))
	}

	if err := util.CheckOGRN(a.Ogrn); err != nil {
		ve["ogrn"] = append(ve["org"], fmt.Sprintf("Некорректный ОРГН(%d): %s", a.Ogrn, err))
	}

	return ve
}
