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

func  (ob *model.FORMSDATAXML) Create() error{
 const sqlstr = `create table FORMSDATAXML  (
                         VIEWNAME CHAR(100),
                         ,
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err := config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.FORMSDATAXML) Delete() error{
   const sqlstr = `DELETE FROM FORMSDATAXML  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.FORMSDATAXML) Save() error{
        const sqlstr = `update or insert into  FORMSDATAXML  (VIEWNAME, DATAXML)
        values ( ob.