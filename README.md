# Project Nyxeon

- Golang backend 
- VueJS frontend 
- MongoDB 
- Redis to store session token and user id mapping 

Database schemas:

- User:
    - Name, avatar from Github, array of quest IDs
- Quest:
    - Category
    - Content 
    - Date 

## Work in progress 

1. Use gin-oauth2 to expose a Oauth redirection and callback API. 
2. Add the login frontend page. 