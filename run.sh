#!/bin/bash
echo 'Run and Create Container'
docker run -d --name mytokoApp -p 8080:8080 docker-mytoko:latest
