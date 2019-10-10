package Tools

import (
	"TipuJson/TipuJson/Modles"
	"reflect"
)

func ScanReflectFeild(waitScanType reflect.Type, waitScanValue reflect.Value) (feilds []*Modles.ReflectField, err error) {
	feilds = []*Modles.ReflectField{}
	for i := 0; i < waitScanType.NumField(); i++ {
		typeFiled := waitScanType.Field(i)
		name := typeFiled.Name
		jsonName := typeFiled.Tag.Get("json")
		if jsonName == "" {
			jsonName = name
		}
		fieldType := typeFiled.Type
		value := waitScanValue.Field(i)
		feilds = append(feilds, &Modles.ReflectField{
			Type:     fieldType,
			Value:    value,
			JsonName: jsonName,
			Name:     name,
		})
	}
	return
}
