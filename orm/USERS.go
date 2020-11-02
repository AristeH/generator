// Code generated. DO NOT EDIT.
package orm

import (
	"encoding/json"
	"my/server/model"
	"my/server/config"
	"io/ioutil"
	"os"
	"fmt"
)

func  (ob *model.USERS) Create() error{
 const sqlstr = `create table USERS  (
                         ID CHAR(36),
                         NAME CHAR(100),
                         CODE CHAR(20),
                         PASSWORD CHAR(20),
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err := config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.USERS) Delete() error{
   const sqlstr = `DELETE FROM USERS  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.USERS) Save() error{
        const sqlstr = `update or insert into  USERS  (ID, NAME, CODE, PASSWORD)
        values ( ob.