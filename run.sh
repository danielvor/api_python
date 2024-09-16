#!/bin/bash

# This script is used to give a user the correct permissions
chmod +x run.sh

# Start Python server
python3 -m pip install -r python_app/requirements.txt
python3 python_app/main.py &

# Start Go server
go run go_app/main.go &

# Serve HTML (using Python HTTP server)
python3 -m http.server 8081 --directory .
