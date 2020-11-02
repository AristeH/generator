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

func  (ob *model.HIERARCHY) Create() error{
 const sqlstr = `create table HIERARCHY  (
                         ID CHAR(36),
                         NAME CHAR(100),
                         ,
                         WEIGHTLEFT CHAR(1),
                         WEIGHTRIGHT CHAR(1),
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err := config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.HIERARCHY) Delete() error{
   const sqlstr = `DELETE FROM HIERARCHY  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.HIERARCHY) Save() error{
        const sqlstr = `update or insert into  HIERARCHY  (ID, NAME, TEXT, WEIGHTLEFT, WEIGHTRIGHT)
        values ( ob.