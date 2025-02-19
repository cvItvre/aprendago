package cap17

import (
	"fmt"
	"slices"
	"sort"
)

type Agent struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type sortByAge []Agent

func (a sortByAge) Len() int           { return len(a) }
func (a sortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a sortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type sortByLastName []Agent

func (a sortByLastName) Len() int           { return len(a) }
func (a sortByLastName) Less(i, j int) bool { return a[i].Last < a[j].Last }
func (a sortByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func ex05() {
	a1 := Agent{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	a2 := Agent{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	a3 := Agent{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	agents := []Agent{a1, a2, a3}

	fmt.Println(agents)

	// your code goes here

	for _, agent := range agents {
		slices.Sort(agent.Sayings)
	}

	sort.Sort(sortByAge(agents))
	fmt.Println(agents)

	sort.Sort(sortByLastName(agents))
	fmt.Println(agents)

	for _, agent := range agents {
		fmt.Println(agent.First, agent.Last)
		for _, saying := range agent.Sayings {
			fmt.Printf("\t- '%v'\n", saying)
		}
		fmt.Println()
	}

}
