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

- [x] Use gin-oauth2 to expose a Oauth redirection and callback API. 
- [x] Add the login frontend page. 
- [] Install MongoDB and Redis. 
- [x] Write the session controller and user controller. 
- [x] Write the quest controller.  