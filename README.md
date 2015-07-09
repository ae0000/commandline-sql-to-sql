This converts command line boxed sql to insert sql....
======================================================

So..

This:
```
mysql > select * from SomeTable where ID = 9086;
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
| pe_id | p_id | pa_id | pd_id | pe_order | pe_statementdescription | pe_timestamp        | pe_status |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
|  1302 | 9046 |     0 |   518 |        1 |                         | 2015-07-06 05:49:10 | active    |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
```

Gets converted into:

```sql
INSERT INTO `SomeTable`
(`pe_id`, `p_id`, `pa_id`, `pd_id`, `pe_order`, `pe_statementdescription`, `pe_timestamp`, `pe_status`)
VALUES
("1302", "9046", "0", "518", "1", "2015-07-06 05:49:10", "active");
```

It doesn't check for errors. It doesn't check that you have the right number of fields to values or anything like that.


It just dumps out what might be a reasonable representation of the entered text. Maybe.


Feel free to do what you want with it. I whipped this up in a night as I was getting sick of doing the same thing by hand.. ugggh.
