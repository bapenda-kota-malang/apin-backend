#!/bin/bash

CURDIR="$(pwd)"
CRONSP="05 0 1 * * /bin/bash $CURDIR/createsp.sh >> $CURDIR/createsp.log"
CRONREALISASI="0 2 * * * /bin/bash $CURDIR/updaterealisasi.sh >> $CURDIR/updaterealisasi.log"

(crontab -l ; echo -e "$CRONSP\n$CRONREALISASI") | crontab -

echo "success add cron to crontab"