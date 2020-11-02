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

func  (ob *model.PERSON) Create() error{
 const sqlstr = `create table PERSON  (
                         ID CHAR(36),
                         NAME CHAR(36),
                         CODE CHAR(14),
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err := config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.PERSON) Delete() error{
   const sqlstr = `DELETE FROM PERSON  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.PERSON) Save() error{
        const sqlstr = `update or insert into  PERSON  (ID, NAME, CODE)
        values ( ob.