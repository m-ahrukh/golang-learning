package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "phone_normalizer"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	// db, err := sql.Open("postgres", psqlInfo)
	// must(err)
	// err = resetDB(db, dbname)
	// must(err)
	// db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()
	must(createPhoneNumbersTable(db))

	_, err = insertPhoneNumbers(db, "1234567890")
	must(err)
	_, err = insertPhoneNumbers(db, "123 456 7891")
	must(err)
	_, err = insertPhoneNumbers(db, "(123) 456 7892")
	must(err)
	_, err = insertPhoneNumbers(db, "(123) 456-7893")
	must(err)
	_, err = insertPhoneNumbers(db, "123-456-7894")
	must(err)
	_, err = insertPhoneNumbers(db, "123-456-7890")
	must(err)
	_, err = insertPhoneNumbers(db, "(123)456-7892")
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
	CREATE TABLE IF NOT EXISTS phone_number (
		id SERIAL,
		value VARCHAR(20)
	)`
	_, err := db.Exec(statement)
	return err
}

func insertPhoneNumbers(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_number(value) VALUES($1) RETURNING id` //Sql injection
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return -1, err
	// }
	return id, nil
}

func normalize(phone string) string {
	//normaize number by iterating string
	//normalize number using regex

	//Method 1.a
	output := ""
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			output = output + string(ch)
		}
	}
	return output

	//Method 1.b
	// var buf bytes.Buffer
	// for _, ch := range phone {
	// 	if ch >= '0' && ch <= '9' {
	// 		buf.WriteRune(ch)
	// 	}
	// }
	// return buf.String()

	//Method 2
	// re := regexp.MustCompile("[0-9]+")
	// matches := re.FindAllString(phone, -1)
	// fmt.Println(strings.Join(matches, ""))
	// return strings.Join(matches, "")
}
