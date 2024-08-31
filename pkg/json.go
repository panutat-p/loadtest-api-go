package pkg

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func PrintJSON(o any) {
	kind := reflect.TypeOf(o).Kind()
	switch kind {
	case reflect.String:
		s := reflect.ValueOf(o).String()
		b := []byte(s)
		if len(b) == 0 {
			fmt.Println("❌  Cannot Unmarshal empty string to JSON")
			return
		}
		var v any
		err := json.Unmarshal([]byte(s), &v)
		if err != nil {
			fmt.Println("❌  Cannot Unmarshal string to JSON, err:", err)
			return
		}
		b, _ = json.MarshalIndent(v, "", "  ")
		fmt.Println(string(b))
	case reflect.Slice:
		if reflect.TypeOf(o).Elem().Kind() == reflect.Uint8 {
			b := reflect.ValueOf(o).Bytes()
			if len(b) == 0 {
				fmt.Println("❌  Cannot Unmarshal empty []byte to JSON")
				return
			}
			var v any
			err := json.Unmarshal(b, &v)
			if err != nil {
				fmt.Println("❌  Cannot Unmarshal []byte to JSON, err:", err)
				return
			}
			b, _ = json.MarshalIndent(v, "", "  ")
			fmt.Println(string(b))
		} else {
			b, err := json.MarshalIndent(o, "", "  ")
			if err != nil {
				fmt.Println("❌  Cannot Marshal to JSON, err:", err)
				return
			}
			fmt.Println(string(b))
		}
	default:
		b, err := json.MarshalIndent(o, "", "  ")
		if err != nil {
			fmt.Println("❌  Cannot Marshal object to JSON, err:", err)
			return
		}
		fmt.Println(string(b))
	}
}
