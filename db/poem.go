package db

import (
	"encoding/json"
	"fmt"
)

type Poem struct {
	Id      int
	Title   string
	Auther  string
	Dynasty string
	Content string
}

func checkError(err error) bool {
	if err != nil {
		return true
	}
	return false
}
func (p *Poem) Insert() bool {
	stmtInsert, err := db.Prepare(" insert into poems (title,auther,dynasty,content) values (?,?,?,?) ")
	if checkError(err) {
		return false
	}
	_, err = stmtInsert.Exec(&p.Title, &p.Auther, &p.Dynasty, &p.Content)
	if checkError(err) {
		return false
	}
	return true
}
func (p *Poem) Save() {
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	p.Insert()

}

func QueryPoemsByAuthor(author string) (poems []Poem, err error) {

	return queryPoems("auther", author)
}

func queryPoems(field string, value string) (poems []Poem, err error) {
	sqlStr := "select id,title,auther,dynasty,content from poems where 1=1  "
	sqlStr += fmt.Sprintf(" and %s = ?", field)
	fmt.Println(sqlStr, value)
	stmtOut, err := db.Prepare(sqlStr)
	if checkError(err) {
		return nil, err
	}
	rows, err := stmtOut.Query(value)
	defer rows.Close()
	if checkError(err) {
		return nil, err
	}
	for rows.Next() {
		p := Poem{}
		err = rows.Scan(&p.Id, &p.Title, &p.Auther, &p.Dynasty, &p.Content)
		if err != nil {
			return nil, err
		}
		poems = append(poems, p)
	}
	fmt.Println(poems)
	return poems, nil
}
