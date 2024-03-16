# Saver

## Description

just downloads a image from a specific url and saves it in `./upload` directory.

## GoLang - Start & Build

-   start app with `go run main.go -imageUrl <url>`
-   for example: `go run main.go -imageUrl https://www.br.de/radio/puls/puls-meme-100~_v-img__16__9__l_-1dc0e8f74459dd04c91a0d45af4972b9069f1135.jpg`
-   build app with `go build`

## environment variables

-   check `.env.examples`:

```
MONGODB_CONNECTION_STRING=mongodb://localhost:27017
MONGODB_DATABASE_STRING=saver-database
MONGODB_COLLECTION_STRING=saver-collection
```

## Using MongoDB with Mongosh

-   don't forget to run `docker-compose up -d` to create a mongo database
-   use `docker exec -it mongodb bash` to connect to mongodb-container
-   use `mongosh` to use the mongodb-shell and then run `use saver-database`

### Read all entires

-   check collection `saver-collection` with inserted docs: `db["saver-collection"].find()`

### Delete all entries

-   `db["saver-collection"].deleteMany({})`
