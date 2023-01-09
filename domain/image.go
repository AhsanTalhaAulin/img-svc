package domain

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Image struct {
	ID         uint
	Name       string
	Uuid       string
	Lat        float64
	Lon        float64
	Created_at string
}

var TimeLayout = "2006-01-02 15:04:05"

type SearchResponse struct {
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Radius    float64 `json:"radius"`
	Unit      string  `json:"unit"`
	Timestamp string  `json:"timeStamp"`

	UrlList []string `json:"urlList"`
}

type SearchRequest struct {
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Radius    float64 `json:"radius"`
	Unit      string  `json:"unit"`
	Timestamp string  `json:"timeStamp"`
}

func (searchRequest SearchRequest) Validate() error {

	return validation.ValidateStruct(&searchRequest,
		validation.Field(&searchRequest.Lat, validation.Required, validation.Min(-85.05112878), validation.Max(85.05112878)),
		validation.Field(&searchRequest.Lon, validation.Required, validation.Min(-180.0), validation.Max(180.0)),
		validation.Field(&searchRequest.Radius, validation.Required, validation.Min(0.0)),
		validation.Field(&searchRequest.Unit, validation.Required, validation.In("km", "m")),
		validation.Field(&searchRequest.Timestamp, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (2[0-3]|[01][0-9]):[0-5][0-9]:[0-5][0-9]$"))),
	)
}
