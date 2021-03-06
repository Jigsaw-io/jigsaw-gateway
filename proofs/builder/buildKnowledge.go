package builder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Jigsaw-io/jigsaw-gateway/model"

	// "github.com/stellar/go/build"
	// "github.com/stellar/go/clients/horizon"
	// "github.com/zeemzo/jigsaw-gateway/constants"
	"github.com/stellar/go/xdr"
	"github.com/Jigsaw-io/jigsaw-gateway/dao"
	"github.com/Jigsaw-io/jigsaw-gateway/proofs/stellarExecuter"

	"github.com/Jigsaw-io/jigsaw-gateway/api/apiModel"
)

type AbstractKnowledgeBuilder struct {
	model.KnowledgeAPI
}

/*BuildUserICOXLM - WORKING MODEL
@author - Azeem Ashraf
@desc - use the parameter to get user public key to build payment transaction for XLM and submit it
@params - ResponseWriter,Request
*/
func (AP *AbstractKnowledgeBuilder) BuildCreateKnowledge(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	object := dao.Connection{}

	var txe xdr.Transaction

	//decode the XDR
	err := xdr.SafeUnmarshalBase64(AP.XDR, &txe)
	if err != nil {
		fmt.Println(err)
	}

	//GET THE TYPE AND IDENTIFIER FROM THE XDR
	AP.PublicKey = txe.SourceAccount.Address()
	AP.Type = strings.TrimLeft(fmt.Sprintf("%s", txe.Operations[0].Body.ManageDataOp.DataValue), "&")
	AP.PreviousTxn = strings.TrimLeft(fmt.Sprintf("%s", txe.Operations[1].Body.ManageDataOp.DataValue), "&")
	AP.KnowledgeHash = strings.TrimLeft(fmt.Sprintf("%s", txe.Operations[2].Body.ManageDataOp.DataValue), "&")

	//SUBMIT THE GATEWAY'S SIGNED XDR
	display1 := stellarExecuter.ConcreteSubmitXDR{XDR: AP.XDR}
	response1 := display1.SubmitXDR(true)
	fmt.Println(AP)

	if response1.Error.Code == 400 {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Failed: " + response1.Error.Message,
		}
		json.NewEncoder(w).Encode(result)
		return
	} else {

		AP.TxnHash = response1.TXNID
		err2 := object.InsertKnowledge(AP.KnowledgeAPI)
		if err2 != nil {
			w.WriteHeader(http.StatusBadRequest)
			result := apiModel.SubmitXDRSuccess{
				Status: "Failed",
			}
			json.NewEncoder(w).Encode(result)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			result := apiModel.SubmitXDRSuccess{
				Status: "Success",
			}
			json.NewEncoder(w).Encode(result)
			return
		}
	}

}
