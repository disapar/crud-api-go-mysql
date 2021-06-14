## rest api in go with mysql


create a database in MySQL with the name restapigo
then run the api and it will take care of migrating the people table

**Request all Users**

GET http://localhost:9000/all

**Request a particular user**

GET http://localhost:9000/find/1

**Insert new user**

POST http://localhost:9000/save

**Delete a user**

DELETE http://localhost:9000/delete/1
