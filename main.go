package merger

import (
	"errors"
	"reflect"
)

func Merge(a map[string]interface{}, b map[string]interface{}) (map[string]interface{}, error) {
	mType := reflect.TypeOf(a)
	for k, v := range b {
		_, p := a[k]
		if p && reflect.TypeOf(v) != reflect.TypeOf(a[k]) {
			return nil, errors.New("failed to merge")
		}
		if p == false {
			a[k] = v
		}
		if reflect.TypeOf(v) == mType && reflect.TypeOf(a[k]) == mType {
			n, e := merge(a[k].(map[string]interface{}), v.(map[string]interface{}))
			if e != nil {
				return nil, e
			}
			a[k] = n
		}

	}
	return a, nil
}
