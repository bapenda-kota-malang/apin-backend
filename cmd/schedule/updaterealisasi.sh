#!/bin/bash

NOW=`date`
echo "start update target realisasi at: $NOW"

source gentoken.sh

curl --location --request PATCH 'localhost:8080/target-realisasi/schedule' --header "Authorization: Bearer $TOKEN"