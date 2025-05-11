#!/bin/sh

echo "Start the Network"
minifab netup -s couchdb -e true -i 2.4.8 -o bank.fin.com

sleep 5

echo "Create the Channel"
minifab create 

sleep 2

echo "Join peers to the channel"
minifab join 

sleep 2

echo "Anchor update"
minifab anchorupdate