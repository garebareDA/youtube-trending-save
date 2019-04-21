package main

import(
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const(

	//Dialect
	Dialect = "mysql"

	//DBUser ユーザー名
	DBUser = "User"

	//DBPass パスワード
	DBPass = "password"

	//DBProtocol プロトコル
	DBProtocol = "tcp(127.0.0.1)"

	//DBName DB名
	DBName = "ranking"

	//DBchar 文字コード
	DBchar = "charset=utf8mb4"
)

type channel struct{
	gorm.Model
	URL string
	Title string
	ChannelID string
	ChannelTitle string
	Description string `gorm:"size:105"`
	Date string `gorm:"size:8"`
}

func connectDB() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s?%s&parseTime=true"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName, DBchar,)

	db, err := gorm.Open(Dialect, connect)
	if err != nil {
		panic(err)
	}

	return db
}

func insert(channels []channel, db *gorm.DB) {
	for _, channel := range channels {
			db.NewRecord(channel)
			db.Create(&channel)
	}
}

func findByDate(date string, db *gorm.DB) []channel {
	var findChannel []channel
	db.Where("Date = ?", date).Find(&findChannel)
	return findChannel
}
