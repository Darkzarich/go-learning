package main

import (
	"github.com/gin-gonic/gin"

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

	r.Run(":3000")
}
