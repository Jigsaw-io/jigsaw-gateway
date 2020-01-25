package routes

import (
	"net/http"

	"github.com/Jigsaw-io/jigsaw-gateway/api/businessFacades"
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
		"userICOJIGXU",
		"POST",
		"/api/transactions/userICOJIGXU",
		businessFacades.UserICOJIGXUHandler, //Calls the UserICOJIGXUHandler
	},
	Route{
		"create knowledge",
		"POST",
		"/api/transactions/genesis",
		businessFacades.CreateKnowledge, //Calls the UserICOJIGXUHandler
	},
	Route{
		"add knowledge",
		"POST",
		"/api/transactions/contribute",
		businessFacades.AddKnowledge, //Calls the UserICOJIGXUHandler
	},
	Route{
		"vote",
		"POST",
		"/api/transactions/vote",
		businessFacades.AddVote, //Calls the UserICOJIGXUHandler
	},
	Route{
		"lastVote",
		"GET",
		"/api/transactions/lastVote/{ContributionID}",
		businessFacades.LastVote, //Calls the UserICOJIGXUHandler
	},
	Route{
		"lastContribution",
		"GET",
		"/api/transactions/lastContribution/{KnowledgeID}",
		businessFacades.LastContribution, //Calls the UserICOJIGXUHandler
	},
	Route{
		"LastKnowledge",
		"GET",
		"/api/transactions/lastKnowledge/{PublicKey}",
		businessFacades.LastKnowledge, //Calls the UserICOJIGXUHandler
	},
	Route{
		"POC",
		"POST",
		"/api/transactions/userICOXLM",
		businessFacades.UserICOXLMHandler, //Calls the UserICOXLMHandler
	},
}
