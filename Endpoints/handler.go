package endpoints

import (
	"PluckyAPI/Models"
	"PluckyAPI/Utilities/misc"
	"encoding/json"
	"fmt"
	"net/http"
)

/*HandleRequest Just a place holder for now, will eventually be called asynchronously by an endpoint and passed instructions
for building an sql query and handing it off to another function for execution. The main job of this func will be to parse a json requests
into a json response containing all of the requested rows etc.
an enpoin */
func (env *Container) HandleRequest(w http.ResponseWriter, r *http.Request) {
	//Ensure we have all we need from the request
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
	}

	//Variables which will be used in the request
	builder := env.Builder
	var request = models.Request{}
	var query = models.Query{}
	var err error

	//Decode the request
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error decoding request", 400)
	}

	//Determine if the request is valid
	//Table name check
	if !env.findTable(request.TableName) {
		panic(fmt.Errorf("Invalid table"))
	}
	table := env.Tables[request.TableName].Columns
	columns := make([]string, 0)

	for _, value := range table {
		columns = append(columns, value.ColumnName)
	}

	//Table Value check
	for key := range request.Values {
		if !misc.Contains(columns, key) {
			panic(fmt.Errorf("Invalid value columns"))
		}
	}

	//Update check

	//Determine the type of request
	switch request.Command {
	case "SELECT":
		query, err = builder.Select(request, query)
		if err != nil {
			//Need a func to handle creating and returning an error to the caller
			panic(err)
		}
	case "INSERT":
		query, err = builder.Insert(request, query)
		if err != nil {
			panic(err)
		}
	case "UPDATE":
		query, err = builder.Update(request, query)
		if err != nil {
			panic(err)
		}
	case "DELETE":
		query, err = builder.Delete(request, query)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("Something went wrong lol")
	}

	if request.Values != nil && request.Command != "INSERT" {
		query, err = builder.Where(request, query)
		if err != nil {
			panic(err)
		}
	}

	query.AppendToQuery(";")

	for _, param := range query.GetParams() {
		fmt.Print(param)
	}
	fmt.Println(query.GetQuery())

	fmt.Println(env.executeQuery(query))
}
