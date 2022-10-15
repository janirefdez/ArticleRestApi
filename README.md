# ArticleRestApi
[![Go](https://img.shields.io/badge/Go-1.17.1-blue.svg)](https://go.dev/doc/)

There are two branches in this project:

- `main`: example where you can see how to create a simple REST API with Go => [Tutorial: Create rest API with Go](https://dev.to/janirefdez/create-a-rest-api-with-go-1j52) 
- `connection-database`: example to know how to connect REST API with PostgreSQL database => [Tutorial: Connect rest API to database](https://dev.to/janirefdez/connect-rest-api-to-database-with-go-d8m)

## How to run the API 
```
go run cmd/main.go
```

## Request endpoints

- `GET /articles`=>  Get all articles. 

Request-response example:
```
GET /articles HTTP/1.1
Host: localhost:8080
Connection: close
User-Agent: Paw/3.4.0 (Macintosh; OS X/12.6.0) GCDHTTPRequest
```

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 15 Oct 2022 15:17:54 GMT
Content-Length: 415
Connection: close

[{"Id":"8617bf49-39a9-4268-b113-7b6bcd189ba2","Title":"Article 1","desc":"Article Description 1","content":"Article Content 1"},{"Id":"9b368c22-bdd9-4b78-b5d2-8efa3cd942f7","Title":"New Article","desc":"New article description","content":"New content description"},{"Id":"38da7ce2-02b5-471a-90b8-c299f2ef132e","Title":"Updated Article","desc":"Updated article description","content":"Updated content description"}]
```

 - `GET /articles/{id}`:  Get article by id.

Request-response example:
```
GET /articles/8617bf49-39a9-4268-b113-7b6bcd189ba2 HTTP/1.1
Host: localhost:8080
Connection: close
User-Agent: Paw/3.4.0 (Macintosh; OS X/12.6.0) GCDHTTPRequest
```
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 15 Oct 2022 15:17:41 GMT
Content-Length: 127
Connection: close

{"Id":"8617bf49-39a9-4268-b113-7b6bcd189ba2","Title":"Article 1","desc":"Article Description 1","content":"Article Content 1"}
```

 -  `POST /articles`: Add new Article.

Request-response example:

```
POST /articles HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
Connection: close
User-Agent: Paw/3.4.0 (Macintosh; OS X/12.6.0) GCDHTTPRequest
Content-Length: 92

{"Title":"New Article","Desc":"New article description","Content":"New content description"}
```

```
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sat, 15 Oct 2022 15:17:43 GMT
Content-Length: 10
Connection: close

"Created"
```

 - `PUT /articles/{id}`: Update article by id.

Request-response example:
```
PUT /articles/38da7ce2-02b5-471a-90b8-c299f2ef132e HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
Connection: close
User-Agent: Paw/3.4.0 (Macintosh; OS X/12.6.0) GCDHTTPRequest
Content-Length: 104

{"Title":"Updated Article","Desc":"Updated article description","Content":"Updated content description"}

```

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 15 Oct 2022 15:17:52 GMT
Content-Length: 10
Connection: close

"Updated"
```

 - `DELETE /articles/{id}`: Delete article by id. 

Request-response example:

```
DELETE /articles/8617bf49-39a9-4268-b113-7b6bcd189ba2 HTTP/1.1
Host: localhost:8080
Connection: close
User-Agent: Paw/3.4.0 (Macintosh; OS X/12.6.0) GCDHTTPRequest
```

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 15 Oct 2022 15:17:57 GMT
Content-Length: 10
Connection: close

"Deleted"
```
