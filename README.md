# Go URL SHORTENING SERVICE

## Running the exmaple

To run this exmaple, from the root of this project:

```
go run *.go

go get github.com/gorilla/mux
```

## Hello Service

```
curl -X GET "http://localhost:8000/"
```

## Get All Url Maps

```
curl -X GET "http://localhost:8000/urls"
```

## Get specific url with url Id

```
curl -X GET "http://localhost:8000/urls/{urlId}"
```

## Create your own short url

```
curl -H "Content-Type: application/json" -d '{"original_url":{original_url}}' http://localhost:8000/urls
```

## To do

- [x] Basic Restful API
- [ ] Dockerize project
- [ ] Add Redis cache makes access performance better
- [ ] Use PostgreSQL as long-term storage
- [ ] Hash URLs to make them shorter
- [ ] Go HTTP Redirect
- [ ] Unit test
- [ ] Refactor
