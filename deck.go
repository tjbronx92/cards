package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, cardx := range d {
		fmt.Println(i, cardx)
	}
}
