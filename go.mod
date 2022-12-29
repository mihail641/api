module my_project

go 1.18

require (
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.6
	example.com/projectApiClient v0.0.0
		)
	replace example.com/projectApiClient => ../projectApiClient