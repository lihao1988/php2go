package array

import (
	"fmt"
	"testing"
)

type keyValue struct {
	key   string
	value string
}

type PubKeyValue struct {
	Key   string
	Value string
}

// test array column
func TestArrayColumn(t *testing.T) {
	/*dataMap := []PubKeyValue{
		{Key: "a", Value: "A"},
		{Key: "b", Value: "B"},
	}*/

	/*dataMap := map[string]map[string]string{
		"a": {"key": "a", "value": "A"},
		"b": {"key": "b", "value": "B"},
	}

	dest := map[string]string{}*/

	dataMap := []map[string]string{
		{"key": "a", "value": "A"},
		{"key": "b", "value": "B"},
	}

	dest := map[string]map[string]string{}

	Column(&dest, dataMap, "", "value")
	fmt.Println(dest)
}

// test array diff
func TestArrayDiff(t *testing.T) {
	dataMap := map[string]keyValue{
		"a": {key: "a", value: "A"},
		"b": {key: "b", value: "B"},
	}

	dataMap1 := []keyValue{
		{key: "c", value: "C"},
		{key: "b", value: "B"},
	}

	dataMap2 := []string{
		"a", "b",
	}

	Diff(&dataMap, dataMap1, dataMap2)
	fmt.Println("dataMap", dataMap)
}

// test array intersect
func TestArrayIntersect(t *testing.T) {
	dataMap := []keyValue{
		{key: "c", value: "C"},
		{key: "b", value: "B"},
	}

	dataMap1 := map[string]keyValue{
		"a": {key: "a", value: "A"},
		"b": {key: "b", value: "B"},
		"c": {key: "c", value: "C"},
	}

	dataMap2 := []string{
		"a", "b",
	}

	Intersect(&dataMap, dataMap1, &dataMap2)
	fmt.Println("dataMap", dataMap)
}

// test in array
func TestInArray(t *testing.T) {
	/*data := map[int]string{1: "1"}
	dataMap := map[string]map[int]string{
		"a": {1: "1"},
		"b": {2: "2"},
	}*/

	/*dataMap := []keyValue{
		{key: "a", value: "A"},
		{key: "b", value: "B"},
	}
	data := keyValue{key: "b", value: "A"}*/

	dataMap := []string{"a", "b", "c"}
	data := "a"

	isExist := In(&data, dataMap)
	fmt.Println(isExist)
}

// test array array_merge
func TestArrayMergeForSlice(t *testing.T) {
	dataMap := []map[int]string{
		{1: "1"},
		{2: "2"},
	}

	oDataMap1 := []map[int]string{
		{3: "3"},
		{4: "4"},
	}

	Merge(&dataMap, &oDataMap1)
	fmt.Println(dataMap)
}

// test map array_merge
func TestArrayMergeForMap(t *testing.T) {
	dataMap := map[string]map[int]string{
		"a": {1: "1"},
		"b": {2: "2"},
	}

	oDataMap1 := map[string]map[int]string{
		"c": {3: "3"},
		"b": {4: "4"},
	}

	oDataMap2 := map[string]map[int]string{}

	oDataMap3 := map[string]map[int]string{
		"e": {5: "5"},
		"f": {6: "6"},
	}

	Merge(&dataMap, &oDataMap1, oDataMap2, oDataMap3)
	fmt.Println(dataMap)
}

// test slice array_values
func TestArrayValuesForSlice(t *testing.T) {
	dataMap := []map[int]string{
		{1: "1", 2: "2"},
		{1: "1", 2: "2"},
	}
	var valueArr []map[int]string
	Values(dataMap, &valueArr)
	fmt.Println(valueArr)
}

// test Map array_values
func TestArrayValuesForMap(t *testing.T) {
	/*dataMap := map[string]keyValue{
		"a":{key: "a", value: "A"},
		"b":{key: "b", value: "B"},
	}*/
	dataMap := map[string]map[int]string{
		"a": {1: "1", 2: "2"},
		"b": {1: "1", 2: "2"},
	}
	var valueArr []map[int]string
	Values(dataMap, &valueArr)
	fmt.Println(valueArr)

	/*fmt.Println(vValue.Elem().Type().Elem().Name(), dValue.Type().Elem().Name())
	fmt.Println(vValue.Elem().Type().Elem().PkgPath(), dValue.Type().Elem().PkgPath())*/
}

// test array_keys
func TestArrayKeysForMap(t *testing.T) {
	dataMap := map[string]keyValue{
		"a": {key: "a", value: "A"},
		"b": {key: "b", value: "B"},
	}
	var keysArr []string
	Keys(dataMap, &keysArr)
	fmt.Println(keysArr)
}

// test array_unique
func TestArrayUnique(t *testing.T) {
	/*dataMap := map[string]keyValue{
		"a": {key: "a", value: "A"},
		"b": {key: "b", value: "B"},
		"c": {key: "a", value: "A"},
	}*/

	dataMap := []keyValue{
		{key: "a", value: "A"},
		{key: "b", value: "B"},
		{key: "a", value: "A"},
	}

	Unique(&dataMap)
	fmt.Println(dataMap)
}
