package models

import (
    "github.com/astaxie/beego/orm"
    "time"
)

type User {
    Id           int       `orm:"column(id);auto"`
    FirstName    string    `orm:"column(firstname);size(50)"`
    LastName     string    `orm:"column(lastname);size(50)"`
    Email        string    `orm:"column(email);size(255)"`
    Password     string    `orm:"column(password);size(128)"`
    ctime        time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
    mtime        time.Time `orm:"column(create_time);`
    Rands        string    `orm:"size(10)"`
}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(User))
}
