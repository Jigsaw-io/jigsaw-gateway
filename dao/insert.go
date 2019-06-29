package dao

import (
	"fmt"
	"github.com/zeemzo/jigsaw-gateway/model"
)

/*InsertICOTransaction Insert a single Transaction Object to TransactionCollection in DB
@author - Azeem Ashraf
*/
func (cd *Connection) InsertICOTransaction(Coc model.UserICOAPI) error {

	session, err := cd.connect()
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	c := session.DB("jigsaw-gateway").C("ICO")
	err1 := c.Insert(Coc)
	if err1 != nil {
		fmt.Println(err1)
	}

	return err
}
