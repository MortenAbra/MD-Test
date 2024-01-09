#!/bin/bash

# Start the docker services
docker compose up -d

# Check the status of the docker services
# Wait until all services are up and running
echo "Waiting for docker services to start..."
until [ $(docker compose ps -q | wc -l) -eq $(docker compose ps -q | xargs docker inspect -f '{{ .State.Running }}' | grep true | wc -l) ]; do
  echo -n "."
  sleep 1
done
echo "Docker services are up and running!"

# Run go mod download to download the necessary modules
go mod download

# Run the Go program
go run main.go
