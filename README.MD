Hello.

## Current status:
Tests were done partially - integration tests, unit tests require more time.
The "optional: JSON in MQ" task requires more time
K8S requires more time

## API endpoints:
POST   /api/v1/sign-in        
POST   /api/v1/companies/     
GET    /api/v1/companies/     
GET    /api/v1/companies/:id  
PUT    /api/v1/companies/     
PUT    /api/v1/companies/:id  
DELETE /api/v1/companies/     
DELETE /api/v1/companies/:id  

## how to start local

need docker & docker-compose 

create data base:
```shell script
 docker-compose -f docker-compose-pg.yaml up -d
```
The database will be initialized automatically by the docker-composer via the /initdb/init.sh

The settings for the service are in configs/config.yaml

run project:
```shell script
 go build cmd/main.go
 go test ./... 
 go run cmd/main.go

```