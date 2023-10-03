package services

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

func IsDuplicatedKeyError(err error) bool {
	const DuplicatedErrorCode = 1062
	var mysqlError *mysql.MySQLError

	if errors.As(err, &mysqlError) {
		return mysqlError.Number == DuplicatedErrorCode
	}

	return false
}
