# Gin Examples

Just trying out Gin Web Framework

Each part of my journey is split into different packages that unite into one API via `RegisterRoutes` function that creates a new corresponding to that package router group.

API can be tested out with simply curl requests:

```shell
# methods
curl "localhost:3000/methods/get"
curl -X POST "localhost:3000/methods/post"
curl -X PUT "localhost:3000/methods/put"
curl -X DELETE "localhost:3000/methods/delete"
curl -X PATCH "localhost:3000/methods/patch"
curl -X HEAD "localhost:3000/methods/head"
curl -X OPTIONS "localhost:3000/methods/options"
# parameters
curl "localhost:3000/parameters/users/1"
curl -X GET "localhost:3000/parameters/users/1/string"
curl -X POST "localhost:3000/parameters/users/1/string"
curl -X POST -w "\nHTTP Status: %{http_code}\n" localhost:3000/parameters/users/5/error
# query
curl "localhost:3000/query/users?page=1&limit=10&sort=asc"
curl -X GET "localhost:3000/query/users?page=1&limit=10&sort=ascc" # invalid sort, applies default asc
curl -g -X GET "localhost:3000/query/posts?filters[rating]=5&filters[user]=me" # QueryMap {"filters":{"rating":"5","user":"me"}}
# multipart_form
curl -X POST -d "message=something&user=me" "localhost:3000/multipart_form/test_form"
curl -X POST -d "message=something" "localhost:3000/multipart_form/test_form" # missing user, applies default anonymous
curl -X POST -d "filters[rating]=5&filters[user]=me" "localhost:3000/multipart_form/posts" # PostFormMap {"filters":{"rating":"5","user":"me"}}
# files
curl -X POST -F "file=@/home/user/go-basics/http_api_services_examples/gin_examples/test.txt" "localhost:3000/files/single/upload"
# middleware (just simple fake token based auth middleware that restricts /middleware/test and /middleware/test2 endpoints)
curl http://localhost:3000/middleware/test -H "Token: 123456"
# redirect
curl "localhost:3000/redirect/test"
curl "localhost:3000/redirect/test2"
curl -X POST "localhost:3000/redirect/test3"
curl -X POST "localhost:3000/redirect/test4"\
# response - an example of how to make consistent response format (shape) and reduce boilerplate code
curl "localhost:3000/response/posts/0"
curl "localhost:3000/response/posts/1"
# header_versioning - an example of how to use middleware to handle different versions of API
curl "localhost:3000/header-versioning/hello"
curl -H "Accept-Version: v2" "localhost:3000/header-versioning/hello"
```
