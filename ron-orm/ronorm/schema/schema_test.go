package schema

import (
	"ronorm/ronorm/dialect"
	"testing"
)

type User struct {
	Name string `ronorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
