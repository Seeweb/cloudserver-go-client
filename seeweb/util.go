package seeweb

import (
	"fmt"
	"reflect"
	"time"
)

// Compares two structs in similiar way of reflect.DeepEqual, but with Date
// (time.Time) is limited to the fields at the root of the structs. This
// function was introduced because dates aren't being evaluated in same way
// between Mac and Linux.
func equalStructWithDatesFn(a interface{}, b interface{}) bool {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	if va.NumField() != vb.NumField() {
		return false
	}

	for i := 1; i < va.NumField(); i++ {
		if va.Field(i).Type() != vb.Field(i).Type() {
			return false
		}
		valueOfA := va.Field(i).Interface()
		valueOfB := vb.Field(i).Interface()
		if fmt.Sprint(va.Field(i).Type()) == "time.Time" {
			if !equalForDatesFn(valueOfA, valueOfB) {
				return false
			}
		} else {
			if !reflect.DeepEqual(valueOfA, valueOfB) {
				return false
			}
		}
	}
	return true
}

func equalForDatesFn(a interface{}, b interface{}) bool {
	da := a.(time.Time)
	db := b.(time.Time)

	return da.Equal(db)
}
