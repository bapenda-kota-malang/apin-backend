#!/bin/bash

TOKEN="$(curl -s --location --request POST 'localhost:8080/auth/login' --header 'Content-Type: application/json' \
--data-raw '{ "name": "admin", "password": "P@SS123"}' | \
python3 -c "import sys, json; print(json.load(sys.stdin)['data']['accessToken'])")"

if test -z "$TOKEN" 
then
   echo "token empty"
   exit 1
fi