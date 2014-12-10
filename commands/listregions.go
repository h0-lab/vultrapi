package commands

import (
	"fmt"
	. "github.com/stephan83/vultrapi/clients"
	"github.com/stephan83/vultrapi/requests"
	"os"
	"sort"
	"text/tabwriter"
)

type listRegions struct{}

func NewListRegions() Command {
	return listRegions{}
}

func (_ listRegions) Desc() string {
	return "List all available regions."
}

func (_ listRegions) Args() string {
	return ""
}

func (_ listRegions) NeedsKey() bool {
	return false
}

func (_ listRegions) PrintOptions() {
	fmt.Println("None.")
}

func (_ listRegions) Exec(c Client, args []string, key string) (err error) {
	r, err := requests.GetRegions(c)
	if err != nil {
		return
	}

	a := r.Array()
	sort.Sort(a)

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(w, "ID\tNAME\tCONTINENT\tCOUNTRY\tSTATE")

	for _, v := range a {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", v.Id, v.Name,
			v.Continent, v.Country, v.State)
	}

	w.Flush()

	return
}
