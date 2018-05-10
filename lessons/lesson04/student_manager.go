package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Id   int
	Name string
}

var stuMap = make(map[string]Student)

func main() {
	for {
		var cmd string
		fmt.Println("Input:")
		fmt.Scan(&cmd)
		switch cmd {
		case "add":
			add()
		case "del":
			del()
		case "list":
			list()
		case "save":
			save()
		case "load":
			load()
		case "exit":
			return
		default:
			usage()
		}
	}
}

func add() {
	var id int
	var name string
	fmt.Scan(&id, &name)
	student := Student{Id: id, Name: name}
	stuMap[name] = student
}

func list() {
	fmt.Printf("id\tname\n")
	for _, stu := range stuMap {
		fmt.Printf("%d\t%s\n", stu.Id, stu.Name)
	}
}

func del() {
	var name string
	fmt.Scan(&name)
	delete(stuMap, name)
}

func save() {
	var filename string
	fmt.Scan(&filename)
	f, _ := os.Create(filename)
	defer f.Close()

	buf, _ := json.Marshal(stuMap)
	f.Write(buf)
}

func load() {
	var filename string
	fmt.Scan(&filename)
	buf, _ := ioutil.ReadFile(filename)
	json.Unmarshal(buf, &stuMap)
}

func usage() {}
