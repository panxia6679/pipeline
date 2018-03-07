package main

import (
	"fmt"
	"net/http"
	"os"
	"math/rand"
	"time"
	"github.com/codegangsta/negroni"
	"github.com/olekukonko/tablewriter"
	"github.com/cloudnativego/go-primer/npcs"

)

func main() {




	data := [][]string{
		[]string{"Alfred", "15", "10/20", "(10.32, 56.21, 30.25)"}, 
		[]string{"Beelzebub", "30", "30/50", "(1,1,1)"}, 
		[]string{"Hortense", "21", "80/80", "(1,1,1)"}, 
		[]string{"Pokey", "8", "30/40", "(1,1,1)"},
	}

	table := tablewriter.NewWriter(os.Stdout) 
	table.SetHeader([]string{"NPC", "Speed", "Power", "Location"}) 
	table.AppendBulk(data)
	table.Render()



	mob := npcs.NonPlayerCharacter{Name: "Alfred"} 
	fmt.Println(mob)


	hortense := npcs.NonPlayerCharacter{Name: "Hortense", Loc: npcs.Location{X: 10.0, Y: 10.0, Z: 10.0}}
	fmt.Println(hortense)

	fmt.Printf("Alfred is %f units from Hortense.\n", mob.DistanceTo(hortense))

	var rolls = getDieRolls()
	for index, rollFunc := range rolls {
		fmt.Printf("Die Roll Attempt #%d, result: %d\n", index, rollFunc(10))
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	n := negroni.Classic()
	n.UseHandler(mux)
	hostString := fmt.Sprintf(":%s", port)
	n.Run(hostString)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello from Go!")
}


type dieRollFunc func(int) int


func fakeDieRoll(size int) int { 
	return 42
}


func dieRoll(size int) int { rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

func getDieRolls() []dieRollFunc { 
	return []dieRollFunc{
		dieRoll, 
		fakeDieRoll,
	} 
}
