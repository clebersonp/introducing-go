package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Person struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}

func (p *Person) String() string {
	return p.Name
}

type Result struct {
	person *Person
	err    error
}
type Controller struct {
	m       sync.Mutex
	results []*Result
}

func (s *Controller) Add(person *Person, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.results = append(s.results, &Result{person, err})
}

func main() {
	var wg sync.WaitGroup
	names := [...]string{
		"NAME1", "NAME2", "NAME3", "NAME4", "NAME5", "NAME6", "NAME7", "NAME8", "NAME9", "NAME10", "NAME11", "NAME12",
		"NAME13", "NAME14", "NAME15", "NAME16", "NAME17", "NAME18", "NAME19", "NAME20", "NAME21", "NAME22", "NAME23",
		"NAME24", "NAME25", "NAME26", "NAME27", "NAME28", "NAME29", "NAME30", "NAME31", "NAME32", "NAME33", "NAME34",
		"NAME35", "NAME36", "NAME37", "NAME38", "NAME39", "NAME40", "NAME41", "NAME42", "NAME43", "NAME44", "NAME45",
		"NAME46", "NAME47", "NAME48", "NAME49", "NAME50", "NAME51", "NAME52", "NAME53", "NAME54", "NAME55", "NAME56",
		"NAME57", "NAME58", "NAME59", "NAME60", "NAME61", "NAME62", "NAME63", "NAME64", "NAME65", "NAME66", "NAME67",
		"NAME68", "NAME69", "NAME70", "NAME71", "NAME72", "NAME73", "NAME74", "NAME75", "NAME76", "NAME77", "NAME78",
		"NAME79", "NAME80", "NAME81", "NAME82", "NAME83", "NAME84", "NAME85", "NAME86", "NAME87", "NAME88", "NAME89",
		"NAME90", "NAME91", "NAME92", "NAME93", "NAME94", "NAME95", "NAME96", "NAME97", "NAME98", "NAME99", "NAME100",
	}
	controller := &Controller{results: make([]*Result, 0, len(names))}
	for _, name := range names {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if name == "NAME3" {
				controller.Add(nil, nil)
				return
			}
			if name == "NAME4" {
				controller.Add(nil, fmt.Errorf("error"))
				return
			}
			strOrder, _ := strings.CutPrefix(name, "NAME")
			order, _ := strconv.Atoi(strOrder)
			controller.Add(&Person{name, order}, nil)
		}()
	}

	wg.Wait()
	people := make([]*Person, 0, len(controller.results))
	for _, result := range controller.results {
		if result.err != nil {
			log.Println(result.err)
		}
		if result.person != nil {
			people = append(people, result.person)
		}
	}

	//sort.Slice(people, func(i, j int) bool {
	//	return people[i].Order < people[j].Order
	//})
	slices.SortFunc(people, func(a, b *Person) int {
		return cmp.Compare(a.Order, b.Order)
	})
	bytesJson, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	fmt.Printf("qty of people: %d, capacity: %d, people: %v\n", len(people), cap(people), string(bytesJson))

}
