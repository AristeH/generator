package main

import (
	"fmt"
	_ "io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/jmoiron/sqlx"
	_ "github.com/nakagami/firebirdsql"
)

type MetaData struct {
	Name        string
	Json        string
	Form        string
	Description string
}

var mapnames = make(map[string]MetaData)

type Fields struct {
	FieldName      string
	FIELD_TYPE     string
	FIELD_LENGTH   int
	FIELD_SCALE    string
	FIELD_SUB_TYPE string
	Json           string
	Form           string
	Create         string
	Query          string
}

type MetaTable struct {
	Fields   []Fields //список полей таблицы
	Name     string   // имя таблицы
	Json     string   // тег json
	Form     string   // имя формы
	Type     string   // тип таблицы
	Spisok   string
	Spisokob string
	Row      string
	Shapka   string
	Vopr     string
	Integer  bool
	Time     bool
}

// SchemaMetaData struct
type SchemaMetaData struct {
	TablesMetaData []MetaTable
	ViewsMetaData  []MetaTable
}

var SMD SchemaMetaData
var templates *template.Template

// инициализация вспомогательных данных
func initmap() {
	// таблицы
	mapnames["CODETEL"] = MetaData{"Коды телефонов", "кодытелефонов", "Коды телефонов", ""}
	mapnames["PERSON"] = MetaData{"Физические лица", "ФизическиеЛица", "Физические лица", ""}
	mapnames["PEOPLE"] = MetaData{"Физические лица", "ФизическиеЛица", "Физические лица", ""}
	mapnames["DEPARTMENT"] = MetaData{"Подразделения", "подразделения", "Подразделения", ""}
	mapnames["POST"] = MetaData{"Должности", "должности", "Должности", ""}
	mapnames["PERSONEL"] = MetaData{"Кадровые данные", "перемещения", "Кадровые данные", ""}
	mapnames["USERS"] = MetaData{"Пользователи", "Пользователи", "Пользователи", ""}
	mapnames["MD"] = MetaData{"Метаданные", "Метаданные", "Метаданные", ""}

	// поля
	mapnames["ID"] = MetaData{"guid", "Ссылка", "Ссылка", ""}
	mapnames["NAME"] = MetaData{"name", "Наименование", "Наименование", ""}
	mapnames["ID_PARENT"] = MetaData{"parent", "Родитель", "Папка", ""}
	mapnames["LEVEL"] = MetaData{"level", "Уровень", "Уровень", ""}
	mapnames["CODE"] = MetaData{"cod", "Код", "Код", ""}
	mapnames["ID_PERSON"] = MetaData{"guid", "физическоелицо", "Физическое лицо", ""}
	mapnames["ID_POST"] = MetaData{"guid", "Должности", "Должности", ""}
	mapnames["ID_DEPARTMENT"] = MetaData{"guid", "Подразделения", "Подразделение", ""}
	mapnames["EVENTSTART"] = MetaData{"eventstart", "видсобытияначало", "Вид события начало", ""}
	mapnames["DATESTART"] = MetaData{"datestart", "датасобытияначало", "Дата события начало", ""}
	mapnames["EVENTEND"] = MetaData{"eventend", "видсобытияокончание", "Вид события окончание", ""}
	mapnames["DATEEND"] = MetaData{"dateend", "датасобытияокончание", "Дата события окончание", ""}
	mapnames["PERSONNUMBER"] = MetaData{"PersonalNumber", "табельныйномер", "Табельный номер", ""}
	mapnames["ID_OPERATOR"] = MetaData{"operator", "оператор", "Оператор", ""}
	mapnames["ID_REGION"] = MetaData{"region", "регион", "Регион", ""}
	mapnames["EMPLOYMENTTYPE"] = MetaData{"guid", "видзанятости", "Вид занятости", ""}
	mapnames["S"] = MetaData{"s", "с", "С", ""}
	mapnames["PO"] = MetaData{"po", "по", "По", ""}
	mapnames["JSON"] = MetaData{"JSON", "JSON", "JSON", ""}
	mapnames["FORM"] = MetaData{"FORM", "FORM", "FORM", ""}
	mapnames["NOTE"] = MetaData{"NOTE", "NOTE", "NOTE", ""}
	mapnames["TABLE"] = MetaData{"TABLE", "TABLE", "TABLE", ""}
}

func SetTable(name string, f Fields) {

	tek := -1
	for i := range SMD.TablesMetaData {
		if SMD.TablesMetaData[i].Name == strings.TrimSpace(name) {
			tek = i
			break
		}
	}
	if tek == -1 {
		tab := new(MetaTable)
		tab.Name = strings.TrimSpace(name)
		tab.Json = mapnames[strings.TrimSpace(name)].Json
		tab.Form = mapnames[strings.TrimSpace(name)].Form
		tab.Fields = make([]Fields, 0, 20)
		SMD.TablesMetaData = append(SMD.TablesMetaData, *tab)
		tek = len(SMD.TablesMetaData) - 1
	}

	fi := new(Fields)
	fi.FIELD_TYPE = f.FIELD_TYPE
	fi.FieldName = strings.TrimSpace(f.FieldName)
	fi.Json = mapnames[strings.TrimSpace(f.FieldName)].Json
	fi.Form = mapnames[strings.TrimSpace(f.FieldName)].Form
	fi.FIELD_LENGTH = f.FIELD_LENGTH / 4
	switch strings.TrimSpace(fi.FIELD_TYPE) {
	case "14":
		fi.FIELD_TYPE = "string"
		fi.Create = fi.FieldName + " CHAR(" + strconv.Itoa(fi.FIELD_LENGTH) + ")"
		fi.Query = "obj." + fi.FieldName
	case "23":
		fi.FIELD_TYPE = "bool"
		//fi.Create = fi.Name+" CHAR("+strconv.Itoa(fi.FIELD_LENGTH)+")"

	case "8":
		fi.FIELD_TYPE = "int"
		fi.Create = fi.FieldName + " CHAR(" + strconv.Itoa(fi.FIELD_LENGTH) + ")"
		fi.Query = "strconv.Itoa(obj." + fi.FieldName + ")"
	case "16":
		fi.FIELD_TYPE = "int"
		fi.Query = "strconv.Itoa(obj." + fi.FieldName + ")"
	case "37":
		fi.FIELD_TYPE = "string"
		fi.Create = fi.FieldName + " CHAR(" + strconv.Itoa(fi.FIELD_LENGTH) + ")"
		fi.Query = "obj." + fi.FieldName
	case "12":
		fi.FIELD_TYPE = "time.Time"
	case "13":
		fi.FIELD_TYPE = "time.Time"
	case "261":
		fi.FIELD_TYPE = "string"
		fi.Query = "obj." + fi.FieldName
	}
	SMD.TablesMetaData[tek].Fields = append(SMD.TablesMetaData[tek].Fields, *fi)

	if SMD.TablesMetaData[tek].Row == "" {
		if fi.FieldName != "ID" {
			SMD.TablesMetaData[tek].Row = "`<td>` +strings.TrimSpace(" + fi.Query + ")+`</td>`"
			SMD.TablesMetaData[tek].Shapka = "<th .sortable>" + fi.Form + "</th>"
		}
	} else {
		if fi.FieldName != "ID" {
			SMD.TablesMetaData[tek].Row = SMD.TablesMetaData[tek].Row + "+`<td>`+strings.TrimSpace(" + fi.Query + ")+`</td>`"
			SMD.TablesMetaData[tek].Shapka = SMD.TablesMetaData[tek].Shapka + "<th .sortable>" + fi.Form + "</th>"
		}
	}
	if SMD.TablesMetaData[tek].Spisok == "" {
		SMD.TablesMetaData[tek].Spisok = fi.FieldName
		SMD.TablesMetaData[tek].Spisokob = "ob." + fi.FieldName
		SMD.TablesMetaData[tek].Vopr = "?"

	} else {
		SMD.TablesMetaData[tek].Spisok = SMD.TablesMetaData[tek].Spisok + ", " + fi.FieldName
		SMD.TablesMetaData[tek].Spisokob = SMD.TablesMetaData[tek].Spisokob + ", " + "ob." + fi.FieldName
		SMD.TablesMetaData[tek].Vopr = SMD.TablesMetaData[tek].Vopr + ", ?"

		if fi.FIELD_TYPE == "string" {

		}
		if fi.FIELD_TYPE == "int" {

			SMD.TablesMetaData[tek].Integer = true
		}
		if fi.FIELD_TYPE == "time.Time" {

			SMD.TablesMetaData[tek].Time = true
		}
	}

}

func main() {
	initmap()
	conn, err := sqlx.Open("firebirdsql", "sysdba:masterkey@localhost:3050/C:/obmen/FIRST.fdb?auth_plugin_name=Legacy_Auth&wire_auth=true&column_name_to_lower=false")
	SMD.TablesMetaData = make([]MetaTable, 0, 20)

	if err != nil {
		fmt.Println(err)
	}

	q := `select R.RDB$RELATION_NAME,  R.RDB$FIELD_NAME,
	F.RDB$FIELD_LENGTH, F.RDB$FIELD_TYPE, F.RDB$FIELD_SCALE
	from RDB$FIELDS F, RDB$RELATION_FIELDS R
	where F.RDB$FIELD_NAME = R.RDB$FIELD_SOURCE and R.RDB$SYSTEM_FLAG = 0
	order by R.RDB$RELATION_NAME, R.RDB$FIELD_POSITION`
	rows, err := conn.Query(q)
	defer rows.Close()

	RELATION_NAME := ""

	for rows.Next() {
		f := Fields{
			FieldName:    "",
			FIELD_TYPE:   "",
			FIELD_LENGTH: 0,
			FIELD_SCALE:  "",
		}
		if err := rows.Scan(&RELATION_NAME, &f.FieldName, &f.FIELD_LENGTH, &f.FIELD_TYPE, &f.FIELD_SCALE); err != nil {
			fmt.Println(err)
		}
		SetTable(RELATION_NAME, f)
	}

	tmpl, err := template.ParseFiles("E:/Aristeh/go/src/пппппп/generator/templates/model.tmpl")
	//	tmpl, err := template.ParseFiles("C:/GOPATH/src/my/generator/templates/model.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	s := SMD.TablesMetaData
	for i := range s {
		if s[i].Name != "" {
			file, _ := os.Create("E:/Aristeh/go/src/пппппп/server/model/" + s[i].Name + ".go")
			//file, _ := os.Create("C:/GOPATH/src/my/server/model/" + s[i].Name + ".go")
			tmpl.Execute(file, SMD.TablesMetaData[i])
		}
	}
}