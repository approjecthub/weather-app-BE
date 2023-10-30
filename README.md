### Weather App Backend

- To run this app an **MYSQL** instance in required, with a database created with same name mentioned in **.env** file
- **.env.example** file is added to specify required contents for .env file
- Postman collection(**weather-app.postman_collection.json**) is added at the root directory
- Apart from login & registration route all routes are protected by JWT middleware
- To get a valid token, create an account with registration route, then at the successful response in login route JWT token can be obtained. 
