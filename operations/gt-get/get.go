package gt_get

import (
	"git.jereileu.ch/gotables/server/gt-server/fs"
	"git.jereileu.ch/gotables/server/gt-server/operations/shared"
)

func Get(table string, db string, config fs.Conf) (fs.Table, error) {
	retTable := fs.Table{}
	var retError error = nil

	if db == "" {
		retTable, retError = getDBs(config.Dir)
	} else if table == "" {
		retTable, retError = getTables(db, config.Dir)
	} else {
		retTable, retError = getTable(table, db, config.Dir)
	}

	return retTable, retError
}

func getDBs(dir string) (fs.Table, error) {
	dbs, err := fs.GetDBs(dir)
	if err != nil {
		return fs.Table{}, err
	}
	column := fs.Column{
		Name: "DBs",
		Type: "string",
	}
	columns := []fs.Column{column}
	rows := make([][]interface{}, 0)
	for i, db := range dbs {
		rows[i] = append(rows[i], db)
	}
	retTable, err := shared.MakeNewTable(columns, rows)
	return retTable, err
}

func getTables(db string, dir string) (fs.Table, error) {
	tables, err := fs.GetTables(db, dir)
	if err != nil {
		return fs.Table{}, err
	}
	column := fs.Column{
		Name: "Tables",
		Type: "string",
	}
	columns := []fs.Column{column}
	rows := make([][]interface{}, 0)
	for i, table := range tables {
		rows[i] = append(rows[i], table)
	}
	retTable, err := shared.MakeNewTable(columns, rows)
	return retTable, err
}

func getTable(table string, db string, dir string) (fs.Table, error) {
	retTable, retError := fs.GetTable(table, db, dir)
	return retTable, retError
}
