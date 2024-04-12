package Projet_GO_Reservation

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

type Db struct {
}

func connectDB() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go_reserv")
	if err != nil {
		log("Impossible de se connecter à la BDD", err)
		return nil
	}
	log("BDD Connecting ok")
	return db
}

//
// ------------------------------------------------------------------------------------------------ //
//

func (d *Db) SelectDB(table string, column []string, condition *string, debug ...bool) ([]map[string]interface{}, error) {

	if checkData(table, column, condition) == false {
		log("Plz check your condition")
		return nil, nil
	}

	var db = connectDB()

	if db == nil {
		log("What da heck bro, l'instance db est nulle ??")
		return nil, nil
	}

	// checking the right format
	var columns = arrayToString(column)

	if columns == NullString {
		log("Impossible to transform the columns array into a string")
		return nil, nil
	}

	var query *sql.Rows
	var queryString string
	var err error

	if condition == nil {
		query, err = db.Query("SELECT " + columns + " FROM " + table)
		queryString = "SELECT " + columns + " FROM " + table
		if err != nil {
			log("ERROR : ", err)
			return nil, nil
		}
	} else {
		query, err = db.Query("SELECT " + columns + " FROM " + table + " WHERE " + *condition)
		queryString = "SELECT " + columns + " FROM " + table + " WHERE " + *condition
		if err != nil {
			log("ERROR : ", err)
			return nil, err
		}
	}

	if len(debug) > 0 && debug[0] {
		log(queryString)
	}

	var result = transformQueryToMap(query)

	if err := query.Err(); err != nil {
		log("Erreur lors de la lecture des résultats :", err)
		return nil, err
	}

	return result, nil
}

//
// ------------------------------------------------------------------------------------------------ //
//

func (d *Db) InsertDB(table string, column []string, values []string, condition *string, debug ...bool) {

	if checkData(table, column, condition) == false {
		return
	}

	/*if condition == nil {
		query, err = db.Query("INSERT INTO " + table + ", " + columns)
		queryString = "INSERT INTO " + table + ", " + columns
		if err != nil {
			log("ERROR : ", err)
			return nil, nil
		}
	} else {
		query, err = db.Query("INSERT INTO " + table + ", " + columns + " WHERE " + *condition)
		queryString = "INSERT INTO " + table + ", " + columns + " WHERE " + *condition
		if err != nil {
			log("ERROR : ", err)
			return nil, err
		}
	}*/

}

//
// ------------------------------------------------------------------------------------------------ //
//

func (d *Db) UpdateDB(table string) {
	fmt.Printf("Get")
}

//
// ------------------------------------------------------------------------------------------------ //
//

func transformQueryToMap(query *sql.Rows) []map[string]interface{} {
	var result []map[string]interface{}

	for query.Next() {

		//Get all the columns
		columns, err := query.Columns()

		if err != nil {
			log("ERROR : Impossible de récupérer le nom des colonnes")
			return nil
		}

		// Create a slice to stock vlaues
		values := make([]interface{}, len(columns))

		// Create a pointer slice to values
		pointers := make([]interface{}, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}

		if err := query.Scan(pointers...); err != nil {
			return nil
		}

		// Create like a json object
		row := make(map[string]interface{})
		for i, name := range columns {
			switch v := values[i].(type) {
			case []byte:
				row[name] = string(v)
			default:
				row[name] = v
			}
		}

		result = append(result, row)
	}
	return result
}

//
// ------------------------------------------------------------------------------------------------ //
//

func checkData(table string, column []string, condition *string) bool {

	if reflect.TypeOf(table) != reflect.TypeOf("") || table == NullString {
		log("Faut donner un nom de table :/ sous forme de chaine de caractère")
		return false
	}

	if reflect.TypeOf(column).Kind() != reflect.Slice || len(column) == 0 {
		log("Faut donner un tableau de string(s)")
		return false
	}

	if condition != nil {
		if reflect.TypeOf(condition) != reflect.TypeOf("") {
			log("Il faut donner une condition sous forme de string")
			return false
		}
	}
	return true
}
