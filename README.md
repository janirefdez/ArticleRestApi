# ArticleRestApi
[![Go](https://img.shields.io/badge/Go-1.17.1-blue.svg)](https://go.dev/doc/)

Basic example where you can see how to create a simple REST API with Go. 

If you really want to know in detail how it is implemented, you can check my post in this [link](https://dev.to/janirefdez/create-a-rest-api-with-go-1j52)

## How to run the app 
```
go run cmd/main.go
```

## Requests 
 - `GET /articles`:  Get all articles.
 - `GET /articles/{id}`:  Get article by id.
 -  `POST /articles`: Add new Article.
    - Body should contain: 
        - `Title`: A string with the title of the article.
        - `Desc`: A string with the description of the article.
        - `Content`: A string with the content of the article. 
 - `PUT /articles/{id}`: Update article by id.
    - Body should contain: 
        - `Title`: A string with the title of the article.
        - `Desc`: A string with the description of the article.
        - `Content`: A string with the content of the article. 
 - `DELETE /articles/{id}`: Delete article by id. 
