package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Student struct {
	Id   int
	Name string
}

func init() {

	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	orm.RegisterDataBase("default", "postgres", "user=postgres password=123456 dbname=postgres host=192.168.13.13 port=5432 sslmode=disable")
	orm.Debug = true
	orm.RegisterModel(new(Student))
	// set default database
	orm.SetMaxOpenConns("default", 10)
}

func Test() {
	o := orm.NewOrm()
	o.Using("postgres")

	err := o.Begin()

	defer func() {
		fmt.Printf("o.Rollback()")
		o.Rollback()
		//fmt.Printf("not o.Rollback()")
	}()

	user := Student{Name: "slene"}

	u := Student{Id: 2}
	// read one
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
	fmt.Println(user.Id, user.Name)

	if true {
		fmt.Println("直接退出了")
		return
	}

	fmt.Println("没有提交事务")

	err = o.Commit()

	fmt.Println("提交了事务")
}
