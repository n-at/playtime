#!/bin/bash

cd assets

npm install

git clone https://github.com/EmulatorJS/EmulatorJS _tmp
mv _tmp/data/* emulatorjs
rm -rf _tmp
