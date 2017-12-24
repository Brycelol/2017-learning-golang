#!/bin/bash

PSQL_SERVICE_NAME="postgres"
PSQL_SERVICE_VERSION="10.1-alpine"
PSQL_SERVICE_IMAGE="$PSQL_SERVICE_NAME:$PSQL_SERVICE_VERSION"
PSQL_SERVICE_NAME_CONTAINER="chitchat-postgres"
PSQL_PORT=5432

PSQL_CONT_DATA=/var/lib/postgresql/data/pgdata

PSQL_READINESS_CHECK="database system is ready to accept connections"

ENV_PSQL_PW="POSTGRES_PASSWORD"
ENV_PSQL_USER="POSTGRES_USER"
ENV_PSQL_PGDATA="PGDATA"
ENV_PSQL_DB="POSTGRES_DB"

ENV_PSQL_DEFAULT="chitchat"

SQL_FILE="chitchat.sql"

function log()
{
  echo "$(date) [INFO]-[init-postgres] $1"
}

function log_and_exit()
{
  echo "$(date) [ERROR]-[init-postgres] $1"
  if [ ! -z $2 ] ; then exit $2 ; else exit 1; fi
}

function is_active()
{
  systemctl is-active "$1" >> /dev/null && return 0 || return 1
}

function pull_docker_image()
{
  log "Pulling Docker image [$1]."

  if ! is_active "docker" ; then
    log_and_exit "Docker not currently running." 2
  fi

  docker pull "$1" || log_and_exit "An error occurred while pulling docker image."
}

function run_container()
{
  log "Starting container..."

  if ! is_active "docker" ; then
    log_and_exit "Docker not currently running." 2
  fi

  docker run "$@" || log_and_exit "An error occurred while starting container."
}

function has_container()
{
  if ! is_active "docker" ; then
    log_and_exit "Docker not currently running." 2
  fi

  docker ps -a -f name=="$1" >> /dev/null 2>&1 && return $?
}

function container_running()
{
  log "Checking if container [$1] is running."

  if ! is_active "docker" ; then
    log_and_exit "Docker not currently running." 2
  fi

  if ! hasContainer "$1" ; then
    log_and_exit "Container $1 does not exist." 3
  fi

  if [[ $(docker inspect "$1" | jq '.[].State.Status') == "\"running\"" ]] ; then
    return 0
  else
    return 1
  fi
}

if ! is_active "docker" ; then
  log_and_exit "Docker not currently running." 2
fi

log "Initializing $PSQL_SERVICE_NAME_CONTAINER..."

pull_docker_image "$PSQL_SERVICE_IMAGE"

run_container -d \
  -e "$PSQL_PORT":"$PSQL_PORT" \
  -e "$ENV_PSQL_DB"="$ENV_PSQL_DEFAULT" \
  -e "$ENV_PSQL_USER"="$ENV_PSQL_DEFAULT" \
  -e "$ENV_PSQL_PW"="$ENV_PSQL_PW" \
  -e "$ENV_PSQL_PGDATA"="$PSQL_CONT_DATA" \
  -v ${PWD}/"$SQL_FILE":/"$SQL_FILE" \
  --restart=unless-stopped \
  --name="$PSQL_SERVICE_NAME_CONTAINER" \
  "$PSQL_SERVICE_IMAGE"

sleep 5s

log "Postgres initialized."

log "Creating database..."

sleep 5s

docker exec "$PSQL_SERVICE_NAME_CONTAINER" psql -U "$ENV_PSQL_DEFAULT" -f "$SQL_FILE"

log "Finished initializing $PSQL_SERVICE_NAME_CONTAINER."
