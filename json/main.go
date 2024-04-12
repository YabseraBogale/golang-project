package main

// how is go slow here ? maybe 10000
// python for 100

// real    0m0.078s
// user    0m0.075s
// sys     0m0.004s

// golang for 100

// real    0m0.467s
// user    0m0.417s
// sys     0m0.140s

// python for 1000

// real    0m0.239s
// user    0m0.195s
// sys     0m0.044s

// golang for 1000

// real    0m0.866s
// user    0m0.662s
// sys     0m0.250s

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	count := 100
	for i := 0; i < count; i++ {
		db, _ := sql.Open("sqlite3", "data.db")
		defer db.Close()
		statment := "select * from data where id=" + strconv.Itoa(rand.Intn(4000))
		row, _ := db.Query(statment)
		for row.Next() {
			var id int
			var year int
			var vulnerability string
			var attackComplexity string
			var baseScore float32
			row.Scan(&id, &vulnerability, &year, &attackComplexity, &baseScore)
			fmt.Println(id, vulnerability, year, attackComplexity, baseScore)
		}
	}
}
