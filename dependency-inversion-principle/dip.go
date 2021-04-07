// Anotations used :-
// DIP: Dependency Inversion Principle
// HLM: High Level Module
// LLM: Low Level Module

package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// LLM
type RelationShips struct {
	relations []Info
}

// abstraction
type RelationshipBrowser interface {
	FindAllChildren(p *Person) []*Person
}

func (rs *RelationShips) AddParentChildRelationShip(parent, child *Person) {
	rs.relations = append(rs.relations, Info{parent, Parent, child})
	rs.relations = append(rs.relations, Info{child, Child, parent})
}

func (rs *RelationShips) FindAllChildren(p *Person) []*Person {
	children := []*Person{}
	for _, relation := range rs.relations {
		if relation.from.name == p.name && relation.relationship == Parent {
			children = append(children, relation.to)
		}
	}
	return children
}

// HLM
type Research struct {
	relationships RelationShips
}

// this breaks DIP & Reseach HLM depends on
func (r *Research) Investigate(parent Person) {
	childCount := 0
	for _, relation := range r.relationships.relations {
		if relation.from.name == parent.name && relation.relationship == Parent {
			childCount++
			fmt.Printf("Child %d of parent %s is %s.\n", childCount, parent.name, relation.to.name)
		}
	}
}

// HLM
type Research2 struct {
	relationshipBrowser RelationshipBrowser
}

// this adheres (follows) DIP as here HLM does not depends on LLM, instead it depends on abstraction
func (r2 *Research2) Investigate(person Person) {
	for index, child := range r2.relationshipBrowser.FindAllChildren(&person) {
		fmt.Printf("Child %d of parent %s is %s.\n", index+1, person.name, child.name)
	}
}

func main() {
	fmt.Println("DIP")
	parent := Person{name: "John"}
	child1 := Person{name: "Jack"}
	child2 := Person{name: "Matt"}

	relationships := RelationShips{}
	relationships.AddParentChildRelationShip(&parent, &child1)
	relationships.AddParentChildRelationShip(&parent, &child2)

	// this breaks DIP
	fmt.Println("Breaks DIP")
	research := Research{relationships: relationships}
	research.Investigate(parent)

	// this adheres (follows) DIP
	fmt.Println("Follows DIP")
	research2 := Research2{relationshipBrowser: &relationships}
	research2.Investigate(parent)
}
