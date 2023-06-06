package main

import (
	"database/sql"
	"fmt"
	"log"

	u "github.com/c-seeger/Golang-HTML-TO-PDF-Converter"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func displayerror(e error) {
	if e != nil {
		panic(e.Error())
	}
}

type Data struct {
	Title       string
	Description string
}

func main() {

	r := u.NewRequestPdf("")

	templatePath := "sample.html"

	outputPath := "example.pdf"

	db, err = sql.Open("mysql", "lemma:admin@tcp(localhost:3306)/PDF")

	displayerror(err)
	defer db.Close()

	err = db.Ping()
	displayerror(err)
	fmt.Println("Successful Connection to Database!")

	stmt := "SELECT * FROM PDF.html ;"

	rows, err := db.Query(stmt)
	displayerror(err)
	defer rows.Close()

	var data Data

	for rows.Next() {
		err = rows.Scan(&data.Title, &data.Description)
		displayerror(err)
	}

	if err := r.ParseTemplate(templatePath, data); err != nil {
		log.Fatal(err)
	}
	if err := r.GeneratePDF(outputPath); err != nil {
		log.Fatal(err)
	}
	fmt.Println("pdf generated successfully")
}
