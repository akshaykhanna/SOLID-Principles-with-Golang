package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Journals struct {
	entries []string
}

func (j *Journals) addJournal(journal string) {
	j.entries = append(j.entries, journal)
}

func (j *Journals) remove(journal string) {
	var newJournals []string
	for _, entry := range j.entries {
		if entry != journal {
			newJournals = append(newJournals, entry)
		}
	}
	j.entries = newJournals
}

func (j *Journals) String() string {
	return strings.Join(j.entries, "\n")
}

type Persistence struct {
	lineSep string
}

func (p *Persistence) saveFile(j Journals, fileName string) {
	_ = ioutil.WriteFile(fileName,
		[]byte(strings.Join(j.entries, p.lineSep)), 0644)
}

func main() {
	myJournal := Journals{}
	myJournal.addJournal("I took 4 interviews today")
	myJournal.addJournal("I am practicing SOLID principals")
	myJournal.addJournal("Sunset")
	fmt.Println(myJournal.String())
	persistence := Persistence{"\n"}
	persistence.saveFile(myJournal, "journals")
}
