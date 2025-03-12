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

apt-get update
apt-get install -yq git openssl

if ! command -v docker &> /dev/null; then
  apt-get install -yq docker.io
fi

if [ -x "$(command -v docker-compose)" ]; then
    DOCKER_COMPOSE="docker-compose"
    DOCKER_COMPOSE_CMD="$(which docker-compose)"
elif $(docker compose &>/dev/null) && [ $? -eq 0 ]; then
    DOCKER_COMPOSE="docker compose"
    DOCKER_COMPOSE_CMD="$(which docker) compose"
else
    apt-get install -yq docker-compose
    DOCKER_COMPOSE="docker-compose"
    DOCKER_COMPOSE_CMD="$(which docker-compose)"
fi

if [ -z "$(docker network ls | grep playtime_network)" ]; then
  docker network create playtime_network
fi

echo ">>>>>>>>>> BUILD PLAYTIME IMAGE <<<<<<<<<<"

if [ ! -d build ]; then
  git clone "https://github.com/n-at/playtime" "build"
fi
cd build
docker image build -t "playtime:latest" .
cd ..

echo ">>>>>>>>>> LET'S ENCRYPT CERTIFICATE <<<<<<<<<<"

mkdir -m 0777 certbot certbot/etc certbot/log certbot/var certbot/webroot
cd certbot
cp "../build/docker/certbot/docker-compose.yml" .
sed -i "s/REPLACE_DOMAIN/${DOMAIN}/g" docker-compose.yml
cd ..

if [ ! -f "certbot/etc/dhparam.pem" ]; then
  openssl dhparam -out "certbot/etc/dhparam.pem" 2048
fi

mkdir -m 0777 router router/logs router/conf
cd router
cp "../build/docker/router/docker-compose.yml" .
cp "../build/docker/router/certbot.conf" "conf/vhost.conf"
$DOCKER_COMPOSE up -d
cd ..

cd certbot
$DOCKER_COMPOSE up
cd ..

#install crontab job
CRON_JOB="${DOCKER_COMPOSE_CMD} -f \"$(realpath $PWD)/certbot/docker-compose.yml\" up && $(which docker) kill -s SIGHUP playtime-router playtime-coturn"
crontab -l > _tmp_crontab
echo "0 0 * * * ${CRON_JOB}" >> _tmp_crontab
crontab _tmp_crontab
rm _tmp_crontab

echo ">>>>>>>>>> SETUP COTURN <<<<<<<<<<"

TURN_USER="turnuser"
TURN_PASSWORD=$(openssl rand -hex 20)

mkdir -m 0777 coturn coturn/data coturn/db coturn/log
cd coturn
cp "../build/docker/coturn/turnserver.conf" .
sed -i "s/REPLACE_EXTERNAL_IP/${EXTERNAL_IP}/g"     "turnserver.conf"
sed -i "s/REPLACE_DOMAIN/${DOMAIN}/g"               "turnserver.conf"
sed -i "s/REPLACE_TURN_LOGIN/${TURN_USER}/g"        "turnserver.conf"
sed -i "s/REPLACE_TURN_PASSWORD/${TURN_PASSWORD}/g" "turnserver.conf"
cp "../build/docker/coturn/docker-compose.yml" .
$DOCKER_COMPOSE up -d
cd ..

echo ">>>>>>>>>> SETUP PLAYTIME <<<<<<<<<<"

mkdir -m 0777 playtime playtime/data playtime/uploads
cd playtime
cp "../build/docker/playtime/docker-compose.yml" .
sed -i "s/REPLACE_TURN_URL/turn:${DOMAIN}:3478/g"   "docker-compose.yml"
sed -i "s/REPLACE_TURN_USER/${TURN_USER}/g"         "docker-compose.yml"
sed -i "s/REPLACE_TURN_PASSWORD/${TURN_PASSWORD}/g" "docker-compose.yml"
$DOCKER_COMPOSE up -d --force-recreate
cd ..

cd router
cp "../build/docker/router/playtime.conf" "conf/vhost.conf"
sed -i "s/REPLACE_DOMAIN/${DOMAIN}/g" "conf/vhost.conf"
$DOCKER_COMPOSE up -d --force-recreate
cd ..

sleep 5
playtime_password=$(cat "playtime/data/admin.password")
echo ">>>>>>>>>> DONE <<<<<<<<<<"
echo ""
echo "https://${DOMAIN}/"
echo ""
echo "Your login: admin"
echo "Your password: ${playtime_password}"
echo ""
