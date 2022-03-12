# Movie Query App
_Version: Week 2_

A command line app to query top 10 movies of IMDB movie ranking list.

## Usage
#### Listing
You can list movies by providing "list" command:

```sh
go run main.go list
```

#### Searching
You can search movies by providing "search" command followed by your search query (not case-sensitive):

```sh
go run main.go search pulp fiction
```
If queried movie is found, full movie title and position in the list will be displayed.
If it is not, a "Movie Not Found" message will be displayed.
