package main

type account struct {
	id    int
	name  string
	money int
}

func QueryRowDemo(id int) (account, error) {
	sqlStr := "select id, name, money from account where id = ?"
	a := account{}
	err := db.QueryRow(sqlStr, id).Scan(&a.id, &a.name, &a.money)
	return a, err
}
