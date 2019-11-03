package Service

import (
	"../Model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"log"
)

var DBEngine *xorm.Engine

func init() {
	fmt.Println("init start")
	driveName := "mysql"
	DBSrcName := "root:root@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DBEngine, err = xorm.NewEngine(driveName, DBSrcName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DBEngine.ShowSQL(true)
	DBEngine.SetMaxOpenConns(2)
	err = DBEngine.Sync2(new(Model.User))
	if err != nil {
		log.Fatal("failed to create table: "+err.Error())
	}

	fmt.Println("init ok")
}
