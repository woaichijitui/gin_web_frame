package ctype

import (
	"database/sql/driver"
	"strings"
)

type StrArray []string

func (s *StrArray) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	if string(bytes) == "" {
		*s = []string{}
		return nil
	}
	*s = strings.Split(string(bytes), "\n")
	return nil

}

func (a *StrArray) Value() (driver.Value, error) {
	//	将数值转化为值
	return strings.Join(*a, "\n"), nil
}
