#!/bin/bash

docker compose -f dev.docker-compose.yml --env-file .env.dev --env-file ../secrets.env up --build