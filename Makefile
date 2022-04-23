N := pq
S := 10

update:
	curl -X POST localhost:3000/players/$(N)/$(S) -w "\n"

get:
	curl localhost:3000/players/$(N) -w "\n"

league:
	curl localhost:3000/league -w "\n"
