Hello.

Welcome aboard!

POST   /api/v1/sign-in       
POST   /api/v1/companies/    
GET    /api/v1/companies/    
GET    /api/v1/companies/:id 
PUT    /api/v1/companies/:id 
DELETE /api/v1/companies/:id 


docker run --name=company-srvc-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres 
docker run --name company-srvc-db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=postgres -d -v "$(pwd)":/docker-entrypoint-initdb.d postgres:14-alpine
## how to start local

get docker
get docker-compose 

data base:
```shell script
 docker-compose -f docker-compose-pg.yaml up -d
```
message queue:
```shell script
 docker-compose -f docker-compose-mq.yaml up -d
```

run project:
```shell script
 go run cmd/main.go
```

entry point for database initialization - docker-entrypoint-initdb.d. Any *.sql or *.sh files in this directory will be treated as database initialization scripts. There are a few nuances:

if the database has already been initialized earlier, no changes will be applied to it;

if there are several files in the directory, they will be sorted by name using the current locale (by default en_US.utf8).

docker build -t postgres-db ./
docker run -d --name postgres-db-cmpny -p 5432:5432 postgres-db