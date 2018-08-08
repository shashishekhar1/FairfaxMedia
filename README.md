# FairfaxMedia
Code for Fairfax media challenge.

I have developed the code in go language as it was preferred also it looked like a good starting point. I have tried
to organize the code into different files and also included some basic unit test.

Atom and Postman has been used for development and testing of this solution .

The starting point for the code is 
./src/main/appmain.go

=> Routes have been defined in
./src/main/routes.go

=> Handlers for the routes have been defined in
./src/main/handlers.go

=> I have created date conversion and filtering utility function which is defined in files
./src/main/utils.go

=> Unit test cases have been defined in
./src/main/app_test.go

Go slice has been used for storing the articles and do search and filtering of articles on that slice.
Some error handling has also been implemented in case of search error or filtering errors.

Installation details of the api.

Download the ArticlesProject from GIT repository.
=> cd to the ArticlesProject directory
Use the below command to Execute the project
=> go run .\src\main\appmain.go .\src\main\handlers.go .\src\main\utils.go .\src\main\routes.go
