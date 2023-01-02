package domain

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
