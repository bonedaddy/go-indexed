#! /bin/bash

# used to backup the postgresql database

if [[ ! -d db_backups ]]; then
    mkdir db_backups
fi

FILENAME=dump_`date +%d-%m-%Y"_"%H_%M_%S`.sql

# https://stackoverflow.com/questions/24718706/backup-restore-a-dockerized-postgresql-database
docker-compose exec -T postgres pg_dumpall -c -U postgres > "db_backups/$FILENAME"

gzip -9 "db_backups/$FILENAME"