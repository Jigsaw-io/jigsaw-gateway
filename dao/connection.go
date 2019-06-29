package dao

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

/*Connection The Mgo Connection
@author - Azeem Ashraf
*/
type Connection struct {
}

func (cd *Connection) connect()(*mgo.Session,error) {
	//mongo connection to Zeemzo Mlab Account
	session, err := mgo.Dial("mongodb://zeemzo:abcd1234@ds243897.mlab.com:43897/jigsaw-gateway")
	

	if err != nil {
		fmt.Println(err)
	}
	return session,err

}
