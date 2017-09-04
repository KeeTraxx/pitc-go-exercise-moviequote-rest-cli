package main

import "fmt"

// MovieQuote represents a quote from a movie
// TODO: Analyze http://pitc-moviequote.ose3-lab.puzzle.ch/v1/moviequotes/random
// and create a proper struct
type MovieQuote struct {
}

func main() {

	// TODO: Make a http GET Request to http://pitc-moviequote.ose3-lab.puzzle.ch/v1/moviequotes/random
	// HINT: Use http.Get

	// TODO: Read HTTP response body
	// HINT: ioutil.ReadAll can read streams

	// TODO: Convert json to your struct
	// HINT: Use json.Unmarshal

	// TODO Print out the result
	fmt.Printf("%v\n", "Hasta la vista baby")

}
