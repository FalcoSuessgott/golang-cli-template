#!/usr/bin/env bash

set -euo pipefail

read -rp "GitHub Username: " user
read -rp "Projectname: " projectname

git clone git@github.com:FalcoSuessgott/golang-cli-template.git "$projectname"
cd "$projectname"
rm -rf .git
find . -type f -exec sed -i "s/golang-cli-template/$projectname/g" {} +
find . -type f -exec sed -i "s/FalcoSuessgott/$user/g" {} +
git init
git add .
git commit -m "initial commit"
git remote add origin "git@github.com:$user/$projectname.git"

echo "template successfully installed."
