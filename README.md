### Weather App Backend

- **.env.example** file is added to specify required contents for .env file
- To run this app an **MYSQL** instance in required, with a database created with same name mentioned in **.env** file
- To download all the dependencies run: `go mod download`
- After the above steps, locate the root director then run `go run main.go`
- Postman collection(**weather-app.postman_collection.json**) is added at the root directory
- Apart from login & registration route all routes are protected by JWT middleware
- To get a valid token, create an account with registration route, then at the successful response in login route JWT token can be obtained. 