package example

import "fmt"

func ExampleHello_stanley() {
	greeting, err := Hello("Stan")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output: Hello, Stan
}

func ExamplePage() {
	checkIns := map[string]bool{
		"Bob":   true,
		"Alice": false,
		"Eve":   false,
		"Janet": true,
		"John":  false,
		"Susan": true,
	}
	Page(checkIns)

	// Unordered output:
	// Paging John; please see the front desk to check in.
	// Paging Eve; please see the front desk to check in.
	// Paging Alice; please see the front desk to check in.
}
