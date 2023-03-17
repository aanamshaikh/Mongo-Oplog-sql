package parser

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

const (
	insert = "i"
	update = "u"
	delete = "d"
)

func CreateSql(tableName, operation string, values, oValue map[string]interface{}) string {
	var sqlStmt string
	switch operation {
	case insert:
		sqlStmt = generateInsertStmt(tableName, values)
	case update:
		// sqlStmt = generateUpdateStmt(tableName, values, oValue)
	case delete:
		sqlStmt = generateDeleteStmt(tableName, values)

	default:
		fmt.Println("Invalid operation type")
	}
	return sqlStmt
}

func generateInsertStmt(tableName string, values map[string]interface{}) string {
	var columns []string
	var valuesList []string

	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		columns = append(columns, k)
		v := reflect.ValueOf(values[k])
		if v.Kind() == reflect.String {
			valuesList = append(valuesList, fmt.Sprintf("'%v'", values[k]))
		} else {
			valuesList = append(valuesList, fmt.Sprintf("%v", values[k]))
		}
	}

	sqlStmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ","), strings.Join(valuesList, ","))
	return sqlStmt
}

func generateDeleteStmt(tableName string, values map[string]interface{}) string {
	var sqlStmt string
	for key, value := range values {
		//create a function to return a '' or val
		v := reflect.ValueOf(value)
		if v.Kind() == reflect.String {
			sqlStmt = fmt.Sprintf("DELETE FROM %s WHERE %v='%v'", tableName, key, value)
		} else {
			sqlStmt = fmt.Sprintf("DELETE FROM %s WHERE %v=%v", tableName, key, value)
		}
		break
	}
	return sqlStmt
}
