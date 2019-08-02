# Find my mall app
Simple Go application

# Running on local

Get dependencies with:
* `dep ensure`

Start mongo service:
* `docker compose up -d`

Build repo:
* `go build`

Start app:
* `./find-my-mall`

Run the tests:
* `go test -v`

Feeding mongodb with data:
* Go into the consume subfolder`cd consumer`

* Run consumer `go test -v`

Using the backend api:

See [this](swagger.yaml) for further information
