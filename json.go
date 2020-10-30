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

		JSONPretty(ret, "", false)
	}
}

func JSONPretty(ret interface{}, indent string, fromMap bool) {
	if ret == nil {
		fmt.Printf("null\n")
		return
	}

	t := reflect.TypeOf(ret)
	switch t.Kind() {
	case reflect.Map:
		displayMap(ret.(map[string]interface{}), indent, fromMap)
	case reflect.Slice:
		displaySlice(ret.([]interface{}), indent, fromMap)
	case reflect.String:
		displayString(ret.(string), indent, fromMap)
	case reflect.Float32:
		displayFloat64(float64(ret.(float32)), indent, fromMap)
	case reflect.Float64:
		displayFloat64(ret.(float64), indent, fromMap)
	case reflect.Int:
		displayInt(ret.(int), indent, fromMap)
	case reflect.Bool:
		displayBool(ret.(bool), indent, fromMap)
	default:
		fmt.Printf("ERROR: invalid type, %T\n", ret)
		os.Exit(1)
	}
}

func displayMap(ret map[string]interface{}, indent string, fromMap bool) {
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

	for _, k := range keys {
		v := ret[k]
		fmt.Printf("%s%q: ", newIndent, k)
		JSONPretty(v, newIndent, true)
	}
	fmt.Printf("%s}\n", indent)
}

func displaySlice(ret []interface{}, indent string, needIndent bool) {
	newIndent := indent + IndentLevel

	fmt.Printf("[\n")
	for _, v := range ret {
		JSONPretty(v, newIndent, false)
	}
	fmt.Printf("%s]\n", indent)
}

func displayString(s string, indent string, fromMap bool) {
	var headIndent string
	if !fromMap {
		headIndent = indent
	}
	fmt.Printf("%s%q\n", headIndent, s)
}

func displayFloat64(f float64, indent string, fromMap bool) {
	var headIndent string
	if !fromMap {
		headIndent = indent
	}

	floatIsInt := func(ff float64) bool {
		return float64(int(ff)) == ff
	}

	if floatIsInt(f) {
		fmt.Printf("%s%d\n", headIndent, int(f))
	} else {
		fmt.Printf("%s%f\n", headIndent, f)
	}
}

func displayInt(f int, indent string, fromMap bool) {
	var headIndent string
	if !fromMap {
		headIndent = indent
	}
	fmt.Printf("%s%d\n", headIndent, f)
}

func displayBool(f bool, indent string, fromMap bool) {
	var headIndent string
	if !fromMap {
		headIndent = indent
	}
	fmt.Printf("%s%v\n", headIndent, f)
}
