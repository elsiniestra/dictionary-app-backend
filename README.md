# Simple dictionary API with Oxford Dictionary integration.

### What about the project?
The simple REST API with the simple idea – get info about the requested word.
By request for info about a word, it refers to Oxford Dictionary data, saves it into DynamoDB to decrease
the count of requests to the external resource, and returns processed data.
Uses DynamoDB as the main database and Redis for caching.

# Set up
0.
- Get credentials from [Oxford Dictionary API](https://developer.oxforddictionaries.com). It has a "freemium" plan, 
   that has to be enough for personal purposes.
- Get AWS connection credentials, also have the "freemium" plan.
1. Fill a `.env` file by the example in `example.env`.
2. Run `make run` or `make build` for local run or `make compose-up` for run using the docker-compose.

Architecture inspired by The Clean Architecture and https://github.com/evt/rest-api-example.

P.S. That's my first project on Golang, so will be so proud of your comments and piece of advice about the 
improvements!