#!/bin/bash

DATABASE="apin-migrate-old-2"
USER="dexwip"
SCHEMA="public"

echo "start dump data from $DATABASE $SCHEMA schema"

for table in $(psql -U $USER -d $DATABASE -t -c "Select table_name From information_schema.tables Where table_type='BASE TABLE' and table_schema ='$SCHEMA' order by table_name");
do pg_dump -t \"$table\" -U $USER $DATABASE --data-only > data-$table.sql;
done;

echo "done"
