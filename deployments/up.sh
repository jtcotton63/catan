#!/bin/bash

# Remove previously copied files
rm -rf tmp

# Copy all files
mkdir tmp
cp -R ../pkg tmp/
cp ../go.mod tmp/

# Run docker compose
docker-compose build
docker-compose up