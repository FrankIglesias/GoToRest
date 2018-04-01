# RESTful API Example with golang
This is simple example restful api server only with **gorilla/mux**.  
For simplifying code, this example uses a mock database that is `map[string]interface{}`

## Install and Run
```shell
$ go build
$ ./GoToRest
```

## API Endpoint
- http://localhost:3000/api/v1/students
    - `GET`: get list of students
    - `POST`: create a student
- http://localhost:3000/api/v1/student/{name}
    - `GET`: gets a student
    - `DELETE`: remove a student
