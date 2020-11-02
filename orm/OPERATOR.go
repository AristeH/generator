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

func  (ob *model.OPERATOR) Create() error{
 const sqlstr = `create table OPERATOR  (
                         ID CHAR(36),
                         NAME CHAR(100),
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err := config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.OPERATOR) Delete() error{
   const sqlstr = `DELETE FROM OPERATOR  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.OPERATOR) Save() error{
        const sqlstr = `update or insert into  OPERATOR  (ID, NAME)
        values ( ob.