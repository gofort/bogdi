#!/usr/bin/env bash
hugo -D --baseURL https://bogdi.xyz/ && scp -r ./public/* cloud-dev:~/hugo
