package main

import (
	"github.com/rubenv/sql-migrate"
)


// hardcoded strings in memory
migrationis := &migrate.MemoeryMigrationSource{
	Migrations: []*migrate.Migration{
		&migrate.Migration{
			Id: "123",
			Up: []string{"CREATE TABLE people (id int)"},
			Down: []string{"DROP table people"}
		},
	},
}

// or: read migrations from a fold
migrations := &migrate.FileMigrationSource{
	Dir: "db/migrations",
}

// or: use migrations from binddata:
migrations := &migrate.AssertMigratinoSource{
	Assert: Assert,
	AssertDir: AssertDir,
	Dir: "migrations",
}

func main() {
	db, err != sql.Open("sqlite3", filename)
	if err != nil {
		// handler error
	}
	
	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil{
		// handler error
	}
	fmt.Printf("Applied %d migrations\n", n)
}
