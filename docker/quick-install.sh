#!/bin/bash

#Tested only on Ubuntu 22.04

echo ">>>>>>>>>> CHECK INPUT ARGUMENTS <<<<<<<<<<"

EXTERNAL_IP=$1
DOMAIN=$2

if [ "$EUID" -ne 0 ]; then
  echo "Run this script as root"
  exit 1
fi

if [ -z "${EXTERNAL_IP}" ]; then
  echo "Pass external IP of your server in first argument"
  exit 1
fi

if [ -z "${DOMAIN}" ]; then
  echo "Pass your domain in second argument"
  exit 1
fi

echo ">>>>>>>>>> INSTALL DOCKER <<<<<<<<<<"

if ! command -v docker &> /dev/null; then
  #https://docs.docker.com/engine/install/ubuntu/
  apt-get install ca-certificates curl gnupg
  install -m 0755 -d /etc/apt/keyrings
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg
  chmod a+r /etc/apt/keyrings/docker.gpg
  echo \
    "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
    "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
    tee /etc/apt/sources.list.d/docker.list > /dev/null
  apt-get update
  apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
else
  echo "docker is already installed"
fi

docker network create --ignore playtime_network

echo ">>>>>>>>>> BUILD PLAYTIME IMAGE <<<<<<<<<<"

if [ ! -x build ]; then
  git clone "https://github.com/n-at/playtime" "build"
fi
cd build
docker image build -t "playtime:latest" .
cd ..

echo ">>>>>>>>>> LET'S ENCRYPT CERTIFICATE <<<<<<<<<<"

apt-get update
apt-get install openssl

mkdir certbot certbot/etc certbot/log certbot/var certbot/webroot
cd certbot
cp "../build/docker/certbot/docker-compose.yml" .
sed -i "s/REPLACE_DOMAIN/${DOMAIN}/g" docker-compose.yml
openssl dhparam -out "etc/dhparam.pem" 2048
cd ..

mkdir router router/logs router/conf
cd router
cp "../build/docker/router/docker-compose.yml" .
cp "../build/docker/router/certbot.conf" "conf/vhost.conf"
docker compose up -d
cd ..

cd certbot
docker compose up
cd ..

#install crontab job
CRON_JOB="$(which docker) compose -f \"${PWD}/docker-compose.yml\" up && $(which docker) kill -s SIGHUP playtime-router"
crontab -l > _tmp_crontab
echo "${CRON_JOB}" >> _tmp_crontab
crontab _tmp_crontab
rm _tmp_crontab

echo ">>>>>>>>>> SETUP COTURN <<<<<<<<<<"

TURN_USER="turnuser"
TURN_PASSWORD=$(openssl rand -base64 20)

mkdir -m 0777 coturn coturn/data coturn/db coturn/log
cd coturn
cp "../build/docker/coturn/turnserver.conf" .
sed -i "s/REPLACE_EXTERNAL_IP/${EXTERNAL_IP}/g"     "turnserver.conf"
sed -i "s/REPLACE_DOMAIN/${DOMAIN}/g"               "turnserver.conf"
sed -i "s/REPLACE_TURN_LOGIN/${TURN_USER}/g"        "turnserver.conf"
sed -i "s/REPLACE_TURN_PASSWORD/${TURN_PASSWORD}/g" "turnserver.conf"
cp "../build/docker/coturn/docker-compose.yml"
docker compose up -d
cd ..

echo ">>>>>>>>>> SETUP PLAYTIME <<<<<<<<<<"

mkdir -m 0777 playtime playtime/data playtime/uploads
cd playtime
cp "../build/docker/playtime/docker-compose.yml" .
sed -i "s/REPLACE_TURN_URL/turn:${DOMAIN}:3478/g"   "docker-compose.yml"
sed -i "s/REPLACE_TURN_USER/${TURN_USER}/g"         "docker-compose.yml"
sed -i "s/REPLACE_TURN_PASSWORD/${TURN_PASSWORD}/g" "docker-compose.yml"
docker compose up -d
cd ..

cd router
cp "../build/docker/router/playtime.conf" "conf/vhost.conf"
sed -i "s/REPLACE_DOMAIN/${DOMAIN}/g" "conf/vhost.conf"
docker compose up -d --force-recreate
