# cts-policy-tes


Create .env file
Create a .env file in the root of your project directory and define the MongoDB URI there:

MONGODB_URI=mongodb://localhost:27017

Running the Application
Make sure the .env file is present in the root directory of your project. When you run the Go application, it will automatically load the MongoDB URI from the .env file.

go run main.go
