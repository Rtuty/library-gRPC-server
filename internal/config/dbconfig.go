package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type dataSource struct{ Host, Port, User, Passwd, Dbname string }

// GetConnection иннициализирует переменные окружения и возвращает данные для подключения MySql
func GetConnection() (cf dataSource, err error) {
	var pstgCon dataSource
	res := &pstgCon

	envVars := []string{"HOST", "PORT", "USER", "PASSWD", "DBNAME"}

	for _, v := range envVars { // если переменная окружения иннициализирована и корректно записана в .env
		value := os.Getenv(v)
		if value == "" {
			return dataSource{}, fmt.Errorf("invalid environment variable %s", v)
		} else { // заполняется и возращается структура pstgCon
			field := reflect.ValueOf(res).Elem().FieldByNameFunc(
				func(fieldName string) bool {
					return strings.EqualFold(fieldName, v)
				})
			if field.IsValid() {
				field.SetString(value)
			}
		}
	}

	return pstgCon, nil
}
