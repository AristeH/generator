// Code generated. DO NOT EDIT.
package orm

import (
	"encoding/json"
	"my/generator/orm"
	"my/generator/model"
	"my/generator/config"
	"io/ioutil"
	"os"
)

func  (ob *model.PEOPLE) Create() error{
 const sqlstr = `create table PEOPLE  ( ID, NAME, CODE,,
                 CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
		_, err = config.DB.Exec(sqlstr)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.PEOPLE) Delete() error{
   const sqlstr = `DELETE FROM PEOPLE  WHERE ID =  ?`
		_, err = config.DB.Exec(sqlstr, PEOPLE.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob *model.PEOPLE) Save() error{
        const sqlstr = `update or insert into  PEOPLE  ( ID, NAME, CODE,)
        values ( ob.ID, ob.NAME, ob.CODE,)
        matching (ID)
     		_, err = config.DB.Exec(sqlstr)
     		if err != nil {
     			return err
     		}`
     		return nil
}

func (ob *model.PEOPLE) Read(id string) error{
   const sqlstr = `select * FROM PEOPLE  WHERE ID =  ?`
   row := config.DB.QueryRow(id)
   err := row.Scan( &ob.ID, &ob.NAME, &ob.CODE,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob *model.PEOPLE) ReadFromJson(file string){
	var recs PEOPLEList

	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &recs)

	for i := 0; i < len(recs.Recs); i++ {
	ob.ID = recs.Recs[i].ID
	ob.NAME = recs.Recs[i].NAME
	ob.CODE = recs.Recs[i].CODE
		ob.Save()
	}

}