package main

import "fmt"

func main() {
	// Creating Maps Using var and :=
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  	fmt.Printf("a\t%v\n", a)
  	fmt.Printf("b\t%v\n", b)

	var c = make(map[string]string) // The map is empty now
  	c["brand"] = "Ford"
  	c["model"] = "Mustang"
	c["year"] = "1964"
                                 // a is no longer empty
  	d := make(map[string]int)
  	d["Oslo"] = 1
  	d["Bergen"] = 2
  	d["Trondheim"] = 3
  	d["Stavanger"] = 4

  	fmt.Printf("c\t%v\n", c)
  	fmt.Printf("d\t%v\n", d)


	// Creating an Empty Map
	var e = make(map[string]string)
	var f map[string]string

	fmt.Println("e", e == nil)
	fmt.Println("f", f == nil)


	// Updating and Adding Map Elements
	var g = make(map[string]string)
	g["brand"] = "Ford"
	g["model"] = "Mustang"
	g["year"] = "1964"

	fmt.Println("g Before", g)

	g["year"] = "1970" // Updating an element
	g["color"] = "red" // Adding an element

	fmt.Println("g After", g)


	// Remove Element from Map
	delete(g, "brand")
	fmt.Println("g After Delete", g)


	// Check For Specific Elements in a Map
	var h = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

	val1, ok1 := h["brand"] // Checking for existing key and its value
	val2, ok2 := h["color"] // Checking for non-existing key and its value
	val3, ok3 := h["day"]   // Checking for existing key and its value
	_, ok4 := h["model"]    // Only checking for existing key and not its value

	fmt.Println("val1, ok1", val1, ok1)
	fmt.Println("val2, ok2", val2, ok2)
	fmt.Println("val3, ok3", val3, ok3)
	fmt.Println("_, ok4", ok4)


	// Maps Are References
	var i = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	j := i

	fmt.Println("i", i)
	fmt.Println("j", j)

	j["year"] = "1970"
	fmt.Println("After change to j:")

	fmt.Println("i", i)
	fmt.Println("j", j)

	
	// Iterating Over Maps
	a2 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range a2 {
		fmt.Printf("a2 --> %v : %v, ", k, v)
	}
	fmt.Println()


	// Iterate Over Maps in a Specific Order
	a3 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b3 []string				// defining the order
	b3 = append(b3, "one", "two", "three", "four")

	for k, v := range a3 {        // loop with no order
		fmt.Printf("a3 --> %v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b3 {  // loop with the defined order
		fmt.Printf("b3 --> %v : %v, ", element, a3[element])
	}
}