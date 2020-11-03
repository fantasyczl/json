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

		JSONPretty(ret, &prettyInfo{})
	}
}

// control output format
type prettyInfo struct {
	indent    string
	fromMap   bool
	needComma bool
}

func (p *prettyInfo) getHeadIndent() string {
	var headIndent string
	if !p.fromMap {
		headIndent = p.indent
	}

	return headIndent
}

func (p *prettyInfo) getTail() string {
	var tail string
	if p.needComma {
		tail = ","
	}
	return tail
}

func JSONPretty(ret interface{}, pretty *prettyInfo) {
	if ret == nil {
		fmt.Printf("null\n")
		return
	}

	t := reflect.TypeOf(ret)
	switch t.Kind() {
	case reflect.Map:
		displayMap(ret.(map[string]interface{}), pretty)
	case reflect.Slice:
		displaySlice(ret.([]interface{}), pretty)
	case reflect.String:
		displayString(ret.(string), pretty)
	case reflect.Float32:
		displayFloat64(float64(ret.(float32)), pretty)
	case reflect.Float64:
		displayFloat64(ret.(float64), pretty)
	case reflect.Int:
		displayInt(ret.(int), pretty)
	case reflect.Bool:
		displayBool(ret.(bool), pretty)
	default:
		fmt.Printf("ERROR: invalid type, %T\n", ret)
		os.Exit(1)
	}
}

func displayMap(ret map[string]interface{}, pretty *prettyInfo) {
	newIndent := pretty.indent + IndentLevel
	fmt.Printf("%s{\n", pretty.getHeadIndent())

	keys := make([]string, 0, len(ret))
	for k := range ret {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		v := ret[k]
		fmt.Printf("%s%q: ", newIndent, k)
		JSONPretty(v, &prettyInfo{indent: newIndent, fromMap: true, needComma: i != len(keys)-1})
	}

	fmt.Printf("%s}%s\n", pretty.indent, pretty.getTail())
}

func displaySlice(ret []interface{}, pretty *prettyInfo) {
	newIndent := pretty.indent + IndentLevel

	fmt.Printf("[\n")
	for i, v := range ret {
		JSONPretty(v, &prettyInfo{indent: newIndent, fromMap: false, needComma: i != len(ret)-1})
	}
	var tail string
	if pretty.needComma {
		tail = ","
	}
	fmt.Printf("%s]%s\n", pretty.indent, tail)
}

func displayString(s string, pretty *prettyInfo) {
	fmt.Printf("%s%q%s\n", pretty.getHeadIndent(), s, pretty.getTail())
}

func displayFloat64(f float64, pretty *prettyInfo) {
	floatIsInt := func(ff float64) bool {
		return float64(int(ff)) == ff
	}

	if floatIsInt(f) {
		fmt.Printf("%s%d%s\n", pretty.getHeadIndent(), int(f), pretty.getTail())
	} else {
		fmt.Printf("%s%f%s\n", pretty.getHeadIndent(), f, pretty.getTail())
	}
}

func displayInt(f int, pretty *prettyInfo) {
	fmt.Printf("%s%d%s\n", pretty.getHeadIndent(), f, pretty.getTail())
}

func displayBool(f bool, pretty *prettyInfo) {
	fmt.Printf("%s%v%s\n", pretty.getHeadIndent(), f, pretty.getTail())
}
