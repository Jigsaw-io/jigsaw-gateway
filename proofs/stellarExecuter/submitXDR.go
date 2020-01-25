package stellarExecuter

import (
	"fmt"
	"net/http"

	"github.com/stellar/go/clients/horizon"

	// "github.com/zeemzo/jigsaw-gateway/api/apiModel"
	"github.com/Jigsaw-io/jigsaw-gateway/model"
)

type ConcreteSubmitXDR struct {
	XDR string
}

/*SubmitXDR - WORKING MODEL
@author - Azeem Ashraf
@desc - Submits the XDR to stellar horizon api and returns the TXN hash.
@params - XDR
*/
func (cd *ConcreteSubmitXDR) SubmitXDR(testnet bool) model.SubmitXDRResponse {

	var response model.SubmitXDRResponse

	if testnet {
		// fmt.Println(cd.XDR)
		resp, err := horizon.DefaultTestNetClient.SubmitTransaction(cd.XDR)
		if err != nil {

			error1 := err.(*horizon.Error)
			TC, _ := error1.ResultCodes()
			for _, element := range TC.OperationCodes {
				response.Error.Message = response.Error.Message + element + "? "
			}

			fmt.Println(response.Error.Message)

			response.Error.Code = http.StatusBadRequest
			// response.Error.Message = TC.OperationCodes
			return response
		}

		// fmt.Println("Successful Transaction:")
		// fmt.Println("Ledger:", resp.Ledger)
		fmt.Println("Hash:", resp.Hash)

		response.Error.Code = http.StatusOK
		response.Error.Message = "Transaction performed in the blockchain."
		response.TXNID = resp.Hash
	} else {
		resp, err := horizon.DefaultPublicNetClient.SubmitTransaction(cd.XDR)
		if err != nil {

			error1 := err.(*horizon.Error)
			TC, _ := error1.ResultCodes()
			for _, element := range TC.OperationCodes {
				response.Error.Message = response.Error.Message + element + "? "
			}

			fmt.Println(response.Error.Message)

			response.Error.Code = http.StatusBadRequest
			// response.Error.Message = TC.OperationCodes
			return response
		}

		// fmt.Println("Successful Transaction:")
		// fmt.Println("Ledger:", resp.Ledger)
		fmt.Println("Hash:", resp.Hash)

		response.Error.Code = http.StatusOK
		response.Error.Message = "Transaction performed in the blockchain."
		response.TXNID = resp.Hash
	}

	return response

}
