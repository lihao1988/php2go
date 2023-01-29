package array

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

var simpleTypes = []string{
	"bool", "int", "int8", "int16", "int32", "int64", "uint", "uint8",
	"uint16", "uint32", "uint64", "uintptr", "float32", "float64", "string",
}

// Column array column to be new array
// php array_column
func Column(dest, input interface{}, columnKey, indexKey string) {
	// data validate
	if columnKey == "" && indexKey == "" {
		panic("columnKey or indexKey must be at least one has value")
	}

	// dest interface
	dValue := reflect.ValueOf(dest)
	if dValue.Kind() != reflect.Ptr || dValue.IsNil() {
		panic(fmt.Sprintf("haystack: d type[%T] error", reflect.TypeOf(dValue)))
	}
	dType, dElemType, dKeyType := indirectForArr(dValue)
	if !In(dKeyType, simpleTypes) {
		// only support 'simpleType'
		panic("haystack: dest key type must be 'simpleType'")
	}
	if indexKey != "" {
		if dType != reflect.Map {
			panic("haystack: dest type must be map")
		}
	} else {
		if dType != reflect.Slice && dType != reflect.Array {
			panic("haystack: dest type must be slice or array ")
		}
	}

	// input interface - columnKey and indexKey
	inValue := reflect.ValueOf(input)
	inPrtEValue, inType := indirectSimple(inValue)
	inElemKind, inKeyType, inElemType := columnElemType(inPrtEValue, inType, columnKey, indexKey)
	if inElemKind == reflect.Struct && ((columnKey != "" && !isFirstLetterUp(columnKey)) ||
		(indexKey != "" && !isFirstLetterUp(indexKey))) {
		panic("columnKey or indexKey must be public field")
	}

	// no relation with pkg_path
	if inKeyType != dKeyType {
		panic("'dest' key type does not consist with 'input[indexKey]' type")
	}

	// no relation with pkg_path
	if inElemType != dElemType {
		panic("'dest' elem type does not consist with 'input[columnKey]' type")
	}

	// data operation
	dValueElem := ptrToElem(dValue)
	var tempKey, tempColumn reflect.Value
	switch inValue.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < inValue.Len(); i++ {
			tempColumn = inValue.Index(i)
			if inElemKind == reflect.Struct {
				if indexKey != "" {
					tempKey = tempColumn.FieldByName(indexKey)
				}

				if columnKey != "" {
					tempColumn = tempColumn.FieldByName(columnKey)
				}
			} else {
				if indexKey != "" {
					tempKey = tempColumn.MapIndex(reflect.ValueOf(indexKey))
				}

				if columnKey != "" {
					tempColumn = tempColumn.MapIndex(reflect.ValueOf(columnKey))
				}
			}

			dValueElem = columnSetValue(dType, dValueElem, tempColumn, tempKey)
		}
	case reflect.Map:
		for _, k := range inValue.MapKeys() {
			tempColumn = inValue.MapIndex(k)
			if inElemKind == reflect.Struct {
				if indexKey != "" {
					tempKey = tempColumn.FieldByName(indexKey)
				}

				if columnKey != "" {
					tempColumn = tempColumn.FieldByName(columnKey)
				}
			} else {
				if indexKey != "" {
					tempKey = tempColumn.MapIndex(reflect.ValueOf(indexKey))
				}

				if columnKey != "" {
					tempColumn = tempColumn.MapIndex(reflect.ValueOf(columnKey))
				}
			}

			dValueElem = columnSetValue(dType, dValueElem, tempColumn, tempKey)
		}
	}

	// set elem data
	dValue.Elem().Set(dValueElem)
}

// columnSetValue set column elem data
func columnSetValue(rKind reflect.Kind, rv, column, key reflect.Value) reflect.Value {
	switch rKind {
	case reflect.Slice, reflect.Array:
		rv = reflect.Append(rv, column)
	case reflect.Map:
		if rv.IsNil() {
			panic("columnSetValue: reflect.Value is nil")
		}
		rv.SetMapIndex(key, column)
	}

	return rv
}

// columnElemType analysis column type and index type
func columnElemType(rv reflect.Value, rt reflect.Type, columnKey, indexKey string) (reflect.Kind, string, string) {
	var err error
	var elemKind reflect.Kind
	var keyType, elemType string
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		// indexKey operation
		if indexKey == "" {
			keyType = reflect.Int.String()
		} else {
			elemKind, keyType, err = fieldElemType(rt.Elem(), indexKey)
			if err != nil {
				panic(err.Error())
			}
		}

		// columnKey operation
		if columnKey == "" {
			elemType = rt.Elem().String()
		} else {
			elemKind, elemType, err = fieldElemType(rt.Elem(), columnKey)
			if err != nil {
				panic(err.Error())
			}
		}
	default:
		panic("haystack: v type must be slice, array or map")
	}

	return elemKind, keyType, elemType
}

// fieldElemType analysis field type
func fieldElemType(elemRt reflect.Type, key string) (reflect.Kind, string, error) {
	var err error
	var elemType string
	elemKind := elemRt.Kind()
	switch elemKind {
	case reflect.Struct:
		field, isExist := elemRt.FieldByName(key)
		if !isExist {
			err = errors.New(fmt.Sprintf("input map has no column[%s]", key))
			return elemKind, elemType, err
		}
		elemType = field.Type.String()
	case reflect.Map:
		elemType = elemRt.Elem().String()
	default:
		panic("haystack: elemRt type must be map or struct")
	}

	return elemKind, elemType, err
}

// set operation type
type setOps bool

const (
	// set operation
	instOps setOps = true
	diffOps setOps = false

	// constant data
	uniqueOps bool = false
)

// Diff array Diff to d
// php array_diff
func Diff(d interface{}, dArr ...interface{}) {
	// d interface
	dValue := reflect.ValueOf(d)
	if dValue.Kind() != reflect.Ptr || dValue.IsNil() {
		panic(fmt.Sprintf("haystack: d type[%T] error", reflect.TypeOf(dValue)))
	}

	// array diff
	diffValue := operateSetValue(dValue, diffOps, dArr...)

	// set elem data
	dValue.Elem().Set(diffValue)
}

// Intersect array Intersect to d
// php array_intersect
func Intersect(d interface{}, dArr ...interface{}) {
	// d interface indirectForArr
	dValue := reflect.ValueOf(d)
	if dValue.Kind() != reflect.Ptr || dValue.IsNil() {
		panic(fmt.Sprintf("haystack: d type[%T] error", reflect.TypeOf(dValue)))
	}

	// array intersect
	instValue := operateSetValue(dValue, instOps, dArr...)

	// set elem data
	dValue.Elem().Set(instValue)
}

// operateSetValue operation set_value
func operateSetValue(dv reflect.Value, ops setOps, dArr ...interface{}) reflect.Value {
	// check dValue
	if dv.Kind() == reflect.Ptr {
		dv = dv.Elem()
	}

	// operation set
	var newValue reflect.Value
	dType, _ := indirect(dv)
	indirectStr := reflect.Indirect(dv)
	for _, arr := range dArr {
		// type compare
		arrValue := reflect.ValueOf(arr)
		if arrValue.Kind() == reflect.Ptr && arrValue.IsNil() {
			continue
		}

		// new data
		switch dType {
		case reflect.Slice, reflect.Array:
			newValue = reflect.MakeSlice(indirectStr.Type(), 0, dv.Len())
			for i := 0; i < dv.Len(); i++ {
				if inByRValue(dv.Index(i), arrValue) == bool(ops) {
					newValue = reflect.Append(newValue, dv.Index(i))
				}
			}
		case reflect.Map:
			newValue = reflect.MakeMap(indirectStr.Type())
			for _, k := range dv.MapKeys() {
				if inByRValue(dv.MapIndex(k), arrValue) == bool(ops) {
					newValue.SetMapIndex(k, dv.MapIndex(k))
				}
			}
		}

		// set elem data
		dv.Set(newValue)
	}

	return dv
}

// In check d is in 'arr'
// php in_array
func In(d interface{}, arr interface{}) bool {
	dValue := reflect.ValueOf(d)
	arrValue := reflect.ValueOf(arr)

	return inByRValue(dValue, arrValue)
}

// inByRValue in data by reflect value
func inByRValue(dV reflect.Value, arrV reflect.Value) bool {
	dVKind, dvType := indirect(dV)
	arrType, arrElemType, _ := indirectForArr(arrV)
	if dvType != arrElemType {
		// d does not consist with arr elem
		return false
	}

	isExist := false
	switch dVKind {
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Struct:
		isExist = inDeepEqual(dV, arrV, arrType)
	default:
		isExist = inEqual(dV, arrV, arrType)
	}

	return isExist
}

// inEqual use array simple data equal
func inEqual(dV reflect.Value, arrV reflect.Value, arrT reflect.Kind) bool {
	isExist := false
	dV = ptrToElem(dV)     // check ptr
	arrV = ptrToElem(arrV) // check ptr
	switch arrT {
	case reflect.Slice, reflect.Array:
		for i := 0; i < arrV.Len(); i++ {
			if isExist = dV.Interface() == arrV.Index(i).Interface(); isExist {
				break
			}
		}
	case reflect.Map:
		for _, k := range arrV.MapKeys() {
			if isExist = dV.Interface() == arrV.MapIndex(k).Interface(); isExist {
				break
			}
		}
	default:
		panic("haystack: arrV type must be slice, array or map")
	}

	return isExist
}

// inDeepEqual use array complex data equal
func inDeepEqual(dV reflect.Value, arrV reflect.Value, arrT reflect.Kind) bool {
	isExist := false
	dV = ptrToElem(dV)     // check ptr
	arrV = ptrToElem(arrV) // check ptr
	switch arrT {
	case reflect.Slice, reflect.Array:
		for i := 0; i < arrV.Len(); i++ {
			if isExist = reflect.DeepEqual(dV.Interface(), arrV.Index(i).Interface()); isExist {
				break
			}
		}
	case reflect.Map:
		for _, k := range arrV.MapKeys() {
			if isExist = reflect.DeepEqual(dV.Interface(), arrV.MapIndex(k).Interface()); isExist {
				break
			}
		}
	default:
		panic("haystack: d type must be slice, array or map")
	}

	return isExist
}

// Merge array merge to d
// php array_merge
func Merge(d interface{}, dArr ...interface{}) {
	// d interface indirectForArr
	dValue := reflect.ValueOf(d)
	if dValue.Kind() != reflect.Ptr || dValue.IsNil() {
		panic(fmt.Sprintf("haystack: d type[%T] error", reflect.TypeOf(dValue)))
	}

	// array merge
	dValueElem := ptrToElem(dValue)
	dType, dElemType, dKeyType := indirectForArr(dValue)
	for _, arr := range dArr {
		// type compare
		arrValue := reflect.ValueOf(arr)
		if dValue.Kind() != reflect.Ptr || dValue.IsNil() {
			continue
		}
		arrType, arrElemType, arrKeyType := indirectForArr(arrValue)
		if arrType != dType {
			panic("'dArr' type does not consist with 'd' type")
		}
		if arrElemType != dElemType {
			panic("'dArr' elem type does not consist with 'd' elem type")
		}
		if dKeyType != arrKeyType {
			panic("'dArr' key type does not consist with 'd' key type")
		}

		// data merge
		arrValue = ptrToElem(arrValue)
		switch dType {
		case reflect.Slice, reflect.Array:
			for i := 0; i < arrValue.Len(); i++ {
				dValueElem = reflect.Append(dValueElem, arrValue.Index(i))
			}
		case reflect.Map:
			for _, k := range arrValue.MapKeys() {
				dValueElem.SetMapIndex(k, arrValue.MapIndex(k))
			}
		default:
			panic("haystack: d type must be slice, array or map")
		}
	}

	// set elem data
	dValue.Elem().Set(dValueElem)
}

// Values get array values
// php array_values
func Values(d interface{}, values interface{}) {
	// values interface indirectForArr
	vValue := reflect.ValueOf(values)
	if vValue.Kind() != reflect.Ptr || vValue.IsNil() {
		panic(fmt.Sprintf("haystack: values type[%T] error", reflect.TypeOf(vValue)))
	}
	vType, vElemType, _ := indirectForArr(vValue)
	if vType != reflect.Slice && vType != reflect.Array {
		panic("haystack: values type must be slice or array")
	}

	// d interface indirectForArr
	dValue := reflect.ValueOf(d)
	dType, dElemType, _ := indirectForArr(dValue)
	if vElemType != dElemType {
		panic("'d' key type does not consist with 'keys' elem type")
	}

	vValueElem := ptrToElem(vValue)
	switch dType {
	case reflect.Slice, reflect.Array:
		for i := 0; i < dValue.Len(); i++ {
			vValueElem = reflect.Append(vValueElem, dValue.Index(i))
		}
	case reflect.Map:
		for _, k := range dValue.MapKeys() {
			vValueElem = reflect.Append(vValueElem, dValue.MapIndex(k))
		}
	default:
		panic("haystack: d type must be slice, array or map")
	}

	// set elem data
	vValue.Elem().Set(vValueElem)
}

// Keys get array keys
// php array_keys
func Keys(d interface{}, keys interface{}) {
	// keys interface indirectForArr
	keysValue := reflect.ValueOf(keys)
	if keysValue.Kind() != reflect.Ptr || keysValue.IsNil() {
		panic(fmt.Sprintf("haystack: keys type[%T] error", reflect.TypeOf(keysValue)))
	}
	keysType, keysElemType, _ := indirectForArr(keysValue)
	if keysType != reflect.Slice && keysType != reflect.Array {
		panic("haystack: keys type must be slice or array")
	}

	// d interface indirectForArr
	dValue := reflect.ValueOf(d)
	dType, _, dKeyType := indirectForArr(dValue)
	if keysElemType != dKeyType {
		panic("'keys' key type does not consist with 'd' elem type")
	}

	keysElem := ptrToElem(keysValue)
	switch dType {
	case reflect.Slice, reflect.Array:
		for i := 0; i < dValue.Len(); i++ {
			keysElem = reflect.Append(keysElem, reflect.ValueOf(i))
		}
	case reflect.Map:
		for _, k := range dValue.MapKeys() {
			keysElem = reflect.Append(keysElem, k)
		}
	default:
		panic("haystack: d type must be slice, array or map")
	}

	// set elem data
	keysValue.Elem().Set(keysElem)
}

// Unique array data de-duplication
// php array_unique
func Unique(d interface{}) {
	// d interface indirectForArr
	dv := reflect.ValueOf(d)
	if dv.Kind() != reflect.Ptr || dv.IsNil() {
		panic(fmt.Sprintf("haystack: d type[%T] error", reflect.TypeOf(dv)))
	}
	dvElem := ptrToElem(dv)

	// unique operation
	var newValue reflect.Value
	dType, _ := indirect(dv)
	indirectStr := reflect.Indirect(dv)
	switch dType {
	case reflect.Slice, reflect.Array:
		newValue = reflect.MakeSlice(indirectStr.Type(), 0, dvElem.Len())
		for i := 0; i < dvElem.Len(); i++ {
			if inByRValue(dvElem.Index(i), newValue) == uniqueOps {
				newValue = reflect.Append(newValue, dvElem.Index(i))
			}
		}
	case reflect.Map:
		newValue = reflect.MakeMap(indirectStr.Type())
		for _, k := range dvElem.MapKeys() {
			if inByRValue(dvElem.MapIndex(k), newValue) == uniqueOps {
				newValue.SetMapIndex(k, dvElem.MapIndex(k))
			}
		}
	}

	// set elem data
	dv.Elem().Set(newValue)
}

// indirect get interface kind and type
func indirect(rv reflect.Value) (reflect.Kind, string) {
	// get reflect value and type
	rvValue, rvType := indirectSimple(rv)

	return rvValue.Kind(), rvType.String()
}

// indirectForArr get slice, array or map type
func indirectForArr(rv reflect.Value) (reflect.Kind, string, string) {
	// get reflect value and type
	rvValue, rvType := indirectSimple(rv)

	vType := rvValue.Kind()
	var vKeyType, vElemType string
	switch vType {
	case reflect.Slice, reflect.Array:
		vKeyType = reflect.Int.String()
		vElemType = rvType.Elem().String()
	case reflect.Map:
		vKeyType = rvType.Key().String()
		vElemType = rvType.Elem().String()
	default:
		panic("haystack: v type must be slice, array or map")
	}

	return vType, vElemType, vKeyType
}

// indirectSimple indirect_simply
func indirectSimple(rv reflect.Value) (reflect.Value, reflect.Type) {
	// check valid
	if !rv.IsValid() {
		panic("indirectSimple: reflect.Value is nil")
	}

	// get reflect value and type
	var rvValue reflect.Value
	var rvType reflect.Type
	switch rv.Kind() {
	case reflect.Ptr:
		rvValue = rv.Elem()
		rvType = rv.Type().Elem()
	default:
		rvValue = rv
		rvType = rv.Type()
	}

	return rvValue, rvType
}

// ptrToElem Ptr to elem
func ptrToElem(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v
}

// isFirstLetterUp the first letter is up
func isFirstLetterUp(s string) bool {
	regObj, _ := regexp.Compile("^[A-Z].*")
	return regObj.MatchString(s)
}
