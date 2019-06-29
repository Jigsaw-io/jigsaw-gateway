package interpreter

import (
	"fmt"
	"github.com/zeemzo/jigsaw-gateway/api/apiModel"
	"github.com/zeemzo/jigsaw-gateway/model"
	"github.com/zeemzo/jigsaw-gateway/proofs/retriever/stellarRetriever"
	"net/http"
	"strings"
)

type POEInterface interface {
	RetrievePOE() model.RetrievePOE
}

type AbstractPOE struct {
	POEStruct apiModel.POEStruct
	// Txn       string
	// ProfileID string
	// Hash      string
}

/*InterpretPOE - Working Model
@author - Azeem Ashraf
@desc - Interprets All the required fields necessary to perform POE
@params - ResponseWriter,Request
*/
func (AP *AbstractPOE) InterpretPOE() model.POE {
	var poeObj model.POE

	object := stellarRetriever.ConcretePOE{POEStruct: AP.POEStruct}

	poeObj.RetrievePOE = object.RetrievePOE()

	if poeObj.RetrievePOE.BCHash == "" {
		return poeObj
	} else {
		poeObj.RetrievePOE.Error = matchingHash(
			poeObj.RetrievePOE.BCHash,
			poeObj.RetrievePOE.DBHash,
			AP.POEStruct.ProfileID,
			poeObj.RetrievePOE.Identifier)
		return poeObj
	}
}

func matchingHash(bcHash string, dbHash string, bcProfile string, dbProfile string) model.Error {
	var Rerr model.Error
	fmt.Println("bcHash", "dbHash")
	fmt.Println(bcHash, dbHash)
	if strings.ToUpper(bcHash) == strings.ToUpper(dbHash) {
		Rerr.Code = http.StatusOK
		Rerr.Message = "BC Hash & DB Hash match."
		
		// if bcProfile == dbProfile {
		// 	Rerr.Code = http.StatusOK
		// 	Rerr.Message = "POE Success! DB & BC Hash and Profile match."
		// } else {
		// 	Rerr.Code = http.StatusOK
		// 	Rerr.Message = "BC Profile & DB Profile didn't match."
		// }
		return Rerr

	} else {
		Rerr.Code = http.StatusOK
		Rerr.Message = "Error! BC Hash & DB Hash din't match."
		return Rerr
	}
}