package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"gin_examples/files"
	"gin_examples/header_versioning"
	"gin_examples/methods"
	"gin_examples/middleware"
	"gin_examples/multipart_form"
	"gin_examples/parameters"
	"gin_examples/query"
	"gin_examples/redirect"
	"gin_examples/response"
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
	middleware.RegisterRoutes(r)
	redirect.RegisterRoutes(r)
	response.RegisterRoutes(r)
	header_versioning.RegisterRoutes(r)

	r.Run(":3000")
}
