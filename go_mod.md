# Go Mod is a tool which helps

## Other tools for example

1. Go build
2. Go mod
3. Go run
4. GO env
5. go mod verify
6. go mod tidy

## best practice to name go modules

hosting/user/modulename
EG: github.com/harshareddi/mymodule

## GO mod file

1. Contains module title
2. GO version
3. Dependencies / packages

## For installing any pkg

go get -u hostedOn/module
go get -u github.com/gorilla/mux
Right after installing the mux it will be added to go mod file
and wil be saved in local dir -> /Users/harshavardhanreddy/go/pkg/mod/cache/download/github.com/gorilla/mux
Next time we request this module wont be fetched from internet again. It will be used from this path

## Go SUM file

To avoid malicious attacks this file will have the mod of libraries for the specific versions

## GO list all

lists all the installed packages
go list -m all -> lists the current module pkgs / dependencies
go list -m -versions github.com/gorilla/mux
Example output: <!-- ithub.com/gorilla/mux v1.2.0 v1.3.0 v1.4.0 v1.5.0 v1.6.0 v1.6.1 v1.6.2 v1.7.0 v1.7.1 v1.7.2 v1.7.3 v1.7.4 v1.8.0 v1.8.1 -->

## go mod why github.com/gorilla/mux

shows where the dependecies are being used

## go mod graph

serves same purpose as above and lists the mapping btwn packages and modules in the project

## go mod edit

If mod file needs to be edited
go mod edit -go -> changes the go versions
go mod edit -module -> if changes required to the module name

## go mod vendor

fetches the modules code and places in the project directoray under a folder called vendor. more of like a node_modules in js instead of default difrectory
