//       __      _   _
//      //\\    ||   ||  ||     ||  |\\  ||
// 	   //  \\   ||   ||  ||     ||  ||\\ ||
//    //====\\  ||   ||  ||     ||  || \\||
//   //      \\  \\_//   |====  ||  ||  \\|
//

package main

import (
	"github.com/labstack/echo"

	"img-svc/cache"
	"img-svc/conn"
	"img-svc/svc"
)

// main function
func main() {
	// create a new echo instance
	conn.ConnectDB()
	conn.ConnectAWS()
	conn.ConnectRedis()
	cache.LoadDbInCache()

	e := echo.New()

	e.POST("/img", svc.SaveImg)
	e.GET("/img", svc.GetImg)
	e.POST("/searchImg", svc.SearchImgInCache)
	e.POST("/searchImg/lua", svc.SearchImgWithLuaScript)
	e.GET("/loadCache", svc.LoadDbInCache)

	e.Logger.Fatal(e.Start(":8080"))

}
