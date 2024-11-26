#!/bin/bash


docker build -t go-my-api

docker run -p 8080:8080 -ti go-my-apy 