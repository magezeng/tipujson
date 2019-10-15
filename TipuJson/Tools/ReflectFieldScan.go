package Tools

import (
	"TipuJson/TipuJson/Modles"
	"fmt"
	"reflect"
)

func ScanStructReflectFeild(waitScanType reflect.Type, waitScanValue reflect.Value) (feilds []*Modles.ReflectField, err error) {
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

func ScanSliceReflectFeild(waitScanType reflect.Type, waitScanValue reflect.Value) (feilds []*Modles.ReflectField, err error) {
	feilds = []*Modles.ReflectField{}
	for i := 0; i < waitScanType.Elem().NumField(); i++ {
		typeFiled := waitScanType.Elem().Field(i)
		name := typeFiled.Name
		jsonName := typeFiled.Tag.Get("json")
		if jsonName == "" {
			jsonName = name
		}
		fieldType := typeFiled.Type
		val := reflect.New(waitScanType.Elem()).Elem() //new一个数组中的元素对象,并拿到新对象的值
		newArr := make([]reflect.Value, 0)
		newArr = append(newArr, val)                       //创建一个新数组并把元素的值追加进去
		resArr := reflect.Append(waitScanValue, newArr...) // 把原数组的值和新的数组合并
		waitScanValue.Set(resArr)                          // 把最终结果返回给原数组
		fmt.Println(waitScanValue.Kind())
		value := reflect.ValueOf(waitScanValue).Field(i)
		feilds = append(feilds, &Modles.ReflectField{
			Type:     fieldType,
			Value:    value,
			JsonName: jsonName,
			Name:     name,
		})
	}
	return
}
