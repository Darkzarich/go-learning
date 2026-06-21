package main

import "fmt"

type APIClient interface {
	GetData(query string) (Response, error)
}

type Response struct {
	Text       string
	StatusCode int
}

type MyAPIClient struct {
}

func NewAPIClient() *MyAPIClient {
	return &MyAPIClient{}
}

func (c *MyAPIClient) GetData(query string) (Response, error) {
	fmt.Println("Imitating real API call...", query)

	return Response{
		Text:       "Hello, World!",
		StatusCode: 200,
	}, nil
}

func handleRequest(c APIClient, query string) (Response, error) {
	res, err := c.GetData(query)
	if err != nil {
		return Response{}, err
	}

	return res, nil
}

func main() {
	c := NewAPIClient()

	res, err := handleRequest(c, "foo")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
