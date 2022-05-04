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

example filtering:

<img width="788" alt="image" src="https://user-images.githubusercontent.com/13572116/166609110-bbce1db3-965b-4494-85cd-e0ab75a69900.png">
<img width="605" alt="image" src="https://user-images.githubusercontent.com/13572116/166609166-c39a0982-b8a0-497a-a54d-0affedef61d8.png">
<img width="580" alt="image" src="https://user-images.githubusercontent.com/13572116/166609205-1ff3a942-c8ed-4cfa-bf3e-c06b0b15c347.png">


<img width="788" alt="image" src="https://user-images.githubusercontent.com/13572116/166609110-bbce1db3-965b-4494-85cd-e0ab75a69900.png">

<img width="450" alt="image" src="https://user-images.githubusercontent.com/13572116/166609329-b9e87a54-ff25-4d0d-8bb4-1c1c1e24a695.png">
<img width="482" alt="image" src="https://user-images.githubusercontent.com/13572116/166609361-37814920-7cfc-4423-a1f4-97cbfed1dfe1.png">

update: filter: "usa", sample: only "phone"

<img width="615" alt="image" src="https://user-images.githubusercontent.com/13572116/166629124-36fb5447-8c29-4a2a-bcad-c07e008b2712.png">

acccess denied without jwt:

<img width="455" alt="image" src="https://user-images.githubusercontent.com/13572116/166631694-1e7a74af-3ee9-428d-9aa8-9db783022afc.png">
and DELETE, if I have one

<img width="491" alt="image" src="https://user-images.githubusercontent.com/13572116/166631817-eed396b5-e7a0-4a76-9b1d-5e68f148da75.png">

Delete the list of companies by filter

<img width="590" alt="image" src="https://user-images.githubusercontent.com/13572116/166632843-901102d1-2fed-4095-806e-4da076217a13.png">



