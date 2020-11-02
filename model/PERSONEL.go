// Code generated. DO NOT EDIT.

package model

import (
	"encoding/json"
  
 "strings"

    "time"
	"my/server/config"
	"io/ioutil"
	"os"
	"fmt"
)



type PERSONEL struct {
	ID int `json:"Ссылка"` // ID
	ID_PERSON string `json:"физическоелицо"` // ID_PERSON
	ID_POST string `json:"Должности"` // ID_POST
	ID_DEPARTMENT string `json:"Подразделения"` // ID_DEPARTMENT
	PERSONNUMBER string `json:"табельныйномер"` // PERSONNUMBER
	EMPLOYMENTTYPE bool `json:"видзанятости"` // EMPLOYMENTTYPE
	EVENTSTART string `json:"видсобытияначало"` // EVENTSTART
	DATESTART time.Time `json:"датасобытияначало"` // DATESTART
	EVENTEND string `json:"видсобытияокончание"` // EVENTEND
	DATEEND time.Time `json:"датасобытияокончание"` // DATEEND
}

type PERSONELList struct {
   Recs []PERSONEL `json:"перемещения"`
}

func  (ob PERSONEL) Create() error{
    sqlstr := `create table PERSONEL  (ID, ID_PERSON, ID_POST, ID_DEPARTMENT, PERSONNUMBER, EMPLOYMENTTYPE, EVENTSTART, DATESTART, EVENTEND, DATEEND)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob PERSONEL) Delete() error{
   const sqlstr = `DELETE FROM PERSONEL  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob PERSONEL) Save() error{
   sqlstr := "update or insert into  PERSONEL  (ID, ID_PERSON, ID_POST, ID_DEPARTMENT, PERSONNUMBER, EMPLOYMENTTYPE, EVENTSTART, DATESTART, EVENTEND, DATEEND) "+
   " values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.ID, ob.ID_PERSON, ob.ID_POST, ob.ID_DEPARTMENT, ob.PERSONNUMBER, ob.EMPLOYMENTTYPE, ob.EVENTSTART, ob.DATESTART, ob.EVENTEND, ob.DATEEND)
   if err != nil {
     return err
   }
   return nil
}

func (ob PERSONEL) Read(id string) error{
   const sqlstr = `select * FROM PERSONEL  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.ID, &ob.ID_PERSON, &ob.ID_POST, &ob.ID_DEPARTMENT, &ob.PERSONNUMBER, &ob.EMPLOYMENTTYPE, &ob.EVENTSTART, &ob.DATESTART, &ob.EVENTEND, &ob.DATEEND,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob PERSONEL) ReadFromJson(file string){
	var recs PERSONELList

	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &recs)

	for i := 0; i < len(recs.Recs); i++ {
	   ob.ID = recs.Recs[i].ID
	   ob.ID_PERSON = recs.Recs[i].ID_PERSON
	   ob.ID_POST = recs.Recs[i].ID_POST
	   ob.ID_DEPARTMENT = recs.Recs[i].ID_DEPARTMENT
	   ob.PERSONNUMBER = recs.Recs[i].PERSONNUMBER
	   ob.EMPLOYMENTTYPE = recs.Recs[i].EMPLOYMENTTYPE
	   ob.EVENTSTART = recs.Recs[i].EVENTSTART
	   ob.DATESTART = recs.Recs[i].DATESTART
	   ob.EVENTEND = recs.Recs[i].EVENTEND
	   ob.DATEEND = recs.Recs[i].DATEEND
	   ob.Save()
	}

}

func  (ob PERSONEL)  TmplElem(id string) string{
   
	v := listform{
		Name:  "listform",
		Title: "ob PERSONEL",
		Stroki: []arrayFieldSection{
			{
	  		Fields: []FieldSection{
					{
						Name:     "Ссылка",
						Value:    ob.ID,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Физическое лицо",
						Value:    ob.ID_PERSON,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Должности",
						Value:    ob.ID_POST,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Подразделение",
						Value:    ob.ID_DEPARTMENT,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Табельный номер",
						Value:    ob.PERSONNUMBER,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Вид занятости",
						Value:    ob.EMPLOYMENTTYPE,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Вид события начало",
						Value:    ob.EVENTSTART,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Дата события начало",
						Value:    ob.DATESTART,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Вид события окончание",
						Value:    ob.EVENTEND,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Дата события окончание",
						Value:    ob.DATEEND,
						Buttons: "",
					},
				},
			},

		
		Buttons: []Button{
			{
				Name:       "Войти",
				Parameters: "login",
				Image:      "",
			},
			{
				Name:       "Отмена",
				Parameters: "cancel",
				Image:      "",
			},
		},
	}

	output, err := xml.Marshal(v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return string(output)


return ret
}

func  (ob PERSONEL)  FormSpisok() string{
ret := `
<html  window-frame="standard" window-resizable>
<head>
 <meta http-equiv="content-type" content="text/html; charset=UTF-8" />
  <title>Кадровые данные</title>
<style>

  html {
        font:system;
        min-width:min-content;  /* content will not overflow, no scrollbars */
        max-width:100%;         /* max width - whole screen */
        width:600dip;            /* preferred/initial width */
        height:480dip;     /* content will not overflow, no vertical scrollbars on the window */
        min-height:min-content;
        max-height: 70%;        /* 70% of screen height */
        background:window;
        overflow:none;
        padding:0;
        }

  table
  {
    font:system;
    border:1px solid #bdbccc;
    flow:table-fixed;
    width:*;
    height:*;
    padding:0;
    prototype:Grid url(grid.tis);
    border-spacing:0;
    overflow-x: auto;
  }

  table > thead {
    behavior:column-resizer;
  }

  table > tbody {
    overflow-y: scroll-indicator;
  }

  table th
  {
    color: white;
    font-family:"Century Gothic","Verdana";
	  font-size:10pt;
    border:none;
    padding:4px;
    background-color:#DDD;
    background-image:url(header.png);
    background-repeat:expand stretch-left stretch-right stretch-middle;
	  background-position:3px 3px 3px 3px;
  }

  table th.sortable
  {
    padding:4px 12px 4px 4px;
    foreground-repeat:no-repeat;
    foreground-position: 50% 3px auto auto;
  }
  table th.sortable[sort=asc]
  {
    foreground-image:url(stock:arrow-down);
  }
  table th.sortable[sort=desc]
  {
    foreground-image:url(stock:arrow-up);
  }


  table th:hover
  {
	  color: #ffe598;
    transition:blend;
	  background-image:url(header-hover.png);
  }

  table tr:nth-child(odd) { background-color:white; } /* each odd row */
  table tr:nth-child(even) { background-color:#F4F3F9; } /* each even row */

  table tr:current /* current row */
  {
    background-color:#bdbccc;
    color:white;
  }

  table td
  {
    padding:2px;
  }

  table td:nth-child(1) { font-weight:bold; text-align:center; width:0.1*; } /* first column */
  table td:nth-child(2) { border-left:1px solid #bdbccc; color:#8380A0; width:0.1*; } /* second column */
  table td:nth-child(3) /* last column */
  {
     text-align:left;
     overflow:hidden;        /* three line below - ellipsis */
     text-overflow:ellipsis;
     white-space:nowrap;
     width:0.8*;
     border-left:1px solid #bdbccc;
  }

</style>
  <script type="text/tiscript">

    include "decorators.tis";

    const table = self.$(#thetable);
    assert table : "Cannot find #thetable";

    @click @on "button#AddNewRow" :
    {
      var row = new Element("tr");
      row.insert( new Element("td","&yuml;".htmlUnescape()) );
      row.insert( new Element("td","&yuml;") );
      row.insert( new Element("td","small y with diaeresis") );
      table.body.insert(row,1);
      table.setCurrentRow(row);
    }

    @click @on "button#RemoveCurrentRow" :
    {
      if(table.body.length <= 0)
        return;
      var row = table.getCurrentRow();
      if(!row)
        return;
      var row_index = row.index;
      row.remove();
      if( row_index < table.body.length)
        table.setCurrentRow(table.body[row_index]);
    }

    @click @on "button#SwapCurrentRowWithFirst" :
    {
      var row = table.getCurrentRow();
      if(!row)
        return;
      if(row === table.body.first)
        return;
      row.swap(table.body.first);
    }

    @click @on "button#Sort2C" :
    {
      table.body.sort( :r1,r2: r1[1].text < r2[1].text? -1:1 );
    }


    // attaching custom handler to onRowClick port.
    table.onRowClick = function( row, reason )
    {
      if(reason == #by_mouse)
        //stdout << "got mouse click on row:" << row << "\n";
        debug row: row, row.index;
    }

    // attaching custom handler to onRowClick port.
    table.onRowDoubleClick = function( row, reason )
    {
      stdout << "got double click on row:" << row << "\n";
    }
    </script>

</head>

<body>

<table id="thetable"  style="margin:20px">
  <thead><tr>
   <th .sortable>Физическое лицо</th><th .sortable>Должности</th><th .sortable>Подразделение</th><th .sortable>Табельный номер</th><th .sortable>Вид занятости</th><th .sortable>Вид события начало</th><th .sortable>Дата события начало</th><th .sortable>Вид события окончание</th><th .sortable>Дата события окончание</th>
   </tr></thead>
  <tbody>`

rows, _ := config.DB.Query( `select * FROM PERSONEL`)
defer rows.Close()
for rows.Next() {
   var obj PERSONEL
   rows.Scan( &obj.