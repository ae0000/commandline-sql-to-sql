package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocraft/web"
)

const (
	defaultRawSQL = `mysql > select * from SomeTable where ID = 9086;
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
| pe_id | p_id | pa_id | pd_id | pe_order | pe_statementdescription | pe_timestamp        | pe_status |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+
|  1302 | 9046 |     0 |   518 |        1 |                         | 2015-07-06 05:49:10 | active    |
+-------+------+-------+-------+----------+-------------------------+---------------------+-----------+`
)

func (c *Context) convertSQL(rw web.ResponseWriter, req *web.Request) {

	// Check for POST
	if req.Request.Method == "POST" {
		req.ParseForm()
		c.RawSQL = req.Request.FormValue("clsql")
		c.ConvertedSQL = Convert(c.RawSQL)
	} else {
		// Add default
		c.RawSQL = defaultRawSQL
	}

	rend.HTML(rw, http.StatusOK, "index", c)
}

// Convert takes command line spewed sql and converts it back into a INSERT sql
// string. Sometimes.
func Convert(raw string) string {
	var table, o, fieldLine, valueLine string
	var sep = ", "

	// Work out table.. we are assuming there is a "select XXX from table" line
	// in the raw string, if not default to "SomeTable". Also get rid of ";"
	ss := strings.Split(strings.Replace(raw, ";", " ", -1), " ")

	for i, v := range ss {
		if strings.ToLower(v) == "from" {
			table = ss[i+1]
			break
		}
	}

	if len(table) == 0 {
		table = "SomeTable"
	}

	o = fmt.Sprintf("INSERT INTO `%s`\n(", table)

	// Work out which lines are the field and value ones. They should be the
	// two lines starting with "|"
	lines := strings.Split(raw, "\n")

	for _, l := range lines {
		if len(l) > 2 && l[0:2] == "| " {
			if len(fieldLine) == 0 {
				// This is the field line
				fieldLine = l
				continue
			}
			if len(fieldLine) > 0 && len(valueLine) == 0 {
				// This is the value line
				valueLine = l
				break
			}
		}
	}

	// Add fields
	if len(fieldLine) > 0 {
		fields := strings.Split(fieldLine, "|")

		for _, v := range fields {
			field := strings.TrimSpace(v)
			if len(field) > 0 {
				o += fmt.Sprintf("`%s`%s", field, sep)
			}
		}

		// Trim ", "
		o = strings.TrimRight(o, sep)
	}

	// Add closing ")"
	o += ")"

	// Values
	o += "\nVALUES\n("

	// Add values
	if len(valueLine) > 0 {
		values := strings.Split(valueLine, "|")

		for _, v := range values {
			value := strings.TrimSpace(v)
			if len(value) > 0 {
				o += strconv.Quote(value) + sep
			}
		}

		// Trim ", "
		o = strings.TrimRight(o, sep)
	}

	// Add closing ")"
	o += ");"

	return o
}
