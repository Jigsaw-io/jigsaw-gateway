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
		"/api/transactions/userICOJIGXU",
		businessFacades.UserICOJIGXUHandler, //Calls the UserICOJIGXUHandler
	},
	Route{
		"POC",
		"POST",
		"/api/transactions/userICOXLM",
		businessFacades.UserICOXLMHandler, //Calls the UserICOXLMHandler
	},
}
