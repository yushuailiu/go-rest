package models

import (
	db "mydatabase"
	"log"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) Add() (id int64, err error) {
	rs, err := db.SqlDB.Exec("insert into person (first_name, last_name) values(?,?)",
		p.FirstName, p.LastName)
	if err != nil {
		log.Println(err)
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) Delete() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("delete from person where id=?", p.Id)
	if err != nil {
		log.Println(err)
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *Person) Update() (ra int64, err error) {
	stmt, err := db.SqlDB.Prepare("update person set last_name=?,first_name=? where id=?")
	if err != nil {
		log.Println(err)
	}
	rs, err := stmt.Exec(p.LastName, p.FirstName, p.Id)
	if err != nil {
		log.Println(err)
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *Person) Get() (err error) {
	err = db.SqlDB.QueryRow("select * from person where id=?", p.Id).Scan(
		&p.Id, &p.FirstName, &p.LastName)
	return
}

func (p Person) List(page int64, pageSize int64) (persons []Person, err error) {
	persons = make([]Person, 0)
	rs, err := db.SqlDB.Query("select * from person limit ?,?", (page-1 )*pageSize, pageSize)
	defer rs.Close()
	if err != nil {
		log.Println(err)
	}
	for rs.Next() {
		var person Person
		rs.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rs.Err(); err != nil {
		return
	}
	return
}
