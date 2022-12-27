//       __      _   _
//      //\\    ||   ||  ||     ||  |\\  ||
// 	   //  \\   ||   ||  ||     ||  ||\\ ||
//    //====\\  ||   ||  ||     ||  || \\||
//   //      \\  \\_//   |====  ||  ||  \\|
//

package main

import (
	"github.com/labstack/echo"

	"img-svc/conn"
	"img-svc/svc"
)

// main function
func main() {
	// create a new echo instance
	conn.ConnectDB()
	conn.ConnectAWS()
	e := echo.New()

	e.POST("/img", svc.SaveImg)
	e.GET("/img", svc.GetImg)

	e.Logger.Fatal(e.Start(":8080"))

}
