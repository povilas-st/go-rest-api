# go-rest-api
A simple REST API written in golang


## About
This is a small REST API for user management.\
Requirements:
* docker
* make

## Setup
1. `git clone` the repo somewhere
2. run `make start` inside the project directory
3. go to `http://localhost:8888/ping` inside your browser, you should see a `pong` response
4. to restart the app, hit `ctrl + C` and run `make reset`
4. run `make stop` after you are done working

## Examples
* List users
```
curl http://localhost:8888/users
```
* Get user
```
curl http://localhost:8888/users/1
```
* Create user
```
curl --header "Content-Type: application/json" --request POST --data '{"email":"darth.vader@empire.com","password":"darkside"}' http://localhost:8888/users
```
* Update user
```
curl --header "Content-Type: application/json" --request PUT --data '{"email":"bruce.wayne@gmail.com","password":"imbatman"}' http://localhost:8888/users/1
```
* Delete user
```
curl --request DELETE  http://localhost:8888/users/4
```
