#!/bin/bash

docker compose -f prod.docker-compose.yml --env-file .env.prod --env-file ../secrets.env up --build