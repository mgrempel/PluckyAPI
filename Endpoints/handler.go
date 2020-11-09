package endpoints

import (
	"PluckyAPI/Models"
	"PluckyAPI/Utilities/misc"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*HandleRequest Just a place holder for now, will eventually be called asynchronously by an endpoint and passed instructions
for building an sql query and handing it off to another function for execution. The main job of this func will be to parse a json requests
into a json response containing all of the requested rows etc.
an enpoin */
func (env *Container) HandleRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Ensure we have all we need from the request
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
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
		return
	}

	//Determine if the request is valid
	//Table name check
	if !env.findTable(request.TableName) {
		http.Error(w, "Table does not exist", 400)
		return
	}
	table := env.Tables[request.TableName].Columns
	columns := make([]string, 0)

	for _, value := range table {
		columns = append(columns, value.ColumnName)
	}

	//Table Value check
	for key := range request.Values {
		if !misc.Contains(columns, key) {
			http.Error(w, "Invalid value fields", 400)
			return
		}
	}

	//Update check
	if request.Updates != nil {
		for key := range request.Updates {
			if !misc.Contains(columns, key) {
				http.Error(w, "Invalid update fields", 400)
				return
			}
		}
	}

	//Determine the type of request
	switch request.Command {
	case "SELECT":
		query, err = builder.Select(request, query)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	case "INSERT":
		query, err = builder.Insert(request, query)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	case "UPDATE":
		query, err = builder.Update(request, query)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	case "DELETE":
		query, err = builder.Delete(request, query)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	default:
		http.Error(w, "Could not determine the request type", 400)
		return
	}

	if request.Values != nil && request.Command != "INSERT" {
		query, err = builder.Where(request, query)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	query.AppendToQuery(";")

	resp, err := env.executeQuery(query)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
