package binlog

import (
	"testing"

	"github.com/actiontech/dtle/driver/common"
	"github.com/actiontech/dtle/driver/mysql/mysqlconfig"
	hclog "github.com/hashicorp/go-hclog"
)

//func Test_loadMapping(t *testing.T) {
//
//	replicateDoDbWithRename := []*common.DataSource{
//		{
//			TableSchema:       "db1",
//			TableSchemaRename: "db1-rename",
//			Tables: []*common.Table{
//				{
//					TableName:         "tb1",
//					TableRename:       "db1-tb1-rename",
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//				},
//			},
//		},
//		{
//			TableSchema:       "db2",
//			TableSchemaRename: "db2-rename",
//			Tables: []*common.Table{
//				{
//					TableName:         "tb1",
//					TableRename:       "db2-tb1-rename",
//					TableSchema:       "db2",
//					TableSchemaRename: "db2-rename",
//				},
//			},
//		},
//	}
//
//	replicateDoDbWithoutRename := []*common.DataSource{
//		{
//			TableSchema: "db1",
//			Tables: []*common.Table{
//				{
//					TableName:   "tb1",
//					TableSchema: "db1",
//				},
//			},
//		},
//		{
//			TableSchema: "db2",
//			Tables: []*common.Table{
//				{
//					TableName:   "tb1",
//					TableSchema: "db2",
//				},
//			},
//		},
//	}
//
//	type args struct {
//		sql             string
//		currentSchema   string
//		newSchemaName   string
//		newTableName    string
//		replicationDoDB []*common.DataSource
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// test with rename config
//		// drop table
//		{name: "schema-rename-drop-table", args: args{
//			sql:             "drop table `db1`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "DROP TABLE `db1-rename`.`db1-tb1-rename`"},
//		{name: "schema-rename-drop-table", args: args{ //drop table without specify schema
//			sql:             "drop table `tb1`,`db2`.tb1",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "DROP TABLE `db1-rename`.`db1-tb1-rename`, `db2-rename`.`db2-tb1-rename`"},
//		{name: "schema-rename-drop-table", args: args{ //drop several tables including table that need no renaming
//			sql:             "drop table `db1`.`tb1`,`db2`.tb1,`db3`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "DROP TABLE `db1-rename`.`db1-tb1-rename`, `db2-rename`.`db2-tb1-rename`, `db3`.`tb1`"},
//		// create/drop database
//		{name: "schema-rename-create-database", args: args{
//			sql:             "create database `db1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "CREATE DATABASE `db1-rename`"},
//		{name: "schema-rename-drop-database", args: args{
//			sql:             "drop database `db1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "DROP DATABASE `db1-rename`"},
//		// alter database
//		{name: "schema-rename-alter-database", args: args{
//			sql:             "alter database `db1` character set utf8 collate utf8_general_ci",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "ALTER DATABASE `db1-rename` CHARACTER SET = utf8 COLLATE = utf8_general_ci"},
//		{name: "schema-rename-alter-database", args: args{ // alter default database
//			sql:             "alter database character set utf8 collate utf8_general_ci",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "ALTER DATABASE CHARACTER SET = utf8 COLLATE = utf8_general_ci"},
//		// create index
//		{name: "schema-rename-create-index", args: args{
//			sql:             "create index idx on `db1`.`tb1` (name(10))",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "CREATE INDEX `idx` ON `db1-rename`.`db1-tb1-rename` (`name`(10))"},
//		// drop index
//		{name: "schema-rename-drop-index", args: args{
//			sql:             "drop index idx on `db1`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "DROP INDEX `idx` ON `db1-rename`.`db1-tb1-rename`"},
//		// create table
//		{name: "schema-rename-create-table", args: args{
//			sql:             "create table `db1`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "CREATE TABLE `db1-rename`.`db1-tb1-rename` "},
//		// alter table
//		{name: "schema-rename-alter-table", args: args{
//			sql:             "alter table `db1`.`tb1` add column a int",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "ALTER TABLE `db1-rename`.`db1-tb1-rename` ADD COLUMN `a` INT"},
//		{name: "schema-rename-alter-table", args: args{
//			sql:             "alter table `db1`.`tb1` rename as `db2`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "ALTER TABLE `db1-rename`.`db1-tb1-rename` RENAME AS `db2-rename`.`db2-tb1-rename`"},
//		{name: "schema-rename-alter-table", args: args{
//			sql:             "alter table `db1`.`tb1` rename to `db2`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "ALTER TABLE `db1-rename`.`db1-tb1-rename` RENAME AS `db2-rename`.`db2-tb1-rename`"},
//		// flush tables
//		{name: "schema-rename-flush-tables", args: args{
//			sql:             "flush tables `tb1`,`db2`.tb1",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "FLUSH TABLES `db1-rename`.`db1-tb1-rename`, `db2-rename`.`db2-tb1-rename`"},
//		// truncate table
//		{name: "schema-rename-truncate-table", args: args{
//			sql:             "truncate table `db1`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "TRUNCATE TABLE `db1-rename`.`db1-tb1-rename`"},
//		// rename table
//		{name: "schema-rename-rename-table", args: args{
//			sql:             "rename table `db1`.`tb1` to `db1`.`tb2`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "RENAME TABLE `db1-rename`.`db1-tb1-rename` TO `db1-rename`.`tb2`"},
//		{name: "schema-rename-rename-table", args: args{
//			sql:             "rename table `tb1` to `tb2`,`db2`.`tb1` to `db2`.`tb2`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "RENAME TABLE `db1-rename`.`db1-tb1-rename` TO `db1-rename`.`tb2`, `db2-rename`.`db2-tb1-rename` TO `db2-rename`.`tb2`"},
//		{name: "schema-rename-grant", args: args{
//			sql:             "grant ALL PRIVILEGES ON `db1`.`tb1` TO `test`@`%`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "GRANT ALL ON `db1-rename`.`db1-tb1-rename` TO `test`@`%`"},
//		{name: "schema-rename-revoke", args: args{
//			sql:             "REVOKE INSERT ON `db1`.`tb1` FROM `test`@`%`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithRename,
//		}, want: "REVOKE INSERT ON `db1-rename`.`db1-tb1-rename` FROM `test`@`%`"},
//
//		// test without rename config
//		// drop table
//		{name: "schema-map-drop-table", args: args{ //drop table without specify schema
//			sql:             "drop table `tb1`,`db2`.tb1",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "DROP TABLE `db1`.`tb1`, `db2`.`tb1`"},
//		// create/drop database
//		{name: "schema-map-create-database", args: args{
//			sql:             "create database `db1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "CREATE DATABASE `db1`"},
//		{name: "schema-map-drop-database", args: args{
//			sql:             "drop database `db1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "DROP DATABASE `db1`"},
//		// alter database
//		{name: "schema-map-alter-database", args: args{
//			sql:           "alter database `db1` character set utf8 collate utf8_general_ci",
//			currentSchema: "db1",
//		}, want: "ALTER DATABASE `db1` CHARACTER SET = utf8 COLLATE = utf8_general_ci"},
//		{name: "schema-map-alter-database", args: args{ // alter default database
//			sql:             "alter database character set utf8 collate utf8_general_ci",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "ALTER DATABASE CHARACTER SET = utf8 COLLATE = utf8_general_ci"},
//		// create index
//		{name: "schema-map-create-index", args: args{
//			sql:             "create index idx on `db1`.`tb1` (name(10))",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "CREATE INDEX `idx` ON `db1`.`tb1` (`name`(10))"},
//		// drop index
//		{name: "schema-map-drop-index", args: args{
//			sql:             "drop index idx on `db1`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "DROP INDEX `idx` ON `db1`.`tb1`"},
//		// create table
//		{name: "schema-map-create-table", args: args{
//			sql:             "create table `db1`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "CREATE TABLE `db1`.`tb1` "},
//		// alter table
//		{name: "schema-map-alter-table", args: args{
//			sql:             "alter table `db1`.`tb1` add column a int",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "ALTER TABLE `db1`.`tb1` ADD COLUMN `a` INT"},
//		{name: "schema-map-alter-table", args: args{
//			sql:             "alter table `db1`.`tb1` rename as `db2`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "ALTER TABLE `db1`.`tb1` RENAME AS `db2`.`tb1`"},
//		{name: "schema-map-alter-table", args: args{
//			sql:             "alter table `db1`.`tb1` rename to `db2`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "ALTER TABLE `db1`.`tb1` RENAME AS `db2`.`tb1`"},
//		// flush tables
//		{name: "schema-map-flush-tables", args: args{
//			sql:             "flush tables `tb1`,`db2`.tb1",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "FLUSH TABLES `db1`.`tb1`, `db2`.`tb1`"},
//		// truncate table
//		{name: "schema-map-truncate-table", args: args{
//			sql:             "truncate table `db1`.`tb1`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "TRUNCATE TABLE `db1`.`tb1`"},
//		// rename table
//		{name: "schema-map-rename-table", args: args{
//			sql:             "rename table `db1`.`tb1` to `db1`.`tb2`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "RENAME TABLE `db1`.`tb1` TO `db1`.`tb2`"},
//		{name: "schema-map-rename-table", args: args{
//			sql:             "rename table `tb1` to `tb2`,`db2`.`tb1` to `db2`.`tb2`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "RENAME TABLE `db1`.`tb1` TO `db1`.`tb2`, `db2`.`tb1` TO `db2`.`tb2`"},
//		{name: "schema-map-grant", args: args{
//			sql:             "grant ALL PRIVILEGES ON `db1`.`tb1` TO `test`@`%`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "GRANT ALL ON `db1`.`tb1` TO `test`@`%`"},
//		{name: "schema-map-revoke", args: args{
//			sql:             "REVOKE INSERT ON `db1`.`tb1` FROM `test`@`%`",
//			currentSchema:   "db1",
//			replicationDoDB: replicateDoDbWithoutRename,
//		}, want: "REVOKE INSERT ON `db1`.`tb1` FROM `test`@`%`"},
//	}
//
//	binlogReader := &BinlogReader{
//		logger: hclog.New(&hclog.LoggerOptions{
//			Level:      hclog.Debug,
//			JSONFormat: true,
//		}),
//		mysqlContext: &common.MySQLDriverConfig{},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			stmt, err := parser.New().ParseOneStmt(tt.args.sql, "", "")
//			if err != nil {
//				t.Error(err)
//				return
//			}
//
//			binlogReader.currentReplicateDoDb = tt.args.replicationDoDB
//			schemaRenameMap, schemaNameToTablesRenameMap := binlogReader.generateRenameMaps()
//			if got, err := binlogReader.loadMapping(tt.args.sql, tt.args.currentSchema, schemaRenameMap, schemaNameToTablesRenameMap, stmt); nil != err {
//				t.Errorf("loadMapping() failed: %v", err)
//			} else if got != tt.want {
//				t.Errorf("loadMapping() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestBinlogReader_resolveQuery(t *testing.T) {
	skipFunc1 := func(schema string, tableName string) bool {
		return schema == "skip" || tableName == "skip"
	}
	b := &BinlogReader{
		lowerCaseTableNames: mysqlconfig.LowerCaseTableNames0,
		logger: hclog.Default(),
	}

	type args struct {
		currentSchema string
		sql           string
	}
	tests := []struct {
		name       string
		args       args
		wantResult parseQueryResult
		wantErr    bool
	}{
		{
			name: "drop-table-1",
			args: args{
				currentSchema: "",
				sql:           "drop table a.b, skip.c, d",
			},
			wantResult: parseQueryResult{
				sql: "DROP TABLE `a`.`b`, `d`",
			},
			wantErr: false,
		}, {
			name: "drop-table-2",
			args: args{
				currentSchema: "",
				sql:           "drop table if exists skip.b, skip.c",
			},
			wantResult: parseQueryResult{
				sql: "DROP TABLE IF EXISTS `skip`.`b`",
			},
			wantErr: false,
		}, {
			name:       "empty",
			args:       args{
				currentSchema: "",
				sql:           "",
			},
			wantResult: parseQueryResult{
				isRecognized: false,
				sql:          "",
			},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := b.resolveQuery(tt.args.currentSchema, tt.args.sql, skipFunc1)
			if (err != nil) != tt.wantErr {
				t.Errorf("resolveQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResult.sql != tt.wantResult.sql {
				t.Errorf("resolveQuery() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

//func Test_generateRenameMaps(t *testing.T) {
//	currentReplicateDoDb := []*common.DataSource{
//		{
//			TableSchema:       "db1",
//			TableSchemaRename: "db1-rename",
//			Tables: []*common.Table{
//				{
//					TableName:         "tb1",
//					TableRename:       "db1-tb1-rename",
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//				},
//				{
//					TableName:         "tb2",
//					TableRename:       "db1-tb2-rename",
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//				},
//				{
//					TableName:         "tb3",
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//				},
//			},
//		},
//		{
//			TableSchema: "db2",
//			Tables: []*common.Table{
//				{
//					TableName:   "tb1",
//					TableRename: "db2-tb1-rename",
//					TableSchema: "db2",
//				},
//			},
//		},
//		{
//			TableSchema:       "db3",
//			TableSchemaRename: "db3-rename",
//			Tables: []*common.Table{
//				{
//					TableName:   "tb1",
//					TableSchema: "db3",
//				},
//			},
//		},
//	}
//
//	wantSchemaRenameMap := map[string]string{
//		"db1": "db1-rename",
//		"db3": "db3-rename",
//	}
//
//	wantSchemaToTablesRenameMap := map[string]map[string]string{
//		"db1": {
//			"tb1": "db1-tb1-rename",
//			"tb2": "db1-tb2-rename",
//		},
//		"db2": {
//			"tb1": "db2-tb1-rename",
//		},
//	}
//
//	binlogReader := &BinlogReader{
//		logger:       nil,
//		mysqlContext: &common.MySQLDriverConfig{},
//	}
//	binlogReader.currentReplicateDoDb = currentReplicateDoDb
//
//	schemaRenameMap, schemaToTablesRenameMap := binlogReader.generateRenameMaps()
//
//	if !assert.Equal(t, wantSchemaRenameMap, schemaRenameMap, "unexpected schemaRenameMap") {
//		t.Errorf("unexpected schemaRenameMap: %v", schemaRenameMap)
//	}
//
//	if !assert.Equal(t, wantSchemaToTablesRenameMap, schemaToTablesRenameMap, "unexpected schemaToTablesRenameMap") {
//		t.Errorf("unexpected schemaToTablesRenameMap: %v", schemaToTablesRenameMap)
//	}
//
//}

func Test_matchTable(t *testing.T) {
	tableConfigs := []*common.Table{
		{
			TableName: "tb1",
		},
		{
			TableRegex: "(\\w*)tb_rex",
		},
	}

	rawReplicateDoDb := []*common.DataSource{
		{
			TableSchema: "db1",
			Tables:      tableConfigs,
		},
		{
			TableSchema: "db2",
		},
		{
			TableSchemaRegex: "(\\w*)db_rex1",
		},
	}

	type args struct {
		schemaName string
		tableName  string
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "match_schema",
			args: args{
				schemaName: "db1",
			},
			wantResult: true,
		},
		{
			name: "match_schema",
			args: args{
				schemaName: "db2",
				tableName:  "",
			},
			wantResult: true,
		},
		{
			name: "match_schema_rex",
			args: args{
				schemaName: "testdb_rex1",
				tableName:  "",
			},
			wantResult: true,
		},
		{
			name: "match_table",
			args: args{
				schemaName: "db1",
				tableName:  "tb1",
			},
			wantResult: true,
		},
		{
			name: "match_table_rex",
			args: args{
				schemaName: "db1",
				tableName:  "testtb_rex",
			},
			wantResult: true,
		},
		{
			name: "match_table",
			args: args{
				schemaName: "db2",
				tableName:  "testtb",
			},
			wantResult: true,
		},
		{
			name: "skip_schema",
			args: args{
				schemaName: "db_not_match",
			},
			wantResult: false,
		},
		{
			name: "skip_table",
			args: args{
				schemaName: "db1",
				tableName:  "tb2",
			},
			wantResult: false,
		},
	}

	binlogReader := &BinlogReader{
		mysqlContext: &common.MySQLDriverConfig{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := binlogReader.matchTable(rawReplicateDoDb, tt.args.schemaName, tt.args.tableName); res != tt.wantResult {
				t.Errorf("matchTable() gotResult = %v, want %v", res, tt.wantResult)
			}
		})
	}
}

func Test_skipQueryDDL(t *testing.T) {
	rawReplicateDoDb := []*common.DataSource{
		{
			TableSchema: "db1",
			Tables: []*common.Table{
				{
					TableName: "tb1",
				},
				{
					TableName: "tb2",
				},
			},
		},
		{
			TableSchema: "db2",
		},
		{
			TableSchema: "db3",
			Tables: []*common.Table{
				{
					TableName: "tb1",
				},
			},
		},
		{
			TableSchema: "db4",
			Tables: []*common.Table{
				{
					TableName: "tb1",
				},
			},
		},
	}

	rawReplicateIgnoreDb := []*common.DataSource{
		{
			TableSchema: "db1",
			Tables: []*common.Table{
				{
					TableName: "tb1",
				},
			},
		},
		{
			TableSchema: "db2",
			Tables: []*common.Table{
				{
					TableName: "tb-skip",
				},
			},
		},
		{
			TableSchema: "db3",
		},
		{
			TableSchema: "db4",
			Tables: []*common.Table{
				{
					TableName: "tb1",
				},
			},
		},
	}

	type args struct {
		schemaName string
		tableName  string
	}
	tests := []struct {
		name                                 string
		args                                 args
		wantResult                           bool
		wantResultWithEmptyReplicateIgnoreDb bool
		wantResultWithEmptyReplicateDoDb     bool
	}{
		{
			name: "replicateDoDb-tables/replicateIgnoreDb-table/input-schema",
			args: args{
				schemaName: "db1",
			},
			wantResult:                           false,
			wantResultWithEmptyReplicateDoDb:     false,
			wantResultWithEmptyReplicateIgnoreDb: false,
		},
		{
			name: "replicateDoDb-tables/replicateIgnoreDb-table/input-table",
			args: args{
				schemaName: "db1",
				tableName:  "tb1",
			},
			wantResult:                           true,
			wantResultWithEmptyReplicateDoDb:     true,
			wantResultWithEmptyReplicateIgnoreDb: false,
		},
		{
			name: "replicateDoDb-tables/replicateIgnoreDb-table/input-table",
			args: args{
				schemaName: "db1",
				tableName:  "tb2",
			},
			wantResult:                           false,
			wantResultWithEmptyReplicateDoDb:     false,
			wantResultWithEmptyReplicateIgnoreDb: false,
		},
		{
			name: "replicateDoDb-schema/replicateIgnoreDb-table/input-schema",
			args: args{
				schemaName: "db2",
			},
			wantResult:                           false,
			wantResultWithEmptyReplicateDoDb:     false,
			wantResultWithEmptyReplicateIgnoreDb: false,
		},
		{
			name: "replicateDoDb-schema/replicateIgnoreDb-table/input-table",
			args: args{
				schemaName: "db2",
				tableName:  "tb-skip",
			},
			wantResult:                           true,
			wantResultWithEmptyReplicateDoDb:     true,
			wantResultWithEmptyReplicateIgnoreDb: false,
		},
		{
			name: "replicateDoDb-table/replicateIgnoreDb-schema/input-schema",
			args: args{
				schemaName: "db3",
			},
			wantResult:                           true,
			wantResultWithEmptyReplicateDoDb:     true,
			wantResultWithEmptyReplicateIgnoreDb: false,
		},
		{
			name: "replicateDoDb-table/replicateIgnoreDb-schema/input-table",
			args: args{
				schemaName: "db3",
				tableName:  "tb1",
			},
			wantResult:                           true,
			wantResultWithEmptyReplicateDoDb:     true,
			wantResultWithEmptyReplicateIgnoreDb: false,
		},
		{
			name: "replicateDoDb-table/replicateIgnoreDb-table/input-table",
			args: args{
				schemaName: "db4",
				tableName:  "tb1",
			},
			wantResult:                           true,
			wantResultWithEmptyReplicateDoDb:     true,
			wantResultWithEmptyReplicateIgnoreDb: false,
		},
		{
			name: "replicateDoDb-table/replicateIgnoreDb-table/input-schema",
			args: args{
				schemaName: "db4",
			},
			wantResult:                           false,
			wantResultWithEmptyReplicateDoDb:     false,
			wantResultWithEmptyReplicateIgnoreDb: false,
		},
		{
			name: "not-defined-in-config/input-schema",
			args: args{
				schemaName: "db5",
			},
			wantResult:                           true,
			wantResultWithEmptyReplicateDoDb:     false,
			wantResultWithEmptyReplicateIgnoreDb: true,
		},
		{
			name: "not-defined-in-config/input-table",
			args: args{
				schemaName: "db5",
				tableName:  "tb1",
			},
			wantResult:                           true,
			wantResultWithEmptyReplicateDoDb:     false,
			wantResultWithEmptyReplicateIgnoreDb: true,
		},
	}

	binlogReader := &BinlogReader{
		mysqlContext: &common.MySQLDriverConfig{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			binlogReader.mysqlContext.ReplicateIgnoreDb = rawReplicateIgnoreDb
			binlogReader.mysqlContext.ReplicateDoDb = rawReplicateDoDb
			res := binlogReader.skipQueryDDL(tt.args.schemaName, tt.args.tableName)
			if res != tt.wantResult {
				t.Errorf("skipQueryDDL() gotResult = %v, want %v", res, tt.wantResult)
			}

			binlogReader.mysqlContext.ReplicateDoDb = []*common.DataSource{}
			res = binlogReader.skipQueryDDL(tt.args.schemaName, tt.args.tableName)
			if res != tt.wantResultWithEmptyReplicateDoDb {
				t.Errorf("skipQueryDDL() with empty replicateDoDb gotResult = %v, want %v", res, tt.wantResult)
			}

			binlogReader.mysqlContext.ReplicateIgnoreDb = []*common.DataSource{}
			res = binlogReader.skipQueryDDL(tt.args.schemaName, tt.args.tableName)
			if res {
				t.Errorf("skipQueryDDL() with empty replicateDoDb and ReplicateIgnoreDb gotResult = %v, want false", res)
			}

			binlogReader.mysqlContext.ReplicateDoDb = rawReplicateDoDb
			res = binlogReader.skipQueryDDL(tt.args.schemaName, tt.args.tableName)
			if res != tt.wantResultWithEmptyReplicateIgnoreDb {
				t.Errorf("skipQueryDDL() with empty ReplicateIgnoreDb gotResult = %v, want %v", res, tt.wantResultWithEmptyReplicateIgnoreDb)
			}
		})
	}
}

//func Test_updateCurrentReplicateDoDb(t *testing.T) {
//	tableConfigs := []*common.Table{
//		{TableName: "tb1", TableRename: "tb1-rename"},
//		{TableName: "tb2", TableRename: ""},
//		{TableRegex: "(\\w*)tb-rex1", TableRename: "tb${1}-rex1-rename"},
//		{TableRegex: "(\\w*)tb-rex2", TableRename: "tb${1}-rex2-rename"},
//		{TableRegex: "(\\w*)tb-rex3", TableRename: ""},
//	}
//
//	rawReplicateDoDb := []*common.DataSource{
//		{TableSchema: "db1", TableSchemaRename: "db1-rename", Tables: tableConfigs},
//		{TableSchema: "db2", TableSchemaRename: "db2-rename", Tables: []*common.Table{}},
//		{TableSchema: "db3", TableSchemaRename: "", Tables: tableConfigs},
//		{TableSchemaRegex: "(\\w*)db-rex1", TableSchemaRename: "db${1}-rex1-rename", Tables: tableConfigs},
//		{TableSchemaRegex: "(\\w*)db-rex2", TableSchemaRename: "db${1}-rex2-rename", Tables: []*common.Table{}},
//		{TableSchemaRegex: "(\\w*)db-rex3", TableSchemaRename: "", Tables: tableConfigs},
//	}
//
//	type args struct {
//		schema    string
//		tableName string
//	}
//
//	tests := []struct {
//		name                 string
//		rawReplicateDoDb     []*common.DataSource
//		currentReplicateDoDb []*common.DataSource
//		args                 args
//		want                 []*common.DataSource
//	}{
//		{
//			name:                 "empty-rawReplicateDoDb/empty-currentReplicateDoDb/input-new-table",
//			rawReplicateDoDb:     []*common.DataSource{},
//			currentReplicateDoDb: []*common.DataSource{},
//			args:                 args{schema: "db1", tableName: "tb1"},
//			want: []*common.DataSource{
//				{
//					TableSchema: "db1",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableSchema: "db1", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:             "empty-rawReplicateDoDb/exists-one-table/input-new-table",
//			rawReplicateDoDb: []*common.DataSource{},
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema: "db1",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableSchema: "db1", Where: "true"},
//					},
//				}},
//			args: args{schema: "db1", tableName: "tb2"},
//			want: []*common.DataSource{
//				{
//					TableSchema: "db1",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableSchema: "db1", Where: "true"},
//						{TableName: "tb2", TableSchema: "db1", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:             "empty-rawReplicateDoDb/exists-one-table/input-existed-table",
//			rawReplicateDoDb: []*common.DataSource{},
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema: "db1",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableSchema: "db1", Where: "true"},
//					},
//				}},
//			args: args{schema: "db1", tableName: "tb1"},
//			want: []*common.DataSource{
//				{
//					TableSchema: "db1",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableSchema: "db1", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:             "empty-rawReplicateDoDb/exists-one-schema/input-existed-schema",
//			rawReplicateDoDb: []*common.DataSource{},
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema: "db1",
//				},
//			},
//			args: args{schema: "db1", tableName: ""},
//			want: []*common.DataSource{
//				{
//					TableSchema: "db1",
//				},
//			},
//		},
//		{
//			name:             "empty-rawReplicateDoDb/exists-one-schema/input-new-table-to-existed-schema",
//			rawReplicateDoDb: []*common.DataSource{},
//			currentReplicateDoDb: []*common.DataSource{{
//				TableSchema: "db1",
//			}},
//			args: args{schema: "db1", tableName: "tb1"},
//			want: []*common.DataSource{
//				{
//					TableSchema: "db1",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableSchema: "db1", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:                 "empty-currentReplicateDoDb/input-table",
//			rawReplicateDoDb:     rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{},
//			args:                 args{schema: "db1", tableName: "tb1"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableRename: "tb1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:                 "empty-currentReplicateDoDb/input-table-match-regex",
//			rawReplicateDoDb:     rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{},
//			args:                 args{schema: "db1", tableName: "testtb-rex1"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "testtb-rex1", TableRegex: "(\\w*)tb-rex1", TableRename: "tbtest-rex1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:                 "empty-currentReplicateDoDb/input-table-match-regex-without-rename",
//			rawReplicateDoDb:     rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{},
//			args:                 args{schema: "testdb-rex3", tableName: "testtb-rex3"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "testdb-rex3",
//					TableSchemaRegex:  "(\\w*)db-rex3",
//					TableSchemaRename: "",
//					Tables: []*common.Table{
//						{TableName: "testtb-rex3", TableRegex: "(\\w*)tb-rex3", TableRename: "", TableSchema: "testdb-rex3", TableSchemaRename: "", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:                 "empty-currentReplicateDoDb/input-schema",
//			rawReplicateDoDb:     rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{},
//			args:                 args{schema: "db1", tableName: ""},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//				},
//			},
//		},
//		{
//			name:                 "empty-currentReplicateDoDb/input-not-match-table",
//			rawReplicateDoDb:     rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{},
//			args:                 args{schema: "db1", tableName: "tb-not-match"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//				},
//			},
//		},
//		{
//			name:                 "empty-currentReplicateDoDb/input-not-match-schema",
//			rawReplicateDoDb:     rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{},
//			args:                 args{schema: "db-not-match", tableName: "tb1"},
//			want:                 []*common.DataSource{},
//		},
//		{
//			name:             "input-existed-table",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableRename: "tb1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//			args: args{schema: "db1", tableName: "tb1"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableRename: "tb1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:             "input-existed-table-match-regex",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "testtb-rex1", TableRegex: "(\\w*)tb-rex1", TableRename: "tbtest-rex1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//			args: args{schema: "db1", tableName: "testtb-rex1"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "testtb-rex1", TableRegex: "(\\w*)tb-rex1", TableRename: "tbtest-rex1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:             "input-new-table",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableRename: "tb1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//			args: args{schema: "db1", tableName: "tb2"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableRename: "tb1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//						{TableName: "tb2", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			// a table has existed and then add new table that match the same regex
//			name:             "input-new-table-match-regex",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "testtb-rex1", TableRegex: "(\\w*)tb-rex1", TableRename: "tbtest-rex1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//			args: args{schema: "db1", tableName: "test2tb-rex1"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "testtb-rex1", TableRegex: "(\\w*)tb-rex1", TableRename: "tbtest-rex1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//						{TableName: "test2tb-rex1", TableRegex: "(\\w*)tb-rex1", TableRename: "tbtest2-rex1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			// a table has existed and then add new table that match the different regex
//			name:             "input-new-table-match-regex",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "testtb-rex1", TableRegex: "(\\w*)tb-rex1", TableRename: "tbtest-rex1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//			args: args{schema: "db1", tableName: "testtb-rex2"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "testtb-rex1", TableRegex: "(\\w*)tb-rex1", TableRename: "tbtest-rex1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//						{TableName: "testtb-rex2", TableRegex: "(\\w*)tb-rex2", TableRename: "tbtest-rex2-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:             "input-not-match-table",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableRename: "tb1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//			args: args{schema: "db1", tableName: "tb-not-match"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableRename: "tb1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//		},
//		{
//			name:             "input-existed-schema",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//				},
//			},
//			args: args{schema: "db1", tableName: ""},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//				},
//			},
//		},
//		{
//			// regex
//			name:             "input-existed-schema-match-regex",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "testdb-rex1",
//					TableSchemaRegex:  "(\\w*)db-rex1",
//					TableSchemaRename: "dbtest-rex1-rename",
//				},
//			},
//			args: args{schema: "testdb-rex1", tableName: ""},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "testdb-rex1",
//					TableSchemaRegex:  "(\\w*)db-rex1",
//					TableSchemaRename: "dbtest-rex1-rename",
//				},
//			},
//		},
//		{
//			name:             "input-new-schema",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//				},
//			},
//			args: args{schema: "db2"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//				},
//				{
//					TableSchema:       "db2",
//					TableSchemaRename: "db2-rename",
//				},
//			},
//		},
//		{
//			// a schema has existed and then add new schema that match the same regex
//			name:             "input-new-schema-match-regex",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "test1db-rex1",
//					TableSchemaRegex:  "(\\w*)db-rex1",
//					TableSchemaRename: "dbtest1-rex1-rename",
//				},
//			},
//			args: args{schema: "test2db-rex1"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "test1db-rex1",
//					TableSchemaRegex:  "(\\w*)db-rex1",
//					TableSchemaRename: "dbtest1-rex1-rename",
//				},
//				{
//					TableSchema:       "test2db-rex1",
//					TableSchemaRegex:  "(\\w*)db-rex1",
//					TableSchemaRename: "dbtest2-rex1-rename",
//				},
//			},
//		},
//		{
//			// a schema has existed and then add new schema that match the different regex
//			name:             "input-new-schema-match-regex",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "test1db-rex1",
//					TableSchemaRegex:  "(\\w*)db-rex1",
//					TableSchemaRename: "dbtest1-rex1-rename",
//				},
//			},
//			args: args{schema: "test1db-rex2"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "test1db-rex1",
//					TableSchemaRegex:  "(\\w*)db-rex1",
//					TableSchemaRename: "dbtest1-rex1-rename",
//				},
//				{
//					TableSchema:       "test1db-rex2",
//					TableSchemaRegex:  "(\\w*)db-rex2",
//					TableSchemaRename: "dbtest1-rex2-rename",
//				},
//			},
//		},
//		{
//			name:             "input-not-match-table",
//			rawReplicateDoDb: rawReplicateDoDb,
//			currentReplicateDoDb: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableRename: "tb1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//			args: args{schema: "db1", tableName: "tb-not-match"},
//			want: []*common.DataSource{
//				{
//					TableSchema:       "db1",
//					TableSchemaRename: "db1-rename",
//					Tables: []*common.Table{
//						{TableName: "tb1", TableRename: "tb1-rename", TableSchema: "db1", TableSchemaRename: "db1-rename", Where: "true"},
//					},
//				},
//			},
//		},
//	}
//
//	binlogReader := &BinlogReader{
//		logger: hclog.New(&hclog.LoggerOptions{
//			Level:      hclog.Debug,
//			JSONFormat: true,
//		}),
//		mysqlContext: &common.MySQLDriverConfig{},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			binlogReader.mysqlContext.ReplicateDoDb = test.rawReplicateDoDb
//			//binlogReader.currentReplicateDoDb = test.currentReplicateDoDb
//			if err := binlogReader.updateCurrentReplicateDoDb(test.args.schema, test.args.tableName); nil != err {
//				t.Error(err)
//				return
//			}
//
//			if !assert.Equal(t, test.want, binlogReader.currentReplicateDoDb, "unexpected currentReplicateDoDb") {
//				printObject := func(db *SchemaContext, prefix string) {
//					t.Errorf("%v: &{TableSchema: %v TableSchemaRename: %v}\n", prefix, db.TableSchema, db.TableSchemaRename)
//					for _, tb := range db.TableMap {
//						t.Errorf("table: %+v\n", tb.Table)
//					}
//				}
//
//				for _, db := range binlogReader.tables {
//					printObject(db, "got current db")
//				}
//
//				for _, db := range test.want {
//					printObject(db, "want db")
//				}
//			}
//		})
//
//	}
//
//}

func Test_isSkipQuery(t *testing.T) {
	type args struct {
		sql string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "create-event",
			args: args{"create event if not exists a.event1 on schedule every 5 second on completion preserve do insert into a.a values (0);"},
			want: true,
		}, {
			name: "create-trigger",
			args: args{`CREATE TRIGGER before_employee_update
	BEFORE UPDATE ON employees FOR EACH ROW
	INSERT INTO employees_audit
	SET action = 'update',
		employeeNumber = OLD.employeeNumber,
		lastname = OLD.lastname,
		changedat = NOW();`},
			want: true,
		}, {
			name: "alter-event",
			args: args{`ALTER EVENT no_such_event ON SCHEDULE EVERY '2:3' DAY_HOUR;`},
			want: true,
		}, {
			name: "create-table",
			args: args{"create table a.`create event on schedule at do insert` (id int)"},
			want: true, // TODO It should be false. Regex cannot handle such a case.
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSkipQuery(tt.args.sql); got != tt.want {
				t.Errorf("isSkipQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isExpandSyntaxQuery(t *testing.T) {
	type args struct {
		sql string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "create-procedure-1",
			args: args{sql: "CREATE DEFINER=`root`@`%` PROCEDURE p1 begin end"},
			want: true,
		},
		{
			name: "create-procedure-qq-20220223",
			args: args{sql: "CREATE DEFINER=`tester`@`%` PROCEDURE `p_tb_test_update`()\nBEGIN\n\n-- 随机取一条记录\nSELECT @id:=t1.id\nFROM tb_test AS t1 \nJOIN (SELECT ROUND( RAND()*(SELECT MAX(id)-MIN(id) FROM tb_test) + (SELECT MIN(id) FROM tb_test) ) AS id) AS t2\nWHERE t1.id >= t2.id\n  AND t1.state = 1\nORDER BY t1.id \nLIMIT 1;\n\n-- SELECT @id;\n\n-- SELECT * FROM tb_test WHERE id=@id;\nUPDATE tb_test SET amount = FLOOR(POWER(RAND(),2)*1000000) WHERE id=@id;\n\nEND\" schema=testdb @module=reader job=job_testdb timestamp=2022-02-23T16:21:48.026+0800\n2022-02-23T16:21:48.035+0800 [WARN]  client.driver_mgr.dtle: mysql.reader: QueryEvent is not recognized. will still execute: driver=dtle @module=reader gno=407506 job=job_testdb query=\"CREATE DEFINER=`tester`@`%` PROCEDURE `p_tb_test_update`()\nBEGIN\n\n-- 随机取一条记录\nSELECT @id:=t1.id\nFROM tb_test AS t1 \nJOIN (SELECT ROUND( RAND()*(SELECT MAX(id)-MIN(id) FROM tb_test) + (SELECT MIN(id) FROM tb_test) ) AS id) AS t2\nWHERE t1.id >= t2.id\n  AND t1.state = 1\nORDER BY t1.id \nLIMIT 1;\n\n-- SELECT @id;\n\n-- SELECT * FROM tb_test WHERE id=@id;\nUPDATE tb_test SET amount = FLOOR(POWER(RAND(),2)*1000000) WHERE id=@id;\n\nEND"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isExpandSyntaxQuery(tt.args.sql); got != tt.want {
				t.Errorf("isExpandSyntaxQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
