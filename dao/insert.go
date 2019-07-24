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


/*InsertKnowledge Insert a single Transaction Object to TransactionCollection in DB
@author - Azeem Ashraf
*/
func (cd *Connection) InsertKnowledge(Coc model.KnowledgeAPI) error {

	session, err := cd.connect()
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	c := session.DB("jigsaw-gateway").C("Knowledge")
	err1 := c.Insert(Coc)
	if err1 != nil {
		fmt.Println(err1)
	}

	return err
}


/*InsertContributions Insert a single Transaction Object to TransactionCollection in DB
@author - Azeem Ashraf
*/
func (cd *Connection) InsertContributions(Coc model.ContributionAPI) error {

	session, err := cd.connect()
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	c := session.DB("jigsaw-gateway").C("Contributions")
	err1 := c.Insert(Coc)
	if err1 != nil {
		fmt.Println(err1)
	}

	return err
}


/*InsertVote Insert a single Transaction Object to TransactionCollection in DB
@author - Azeem Ashraf
*/
func (cd *Connection) InsertVote(Coc model.VoteAPI) error {

	session, err := cd.connect()
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	c := session.DB("jigsaw-gateway").C("Votes")
	err1 := c.Insert(Coc)
	if err1 != nil {
		fmt.Println(err1)
	}

	return err
}