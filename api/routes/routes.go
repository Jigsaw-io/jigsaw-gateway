package routes

import (
	"net/http"
	"github.com/zeemzo/jigsaw-gateway/api/businessFacades"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes An Array of type Route
type Routes []Route

/*routes contains all the routes
@author Azeem Ashraf
*/
var routes = Routes{

	Route{
		"POC",
		"POST",
		"/api/transactions/userICO",
		businessFacades.UserICOHandler, //Calls the UserICOHandler
	},
}
