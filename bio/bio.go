package bio

import "fmt"

type Node struct {
	Key   string
	Value string
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) PrintBio() {
	fmt.Printf("Bio Information\n")
	current := ll.Head
	for current != nil {
		fmt.Printf("%s: %s\n", current.Key, current.Value)
		current = current.Next
	}
}

func BimaBio() LinkedList {
	return LinkedList{
		Head: &Node{
			Key:   "Name",
			Value: "Banggi Bima Edriantino",
			Next: &Node{
				Key:   "Username",
				Value: "banggibima",
				Next: &Node{
					Key:   "Location",
					Value: "Malang, Indonesia",
					Next: &Node{
						Key:   "Instagram",
						Value: "instagram.com/edriantino",
					},
				},
			},
		},
	}
}
