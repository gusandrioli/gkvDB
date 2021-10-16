#!/bin/bash

if [[ -z "$1" ]]; then
    echo "No arguments provided"
elif [[ "$1" == "all" ]]; then
    rm -rf ./.tmp
else
    rm -rf ./.tmp/$1.db
fi
