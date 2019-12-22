#!/bin/bash
docker build -t backend .
docker run -d -p 8080:8080 --restart=always --name backend_cinemavilyui backend