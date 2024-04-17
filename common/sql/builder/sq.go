package builder

import sq "github.com/Masterminds/squirrel"

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func ToUpsertSQL(data interface{}, table string, autoSet []string, constraint []string, notUpdate []string) (string, []interface{}, error) {
	cols, vals := FieldColsAndValues(data, autoSet...)
	return psql.Insert(table).
		Columns(cols...).
		Values(vals...).
		Suffix(PostgreSqlJoinDoUpdatesetExclude(constraint, cols, notUpdate)).
		Suffix("RETURNING \"id\"").
		ToSql()
}
