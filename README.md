# Go REST API Project for Employee
This repository contains an Employee REST API project developed by Go that can be run using Docker.

## Prerequisites
Docker installed on your machine
Docker Compose installed on your machine

## Running the Project
To run the project, execute the following command:

`docker-compose up -d`

This command will build and start the Docker containers defined in the docker-compose.yml file in detached mode, allowing them to run in the background.

## Database Migrations
### Creating Migration
To create a migration, execute the following command:

```docker-compose exec edgedb /bin/bash -c "edgedb --tls-security=insecure -H localhost -P 5656 --password -u {{username}} -d {{dbname}} migration create"```

This command will create a new migration for the database.

### Applying Migration
To apply the migration, execute the following command:

```docker-compose exec edgedb /bin/bash -c "edgedb --tls-security=insecure -H localhost -P 5656 --password -u {{username}} -d {{dbname}} migrate"```

This command will apply the migration to the database.
