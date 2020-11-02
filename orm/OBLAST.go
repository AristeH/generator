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

func  (ob *model.OBLAST) Create() error{
 const sqlstr = `create table OBLAST  (
                         ID CHAR(36),
                         NAME CHAR(100),
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err := config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.OBLAST) Delete() error{
   const sqlstr = `DELETE FROM OBLAST  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.OBLAST) Save() error{
        const sqlstr = `update or insert into  OBLAST  (ID, NAME)
        values ( ob.