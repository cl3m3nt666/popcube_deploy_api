# popcube_deploy_api

[![Build Status](https://travis-ci.com/cl3m3nt666/popcube_deploy_api.svg?token=pQ5JuFHLtUEwNb123zaH&branch=master)](https://travis-ci.com/cl3m3nt666/popcube_deploy_api)

## Descriptions

API to deploy an organization for the popcube project.

## Routes

| Methodes | path |
| :------- |:------|
| GET | / |
| POST | /deploy/stack/create |
| POST | /deploy/stack/{name}/remove |
| POST | /deploy/stack/{name}/status |

## Environment variable

| Name | default value |
| :-------: |:------|
| GOPATH | $GOPATH:/go/api |
| GOBIN | $GOPATH/bin |
| XTOKEN | 1234 |
| DEFAULT_DOMAIN | popcube.xyz |
| DEFAULT_DATABASE | popcube |
| BASE_NAME_HOST_DB | database |
| DEFAULT_ORG_PATH | /organisation |
| ORGANISATION_TEMPLATE | /organisation_template |
| DOCKER_USER | |
| DOCKER_PWD | |
| DOCKER_REGISTRY | |
