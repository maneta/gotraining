// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import "fmt"

func main() {

	// Declare and make a map of integer type values.
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	numbers["four"] = 4
	numbers["five"] = 5
	for key, value := range numbers {
		fmt.Printf("Key: %s Value: %d\n", key, value)
	}
}
