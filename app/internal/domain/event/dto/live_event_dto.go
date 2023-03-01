package dto

import validation "github.com/go-ozzo/ozzo-validation"

type LiveEventDTO struct {
	Locale   string `json:"locale"`
	TimeZone string `json:"timeZone"`
	SportId  string `json:"sportId"`
}

func (l LiveEventDTO) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Locale, validation.Length(2, 10)),
		validation.Field(&l.TimeZone, validation.Length(1, 3)),
		validation.Field(&l.SportId, validation.Required, validation.Length(0, 100)),
	)
}
