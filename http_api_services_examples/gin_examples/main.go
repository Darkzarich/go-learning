package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"gin_examples/files"
	"gin_examples/methods"
	"gin_examples/multipart_form"
	"gin_examples/parameters"
	"gin_examples/query"
)

func main() {
	r := gin.Default()

	query.RegisterRoutes(r)
	methods.RegisterRoutes(r)
	parameters.RegisterRoutes(r)
	multipart_form.RegisterRoutes(r)
	if err := files.RegisterRoutes(r); err != nil {
		log.Fatal(err)
	}

	r.Run(":3000")
}
