 #!/bin/bash

 echo "###### INIT SCRIPT STARTED ONLY WITH EMPTY BASE ######"

set -e

psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "postgres" <<-EOSQL
CREATE TABLE Country 
(
    id serial NOT NULL unique,
    name varchar(255) NOT NULL unique
);

CREATE TABLE Company 
(
    code serial NOT NULL unique,
    name varchar(255) NOT NULL unique,
    country_id int references country(id) on delete cascade NOT NULL,
    website varchar(255) NOT NULL,
    Phone varchar(255) NOT NULL
);
EOSQL