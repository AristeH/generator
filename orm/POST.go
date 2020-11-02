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

func  (ob *model.POST) Create() error{
 const sqlstr = `create table POST  (
                         ID CHAR(36),
                         NAME CHAR(150),
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err := config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.POST) Delete() error{
   const sqlstr = `DELETE FROM POST  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.POST) Save() error{
        const sqlstr = `update or insert into  POST  (ID, NAME)
        values ( ob.