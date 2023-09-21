#!/bin/bash

VERSION=0.2.8
APPNAME=util

git add .
git commit -m "debug"
#git remote rename origin gitee
git push -u gitee main

#git remote add github git@github.com:clin003/util.git
#git branch -M main
git push -u github main

git tag "v${VERSION}"
git push --tags -u github main
git push --tags -u gitee main