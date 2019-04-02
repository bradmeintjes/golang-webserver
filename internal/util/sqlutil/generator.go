package sqlutil

import (
	"fmt"
	"reflect"
	"strings"
)

const sqlxTag = "db"

// GenerateInsert generates a named insert statement from a sqlx tagged struct
func GenerateInsert(table string, datastruct interface{}, excl []string) string {
	fields := fieldsFromStruct(datastruct, excl)

	var builder strings.Builder
	builder.WriteString("insert into ")
	builder.WriteString(table)
	builder.WriteString(" (")

	lastFieldIdx := len(fields) - 1
	for i, fld := range fields {
		builder.WriteString(fld)
		if i != lastFieldIdx {
			builder.WriteString(",")
		}
	}

	builder.WriteString(") values (")
	for i := range fields {
		builder.WriteString(fmt.Sprintf(":%s", fields[i]))
		if i != lastFieldIdx {
			builder.WriteString(",")
		}
	}
	builder.WriteString(")")
	return builder.String()
}

// GenerateUpdate generates a named update statement from a sqlx tagged struct
func GenerateUpdate(table string, datastruct interface{}, excl []string) string {
	fields := fieldsFromStruct(datastruct, excl)

	var builder strings.Builder
	builder.WriteString("update ")
	builder.WriteString(table)
	builder.WriteString(" set ")
	for i, fld := range fields {
		builder.WriteString(fmt.Sprintf("%s=:%s", fld, fld))
		if i != len(fields)-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

// GenerateUpsert generates a named upsert statement from a sqlx tagged struct
func GenerateUpsert(table, idField string, datastruct interface{}, excl []string) string {
	var builder strings.Builder
	builder.WriteString(GenerateInsert(table, datastruct, excl))
	builder.WriteString(" on conflict(")
	builder.WriteString(idField)
	builder.WriteString(") do ")
	builder.WriteString(GenerateUpdate("", datastruct, excl))
	return builder.String()
}

// GenerateSelect generates a select statement from a sqlx tagged struct
func GenerateSelect(table string, datastruct interface{}, excl []string) string {
	fields := fieldsFromStruct(datastruct, excl)

	var builder strings.Builder
	builder.WriteString("select ")

	lastFieldIdx := len(fields) - 1
	for i, fld := range fields {
		builder.WriteString(fld)
		if i != lastFieldIdx {
			builder.WriteString(",")
		}
	}

	builder.WriteString(" from ")
	builder.WriteString(table)

	return builder.String()
}

func fieldsFromStruct(datastruct interface{}, excl []string) []string {
	st := reflect.TypeOf(datastruct)
	numFields := st.NumField()
	fields := make([]string, 0, numFields)

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		name := field.Name

		if val, ok := field.Tag.Lookup(sqlxTag); ok {
			if val != "" {
				name = val
			}
		}

		if !contains(excl, name) {
			fields = append(fields, name)
		}
	}
	return fields
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
