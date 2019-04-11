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
* `./find_my_mall`

Run the tests:
* `go test -v`

Feeding mongodb with data:
* Go into the consume subfolder`cd consumer`

* Run consumer `go test -v`

Using the backend api:

**GET** -> `localhost/shoppingMalls?city={CITY_NAME}&score={SCORE}&magaza={SHOPS}`

e.g. `http://localhost:3000/shoppingMalls?city=istanbul&score=1,0&magaza=Atasay,Gratis`

**POST** -> `localhost/shoppingMalls`

**PUT** -> `localhost/shoppingMalls`

**GET** -> `localhost/shoppingMalls/{id}`

**DELETE** -> `localhost/shoppingMalls/{id}`
