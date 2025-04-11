#!/bin/sh

echo "Network Up"
minifab netup -s couchdb -e true -i 2.4.8 -o bank.org.com

sleep 5

echo "Create Channel"
minifab create -c tritrustchannel

sleep 2

echo "Join Channel"
minifab join -c tritrustchannel