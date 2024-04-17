package builder

import (
	"fmt"
	"reflect"
	"strings"
)

// RawFieldNames converts golang struct field into slice string.
//func RawFieldNames(in interface{}, postgresSql ...bool) []string {
//	return builder.RawFieldNames(in, postgresSql...)
//}

func PostgreSqlJoinDoUpdatesetExclude(uqs []string, elems []string, exclude []string) string {
	if len(uqs) == 0 || len(elems) == 0 {
		return ""
	}

	b := new(strings.Builder)
	b.WriteString("ON CONFLICT (")
	b.WriteString(strings.Join(uqs, ","))
	b.WriteString(") DO UPDATE SET ")
	for _, e := range elems {
		var pass bool
		for _, ex := range exclude {
			if e == ex {
				pass = true
				break
			}
		}
		if !pass {
			b.WriteString(fmt.Sprintf("%s = excluded.%s, ", e, e))
		}
	}

	return b.String()[0 : b.Len()-2]
}

func PostgreSqlJoinDOUPDATESET(uqs []string, elems []string) string {
	if len(uqs) == 0 || len(elems) == 0 {
		return ""
	}

	b := new(strings.Builder)
	b.WriteString("ON CONFLICT (")
	b.WriteString(strings.Join(uqs, ","))
	b.WriteString(") DO UPDATE SET ")
	for _, e := range elems {
		b.WriteString(fmt.Sprintf("%s = excluded.%s, ", e, e))
	}

	return b.String()[0 : b.Len()-2]
}

const dbTag = "db"

// FieldColsAndValues converts golang struct field into slice string removes given strs from strings
func FieldColsAndValues(in interface{}, strs ...string) ([]string, []interface{}) {
	out := make([]string, 0)
	values := make([]interface{}, 0)
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	//var pg bool = true
	//if len(postgresSql) > 0 {
	//	pg = postgresSql[0]
	//}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		tagv := fi.Tag.Get(dbTag)
		switch tagv {
		case "-":
			continue
		case "":
			var passing bool
			for _, vv := range strs {
				if vv == tagv {
					passing = true
					break
				}
			}

			if !passing {
				out = append(out, fi.Name)
				values = append(values, v.Field(i).Interface())
			}

		default:

			if strings.Contains(tagv, ",") {
				tagv = strings.TrimSpace(strings.Split(tagv, ",")[0])
			}
			if len(tagv) == 0 {
				tagv = fi.Name
			}
			var passing bool
			for _, vv := range strs {
				if vv == tagv {
					passing = true
					break
				}
			}

			if !passing {
				out = append(out, tagv)
				values = append(values, v.Field(i).Interface())
			}
		}
	}

	return out, values
}

// FieldsToSetMap converts golang struct field into slice string removes given strs from strings
func FieldsToSetMap(in interface{}, strs ...string) map[string]interface{} {
	res := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		tagv := fi.Tag.Get(dbTag)
		switch tagv {
		case "-":
			continue
		case "":
			var passing bool
			for _, vv := range strs {
				if vv == tagv {
					passing = true
					break
				}
			}
			if !passing {
				//out = append(out, fi.Name)
				//values = append(values, v.Field(i).Interface())
				res[fi.Name] = v.Field(i).Interface()
			}

		default:
			// get tag name with the tag opton, e.g.:
			// `db:"id"`
			// `db:"id,type=char,length=16"`
			// `db:",type=char,length=16"`
			if strings.Contains(tagv, ",") {
				tagv = strings.TrimSpace(strings.Split(tagv, ",")[0])
			}
			if len(tagv) == 0 {
				tagv = fi.Name
			}
			var passing bool
			for _, vv := range strs {
				if vv == tagv {
					passing = true
					break
				}
			}

			if !passing {
				res[tagv] = v.Field(i).Interface()
			}
		}
	}

	return res
}
