package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

func printProducts(products []Product, prefix string) {
	fmt.Printf("%s Products:- \n", prefix)
	for _, product := range products {
		fmt.Println(product.name)
	}
}

// breaks ocp
type Filter struct{}

func (*Filter) filterGreenProducts(products []Product) []Product {
	var filteredProducts []Product
	for _, product := range products {
		if product.color == green {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}
func (*Filter) filterLargeProducts(products []Product) []Product {
	var filteredProducts []Product
	for _, product := range products {
		if product.size == large {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

// now below filter follows ocp
type Specification interface {
	isSatisfied(product *Product) bool
}

type ColorSpec struct {
	color Color
}

func (cs ColorSpec) isSatisfied(product *Product) bool {
	return cs.color == product.color
}

type SizeSpec struct {
	size Size
}

func (s SizeSpec) isSatisfied(product *Product) bool {
	return s.size == product.size
}

type AndSpec struct {
	first, second Specification
}

func (as AndSpec) isSatisfied(product *Product) bool {
	return as.first.isSatisfied(product) && as.second.isSatisfied(product)
}

type OrSpec struct {
	first, second Specification
}

func (as OrSpec) isSatisfied(product *Product) bool {
	return as.first.isSatisfied(product) || as.second.isSatisfied(product)
}

type BetterFilter struct{}

func (*BetterFilter) filterProducts(products []Product, spec Specification) []Product {
	var filteredProducts []Product
	for _, product := range products {
		if spec.isSatisfied(&product) {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func main() {
	apple := Product{"Apple", green, small}
	house := Product{"House", blue, large}
	tree := Product{"Tree", green, large}

	products := []Product{apple, house, tree}
	printProducts(products, "All")

	// breaks ocp
	fmt.Println(" --- Output from filter which breaks OCP ---")
	filter := Filter{}
	printProducts(filter.filterGreenProducts(products), "Green")
	printProducts(filter.filterLargeProducts(products), "Large")

	// follows ocp
	fmt.Println(" --- Output from filter which follows OCP ---")
	betterFilter := BetterFilter{}
	printProducts(betterFilter.filterProducts(products, ColorSpec{green}), "Green")
	printProducts(betterFilter.filterProducts(products, SizeSpec{large}), "Large")
	printProducts(betterFilter.filterProducts(products, AndSpec{ColorSpec{green}, SizeSpec{large}}), "Green and Large")
	printProducts(betterFilter.filterProducts(products, OrSpec{ColorSpec{green}, SizeSpec{large}}), "Green or Large")
}
