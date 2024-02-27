#!/bin/bash

echo "${1}"  "${1}_tmp"
mv "./${1}/${1}_tmp"  "${1}"
git add "${1}"