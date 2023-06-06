package main

import (
	"fmt"
	"log"

	u "github.com/c-seeger/Golang-HTML-TO-PDF-Converter"
)

type Data struct {
	Roll int
	Name string
}

type Page struct {
	Class    string
	Students []Data
}

type Report struct {
	Pages []Page
}

func main() {
	stu := []Data{
		{Roll: 1, Name: "andrew"},
		{Roll: 2, Name: "andrew1"},
		{Roll: 3, Name: "andrew2"},
		{Roll: 14, Name: "andrew3"},
		{Roll: 15, Name: "andrew4"},
		{Roll: 16, Name: "andrew5"},
		{Roll: 17, Name: "andrew6"},
		{Roll: 18, Name: "andrew7"},
		{Roll: 19, Name: "andrew8"},
		{Roll: 121, Name: "andrew9"},
		{Roll: 122, Name: "andrew11"},
		{Roll: 123, Name: "andrew12"},
		{Roll: 124, Name: "andrew13"},
		{Roll: 304, Name: "Stephen13"},
	}
	no_of_rows := 3
	array_size := len(stu)

	pages := int(array_size / no_of_rows)

	var pdf []Page
	var p Page

	for i := 0; i < pages; i++ {
		var temp []Data
		for j := 0; j < no_of_rows; j++ {
			temp = append(temp, stu[i*no_of_rows+j])
		}
		p = Page{Students: temp}
		pdf = append(pdf, p)
	}
	if array_size%no_of_rows > 0 {
		var temp []Data
		for j := 0; j < array_size%no_of_rows; j++ {
			temp = append(temp, stu[pages*no_of_rows+j])
		}
		p := Page{
			Class:    "tenth",
			Students: temp}
		pdf = append(pdf, p)
	}

	rep := Report{
		Pages: pdf,
	}
	fmt.Println(rep)

	new_pdf := u.NewRequestPdf("")

	templatePath := "report.html"

	outputPath := "example.pdf"

	if err := new_pdf.ParseTemplate(templatePath, rep); err != nil {
		log.Fatal(err)
	}
	if err := new_pdf.GeneratePDF(outputPath); err != nil {
		log.Fatal(err)
	}
	fmt.Println("pdf generated successfully")
}
