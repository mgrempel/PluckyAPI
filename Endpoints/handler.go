package endpoints

import (
	"PluckyAPI/Models"
	"PluckyAPI/Utilities/misc"
	"encoding/json"
	"fmt"
)

/*HandleRequest Just a place holder for now, will eventually be called asynchronously by an endpoint and passed instructions
for building an sql query and handing it off to another function for execution. The main job of this func will be to parse a json requests
into a json response containing all of the requested rows etc.
an enpoin */
func (env *Container) HandleRequest() {
	builder := env.Builder
	//This will be kinda what incoming JSON will look like.
	b := []byte(`{"Command":"UPDATE","TableName":"TestTable1","Values":{"UserName":"mackenzie","Password":null},"Updates":{"Username":"mackenzie","Password":"rempel"}}`)

	//UnMarshall query into a request struct - Might couple this off later but idk
	var request = models.Request{}
	var query = models.Query{}
	var err error
	json.Unmarshal(b, &request)

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

	//Table column check
	for key := range request.Values {
		fmt.Println(key)
		if !misc.Contains(columns, key) {
			panic(fmt.Errorf("Invalid value columns"))
		}
	}

	//Determine the type of request
	switch request.Command {
	case "SELECT":
		query, err = builder.Select(request, query)
		if err != nil {
			//Need a func to handle creating and returning an error to the caller
			panic(err)
		}
		//fmt.Println(query.GetQuery())
	case "INSERT":
		query, err = builder.Insert(request, query)
		if err != nil {
			panic(err)
		}
		fmt.Println(query.GetQuery())
	case "UPDATE":
		query, err = builder.Update(request, query)
		if err != nil {
			panic(err)
		}
		fmt.Println(query.GetQuery())
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
	fmt.Println(query.GetQuery())
	//fmt.Println("\n")
	for _, param := range query.GetParams() {
		fmt.Print(param + " ")
	}
	//env.executeQuery(query)
	//Evaluate where constraint

	//fmt.Println(request)
}
