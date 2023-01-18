#!/bin/bash

NOW=`date`
echo "start create surat pemberitahuan at: $NOW"

source gentoken.sh

curl --location --request PATCH 'localhost:8080/suratpemberitahuan' --header "Authorization: Bearer $TOKEN"