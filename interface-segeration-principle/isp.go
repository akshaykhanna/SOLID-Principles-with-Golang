package main

import (
	"fmt"
)

type Document struct {
	name string
}

// this interface voids ISP
type Machine interface {
	Print(d Document)
	Scan(d Document)
	Fax(d Document)
}

type MultiFuncPrinter struct {
}

func (mfp *MultiFuncPrinter) Print(d Document) {
	fmt.Println("Print doc")
}

func (mfp *MultiFuncPrinter) Scan(d Document) {
	fmt.Println("Scan doc")
}

func (mfp *MultiFuncPrinter) Fax(d Document) {
	fmt.Println("Fax doc")
}

type OldFashionPrinter struct{}

func (ofp *OldFashionPrinter) Print(d Document) {
	fmt.Println("Print doc")
}

func (ofp *OldFashionPrinter) Scan(d Document) {
	fmt.Println("Can't scan !!!")
}

func (ofp *OldFashionPrinter) Fax(d Document) {
	fmt.Println("Can't fax")
}

// better approach segerate interface as suggested by ISP
type Printer interface {
	Print(d Document)
}

type JustPrinter struct{}

func (jp *JustPrinter) Print(d Document) {
	fmt.Println("Just print doc name " + d.name)
}

type Scanner interface {
	Scan(d Document)
}

type JustScan struct{}

func (jp *JustScan) Scan(d Document) {
	fmt.Println("Just scan doc name " + d.name)
}

type Fax interface {
	Fax(d Document)
}

type JustFax struct{}

func (jp *JustFax) Fax(d Document) {
	fmt.Println("Just fax doc name " + d.name)
}

type Machine2 interface {
	Printer
	Scanner
}

type MultiFuncPrinter2 struct {
	printer Printer
	scanner Scanner
}

func (mfp2 *MultiFuncPrinter2) Print(d Document) {
	mfp2.printer.Print(d)
}

func (mfp2 *MultiFuncPrinter2) Scan(d Document) {
	mfp2.scanner.Scan(d)
}

func main() {
	fmt.Println("ISP")
	myDoc := Document{name: "TestDoc"}
	multiFuncPrinter2 := &MultiFuncPrinter2{printer: &JustPrinter{}, scanner: &JustScan{}}
	multiFuncPrinter2.Print(myDoc)
	multiFuncPrinter2.Scan(myDoc)
}
