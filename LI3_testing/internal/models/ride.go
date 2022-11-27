package models

import (
	"LI3_testing/lib"
	"fmt"
	"strconv"
)

type Ride struct {
	ID          uint
	Date        string
	DriverId    uint
	Driver      Driver
	Username    string
	User        User `gorm:"foreignKey:Username"`
	City        string
	Distance    uint
	ScoreUser   uint
	ScoreDriver uint
	Tip         float64
	Comment     string
}

func (*Ride) FromLine(line string) (OBJ, error) {
	// fmt.Printf("parsing:\"%s\" as ride\n", line)
	ride := &Ride{}
	fields := lib.SplitLine(line, ";")
	if len(fields) < 9 {
		return nil, fmt.Errorf("insuficient fields for ride initialization")
	}
	id, _ := strconv.Atoi(fields[0])
	ride.ID = uint(id)
	ride.Date = fields[1]
	driver_id, _ := strconv.Atoi(fields[2])
	ride.DriverId = uint(driver_id)
	ride.Username = fields[3]
	ride.City = fields[4]
	dist, _ := strconv.Atoi(fields[5])
	ride.Distance = uint(dist)
	u_score, _ := strconv.Atoi(fields[6])
	ride.ScoreUser = uint(u_score)
	d_score, _ := strconv.Atoi(fields[7])
	ride.ScoreDriver = uint(d_score)
	tip, _ := strconv.ParseFloat(fields[8], 64)
	ride.Tip = tip
	if len(fields) == 10 {
		ride.Comment = fields[9]
	}

	return ride, nil
}

// 000000000002;19/10/2019;000000002536;LoSousa98;Faro;1;4;2;5.0;Image outside north effect than though sport.
