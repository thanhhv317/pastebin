Docker MySQL master-slave replication
========================

MySQL master-slave replication with using Docker.

## Run

To run this examples you will need to start containers with "docker-compose"
and after starting setup replication. See commands inside ./build.sh.

#### Create 2 MySQL containers with master-slave row-based replication

```bash
./build.sh
```

#### Make changes to master

```bash
docker exec mysql_master sh -c "export MYSQL_PWD=Aa1234567890; mysql -u root shorten_link -e 'CREATE TABLE shortlinks( shortlink varchar(7) not null, link varchar(255) not null, created_at timestamp default current_timestamp                             null, updated_at timestamp default current_timestamp on update current_timestamp null, expired_at timestamp default current_timestamp null, index(shortlink) );'"
```

#### Read changes from slave

```bash
docker exec mysql_slave sh -c "export MYSQL_PWD=Aa1234567890; mysql -u root shorten_link -e 'select * from shortlinks \G'"
```

## Troubleshooting

#### Check Logs

```bash
docker-compose logs
```

#### Start containers in "normal" mode

> Go through "build.sh" and run command step-by-step.

#### Check running containers

```bash
docker-compose ps
```

#### Clean data dir

```bash
rm -rf ./master/data/*
rm -rf ./slave/data/*
```

#### Run command inside "mysql_master"

```bash
docker exec mysql_master sh -c 'mysql -u root -p111 -e "SHOW MASTER STATUS \G"'
```

#### Run command inside "mysql_slave"

```bash
docker exec mysql_slave sh -c 'mysql -u root -p111 -e "SHOW SLAVE STATUS \G"'
```

#### Enter into "mysql_master"

```bash
docker exec -it mysql_master bash
```

#### Enter into "mysql_slave"

```bash
docker exec -it mysql_slave bash
```
