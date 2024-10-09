package main

import (
	"fmt"
	"time"
)

func main() {

	userTtlMap := New(10)

	userTtlMap.Put("user:1", `{"id": 1, "name": "John", "age": 25, "email": "j@j.com"}`, time.Second*5)
	userTtlMap.Put("user:2", `{"id": 2, "name": "Ferry", "age": 18, "email": "f@f.com"}`, time.Second*10)
	userTtlMap.Put("user:3", `{"id": 3, "name": "Mary", "age": 31, "email": "mary@mary.com"}`, time.Second*2)
	userTtlMap.Put("user:4", `{"id": 4, "name": "Mark", "age": 15, "email": "mark@mark.com"}`, time.Second*3)

	if it, ok := userTtlMap.Get("user:1"); ok {
		fmt.Printf("User:1: %s\n", it)
	} else {
		fmt.Println("User:1 not found")
	}
	if it, ok := userTtlMap.Get("user:2"); ok {
		fmt.Printf("User:2: %s\n", it)
	} else {
		fmt.Println("User:2 not found")
	}
	if it, ok := userTtlMap.Get("user:3"); ok {
		fmt.Printf("User:3: %s\n", it)
	} else {
		fmt.Println("User:3 not found")
	}
	if it, ok := userTtlMap.Get("user:4"); ok {
		fmt.Printf("User:4: %s\n", it)
	} else {
		fmt.Println("User:4 not found")
	}

	time.Sleep(time.Second * 8)

	if it, ok := userTtlMap.Get("user:1"); ok {
		fmt.Printf("User:1: %s\n", it)
	} else {
		fmt.Println("User:1 not found")
	}
	if it, ok := userTtlMap.Get("user:2"); ok {
		fmt.Printf("User:2: %s\n", it)
	} else {
		fmt.Println("User:2 not found")
	}
	if it, ok := userTtlMap.Get("user:3"); ok {
		fmt.Printf("User:3: %s\n", it)
	} else {
		fmt.Println("User:3 not found")
	}
	if it, ok := userTtlMap.Get("user:4"); ok {
		fmt.Printf("User:4: %s\n", it)
	} else {
		fmt.Println("User:4 not found")
	}

	time.Sleep(time.Second * 15)

	if it, ok := userTtlMap.Get("user:1"); ok {
		fmt.Printf("User:1: %s\n", it)
	} else {
		fmt.Println("User:1 not found")
	}
	if it, ok := userTtlMap.Get("user:2"); ok {
		fmt.Printf("User:2: %s\n", it)
	} else {
		fmt.Println("User:2 not found")
	}
	if it, ok := userTtlMap.Get("user:3"); ok {
		fmt.Printf("User:3: %s\n", it)
	} else {
		fmt.Println("User:3 not found")
	}
	if it, ok := userTtlMap.Get("user:4"); ok {
		fmt.Printf("User:4: %s\n", it)
	} else {
		fmt.Println("User:4 not found")
	}

	userTtlMap.Put("user:1", `{"id": 1, "name": "John", "age": 25, "email": "j@j.com"}`, time.Second*5)
	userTtlMap.Put("user:2", `{"id": 2, "name": "Ferry", "age": 18, "email": "f@f.com"}`, time.Second*10)
	userTtlMap.Put("user:3", `{"id": 3, "name": "Mary", "age": 31, "email": "mary@mary.com"}`, time.Second*2)
	userTtlMap.Put("user:4", `{"id": 4, "name": "Mark", "age": 15, "email": "mark@mark.com"}`, time.Second*5)

	if it, ok := userTtlMap.Get("user:1"); ok {
		fmt.Printf("User:1: %s\n", it)
	} else {
		fmt.Println("User:1 not found")
	}
	if it, ok := userTtlMap.Get("user:2"); ok {
		fmt.Printf("User:2: %s\n", it)
	} else {
		fmt.Println("User:2 not found")
	}
	if it, ok := userTtlMap.Get("user:3"); ok {
		fmt.Printf("User:3: %s\n", it)
	} else {
		fmt.Println("User:3 not found")
	}
	if it, ok := userTtlMap.Get("user:4"); ok {
		fmt.Printf("User:4: %s\n", it)
	} else {
		fmt.Println("User:4 not found")
	}

	userTtlMap.Delete("user:1")
	userTtlMap.Delete("user:2")
	userTtlMap.Delete("user:3")

	if it, ok := userTtlMap.Get("user:1"); ok {
		fmt.Printf("User:1: %s\n", it)
	} else {
		fmt.Println("User:1 not found")
	}
	if it, ok := userTtlMap.Get("user:2"); ok {
		fmt.Printf("User:2: %s\n", it)
	} else {
		fmt.Println("User:2 not found")
	}
	if it, ok := userTtlMap.Get("user:3"); ok {
		fmt.Printf("User:3: %s\n", it)
	} else {
		fmt.Println("User:3 not found")
	}
	if it, ok := userTtlMap.Get("user:4"); ok {
		fmt.Printf("User:4: %s\n", it)
	} else {
		fmt.Println("User:4 not found")
	}

	time.Sleep(time.Second * 10)

	if it, ok := userTtlMap.Get("user:1"); ok {
		fmt.Printf("User:1: %s\n", it)
	} else {
		fmt.Println("User:1 not found")
	}
	if it, ok := userTtlMap.Get("user:2"); ok {
		fmt.Printf("User:2: %s\n", it)
	} else {
		fmt.Println("User:2 not found")
	}
	if it, ok := userTtlMap.Get("user:3"); ok {
		fmt.Printf("User:3: %s\n", it)
	} else {
		fmt.Println("User:3 not found")
	}
	if it, ok := userTtlMap.Get("user:4"); ok {
		fmt.Printf("User:4: %s\n", it)
	} else {
		fmt.Println("User:4 not found")
	}

	fmt.Println("Done")
}
