package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassward := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassward, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	/*
		// select
		const sqlStr = `
			select *
			from articles
			where article_id = ?;
		`
		articleID := 2
		row := db.QueryRow(sqlStr, articleID)
		if err := row.Err(); err != nil {
			fmt.Println(err)
			return
		}

		var article models.Article
		var createdAt sql.NullTime

		err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName,
			&article.NiceNum, &createdAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		if createdAt.Valid {
			article.CreatedAt = createdAt.Time
		}

		fmt.Printf("%+v\n", article)
	*/

	/*
		// insert
		article := models.Article{
			Title:    "insert test",
			Contents: "Can I insert data correctly?",
			UserName: "kuro",
		}
		const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values
		(?, ?, ?, 0, now());
		`

		result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	*/

	// transaction
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	article_id := 1
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`
	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	const sqlUpdateNice = `
		update articles set nice = ? where article_id = ?;
	`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	tx.Commit()
}
