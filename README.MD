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

entry point for database initialization - docker-entrypoint-initdb.d. 

Any *.sql or *.sh files in this directory will be treated as database initialization scripts. There are a few nuances:

if the database has already been initialized earlier, no changes will be applied to it;

if there are several files in the directory, they will be sorted by name using the current locale (by default en_US.utf8).


docker build -t postgres-db ./
docker run -d --name postgres-db-cmpny -p 5432:5432 postgres-db

<img width="788" alt="image" src="https://user-images.githubusercontent.com/13572116/166609110-bbce1db3-965b-4494-85cd-e0ab75a69900.png">
<img width="605" alt="image" src="https://user-images.githubusercontent.com/13572116/166609166-c39a0982-b8a0-497a-a54d-0affedef61d8.png">
<img width="580" alt="image" src="https://user-images.githubusercontent.com/13572116/166609205-1ff3a942-c8ed-4cfa-bf3e-c06b0b15c347.png">


<img width="788" alt="image" src="https://user-images.githubusercontent.com/13572116/166609110-bbce1db3-965b-4494-85cd-e0ab75a69900.png">

<img width="450" alt="image" src="https://user-images.githubusercontent.com/13572116/166609329-b9e87a54-ff25-4d0d-8bb4-1c1c1e24a695.png">
<img width="482" alt="image" src="https://user-images.githubusercontent.com/13572116/166609361-37814920-7cfc-4423-a1f4-97cbfed1dfe1.png">

update: filter: "usa", sample: only "phone"
<img width="615" alt="image" src="https://user-images.githubusercontent.com/13572116/166629124-36fb5447-8c29-4a2a-bcad-c07e008b2712.png">

```
