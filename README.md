# Rock Paper Scissor

This app is a Rock, Paper, Scissor game which is written in Golang and has a REST API interface created with gorilla/mux.

## Game Options

The options of the game are:

* Number 0 is mapped to Rock
* Number 1 is mapped to Paper
* Number 2 is mapped to Scissor

## Game Rules

* Paper Vs Rock --> PAPER WINS because paper covers rock
* Rock Vs Scissor --> ROCK WINS because rock smashes scissor
* Scissor Vs Paper --> SCISSOR WINS because scissor cuts paper

## How to run locally

```
>> go run main.go

Serving App: Rock Paper Scissor
Rest API: Mux Routers
Running on: http://127.0.0.1:8080/
```

The `init()` initializes the game options. In Golang, if there is an `init()` then that function is called as soon as the app is executed.

For instance: The entry point to the app is `main()`, however, if there is an `init()` then upon executing the app `>> go run main.go` the `init()` will be called first and then the `main()`.

More info can be found here: https://stackoverflow.com/questions/20508356/non-declaration-statement-outside-function-body-in-go

## REST API Routes

The routes are created with the `gorilla/mux` package. The routes are placed in the `handleRequests()` function in `main.go`.

### Play with computer

The route to play with the computer is `/start-game-with-computer` and is mapped to the `startGameWithComputer()`. The player can input their option in the data flag `-d` in the curl command. The computer selects an option at random.


```bash
>> curl -X POST "http://127.0.0.1:8080/start-game-with-computer" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"UserInput\": 0}"

>> curl -X POST "http://127.0.0.1:8080/start-game-with-computer" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"UserInput\": 1}"

>> curl -X POST "http://127.0.0.1:8080/start-game-with-computer" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"UserInput\": 2}"

```

The output of the game is as follows:

````
{
    "PlayerPick":"Rock",
    "ComputerPick":"Paper",
    "Winner":"Computer"
}
````

### Play with Player

The route to play with another player is `/start-game-with-player` and is mapped to the `startGameWithPlayer()`. Both players can input their options in the data flag `-d` in the curl command.

```bash
>> curl -X POST "http://127.0.0.1:8080/start-game-with-player" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"UserInputOne\": 0, \"UserInputTwo\": 1}"

>> curl -X POST "http://127.0.0.1:8080/start-game-with-player" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"UserInputOne\": 2, \"UserInputTwo\": 1}"

```

The output of the game is as follows:

```
{
    "PlayerOnePick":"Scissor",
    "PlayerTwoPick":"Paper",
    "Winner":"PlayerOne"
}
```

### Select between playing with another Player or Computer

Users can select between playing with another player or with a computer. The route to select between a player or computer is under `/select-player-or-computer` and is mapped to the `selectPlayerOrComputer()`.

* Number 1 is mapped to play with another player
* Number 2 is mapped to play with a computer

```bash
>> curl -X POST "http://127.0.0.1:8080/select-player-or-computer" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"UserInput\": 1}"

>> curl -X POST "http://127.0.0.1:8080/select-player-or-computer" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"UserInput\": 2}"
```

## Dependencies

The dependencies of the app are in the `go.mod` file. It can also be installed via the following command.

```bash
>> go get github.com/gorilla/mux
```

## Testing

The tests are in the `main_test.go` file and can be executed with the following command:

```
>> go test -v

=== RUN   TestGetPlayerSelection
--- PASS: TestGetPlayerSelection (0.00s)
=== RUN   TestPlayGame
--- PASS: TestPlayGame (0.00s)
PASS
ok      github.com/StephenDsouza90/RockPaperScissor     0.195s
```

## Initialized module

To initialize the module, run the following command:

```bash
>> go mod init
```

## Docker

Dockerfile contains a multi-stage build process to reduce the size of the image.

In the first stage of the multi build, the Go program is compiled and a Binary is produce. All the dependencies and the main program are compiled to a single Binary file that can be executed. In the second stage, the Binary is copied to the new image. This will reduce the size of the final image because it will contain the Binary.

To build the Docker image, run the following command:

```bash
>> docker build -t rock-paper-scissor .
```

To run the Docker image, run the following command:

```bash
>> docker run -it --rm -p 8080:8080 rock-paper-scissor
```
