package models

import (
	"LI3_testing/lib"
	"fmt"
	"strconv"
)

type Driver struct {
	ID                  uint
	Name                string
	BirthDate           string
	Gender              string
	CarClass            string
	LicensePlate        string
	City                string
	AccountCreationDate string
	AccountStatus       string
}

func (*Driver) FromLine(line string) (OBJ, error) {
	driver := &Driver{}
	fields := lib.SplitLine(line, ";")
	if len(fields) < 9 {
		return nil, fmt.Errorf("insuficient fields for Driver initialization")
	}
	id, _ := strconv.Atoi(fields[0])
	driver.ID = uint(id)
	driver.Name = fields[1]
	driver.BirthDate = fields[2]
	driver.Gender = fields[3]
	driver.CarClass = fields[4]
	driver.LicensePlate = fields[5]
	driver.City = fields[6]
	driver.AccountCreationDate = fields[7]
	driver.AccountStatus = fields[8]

	return driver, nil
}
