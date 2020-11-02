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



func (ob  DEPARTMENT) ElementData(id string) error, string{
	   const sqlstr = `select * FROM DEPARTMENT  WHERE ID5 =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.