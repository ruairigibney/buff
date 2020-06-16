# buff
Buff Video Stream API Project

## Getting Started

The following instructions will step you through setting up and running the Buff Video Stream API Project

## Prerequisites

Ensure you have docker-compose and Go (version 12 or later) set-up before continuing

## Installing

1. clone this repo to your local - `git clone https://github.com/ruairigibney/buff.git`
2. use docker-compose to create a mysql Docker DB container - `sudo docker-compose up`
3. copy pre-push hooks to .git/hooks/pre-push - `cp pre-push .git/hooks/`

## Restore sample DB (for dev testing)

Run the following from the project root folder to restore sample DB for testing

```cat buffdump.sql | sudo docker exec -i buff_db_1 /usr/bin/mysql -u root --password=buffroot1 buffdb```

## Starting the service

To start the service, run the below command. The service will start running on port 8080.

```go run cmd/main.go```

## APIs

The service has two API routes. Details are below.

### ```GET http://localhost:8080/api/videoStreams```

This endpoint will return video streams and any question IDs associated with a stream. It has two optional query params - `limit` and `offset`. These are used for pagination. 

Examples:

```
GET http://localhost:8080/api/videoStreams?limit=10
GET http://localhost:8080/api/videoStreams?limit=2&offset=2
```

Example response:

```
[
  {
    "ID": 1,
    "title": "Buff Example Stream 1",
    "createdTime": 1591728525,
    "updatedTime": 1591728525,
    "questionIDs": [
      1,
      2
    ]
  },
  {
    "ID": 2,
    "title": "Buff Example Stream 2",
    "createdTime": 1591732125,
    "updatedTime": 1591732125
  },
  {
    "ID": 3,
    "title": "Buff Example Stream 3",
    "createdTime": 1591735725,
    "updatedTime": 1591735725
  }
]
```

### ```GET http://localhost:8080/api/videoStreams/questions/:id```

This endpoint will return a question with all associated answers, and will mark which answer is correct. `id` must be specified and must be a valid question ID.

Example: 

```http://localhost:8080/api/videoStreams/questions/1```

Example response:

```
{
  "ID": 1,
  "VideoStreamID": 1,
  "Text": "Is this an example question?",
  "Answers": [
    {
      "ID": 1,
      "Text": "This is a correct answer",
      "CorrectAnswer": true
    },
    {
      "ID": 2,
      "Text": "This is an incorrect answer"
    }
  ]
}
```

## Running the tests

To run unit tests locally, run the below command. This will give you the current code coverage for the project. If a breakdown of the tests being run is reuqired, add `-v` to the command string.

```go test ./... -cover```
