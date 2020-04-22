package application

import (
	"bytes"
	"github.com/alexey-zayats/claim-handler/internal/form"
	"github.com/alexey-zayats/claim-handler/internal/util"
	"regexp"
	"strconv"
	"strings"
)

// AppKind ...
type AppKind int

const (
	// KindVehicle ...
	KindVehicle AppKind = iota
	// KindPeople ...
	KindPeople
)

// ValidationErrors ...
type ValidationErrors map[string][]string

func (ve ValidationErrors) Error() string {

	buff := bytes.NewBufferString("")

	for key, s := range ve {
		buff.WriteString(key)
		buff.WriteString(": ")
		buff.WriteString(strings.Join(s, ", "))
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// Pass ...
type Pass struct {
	Car        string
	Lastname   string
	Firstname  string
	Middlename string
}

// Application ..
type Application struct {
	Dirty        bool
	Kind         AppKind
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

var re = []*regexp.Regexp{
	// Е 100 EE 123 RUS
	regexp.MustCompile(`((?:\p{L}{1})(?:\s+)?(?:\d{3})(?:\s+)?(?:\p{L}{2})(?:\s+)?(?:\d{2,3})(?:\s+)?(?i:rus?)?)`),
	/*
		// АО 308 26 RUS
		regexp.MustCompile(`((?:\p{L}{2})(?:\s+)?(?:\d{3})(?:\s+)?(?:\d{2,3})(?:\s+)?(?i:rus?)?)`),
		// УМ 5577 34, ТУ 6644 18
		regexp.MustCompile(`((?:\p{L}{2})(?:\s+)?(?:\d{4})(?:\s+)?(?:\d{2,3})(?:\s+)?(?i:rus?)?)`),
		// АЕ 618 Р 29
		regexp.MustCompile(`((?:\p{L}{2})(?:\s+)?(?:\d{3})(?:\s+)?(?:\p{L}{1})(?:\s+)?(?:\d{2,3})(?:\s+)?(?i:rus?)?)`),
		// 0133 ОХ 77
		// 8797 СА 50
		// 2302 КУ 87
		// 4400 РН 50
		regexp.MustCompile(`((?:\d{4})(?:\s+)?(?:\p{L}{2})(?:\s+)?(?:\d{2,3})(?:\s+)?(?i:rus?)?)`),
		// 3456 С 06
		regexp.MustCompile(`((?:\d{4})(?:\s+)?(?:\p{L}{1})(?:\s+)?(?:\d{2,3})(?:\s+)?(?i:rus?)?)`),
		// ТВР 499 91
		regexp.MustCompile(`(?:\p{L}{3})(?:\s+)?((?:\d{3})(?:\s+)?(?:\d{2,3})(?:\s+)?(?i:rus?)?)`),
		// 612 К 36
		regexp.MustCompile(`((?:\d{3})(?:\s+)?(?:\p{L}{1})(?:\s+)?(?:\d{2,3})(?:\s+)?(?i:rus?)?)`),
		// Н 2266 50
		regexp.MustCompile(`(?:\p{L}{1})(?:\s+)?((?:\d{4})(?:\s+)?(?:\d{2,3})(?:\s+)?(?i:rus?)?)`),
		// 009 D 200 77
	*/
}

// Vehicle ...
func Vehicle(form *form.Vehicle) *Application {
	app := &Application{
		Kind: KindVehicle,
	}

	app.DistrictID = app.parseInt64(form.DistrictID)
	app.PassType = int(app.parseInt64(form.PassType))
	app.Title = form.Title
	app.Address = form.Address
	app.Inn = app.parseInt64(form.Inn)
	app.Ogrn = app.parseInt64(form.Ogrn)
	app.CeoName = form.CeoName
	app.CeoPhone = form.CeoPhone
	app.CeoEmail = form.CeoEmail
	app.ActivityKind = app.parseInt64(form.ActivityKind)
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

	app.DistrictID = app.parseInt64(form.DistrictID)
	app.PassType = int(app.parseInt64(form.PassType))
	app.Title = form.Title
	app.Address = form.Address
	app.Inn = app.parseInt64(form.Inn)
	app.Ogrn = app.parseInt64(form.Ogrn)
	app.CeoName = form.CeoName
	app.CeoPhone = form.CeoPhone
	app.CeoEmail = form.CeoEmail
	app.ActivityKind = app.parseInt64(form.ActivityKind)
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
func (a *Application) Validate() ValidationErrors {
	ve := make(ValidationErrors)

	if a.Kind == KindVehicle {
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
	}

	if util.CheckINN(a.Inn) == false {
		ve["inn"] = append(ve["inn"], "Некорректный ИНН")
	}

	if util.CheckOGRN(a.Ogrn) == false {
		ve["ogrn"] = append(ve["org"], "Некорректный ОРГН")
	}

	return ve
}
