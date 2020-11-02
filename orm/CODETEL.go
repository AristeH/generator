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

func  (ob *model.CODETEL) Create() error{
 const sqlstr = `create table CODETEL  (
                         ID CHAR(36),
                         NAME CHAR(50),
                         CODE CHAR(50),
                         ID_OPERATOR CHAR(36),
                         ID_REGION CHAR(36),
                         S CHAR(9),
                         PO CHAR(9),
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err := config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.CODETEL) Delete() error{
   const sqlstr = `DELETE FROM CODETEL  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.CODETEL) Save() error{
        const sqlstr = `update or insert into  CODETEL  (ID, NAME, CODE, ID_OPERATOR, ID_REGION, S, PO)
        values ( ob.