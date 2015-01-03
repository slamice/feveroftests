package models

import (
    "github.com/astaxie/beego/orm"
)

type User {
    Id           int
    FirstName    string
    LastName     string

}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(User))
}
