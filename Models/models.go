package models

//Table struct representing a database table
type Table struct {
	Columns []Column
}

//Column struct representing a database column
type Column struct {
	ColumnName string
	ColumnType string
}

//------------------------------------------------------------------------------

//Query supporting struct for query generation
type Query struct {
	queryString string
	//parameters  []string
	parameters []interface{}
}

//AddParameter Adds a new parameter to the query for binding later on
func (q *Query) AddParameter(newParam string) {
	q.parameters = append(q.parameters, newParam)
}

//AppendToQuery Adds a new clause (or part of one) to the query.
func (q *Query) AppendToQuery(querySection string) {
	q.queryString += querySection
}

//GetQuery returns the query
func (q Query) GetQuery() string {
	return q.queryString
}

//GetParams returns the query parameters.
func (q Query) GetParams() []interface{} {
	return q.parameters
}

//------------------------------------------------------------------------------

//Request embodies the structure of a json request sent to the API
type Request struct {
	Command   string
	TableName string
	Values    map[string]string
	Updates   map[string]string
}

// type QueryRequest struct {
//
// }
