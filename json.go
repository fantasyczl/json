package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
)

const IndentLevel = "    "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		var ret interface{}
		if err := json.Unmarshal([]byte(text), &ret); err != nil {
			log.Fatalf("unmarshal json failed, %s", err)
		}

		JSONPretty(ret, "", false, false)
	}
}

func JSONPretty(ret interface{}, indent string, fromMap bool, needComma bool) {
	if ret == nil {
		fmt.Printf("null\n")
		return
	}

	t := reflect.TypeOf(ret)
	switch t.Kind() {
	case reflect.Map:
		displayMap(ret.(map[string]interface{}), indent, fromMap, needComma)
	case reflect.Slice:
		displaySlice(ret.([]interface{}), indent, fromMap, needComma)
	case reflect.String:
		displayString(ret.(string), indent, fromMap, needComma)
	case reflect.Float32:
		displayFloat64(float64(ret.(float32)), indent, fromMap, needComma)
	case reflect.Float64:
		displayFloat64(ret.(float64), indent, fromMap, needComma)
	case reflect.Int:
		displayInt(ret.(int), indent, fromMap, needComma)
	case reflect.Bool:
		displayBool(ret.(bool), indent, fromMap, needComma)
	default:
		fmt.Printf("ERROR: invalid type, %T\n", ret)
		os.Exit(1)
	}
}

func displayMap(ret map[string]interface{}, indent string, fromMap bool, needComma bool) {
	newIndent := indent + IndentLevel
	var headIndent string
	if !fromMap {
		headIndent = indent
	}

	fmt.Printf("%s{\n", headIndent)

	keys := make([]string, 0, len(ret))
	for k := range ret {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		v := ret[k]
		fmt.Printf("%s%q: ", newIndent, k)
		JSONPretty(v, newIndent, true, i != len(keys)-1)
	}

	var tail string
	if needComma {
		tail = ","
	}
	fmt.Printf("%s}%s\n", indent, tail)
}

func displaySlice(ret []interface{}, indent string, needIndent bool, needComma bool) {
	newIndent := indent + IndentLevel

	fmt.Printf("[\n")
	for i, v := range ret {
		JSONPretty(v, newIndent, false, i != len(ret)-1)
	}
	var tail string
	if needComma {
		tail = ","
	}
	fmt.Printf("%s]%s\n", indent, tail)
}

func displayString(s string, indent string, fromMap bool, needComma bool) {
	var headIndent, tail string
	if !fromMap {
		headIndent = indent
	}
	if needComma {
		tail = ","
	}
	fmt.Printf("%s%q%s\n", headIndent, s, tail)
}

func displayFloat64(f float64, indent string, fromMap bool, needComma bool) {
	var headIndent, tail string
	if !fromMap {
		headIndent = indent
	}
	if needComma {
		tail = ","
	}

	floatIsInt := func(ff float64) bool {
		return float64(int(ff)) == ff
	}

	if floatIsInt(f) {
		fmt.Printf("%s%d%s\n", headIndent, int(f), tail)
	} else {
		fmt.Printf("%s%f%s\n", headIndent, f, tail)
	}
}

func displayInt(f int, indent string, fromMap bool, needComma bool) {
	var headIndent, tail string
	if !fromMap {
		headIndent = indent
	}
	if needComma {
		tail = ","
	}
	fmt.Printf("%s%d%s\n", headIndent, f, tail)
}

func displayBool(f bool, indent string, fromMap bool, needComma bool) {
	var headIndent, tail string
	if !fromMap {
		headIndent = indent
	}
	if needComma {
		tail = ","
	}
	fmt.Printf("%s%v%s\n", headIndent, f, tail)
}
