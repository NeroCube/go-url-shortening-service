# Go URL SHORTENING SERVICE

## Running the exmaple

To run this exmaple, from the root of this project:

```
docker-compose up
```

## Hello Service

```
curl -X GET "http://localhost:8000/"
```

## Create your own short url

```
curl -H "Content-Type: application/json" -d '{"original_url":{original_url}}' http://localhost:8000/urls
```

## Redirect to your original url with tinyURL

```
curl -X GET "http://localhost:8000/{tinyURL}"
```

## To do

- [x] Basic Restful API
- [x] Dockerize project
- [x] Add Redis cache makes access performance better
- [x] Use PostgreSQL as long-term storage
- [x] Hash URLs to make them shorter
- [x] Go HTTP Redirect
- [ ] Unit test
- [ ] Refactor
- [ ] Golang vendor
