#!/usr/bin/env bash

# Autor: vkoster
# Zweck: Command bauen und ins Hauptverzeichnis legen.
#
# Beschreibung
# Decorator für Docker-Compose
#
if [ -z "$1" ]; then
  echo "Bitte einen docker-compose Befehl eingeben!"
  docker compose -h
  exit 1
fi
docker compose -p cmd -f linuxbuild-cmd.yaml $1 $2 $3