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



func (ob  PERSON) ElementData(id string) error, string{
	   const sqlstr = `select * FROM PERSON  WHERE ID5 =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.