#!/bin/bash

echo "Enter the repo url: "
read repo_url

echo "Enter app name: "
read app_name

echo "Enter docker image name: "
read docker_image_name

grep -rl 'github.com/yoshino-s/go-template' . | xargs sed -i "s|github.com/yoshino-s/go-template|$repo_url|g"
sed -i "s|ghcr.io/yoshino-s/go-template|$docker_image_name|g" Taskfile.yml
grep -rl 'go-template' . | xargs sed -i "s|go-template|$app_name|g"

rm -rf init.sh