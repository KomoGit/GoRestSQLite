# GoRestSQLite
REST Api for student registry that allows you to create,read,update and delete data from a local SQLite Database hosted on localhost:8080. At the moment the data is retrieved from a terminal so run the application with one. In the future I will add the front-end GUI application for ease of access to data.

## REST Links
- CREATE 
- ```/api/students/{fName},{email},{pNum}``` Allows you to insert a new user into the database.
- READ ("GET")
- ```/api/students``` Allows you to get all data from database.
- ```/api/students/{id}``` Allows you to get specific user with uuid.
- UPDATE ("PUT")
- ```/api/students/update/name/{id},{fName}/``` Update Fullname
- ```/api/students/update/email/{id},{email}/``` Update email
- ```/api/students/update/number/{id},{pNum}/``` Update phone number
- DELETE ("DELETE")
- ```/api/students/{id}``` Delete student with id.
- ```/api/students/``` Delete all from the database.

## I have used these open-source projects to build mine.
| Names | Links |
| ------ | ------ |
| go-sqlite3 | https://github.com/mattn/go-sqlite3[PlDb] |
| uuid | github.com/google/uuid[PlGh] |
| Gorilla Mux | github.com/gorilla/mux[PlGd] |
| Dillinger.io | dillinger.io [PlOd] |
