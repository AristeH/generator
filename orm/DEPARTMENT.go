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

func  (ob *model.DEPARTMENT) Create() error{
 const sqlstr = `create table DEPARTMENT  (
                         ID CHAR(36),
                         NAME CHAR(200),
                         LEVEL CHAR(1),
                         ID_PARENT CHAR(36),
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err := config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.DEPARTMENT) Delete() error{
   const sqlstr = `DELETE FROM DEPARTMENT  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.DEPARTMENT) Save() error{
        const sqlstr = `update or insert into  DEPARTMENT  (ID, NAME, LEVEL, ID_PARENT)
        values ( ob.