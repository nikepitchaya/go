package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Gender struct {
	Id       uint
	Name     string `grom:"column:keroro;type:varchar(50)"`
}

//ต้องการ log ข้อมูลออกมาดูว่า ตัว gorm genarate Sql Statement อะไรให้เราบ้าง
type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v \n __________________________________ \n", sql)
}

func main() {
	dsn := "root:@tcp(localhost:3306)/gorm?parseTime=true"
	dial := mysql.Open(dsn) //ต้องการ data source name
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
	}) //ต้องการ go driver ของดาต้าเบสที่ต้องการเชื่อมต่อและสามารถเพิ่ม option ของ gorm ได้โดยอยู่ภายใต้ &gorm.Config
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Gender{})
}
