## Run Application
### Prerequisite 
#### Env variables
```bash
# On linux or unix systems set below envs
export mysql_users_username=root
export mysql_users_password=root1234
export mysql_users_host=localhost:53306
export mysql_users_schema=todo_app
export memcache_servers=localhost:11211,localhost:11212,localhost:11213
```

#### Run postgres as docker container
```bash
docker run -d \
--name task-db-postgres \
-e POSTGRES_PASSWORD=mysecretpassword \
-e PGDATA=/var/lib/postgresql/data/pgdata \
-v /custom/mount:/var/lib/postgresql/data \
-p 55432:5432 \
postgres
```

#### Run mysql as docker container
```bash
docker run -d \
--name task-db-mysql \
-e MYSQL_ROOT_PASSWORD=root1234 \
-p 53306:3306 \
mysql
```

#### Run memcached
```bash
docker run -d \
--name task-db-memcache1 \
-p 11211:11211 \
memcached memcached -m 64

docker run -d \
--name task-db-memcache2 \
-p 11212:11211 \
memcached memcached -m 64

docker run -d \
--name task-db-memcache3 \
-p 11213:11211 \
memcached memcached -m 64
# 64 megabytes for storage.
# -p <host_port>:<container_port>
```

### Run Application
```bash
go run main.go
```

### Misc
```bash
# copy, paste below in goland configuration's environment
mysql_users_username=root
mysql_users_password=root1234
mysql_users_host=localhost:53306
mysql_users_schema=todo_app
memcache_servers=localhost:11211,localhost:11212,localhost:11213
```