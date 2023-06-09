# golang-clean-arch-api

This project is a microservice using <b>Golang, </b><b>Clean Architecture, </b><b>MongoDB, </b><b>MySQL</b> and <b>PostgreSQL</b>.
<br /><br />

## Requirements to run the project

* <b>Golang</b>
* <b>Docker</b>
<br /><br />

## Configure

<br />

1. In the root of the project run the command below:
<br /><br />

```
docker-compose up
```
<br />

2. Enter in the postgres database and create the database with the name: <b>varejao</b>
<br /><br />

3. Also enter in the MySQL database and create the database with the name: <b>macapa</b>
<br /><br />

4. In the postgres database run the script that is in the path: <b>/scripts/create-table-varejao.sql</b>
<br /><br />

5. In the MySQL database run the script that is in the path: <b>/scripts/create-table-macapa.sql</b>
<br /><br />

## How to run

<br />

1. In the root of the project run the command below:
<br /><br />

```
docker-compose up
```
<br />

2. After starting the docker container run the command below:
<br /><br />

```
go run main.go
```
<br />

<b>Note:</b> 
In the docker-compose.yml file there is a section where the two test users (varejao, macapa) are imported, the file with the import data is in mongo-seed/init.json
<br /><br />

## How to Use

<br />

1. If you run the command docker-compose up, two users will be added to the mongodb database, you need to use a database user to authenticate to the api and generate a Token, send a post request to <b>http://localhost:9000/users/login</b> url with a json as in the example below
<br /><br />

```json
    {
        "username": "macapa",
        "password": "12345678"
    }
```
<b>Or</b>

```json
    {
        "username": "varejao",
        "password": "12345678"
    }
```
<br />

2. Copy the generated token and add it in the header of your requests to the api, add as <b>Bearer Token</b>
<br /><br />

3. After configuring the token it will be possible to make requests for all routes
<br /><br />

## Swagger

To see all the routes access swagger use this route: http://localhost:9000/docs/index.html
<br /><br />

## Tests

To run all unit tests, run the command below:
<br /><br />

```
go test -v ./...
```
<br />