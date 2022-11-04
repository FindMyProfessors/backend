#!/bin/bash

if [ "$1" = "development" ] || [ "$1" = "main" ]; then
  git submodule foreach "git checkout $1 && git pull || :"
else
  printf "You must supply development or main as arguments to the script\n"
fi