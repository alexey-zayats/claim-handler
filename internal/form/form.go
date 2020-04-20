package form

// Form ...
type Form struct {
	DistrictID   string `json:"district_id" validate:"required,numeric"`
	PassType     string `json:"pass_type" validate:"required,numeric"`
	Title        string `json:"title" validate:"required"`
	Address      string `json:"address" validate:"required"`
	Inn          string `json:"inn" validate:"required,numeric,min=10"`
	Ogrn         string `json:"ogrn" validate:"required,numeric,min=10"`
	CeoName      string `json:"ceo_name" validate:"required"`
	CeoPhone     string `json:"ceo_phone" validate:"required"`
	CeoEmail     string `json:"ceo_email" validate:"required,email"`
	ActivityKind string `json:"activity_kind" validate:"required"`
	Personal     string `json:"personal" validate:"required"`
	Authenticity string `json:"authenticity" validate:"required"`
}

// FIO ...
type FIO struct {
	Lastname   string `json:"lastname" validate:"required"`
	Firstname  string `json:"firstname" validate:"required"`
	Middlename string `json:"middlename" validate:"required"`
}

// Car ...
type Car struct {
	FIO
	Car string `json:"car,omitempty"`
}

// Vehicle ...
type Vehicle struct {
	Form
	Passes []Car `json:"people" validate:"required,dive,required"`
}

// People ...
type People struct {
	Form
	Passes []FIO `json:"people" validate:"required,dive,required"`
}
