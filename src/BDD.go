package Projet_GO_Reservation

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

type Db struct {
}

func connectDB() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go_reserv")
	if err != nil {
		Log.Error("Impossible de se connecter à la BDD", err)
		return nil
	}
	Log.Infos("BDD Connecting ok")
	return db
}

//
// ------------------------------------------------------------------------------------------------ //
//

func (d *Db) SelectDB(table string, column []string, condition *string, debug ...bool) ([]map[string]interface{}, error) {

	if checkData(table, column, nil, condition) == false {
		Log.Error("Plz check your condition")
		return nil, nil
	}

	var db = connectDB()

	if db == nil {
		Log.Error("What da heck bro, l'instance db est nulle ??")
		return nil, nil
	}

	// checking the right format
	var columns = arrayToString(column)

	if columns == NullString {
		Log.Error("Impossible to transform the columns array into a string")
		return nil, nil
	}

	var query *sql.Rows
	var queryString string
	var err error

	if condition == nil {
		query, err = db.Query("SELECT " + columns + " FROM " + table)
		queryString = "SELECT " + columns + " FROM " + table
		if err != nil {
			ILog("ERROR : ", err)
			Log.Debug(queryString)
			return nil, nil
		}
	} else {
		query, err = db.Query("SELECT " + columns + " FROM " + table + " WHERE " + *condition)
		queryString = "SELECT " + columns + " FROM " + table + " WHERE " + *condition
		if err != nil {
			ILog("ERROR : ", err)
			Log.Debug(queryString)
			return nil, err
		}
	}

	if len(debug) > 0 && debug[0] {
		Log.Debug(queryString)
	}

	var result = transformQueryToMap(query)

	if err := query.Err(); err != nil {
		Log.Error("An error Occured : ", err)
		return nil, err
	}

	return result, nil
}

//
// ------------------------------------------------------------------------------------------------ //
//

func (d *Db) InsertDB(table string, column []string, value []string, condition *string, debug ...bool) {

	if checkData(table, column, value, condition) == false {
		return
	}

	var db = connectDB()

	if db == nil {
		Log.Error("What da heck bro, l'instance db est nulle ??")
		return
	}

	var columns = arrayToString(column, true)

	var values = arrayToString(value)

	if columns == NullString {
		Log.Error("Impossible to transform the columns array into a string")
		return
	}

	if values == NullString {
		Log.Error("Impossible to transform the columns array into a string")
		return
	}

	var query *sql.Rows
	var queryString string
	var err error

	if condition == nil {
		query, err = db.Query("INSERT INTO " + table + " (" + columns + ") VALUES (" + values + ")")
		queryString = "INSERT INTO " + table + " (" + columns + ") VALUES (" + values + ")"
		if err != nil {
			ILog("ERROR : ", err)
			Log.Debug(queryString)
			return
		}
	} else {
		query, err = db.Query("INSERT INTO " + table + " (" + columns + ") VALUES (" + values + ") WHERE " + *condition)
		queryString = "INSERT INTO " + table + " (" + columns + ") VALUES (" + values + ") WHERE " + *condition
		if err != nil {
			ILog("ERROR : ", err)
			Log.Debug(queryString)
			return
		}
	}

	if err := query.Err(); err != nil {
		Log.Error("An error Occured : ", err)
		return
	}

	if len(debug) > 0 && debug[0] {
		Log.Debug(queryString)
	}

	return

}

//
// ------------------------------------------------------------------------------------------------ //
//

func (d *Db) UpdateDB(table string, column []string, value []string, condition *string, debug ...bool) {

	if checkData(table, column, value, condition) == false {
		return
	}

	if condition == nil {
		Log.Error("Plz enter a condition to update the table. If you don't want to enter condition put a \"-1\" instead")
		return
	}

	var db = connectDB()

	if db == nil {
		Log.Error("What da heck bro, l'instance db est nulle ??")
		return
	}

	var query *sql.Rows
	var queryString string
	var err error

	var set = concatColumnWithValues(column, value)

	if set == NullString {
		return
	}

	if condition != nil {
		query, err = db.Query("UPDATE " + table + " SET " + set + " WHERE " + *condition)
		queryString = "UPDATE " + table + " SET " + set + " WHERE " + *condition
		if err != nil {
			ILog("ERROR : ", err)
			Log.Debug(queryString)
			return
		}
	} else if *condition == "-1" {
		query, err = db.Query("UPDATE " + table + " SET " + set)
		queryString = "UPDATE " + table + " SET " + set
		if err != nil {
			ILog("ERROR : ", err)
			Log.Debug(queryString)
			return
		}
	}

	if err := query.Err(); err != nil {
		Log.Error("An error Occured : ", err)
		return
	}

	if len(debug) > 0 && debug[0] {
		ILog("DEBUG : " + queryString)
	}

	return

}

//
// ------------------------------------------------------------------------------------------------ //
//

func (d *Db) DeleteDB(table string, condition *string, debug ...bool) {
	// DELETE FROM table WHERE condition

	if reflect.TypeOf(table) != reflect.TypeOf("") || table == NullString {
		Log.Error("Faut donner un nom de table :/ sous forme de chaine de caractère")
	}

	if condition == nil {
		Log.Error("Plz enter a condition to delete a row from a the table. If you don't want to enter condition put a \"-1\" instead")
		return
	}

	var db = connectDB()

	if db == nil {
		Log.Error("What da heck bro, l'instance db est nulle ??")
		return
	}

	var query *sql.Rows
	var queryString string
	var err error

	if condition != nil {
		query, err = db.Query("DELETE FROM " + table + " WHERE " + *condition)
		queryString = "DELETE FROM " + table + " WHERE " + *condition
		if err != nil {
			ILog("ERROR : ", err)
			Log.Debug(queryString)
			return
		}
	} else if *condition == "-1" {
		query, err = db.Query("DELETE FROM " + table)
		queryString = "DELETE FROM " + table
		if err != nil {
			ILog("ERROR : ", err)
			Log.Debug(queryString)
			return
		}
	}

	if err := query.Err(); err != nil {
		Log.Error("An error Occured : ", err)
		return
	}

	if len(debug) > 0 && debug[0] {
		Log.Debug(queryString)
	}

	return
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
			Log.Error("Impossible de récupérer le nom des colonnes")
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
			Log.Error("Impossible to determine the pointer when creating the map")
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

func checkData(table string, column []string, values []string, condition *string) bool {

	if reflect.TypeOf(table) != reflect.TypeOf("") || table == NullString {
		Log.Error("Faut donner un nom de table :/ sous forme de chaine de caractère")
		return false
	}

	if column == nil || reflect.TypeOf(column).Kind() != reflect.Slice || len(column) == 0 {
		Log.Error("Faut donner un tableau de string(s)")
		return false
	}

	if values == nil || reflect.TypeOf(values).Kind() != reflect.Slice || len(column) == 0 {
		Log.Error("Faut donner un tableau de string(s)")
		return false
	}

	if condition != nil && reflect.TypeOf(*condition) != reflect.TypeOf("") {
		Log.Error("Il faut donner une condition sous forme de string")
		return false
	}

	return true
}
