package svc

import (
	"context"
	"encoding/json"
	"img-svc/aws"
	"img-svc/conn"
	"img-svc/domain"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo"
)

func SearchImgInCache(c echo.Context) error {
	start := time.Now()
	log.Printf("Search Request Received at %v", start)

	var searchRequest domain.SearchRequest

	if err := c.Bind(&searchRequest); err != nil {
		return c.String(http.StatusBadRequest, "could not bind request")
	}

	err := searchRequest.Validate()

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	log.Println(searchRequest)

	images := runLuaScript(searchRequest)

	var urlList []string
	for _, name := range images {
		url, err := aws.GetPresignedUrl(name)
		if err != nil {
			log.Println(err)
			url = err.Error()
		}
		urlList = append(urlList, url)
	}

	var response domain.SearchResponse

	response.Lat = searchRequest.Lat
	response.Lon = searchRequest.Lon
	response.Radius = searchRequest.Radius
	response.Unit = searchRequest.Unit
	response.Timestamp = searchRequest.Timestamp
	response.UrlList = urlList

	result, _ := json.Marshal(response)

	log.Printf("Search Request Served From Cache. Time taken: %v", time.Since(start))
	return c.String(http.StatusOK, string(result))

}

func runLuaScript(searchRequest domain.SearchRequest) []string {

	log.Println("Running Lua script ")

	script := redis.NewScript(`
	local radius = ARGV[1]
	local unit = ARGV[2]
	local lon = ARGV[3]
	local lat = ARGV[4]
	local timeStamp = ARGV[5]

	local imgByLocation = redis.call('GEOSEARCH', 'imageLocations', 'FROMLONLAT', lon, lat, 'BYRADIUS', radius, unit)
	
	local result = {}

	for index, value in pairs(imgByLocation) do
		local time, name = string.match(value, "(.+)::(.+)")
		time = tonumber(time)
	
		if(time >= timeStamp-300 and time <=timeStamp+300) then
			table.insert(result, name)
		end
		
	end
	 
	return result
	
	`)

	var ctx = context.Background()
	time, _ := time.Parse(domain.TimeLayout, searchRequest.Timestamp)

	res, err := script.Run(ctx, conn.RedisClient.Rdb, []string{}, searchRequest.Radius, searchRequest.Unit, searchRequest.Lon, searchRequest.Lat, time.Unix()).StringSlice()
	if err != nil {
		log.Println(err)
	}

	log.Println(res)

	log.Println("Lua function ends ")

	return res
}
