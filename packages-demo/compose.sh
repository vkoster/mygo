#!/usr/bin/env bash

# Autor: vkoster
# Zweck: Steuerung für alle Build Varianten für ein Thema (z.B. generics).
#
# Beschreibung
# Decorator für Docker-Compose
#
if [ -z "$1" ]; then
  echo "Bitte einen docker-compose Befehl eingeben!"
  docker compose -h
  exit 1
fi
docker compose -p cmd -f compose.yaml $1 $2 $3