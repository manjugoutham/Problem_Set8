/* I received following json data as Apple Pie Recipe ingredients
{
"Thinly Sliced Peeled Apples" : "6 Cups",
"sugar" : "3/4 cup",
"flour" : "2 tablespoons",
"cinamon" : "1/4 teaspoon",
"nutmeg" : "1/8 tablespoon",
"lemon juice": "1 tablespoon"
}
Please convert this into a map and print the same as the key value pair in each line. Write this data in a file. */

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	ingredientsJson := `
		{
			"Thinly Sliced Peeled Apples" : "6 Cups",
			"sugar" : "3/4 cup",
			"flour" : "2 tablespoons",
			"cinamon" : "1/4 teaspoon",
			"nutmeg" : "1/8 tablespoon",
			"lemon juice": "1 tablespoon"
		}
	`

	var ingredientsMap map[string]string
	err := json.Unmarshal([]byte(ingredientsJson), &ingredientsMap)
	if err != nil {
		fmt.Println(err)
	}

	// map key and value
	for key, value := range ingredientsMap {
		fmt.Printf("%s: %s\n", key, value)
	}

	file, err := os.Create("applepie.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for key, value := range ingredientsMap {
		fmt.Fprintf(file, "%s: %s\n", key, value)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Applepie file is cretaed successfully...")
}
