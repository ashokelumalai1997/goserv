package main

import (
	"context"
	"fmt"

	wire "github.com/ashokelumalai1997/psql-wire"
	"github.com/lib/pq/oid"
)

func main() {

	parameters := make(wire.Parameters)
	parameters[wire.ParamServerVersion] = "13.6"

	wire.ListenAndServe("127.0.0.1:5432", handler)
}

/*
SELECT t.oid, typarray
FROM pg_type t JOIN pg_namespace ns

	ON typnamespace = ns.oid

WHERE typname = 'hstore';
*/
var pgtype_table = wire.Columns{
	{
		Table:  0,
		Name:   "search_path",
		Oid:    oid.T_text,
		Width:  256,
		Format: wire.TextFormat,
	},
}

func handler(ctx context.Context, query string) (wire.PreparedStatementFn, []oid.Oid, wire.Columns, error) {
	fmt.Println(query)
	fmt.Println("COUNT")
	statement := func(ctx context.Context, writer wire.DataWriter, parameters []string) error {
		return nil
	}

	// nullstatement := func(ctx context.Context, writer wire.DataWriter, parameters []string) error {
	// 	return nil
	// }

	// statementDateStyle := func(ctx context.Context, writer wire.DataWriter, parameters []string) error {
	// 	return writer.Complete("ST")
	// }

	// if strings.Contains(query, "datestyle") {
	// 	return statementDateStyle, nil, nil, nil
	// }

	// if strings.Contains(query, "BEGIN") {
	// 	return statementDateStyle, nil, nil, nil
	// }

	return statement, wire.ParseParameters(query), pgtype_table, nil
}
