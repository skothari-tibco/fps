package reshape

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Removing ToArray from core/coerce. Reason being with 1.0.0
// updated function breaks the code. The current version
// here is from 0.9.4
func toArray(val interface{}) ([]interface{}, error) {
	switch t := val.(type) {
	case []interface{}:
		return t, nil

	case []map[string]interface{}:
		var a []interface{}
		for _, v := range t {
			a = append(a, v)
		}
		return a, nil
	case string:
		a := make([]interface{}, 0)
		if t != "" {
			err := json.Unmarshal([]byte(t), &a)
			if err != nil {
				a = append(a, t)
			}
		}
		return a, nil
	case nil:
		return nil, nil
	default:
		s := reflect.ValueOf(val)
		if s.Kind() == reflect.Slice {
			a := make([]interface{}, s.Len())

			for i := 0; i < s.Len(); i++ {
				a[i] = s.Index(i).Interface()
			}
			return a, nil
		}
		return nil, fmt.Errorf("unable to coerce %#v to []interface{}", val)
	}
}
