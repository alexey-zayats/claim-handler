package form

// Single ...
type Single struct {
	DistrictID        string `json:"district_id" validate:"required,numeric"`
	PassType          string `json:"pass_type" validate:"required,numeric"`
	Title             string `json:"title" validate:"required"`
	Address           string `json:"address" validate:"required"`
	Inn               string `json:"inn" validate:"required,numeric,min=10"`
	CeoName           string `json:"ceo_name" validate:"required"`
	CeoPhone          string `json:"ceo_phone" validate:"required"`
	CeoEmail          string `json:"ceo_email" validate:"required,email"`
	ActivityKind      string `json:"activity_kind" validate:"required"`
	Personal          string `json:"personal" validate:"required"`
	Authenticity      string `json:"authenticity" validate:"required"`
	Passes            []Car  `json:"people" validate:"required,dive,required"`
	CityFrom          string `json:"city_from" validate:"required"`
	CityTo            string `json:"city_to" validate:"required"`
	AddressDest       string `json:"address_dest" validate:"required"`
	DateFrom          string `json:"date_from" validate:"required"`
	DateTo            string `json:"date_to" validate:"required"`
	Reason            string `json:"reason" validate:"required"`
	OtherReason       string `json:"other_reason" validate:"required"`
	WhoNeedsHelp      string `json:"who_needs_help" validate:"required"`
	WhoNeedsHelpPhone string `json:"who_needs_help_phone" validate:"required"`
	DocLinks          string `json:"doc_links" validate:"required"`
}
