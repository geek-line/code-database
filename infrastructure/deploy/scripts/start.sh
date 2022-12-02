#!/bin/bash
APP_ROOT=/home/ubuntu/code-database
export $(cat $APP_ROOT/build/.env | grep -v '^#' | xargs)
cd $APP_ROOT/build
nohup ./code-database > $APP_ROOT/code-database.out.log 2> $APP_ROOT/code-database.err.log < /dev/null &