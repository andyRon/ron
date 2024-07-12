package session

import (
	"ronorm/ronorm"
	"testing"
)

type User struct {
	Name string `ronorm:"PRIMARY KEY"`
	Age  int
}

// TODO
func TestSession_CreateTable(t *testing.T) {
	s := ronorm.Engine.NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
