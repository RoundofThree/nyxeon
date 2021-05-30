# Project Nyxeon

This is my first project using Golang. The 

## Tech stack 

- Golang + Gin Framework for Backend
- VueJS for Frontend 
- MongoDB as primary database
- Redis to store session 

## Database collections

- User:
  - Email
  - Array of quest object IDs. 
- Quest:
  - Array of categories
  - Content
  - Date

## Configuration 

Edit the `config/<environment>.yaml` file and run the server with the `m` flag set. 

```bash
./nyxeon -m development  # to use the configs in config/development.yaml 
```

You must provide:

- `server`: The IP address and port the server will listen to. 
- `client.domain`: The domain of the front end. This is set in the cookie. 
- `cors.origin`
- `mongo`: The database and credentials 
- `redis`: The Redis cache and credentials 
- `oauth`: The client ID and client secret and scopes. Set the redirect url after callback is successful. 

## Deployment 

## Work in progress 

- [x] Use gin-oauth2 to expose a Oauth redirection and callback API. 
- [x] Add the login frontend page. 
- [x] Install MongoDB and Redis. 
- [x] Write the session controller and user controller. 
- [x] Write the quest controller.  
- [x] Add calendar heatmap. 
- [x] Add state to VueJS.
- [x] Add pie chart. 
- [x] Lint code. 
- [x] Add configuration instructions.
- [x] Deploy. 
- [x] Write Dockerfile. 