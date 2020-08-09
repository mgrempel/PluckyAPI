package endpoints

import (
	"PluckyAPI/Models"
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
	b := []byte(`{"Command":"DELETE",
							  "TableName":"TestTable1",
								"Values":{"UserName":"mackenzie","Password":"remple"},
								"Updates":{"UserName":"mackenzie","Password":"rempel"}}`)

	//UnMarshall query into a request struct - Might couple this off later but idk
	var request = models.Request{}
	query := models.Query{}
	json.Unmarshal(b, &request)

	//Determine the type of request
	switch request.Command {
	case "SELECT":
		query, err := builder.Select(request, query)
		if err != nil {
			//Need a func to handle creating and returning an error to the caller
			panic(err)
		}
		fmt.Println(query.GetQuery())
	case "INSERT":
		query, err := builder.Insert(request, query)
		if err != nil {
			panic(err)
		}
		fmt.Println(query.GetQuery())
	case "UPDATE":
		query, err := builder.Update(request, query)
		if err != nil {
			panic(err)
		}
		fmt.Println(query.GetQuery())
	case "DELETE":
		query, err := builder.Delete(request, query)
		if err != nil {
			panic(err)
		}
		fmt.Println(query.GetQuery())
	default:
		fmt.Println("Something went wrong lol")
	}

	//Evaluate where constraint

	//fmt.Println(request)
}
