package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Image struct {
	ID   uint
	Name string
	Uuid string
	Lat  float64
	Lon  float64
}

type SearchResponse struct {
	Lat    float64
	Lon    float64
	Radius float64

	UrlList []string
}

type SearchRequest struct {
	Lat    float64
	Lon    float64
	Radius float64
	Unit   string
}

func (searchRequest SearchRequest) Validate() error {

	return validation.ValidateStruct(&searchRequest,
		validation.Field(&searchRequest.Lat, validation.Required, validation.Min(-85.05112878), validation.Max(85.05112878)),
		validation.Field(&searchRequest.Lon, validation.Required, validation.Min(-180.0), validation.Max(180.0)),
		validation.Field(&searchRequest.Radius, validation.Required, validation.Min(0.0)),
		validation.Field(&searchRequest.Unit, validation.Required, validation.In("km", "m")),
	)
}
