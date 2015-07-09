package main

import "fmt"

func ExampleBasicConverter() {
	raw := `
mysql > select * from SomeTable where ID = 9086;
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
| pe_id | p_id | pa_id | pd_id | pe_order | pe_statementdescription | pe_timestamp        | pe_status |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
|  1302 | 9046 |     0 |   518 |        1 |                         | 2015-07-06 05:49:10 | active    |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
`
	fmt.Println(Convert(raw))

	// Output:
	// INSERT INTO `SomeTable`
	// (`pe_id`, `p_id`, `pa_id`, `pd_id`, `pe_order`, `pe_statementdescription`, `pe_timestamp`, `pe_status`)
	// VALUES
	// ("1302", "9046", "0", "518", "1", "2015-07-06 05:49:10", "active");
}

func ExampleScrewyTableNameConverter() {
	raw := `
mysql > select * FroM SomeTable;
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
| pe_id | p_id | pa_id | pd_id | pe_order | pe_statementdescription | pe_timestamp        | pe_status |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
|  1302 | 9046 |     0 |   518 |        1 |                         | 2015-07-06 05:49:10 | active    |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
`
	fmt.Println(Convert(raw))

	// Output:
	// INSERT INTO `SomeTable`
	// (`pe_id`, `p_id`, `pa_id`, `pd_id`, `pe_order`, `pe_statementdescription`, `pe_timestamp`, `pe_status`)
	// VALUES
	// ("1302", "9046", "0", "518", "1", "2015-07-06 05:49:10", "active");
}

func ExampleMultipleRows() {
	raw := `
mysql > select * FroM SomeTable;
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
| pe_id | p_id | pa_id | pd_id | pe_order | pe_statementdescription | pe_timestamp        | pe_status |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
|  1302 | 9046 |     0 |   518 |        1 |                         | 2015-07-06 05:49:10 | active    |
|  x    | x    |   x   |   x   |        x |             xx          |   x x x x xx x x xx | zactive   |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
`
	fmt.Println(Convert(raw))

	// Output:
	// INSERT INTO `SomeTable`
	// (`pe_id`, `p_id`, `pa_id`, `pd_id`, `pe_order`, `pe_statementdescription`, `pe_timestamp`, `pe_status`)
	// VALUES
	// ("1302", "9046", "0", "518", "1", "2015-07-06 05:49:10", "active");
}

func ExampleNoSelectTableConverter() {
	raw := `
PAH
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
| pe_id | p_id | pa_id | pd_id | pe_order | pe_statementdescription | pe_timestamp        | pe_status |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
|  1302 | 9046 |     0 |   518 |        1 |                         | 2015-07-06 05:49:10 | active    |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
`
	fmt.Println(Convert(raw))

	// Output:
	// INSERT INTO `SomeTable`
	// (`pe_id`, `p_id`, `pa_id`, `pd_id`, `pe_order`, `pe_statementdescription`, `pe_timestamp`, `pe_status`)
	// VALUES
	// ("1302", "9046", "0", "518", "1", "2015-07-06 05:49:10", "active");
}

func ExampleMissingThings() {
	raw := `
| pe_id | p_id | pa_id | pd_id | pe_order | pe_statementdescription | pe_timestamp        | pe_status |
|  1302 | 9046 |     0 |   518 |        1 |                         | 2015-07-06 05:49:10 | active    |
`
	fmt.Println(Convert(raw))

	// Output:
	// INSERT INTO `SomeTable`
	// (`pe_id`, `p_id`, `pa_id`, `pd_id`, `pe_order`, `pe_statementdescription`, `pe_timestamp`, `pe_status`)
	// VALUES
	// ("1302", "9046", "0", "518", "1", "2015-07-06 05:49:10", "active");
}
