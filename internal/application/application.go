package application

import (
	"github.com/alexey-zayats/claim-handler/internal/form"
	"regexp"
	"strconv"
)

// AppKind ...
type AppKind int

const (
	// KindVehicle ...
	KindVehicle AppKind = iota
	// KindPeople ...
	KindPeople
)

// Error ...
type Error struct {
	Name  string
	Value string
}

// Validation ...
type Validation []Error

// Pass ...
type Pass struct {
	Car        string
	Lastname   string
	Firstname  string
	Middlename string
}

// Application ..
type Application struct {
	Kind         AppKind
	DistrictID   int
	PassType     int
	Title        string
	Address      string
	Inn          int64
	Ogrn         int64
	CeoName      string
	CeoPhone     string
	CeoEmail     string
	ActivityKind int
	Agreement    int
	Reliability  int
	Passes       []Pass
}

var re = regexp.MustCompile(`((?:\p{L}{1})(?:\s+)?(?:\d{3})(?:\s+)?(?:\p{L}{2})(?:\s+)?(?:\d{2,3}))`)

// Vehicle ...
func Vehicle(form *form.Vehicle) *Application {
	app := &Application{
		Kind: KindVehicle,
	}

	app.DistrictID = int(app.parseInt64(form.DistrictID))
	app.PassType = int(app.parseInt64(form.PassType))
	app.Title = form.Title
	app.Address = form.Address
	app.Inn = app.parseInt64(form.Inn)
	app.Ogrn = app.parseInt64(form.Ogrn)
	app.CeoName = form.CeoName
	app.CeoPhone = form.CeoPhone
	app.CeoEmail = form.CeoEmail
	app.ActivityKind = int(app.parseInt64(form.ActivityKind))
	app.Agreement = int(app.parseInt64(form.Personal))
	app.Reliability = int(app.parseInt64(form.Authenticity))

	for _, p := range form.Passes {
		app.Passes = append(app.Passes, Pass{
			Car:        p.Car,
			Lastname:   p.Lastname,
			Firstname:  p.Firstname,
			Middlename: p.Middlename,
		})
	}

	return app
}

// People ...
func People(form *form.People) *Application {
	app := &Application{
		Kind: KindPeople,
	}

	app.DistrictID = int(app.parseInt64(form.DistrictID))
	app.PassType = int(app.parseInt64(form.PassType))
	app.Title = form.Title
	app.Address = form.Address
	app.Inn = app.parseInt64(form.Inn)
	app.Ogrn = app.parseInt64(form.Ogrn)
	app.CeoName = form.CeoName
	app.CeoPhone = form.CeoPhone
	app.CeoEmail = form.CeoEmail
	app.ActivityKind = int(app.parseInt64(form.ActivityKind))
	app.Agreement = int(app.parseInt64(form.Personal))
	app.Reliability = int(app.parseInt64(form.Authenticity))

	for _, p := range form.Passes {
		app.Passes = append(app.Passes, Pass{
			Lastname:   p.Lastname,
			Firstname:  p.Firstname,
			Middlename: p.Middlename,
		})
	}

	return app
}

func (a *Application) parseInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		n = 0
	}
	return n
}

// Validate ...
func (a *Application) Validate() Validation {
	var v Validation

	v = append(v, Error{
		Name:  "aaa",
		Value: "bbb",
	})

	return v
}
