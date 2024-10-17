package utils

import (
	"errors"
	"reflect"
)

func MustBePointer(val interface{}, key string) error {
	if reflect.TypeOf(val).Kind() != reflect.Ptr {
		getKey := key
		if key != "" {
			getKey = key + " "
		}
		return errors.New(getKey + "must be pointer")
	}
	return nil
}
