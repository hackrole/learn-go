package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	gorp "gopkg.in/gorp.v1"
)

func main() {
	// init dbmap
	dbmap := initDb()
	defer dbmap.Db.Close()

	// delete existing rows
	err := dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	// create two posts
	p1 := newPost("Go 1.1 released!", "lorem ipsum lorem ipsum")
	p2 := newPost("Go 1.2 released!", "lorem ipsum lorem ipsum")

	// insert rows - auto increment pks
	err = dbmap.Insert(&p1, &p2)
	checkErr(err, "inert failed")

	// use converience selectInt
	count, err := dbmap.SelectInt("select count(*) from posts")
	checkErr(err, "select count(*) failed")

	// upate a row
	p2.Title = "Go 1.2 is better than ever"
	count, err = dbmap.Update(&p2)
	checkErr(err, "update failed")
	log.Println("rows updated:", count)

	// fetch one row
	err = dbmap.SelectOne(&p2, "select * from posts where post_id=?", p2.Id)
	checkErr(err, "selectone failed")
	log.Println("p2 row:", p2)

	// fetch all rows
	var posts []Post
	_, err = dbmap.Select(&posts, "select * from posts order by post_id")
	checkErr(err, "select failed")
	log.Println("All rows:")
	for x, p := range posts {
		log.Printf("   %d: %v\n", x, p)
	}

	// delete row by pk
	count, err = dbmap.Delete(&p1)
	checkErr(err, "delete failed")
	log.Println("rows deleted:", count)

	// delete row manully via exec
	_, err = dbmap.Exec("delete from posts where post_id=?", p2.Id)
	checkErr(err, "Exec failed")

	// confirm count is zero
	count, err = dbmap.SelectInt("select count(*) from posts")
	checkErr(err, "select count(*) failed")
	log.Println("row count - should be zero:", count)
	log.Println("Done!")
}

type Post struct {
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `db:",size:50"`
	Body    string `db:"article_body,size:1024"`
}

func newPost(title, body string) Post {
	return Post{
		Created: time.Now().UnixNano(),
		Title:   title,
		Body:    body,
	}
}

func initDb() *gorp.DbMap {
	// connect to db using stanard go database/sql API
	db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
	checkErr(err, "sql.open failed")

	// construct a gorp dbmap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts'
	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	// create table
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
