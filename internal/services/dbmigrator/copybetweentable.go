package dbmigrator

import "database/sql"

// this struct implement pgx.CopyFromSource to avoid buffering entire data set to memory,
// see https://pkg.go.dev/github.com/jackc/pgx/v5#hdr-Copy_Protocol
// and https://pkg.go.dev/github.com/jackc/pgx/v5#CopyFromSource
type copyBeetweenTable struct {
	columns []string
	srcRows *sql.Rows
	err     error
}

func (cbt *copyBeetweenTable) Next() bool {
	return cbt.srcRows.Next()
}

func (cbt *copyBeetweenTable) Values() ([]interface{}, error) {
	// TODO: i think this this code potential make memory leak, maybe search another method to get data
	values := make([]any, len(cbt.columns))
	for i := range values {
		values[i] = new(any)
	}
	err := cbt.srcRows.Scan(values...)
	if err != nil {
		return nil, err
	}
	return values, err
}

func (cbt *copyBeetweenTable) Err() error {
	return cbt.err
}
