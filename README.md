# GlobalWebIndex Engineering Challenge

## TL;DR
```
//Update .env file and execute:
//NOTE: Normally .env would not been pushed to the repo
docker-compose up -d --build
```

or

```
//Update config.yml or .env file and execute:
make api-build
make api-run
```

## Introduction

The main purpose of this challenge is to create some basic endpoints for the user and its assets. In order to achieve that the following endpoints have been created:

#### User
```{POST}``` /auth/login: Authenticate  
```{POST}``` /auth/refresh: Refresh token  
```{POST}``` /auth/register: Register  
#### Asset
```{GET}``` /asset: Get all assets  
```{GET}``` /asset/favorite: Get all user's favorite assets  
#### Audience
```{GET}``` /audience: Get all audiences  
```{POST}``` /audience: Create an audience   
```{GET}``` /audience/{audienceId}: Get an audience  
```{PUT}``` /audience/{audienceId}: Updates an audience   
```{DELETE}``` /audience/{audienceId}: Deletes an audience  
```{GET}``` /audience/favorite: Get all user's favorite audiences  
```{POST}``` /audience/{audienceId}/favorite: Make an audience favorite  
```{DELETE}``` /audience/{audienceId}/favorite: Remove an audience from favorites  
#### Chart
```{GET}``` /chart: Get all charts  
```{POST}``` /chart: Create an chart   
```{GET}``` /chart/{chartId}: Get an chart  
```{PUT}``` /chart/{chartId}: Updates an chart   
```{DELETE}``` /chart/{chartId}: Deletes an chart  
```{GET}``` /chart/favorite: Get all user's favorite charts  
```{POST}``` /chart/{chartId}/favorite: Make an chart favorite  
```{DELETE}``` /chart/{chartId}/favorite: Remove an chart from favorites  
#### Chart
```{GET}``` /insight: Get all insights  
```{POST}``` /insight: Create an insight   
```{GET}``` /insight/{insightId}: Get an insight  
```{PUT}``` /insight/{insightId}: Updates an insight   
```{DELETE}``` /insight/{insightId}: Deletes an insight  
```{GET}``` /insight/favorite: Get all user's favorite insights  
```{POST}``` /insight/{insightId}/favorite: Make an insight favorite  
```{DELETE}``` /insight/{insightId}/favorite: Remove an insight from favorites  

**Navigate to {ip}:{port}/swagger/index.html for more**

## Libraries being used
* [chi](https://github.com/go-chi/chi) - HTTP Router & Auth
* [gorm](https://github.com/go-gorm/gorm) - ORM
* [uuid](https://github.com/google/uuid) - For generating UUIDs
* [envconfig](https://github.com/kelseyhightower/envconfig) - For reading environmental variables
* [swag](https://github.com/swaggo/swag) - For generating swagger docs
