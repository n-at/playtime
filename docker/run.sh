#!/bin/bash

PLAYTIME_ARGS=""

#DEBUG

if [ -n "${PLAYTIME_DEBUG_EMULATOR}" ]; then
  PLAYTIME_ARGS="${PLAYTIME_ARGS} -debug-emulator"
fi

if [ -n "${PLAYTIME_DEBUG_TEMPLATES}" ]; then
  PLAYTIME_ARGS="${PLAYTIME_ARGS} -debug-templates"
fi

if [ -n "${PLAYTIME_DEBUG_NETPLAY}" ]; then
  PLAYTIME_ARGS="${PLAYTIME_ARGS} -debug-netplay"
fi

if [ -n "${PLAYTIME_VERBOSE}" ]; then
  PLAYTIME_ARGS="${PLAYTIME_ARGS} -verbose"
fi

#TURN

if [ -n "${PLAYTIME_TURN_URL}" ]; then
  PLAYTIME_ARGS="${PLAYTIME_ARGS} -turn-server-url ${PLAYTIME_TURN_URL}"
fi

if [ -n "${PLAYTIME_TURN_USER}" ]; then
  PLAYTIME_ARGS="${PLAYTIME_ARGS}  -turn-server-user ${PLAYTIME_TURN_USER}"
fi

if [ -n "${PLAYTIME_TURN_PASSWORD}" ]; then
  PLAYTIME_ARGS="${PLAYTIME_ARGS}  -turn-server-password ${PLAYTIME_TURN_PASSWORD}"
fi

#

./app ${PLAYTIME_ARGS}
