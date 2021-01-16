#!/usr/bin/env bash
rm -rf ./public
hugo -D --baseURL https://bogdi.xyz/ && scp -r ./public/* cloud-dev:~/hugo
