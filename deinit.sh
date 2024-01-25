#!/bin/bash

echo ${1}
mv ./${1} ${1}_tmp
git submodule deinit ${1}
git rm --cached ${1}
mv ${1}_tmp ${1}
git add ${1}