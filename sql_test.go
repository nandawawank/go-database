package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO job (title, organization, sequence, winner) VALUES ('test', 2, 2, 2);"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("successful insert data new job")
}

func TestQuerySql(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT title FROM job"
	rows, err := db.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var title string
		err := rows.Scan(&title)
		if err != nil {
			panic(err)
		}

		fmt.Println("title:", title)
	}
	rows.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id_job, title, organization, sequence, winner FROM job"
	rows, err := db.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id_job, organization, sequence, winner sql.NullInt16 // if data can be null
		var title string

		err := rows.Scan(&id_job, &title, &organization, &sequence, &winner)
		if err != nil {
			panic(err)
		}

		fmt.Println("=================")
		if id_job.Valid {
			fmt.Println("id_job:", id_job.Int16)
		}
		fmt.Println("title:", title)
		fmt.Println("organization:", organization)
		fmt.Println("sequence:", sequence)
		fmt.Println("winner:", winner)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	title := "Ketua"
	query := "SELECT organization, title FROM job WHERE title = '" + title + "'"

	rows, err := db.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var title string
		var organization int

		err := rows.Scan(&organization, &title)
		if err != nil {
			panic(err)
		}

		fmt.Println("title:", title)
		fmt.Println("organization:", organization)
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	title := "Ketua"
	query := "SELECT organization, title FROM job WHERE title = ?"

	rows, err := db.QueryContext(ctx, query, title)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var title string
		var organization int16
		err := rows.Scan(&organization, &title)

		if err != nil {
			panic(err)
		}

		fmt.Println("title:", title)
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	title := "Ketua"
	organization := 2
	sequence := 2
	winner := 2

	query := "INSERT INTO job (title, organization, sequence, winner) VALUES (?, ?, ?, ?)"
	result, err := db.ExecContext(ctx, query, title, organization, sequence, winner)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully inserted new data", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO job (title, organization, sequence, winner) VALUES (?, ?, ?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	defer statement.Close()

	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		title := "title" + strconv.Itoa(i)
		organization := i
		sequence := i
		winner := i

		result, err := statement.ExecContext(ctx, title, organization, sequence, winner)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Job ID ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO job (title, organization, sequence, winner) VALUES (?, ?, ?, ?)"
	for i := 0; i < 10; i++ {
		title := "title" + strconv.Itoa(i)
		organization := i
		sequence := i
		winner := i

		result, err := tx.ExecContext(ctx, script, title, organization, sequence, winner)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Job ID ", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}

func TestRepository(t *testing.T) {

}
