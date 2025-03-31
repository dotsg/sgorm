package structs

import "sort"

/*
desc users;
+------------+--------------+------+------+---------+----------------+
| Field      | Type         | Null | Key  | Default | Extra          |
+------------+--------------+------+------+---------+----------------+
| id         | int          | NO   | PRI  |         | auto_increment |
| name       | varchar(255) | NO   |      | ""      |                |
| email      | varchar(255) | NO   | MUL  | ""      |                |
| created_at | int unsigned | NO   |      | "0"     |                |
| updated_at | int unsigned | NO   |      | "0"     |                |
+------------+--------------+------+------+---------+----------------+
*/

// RowDesc defines a row in mysql's desc command
type RowDesc struct {
	Field, Type, Null, Key, Default, Extra string
}

/*
show index from users;
+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
| Table | Non_unique | Key_name | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment | Visible | Expression |
+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
| users |          0 | email    |            1 | email       | NULL      |           0 |     NULL | NULL   |      | BTREE      |         |               | YES     | NULL       |
| users |          1 | name     |            1 | name        | NULL      |           0 |     NULL | NULL   |      | BTREE      |         |               | YES     | NULL       |
+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
*/

// IndexDesc defines a row in mysql's `show index from xxx` command
type IndexDesc struct {
	Table, KeyName, ColumnName, Collation, SubPart, Packed, Null, IndexType, Comment, IndexComment, Visible, Expression string
	NonUnique, SeqInIndex, Cardinality                                                                                  int
}

func GroupIndex(indexDesc []*IndexDesc) map[string][]*IndexDesc {
	data := make(map[string][]*IndexDesc, 0)

	for _, idx := range indexDesc {
		key := idx.KeyName
		if _, ok := data[key]; !ok {
			data[key] = []*IndexDesc{}
		}
	}

	for name := range data {
		items := FilterBy(indexDesc, func(item *IndexDesc) bool {
			return item.KeyName == name
		})

		sort.Slice(items, func(i, j int) bool {
			return items[i].SeqInIndex < items[j].SeqInIndex
		})

		data[name] = items

	}

	return data
}
