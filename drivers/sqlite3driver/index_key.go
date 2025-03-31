package sqlite3driver

import (
	"strings"

	"github.com/olachat/gola/structs"
)

func (s SQLiteDriver) SetIndexAndKey(tables []*structs.Table) (err error) {
	for _, t := range tables {
		idxs, err := s.indexes(t.Name)
		if err != nil {
			return err
		}
		n := 0
		for _, idx := range idxs {
			n += len(idx.Columns)
		}

		indexDesc := make([]*structs.IndexDesc, 0, n)
		for _, idx := range idxs {
			// skip primary key
			if strings.HasPrefix(idx.Name, "sqlite_autoindex") {
				continue
			}
			for _, col := range idx.Columns {
				item := new(structs.IndexDesc)
				item.Table = t.Name
				item.SeqInIndex = idx.SeqNum
				if idx.Unique == 1 {
					item.NonUnique = 0
				} else {
					item.NonUnique = 1
				}
				item.ColumnName = col
				item.KeyName = idx.Name
				indexDesc = append(indexDesc, item)
			}
		}
		t.Indexes = structs.GroupIndex(indexDesc)
	}

	return nil
}
