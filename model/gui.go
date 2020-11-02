// Формат передаваемых данных
// гуиэлемент.типзначения, заголовок, значение

package externalserver

import (
	"encoding/xml"
	"fmt"
)

// Gui Fom
const (
	ListForm = 1
)

// Типы передаваемых данных
const (
	Text    = "1"
	Int     = "2"
	Real    = "3"
	Obj     = "4"
	Data    = "5"
	Command = "6"
)

//Component Список элементов гуи
type Component struct {
	Section interface{} `xml:"section"`
}

// ID - опискание гуи эелемента его тип и представление
type ID struct {
	Gui  int `xml:"gui,attr,omitempty"`
	Type int `xml:"type,attr,omitempty"`
}

//listform
type listform struct {
	Name    string              `xml:"name"`
	Title   string              `xml:"title"`
	Stroki  []arrayFieldSection `xml:"stroki"`
	Buttons []Button            `xml:"buttons"`
}

//Button - описание кнопки
type Button struct {
	Name       string `xml:"name"`
	Parameters string `xml:"parameters"`
	Image      string `xml:"image"`
}

//FieldSection - описание поля со значением и кнопками
type FieldSection struct {
	Name    string `xml:"name"`
	Value   string `xml:"value"`
	Buttons string `xml:"buttons"`
}

type arrayFieldSection struct {
	Fields []FieldSection `xml:"fields"`
}

//Login функция доступа к серверу
func Login(param []string) string {
	v := listform{
		Name:  "listform",
		Title: "Введите регистрационные данные",
		Stroki: []arrayFieldSection{
			{
				Fields: []FieldSection{
					{
						Name:    "Имя пользователя",
						Value:   "Ведите имя",
						Buttons: "",
					},
				},
			},
			{
				Fields: []FieldSection{
					{
						Name:    "Пароль",
						Value:   "Ведите пароль",
						Buttons: "",
					},
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
}
