# GoREST
My playground for creating REST API using Golang.

# Last changes
* simple (CRUD) user management 
* simple JSON Web Token (JWT) authentication *thanks to [appleboy](github.com/appleboy/gin-jwt)
* add auto migrate using cli

# How to use
* clone this repo
* install the cli
    * go to gcli directory `cd gcli`
    * `go install`
* run the `gcli` from the root of this project
    * do migration `gcli db migrate`
    * make sure there's no error on your migration 
* run the app `go run main.go`

# Used things
* [Gin Gonic](https://github.com/gin-gonic/gin)
* [Gorm](https://github.com/jinzhu/gorm)

# Wanna discuss with me?
Very open for suggestions. Sharing is caring is learning.
Just mail me at awcodify@gmail.com
