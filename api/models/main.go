package main

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@(127.0.0.1:3306)/authtutorial")
}
