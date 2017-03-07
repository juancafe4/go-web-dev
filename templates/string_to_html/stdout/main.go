package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatalln(err)
	}
}

// func main() {
// 	name := "Juan Carlos Ferrel"

// 	tpl := `
//   <DOCTYPE html>
//   <html lang="eng"
//   <head>
//   <meta charset="UTF-8"
//   <title>Hello World!</title>
//   </head>
//   <body>
//   <h1>` + name + `</h1>
//   </body>
//   </html>
//   `
// 	fmt.Println(tpl)
// }
