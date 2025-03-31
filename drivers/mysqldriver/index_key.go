package mysqldriver

import (
	"github.com/olachat/gola/structs"
)

// SetIndexAndKey sets indexes & key information to given tables
func (m *MySQLDriver) SetIndexAndKey(tables []*structs.Table) (err error) {
	for _, t := range tables {
		var tableDesc []*structs.RowDesc
		rows, err := m.conn.Query("desc `" + t.Name + "`")
		if err != nil {
			return err
		}

		for rows.Next() {
			rd := new(structs.RowDesc)
			rows.Scan(&rd.Field, &rd.Type, &rd.Null, &rd.Key, &rd.Default, &rd.Extra)
			tableDesc = append(tableDesc, rd)
		}

		for _, rd := range tableDesc {
			if rd.Key == "PRI" {
				if t.PKey == nil {
					t.PKey = &structs.PrimaryKey{}
					t.PKey.Name = rd.Field
					t.PKey.Columns = []string{rd.Field}
				} else {
					t.PKey.Columns = append(t.PKey.Columns, rd.Field)
				}
			}
		}

		var indexDesc []*structs.IndexDesc
		rows, err = m.conn.Query("show index from `" + t.Name + "`")
		if err != nil {
			return err
		}

		hasKeyInShowIndex := false
		for rows.Next() {
			id := new(structs.IndexDesc)
			rows.Scan(&id.Table, &id.NonUnique, &id.KeyName, &id.SeqInIndex, &id.ColumnName, &id.Collation, &id.Cardinality,
				&id.SubPart, &id.Packed, &id.Null, &id.IndexType, &id.Comment, &id.IndexComment, &id.Visible, &id.Expression)

			indexDesc = append(indexDesc, id)

			if id.KeyName == "PRIMARY" {
				hasKeyInShowIndex = true
			}
		}

		// Work around for bug in github.com/dolthub/go-mysql-server v0.12.0
		// Which doesn't return PRIMARY info in show index from XXX query
		// https://github.com/dolthub/go-mysql-server/issues/1157
		if !hasKeyInShowIndex && t.PKey != nil && len(t.PKey.Columns) > 1 {
			for i, col := range t.PKey.Columns {
				id := new(structs.IndexDesc)
				id.Table = t.Name
				id.NonUnique = 1
				id.KeyName = "PRIMARY"
				id.ColumnName = col
				id.SeqInIndex = i + 1
				indexDesc = append(indexDesc, id)
			}
		}

		t.Indexes = structs.GroupIndex(indexDesc)

		// Hack to handle table without primary key, but has only one unique key
		// Just consider that unique key as primary
		if t.PKey == nil {
			uniqueIdxCount := 0
			var uniqueIdx []*structs.IndexDesc
			var uniqueIdxName string
			for idxName, idx := range t.Indexes {
				isUnique := true
				for _, node := range idx {
					if node.NonUnique > 0 {
						isUnique = false
					}
				}
				if isUnique {
					uniqueIdxCount += 1
					uniqueIdx = idx
					uniqueIdxName = idxName
				}
			}

			if uniqueIdxCount == 1 {
				t.PKey = &structs.PrimaryKey{}
				t.PKey.Name = uniqueIdxName
				t.PKey.Columns = make([]string, len(uniqueIdx))
				for i, node := range uniqueIdx {
					t.PKey.Columns[i] = node.ColumnName
				}
			}
		}
	}

	return nil
}
