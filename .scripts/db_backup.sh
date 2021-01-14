#! /bin/bash

# used to backup the postgresql database

MC_COPY="yes"
MC_INSECURE="yes"
MC_HOST="..."
MC_BUCKET="indexed-db-backups"
FILENAME=dump_`date +%d-%m-%Y"_"%H_%M_%S`.sql

if [[ ! -d db_backups ]]; then
    mkdir db_backups
fi

# https://stackoverflow.com/questions/24718706/backup-restore-a-dockerized-postgresql-database
docker-compose exec -T postgres pg_dumpall -c -U postgres > "db_backups/$FILENAME"

gzip -9 "db_backups/$FILENAME"

# used to copy backups to remote minio host
if [[ "$MC_COPY" == "yes" ]]; then
    if [[ "$MC_INSECURE" == "yes" ]]; then
        mc cp --insecure "db_backups/$FILENAME.gz" "$MC_HOST/$MC_BUCKET"
    else
        mc cp "db_backups/$FILENAME.gz" "$MC_HOST/$MC_BUCKET"
    fi
fi
