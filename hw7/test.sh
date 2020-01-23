#!/bin/env bash

res1=`./go-envdir ./env env | grep FOO`
res2=`./go-envdir ./env env | grep BAR`

if [[ "$res1" == "FOO=HELLO" ]] &&  [[ "$res2" == "BAR=WORLD" ]]; then
  echo "Ok"
else
  echo "Not ok"
fi