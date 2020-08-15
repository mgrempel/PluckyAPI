package endpoints

import (
	models "PluckyAPI/Models"
	"database/sql"
	"encoding/json"
	"fmt"
)

//Perform any binding if nessecary then pass the query off to
func (env Container) executeQuery(query models.Query) ([]byte, error) {
	db := env.Db
	queryString := query.GetQuery()
	fmt.Println(queryString)

	err := db.Ping()
	if err != nil {
		fmt.Println("Can't connect to the database")
	}

	args := []interface{}{query.GetParams()}
	fmt.Println(args)

	var rows *sql.Rows

	// if len(args) > 0 {
	//
	// }
	//This is throwing an issue
	//	sqlParams := []string{"Mackenzie", "Rempel"}
	//var testToo = []interface{}{sqlParams}

	rows, err = db.Query(query.GetQuery(), (query.GetParams())...)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	//Solution for this grabbed from stackoverflow.com/a/60386531/12849275
	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		return nil, err
	}

	count := len(columnTypes)
	finalRows := []interface{}{}

	for rows.Next() {

		scanArgs := make([]interface{}, count)

		for i, v := range columnTypes {

			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
				break
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
				break
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
				break
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := rows.Scan(scanArgs...)

		if err != nil {
			panic(err)
		}

		masterData := map[string]interface{}{}

		for i, v := range columnTypes {

			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				masterData[v.Name()] = z.Bool
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				masterData[v.Name()] = z.String
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				masterData[v.Name()] = z.Int64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				masterData[v.Name()] = z.Float64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				masterData[v.Name()] = z.Int32
				continue
			}

			masterData[v.Name()] = scanArgs[i]
		}

		finalRows = append(finalRows, masterData)
	}

	z, err := json.Marshal(finalRows)
	if err != nil {
		return nil, err
	}

	return z, nil
}
