#!/bin/bash
cd ..

docker stack deploy -c <(docker-compose -f app.yml config) app
docker stack deploy -c <(docker-compose -f db.yml config) app

cd shell