# shoten10-diary-application
This repository provides the downloadable example programs for chapter 6 of the book, “コタツと蜜柑とゴーファーとわたし”.


### How to launch the application
```
$ docker-compose up --build
```

### How to connect the database
After launching your application with docker-compose, connect to the `wwgt-diary-db` container with the `wwgt-diary` user.
```
$ docker exec -it wwgt-diary-db psql -U wwgt-diary
```

### How to migrate
Exit from the container, perform migration.

```
$ export POSTGRESQL_URL='postgres://wwgt-diary:wwgt-diary@0.0.0.0:5435/wwgt-diary?sslmode=disable'
$ migrate -database ${POSTGRESQL_URL} -path migrations up
```
