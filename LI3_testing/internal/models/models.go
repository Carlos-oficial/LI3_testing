package models

import (
	"bufio"
	"os"

	"gorm.io/gorm"
)

type OBJ interface {
	FromLine(string) (OBJ, error)
}

func LoadFile(db *gorm.DB, mapper OBJ, file_path string) error {
	readFile, err := os.Open(file_path)
	if err != nil {
		return nil
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var objs []OBJ
	// var i uint = 1
	for fileScanner.Scan() {
		mapper, err = mapper.FromLine(fileScanner.Text())
		if err != nil {
			return err
		}
		objs = append(objs, mapper)
		// if i%1000 == 0 {
		// 	// inserts into the db every 1000 rows
		// 	println(db.Create(objs))
		// 	objs = []OBJ{}
		// }
		// i++
	}
	db.Create(objs)
	return err
}
