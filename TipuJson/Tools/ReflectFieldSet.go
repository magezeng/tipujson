package Tools

import (
	"TipuJson/TipuJson/Modles"
	"fmt"
)

func SetStructReflectField(reflectFilde []*Modles.ReflectField, expression *Modles.JsonExpression) (err error) {
	for filedName := range reflectFilde {
		name := reflectFilde[filedName].Name
		fmt.Println(name)
	}
	return
}

func SetSliceReflectField(reflectFilde []*Modles.ReflectField) (err error) {
	return nil
}

func SetMapReflectField(reflectFilde []*Modles.ReflectField) (err error) {
	return nil
}
