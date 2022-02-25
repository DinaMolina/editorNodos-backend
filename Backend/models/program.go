package models

type Node struct {
	Id   int64
	Type int16
	Next int64
}

type Program struct {
	Id    int64
	Name  string
	Nodes []Node
}
