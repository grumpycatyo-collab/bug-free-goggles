## Structure
1.  `core/` - the main directory where all the handlers are located
2. `models/` - location of models

## Launching
Since it's a primitive project, to launch it you will require a simple command:
```go
go run main.go
```
And then as a output you should receive something like:
```bash
Started fetching data every 5 minutes
Started fetching data every 30 seconds
[{1207270196287897 Project 2 project} {1207269063496742 Project 1 project}]
[{1207269034735706 user3123@gmail.com user} {1207269038212633 user313@gmail.com user} {1207269038333720 user312f93@gmail.com user} {1207269339457428 Max Plămădeală user}]
[{1207270196287897 Project 2 project} {1207269063496742 Project 1 project}]
[{1207269034735706 user3123@gmail.com user} {1207269038212633 user313@gmail.com user} {1207269038333720 user312f93@gmail.com user} {1207269339457428 Max Plămădeală user}]
```
At this point I am returning every user and project as a part of a slice since I didn't manage to deal with jsons.

## Handler Structure
In the `core/` directory you will have the main 2 handlers, `GetUsers` and `GetProjects`. Every one of them is just connecting 
to the ASANA API and getting the structures which are after converted to models in go.

### Rate Limiting
In this case, I implemented a primitive rate limiter which is not compatible with massive high loads:
```go
limiter := rate.NewLimiter(15, 2)
	if !limiter.Allow() {
		return nil, errors.New("the API is at capacity, try again later")
	}
```
## Data Extraction
Data Extraction was realized with the help of channels since I wanted a nice structure and a high load compatibility code. You can see the code in `main.go`.
Additionaly, since I didn't have much time realizing the high load extraction, I can explain what I wanted to do. I wanted to make a simple caching that caches the responses and returns them if the returns are remaining the same. As you can see I just started doing it right here:
```go
type CachedItems struct {
	lock     sync.RWMutex
	Projects []models.Project
	Users    []models.User
}
```
I can give you a basic outline of what I've wanted to do here:
```go
ctx := context.Background()

	// initializing cache and fill
	cache := CachedPopularItems{}
	cache.Movies = getPopularMoviesFromDB()
	go func() {
		timer := time.NewTicker(1 * time.Second)
		defer timer.Stop()

		// initializing background job
		for {
			select {
			// refreshing cache
			case <-timer.C:
				movies := getPopularMoviesFromDB()

				// updating cache struct
				cache.lock.Lock()
				cache.Movies = movies
				cache.lock.Unlock()

			// app is terminating
			case <-ctx.Done():
				break
			}
		}
	}()

```
Of course, the example is took from the internet, however as a referrence it was nice.

## Tests
Tests are written in the last 30 minutes of the time so I managed to make the most primitive tests.
First, I wanted to make a benchmarking test using `testing.B`, to see multiple users handling and using WaitGroups. It can be tested directly in the GoLand IDE.
For the rest of the two tests, that check the basic matching of two structs, one which is returned and one which is awaited you should see the following response:
```bash
go test ./tests -v -bench=.
=== RUN   TestProjects
    core_test.go:61: got [{"1207270191287897" "Project 2" "project"} {"1207269063496742" "Project 1" "project"}], wanted [{"1207270196287897" "Project 2" "project"} {"1207269063496742" "Project 1" "project"}]
--- FAIL: TestProjects (1.17s)
=== RUN   TestUsers
--- PASS: TestUsers (0.41s)
```
Thus, as you see, TestUsers has passed the checks but TestProjects hasn't.
Also, for launching tests, you can use the commands in the makefile or simply input:
```makefile
make test
```
## What I would've done if I had more time
1. Implemented caching for high load tolerance.
2. Graceful shutdown and better error handling (panics) for better service managing.
3. Better tests and benchmarking

## What I would've done if I had more time
1. Implemented caching for high load tolerance.
2. Graceful shutdown and better error handling (panics) for better service managing.
3. Better tests and benchmarking
## What I would've done if I had more time
1. Implemented caching for high load tolerance.
2. Graceful shutdown and better error handling (panics) for better service managing.
3. Better tests and benchmarking
4. ## What I would've done if I had more time
1. Implemented caching for high load tolerance.
2. Graceful shutdown and better error handling (panics) for better service managing.
3. Better tests and benchmarking


## What I would've done if I had more time
1. Implemented caching for high load tolerance.
2. Graceful shutdown and better error handling (panics) for better service managing.
3. Better tests and benchmarking
## What I would've done if I had more time
1. Implemented caching for high load tolerance.
2. Graceful shutdown and better error handling (panics) for better service managing.
3. Better tests and benchmarking
4. ## What I would've done if I had more time
1. Implemented caching for high load tolerance.
2. Graceful shutdown and better error handling (panics) for better service managing.
3. Better tests and benchmarking


## What I would've done if I had more time
1. Implemented caching for high load tolerance.
2. Graceful shutdown and better error handling (panics) for better service managing.
3. Better tests and benchmarking
## What I would've done if I had more time
1. Implemented caching for high load tolerance.
2. Graceful shutdown and better error handling (panics) for better service managing.
3. Better tests and benchmarking
4. ## What I would've done if I had more time
1. Implemented caching for high load tolerance.
2. Graceful shutdown and better error handling (panics) for better service managing.
3. Better tests and benchmarking