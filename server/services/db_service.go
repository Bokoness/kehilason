package services

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

func IsDuplicatedKeyError(err error) bool {
	var mysqlError *mysql.MySQLError
	if errors.As(err, &mysqlError) {
		return mysqlError.Number == 1062
	}
	return false
}
