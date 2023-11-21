# Changelog
## 2023-10-30 @derchrischkya
- added `post-helm-job.yaml` to add inital client-id and client-secret after the openid is up and running
- added `libary.postman_collection.json` to test the api endpoints

## 2023-10-01 @derchrischkya
- added inital_containers folder to store components required for the initial setup of the project

## 2023-09-21 @derchrischkya
- added startup shell script for persistent-database infrastructure `database_container/db/*`
- updated Dokerfile for Copy folder to container on build-path
- API updated - added working capture-book endpoint functionality

## 2023-09-17 @derchrischkya
- changed configuration of postgres db to persistent storage using

## 2023-09-12 @derchrischkya
- inital repository
- added api_container with authentication.go and mockup endpoints
- added working openid_container configurations
