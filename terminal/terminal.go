package main

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {

	t := table.NewWriter()
	t.SetCaption("Users")

	t.AppendHeader(table.Row{"#", "Name", "Occupation"})
	t.AppendRow(table.Row{1, "John Doe", "gardener"})
	t.AppendRow(table.Row{2, "Roger Roe", "driver"})
	t.AppendRows([]table.Row{{3, "Paul Smith", "trader"},
		{4, "Lucy Smith", "teacher"}})

	fmt.Println(t.Render())
}

//---------------------------------------------->
/*package main

import (
    "fmt"

    "github.com/jedib0t/go-pretty/v6/table"
    "github.com/jedib0t/go-pretty/v6/text"
)

func main() {

    t := table.NewWriter()
    t.SetTitle("Users")
    t.SetAutoIndex(true)
    t.Style().Format.Header = text.FormatTitle

    t.AppendHeader(table.Row{"Name", "Occupation"})
    t.AppendRow(table.Row{"John Doe", "gardener"})
    t.AppendRow(table.Row{"Roger Roe", "driver"})
    t.AppendRows([]table.Row{{"Paul Smith", "trader"},
        {"Lucy Smith", "teacher"}})

    fmt.Println(t.Render())
}*/
//------------------------------>
/*package main

import (
    "fmt"

    "github.com/jedib0t/go-pretty/v6/table"
    "github.com/jedib0t/go-pretty/v6/text"
)

type User struct {
    Name       string
    Occupation string
    Salary     int
}

func main() {

    users := []User{{"John Doe", "gardener", 1250}, {"Roger Roe", "driver", 950},
        {"Paul Smith", "trader", 2100}, {"Lucy Smith", "teacher", 880}}

    var total int

    for _, u := range users {
        total += u.Salary
    }

    t := table.NewWriter()
    t.SetCaption("Users")
    t.SetAutoIndex(true)
    t.Style().Format.Header = text.FormatTitle
    t.Style().Format.Footer = text.FormatTitle

    t.AppendHeader(table.Row{"Name", "Occupation", "Salary"})

    for _, u := range users {
        t.AppendRow(table.Row{u.Name, u.Occupation, u.Salary})
    }

    t.AppendFooter(table.Row{"", "Total", total})

    fmt.Println(t.Render())
}
//--------------------------->
https://zetcode.com/golang/terminal-table/


*/
