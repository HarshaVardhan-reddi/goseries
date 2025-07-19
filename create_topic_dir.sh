#!/bin/bash
VERSION=$1
NAME=$2

FOLDER="${VERSION}my${NAME}"

echo "creating ${FOLDER}"
mkdir $FOLDER
cd $FOLDER
go mod init "my${NAME}"
touch main.go