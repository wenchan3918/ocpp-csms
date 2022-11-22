#!/bin/bash

git add .
git commit -m "script deploy"
git push

#login remote host then git pull and restart container
ssh centos@sslmmsdev2 <<EOF
sudo su -
id
cd /data/ocpp-csms
git pull

docker-compose restart
EOF
