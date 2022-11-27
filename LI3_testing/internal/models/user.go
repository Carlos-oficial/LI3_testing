package models

import (
	"LI3_testing/lib"
	"fmt"
)

type User struct {
	Username            string `gorm:"primaryKey"`
	Name                string
	Gender              string
	BirthDate           string
	AccountCreationDate string
	PayMethod           string
	AccountStatus       string
}

func (*User) FromLine(line string) (OBJ, error) {
	user := &User{}
	fields := lib.SplitLine(line, ";")
	if len(fields) < 7 {
		return nil, fmt.Errorf("insuficient fields for user initialization")
	}
	user.Username = fields[0]
	user.Name = fields[1]
	user.Gender = fields[2]
	user.BirthDate = fields[3]
	user.AccountCreationDate = fields[4]
	user.PayMethod = fields[5]
	user.AccountStatus = fields[6]

	return user, nil
}
