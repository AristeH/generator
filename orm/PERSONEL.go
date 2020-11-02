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

func  (ob *model.PERSONEL) Create() error{
 const sqlstr = `create table PERSONEL  (
                         ,
                         ID_PERSON CHAR(36),
                         ID_POST CHAR(36),
                         ID_DEPARTMENT CHAR(36),
                         PERSONNUMBER CHAR(15),
                         ,
                         EVENTSTART CHAR(36),
                         ,
                         EVENTEND CHAR(36),
                         ,
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err := config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.PERSONEL) Delete() error{
   const sqlstr = `DELETE FROM PERSONEL  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.PERSONEL) Save() error{
        const sqlstr = `update or insert into  PERSONEL  (ID, ID_PERSON, ID_POST, ID_DEPARTMENT, PERSONNUMBER, EMPLOYMENTTYPE, EVENTSTART, DATESTART, EVENTEND, DATEEND)
        values ( ob.