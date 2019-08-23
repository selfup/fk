#!/usr/bin/env bash

set -e

cat .gitignore | grep -q "fk.exe"

if [[ $? == "0" ]]
then
    go install

    fk ls . 
fi
