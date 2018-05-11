package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var cmd string
		var name string
		var id int
		fmt.Print(">")
		scanner.Scan()
		line := scanner.Text()
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "add":
			fmt.Sscan(line, &cmd, &id, &name)
			add(id, name)
		case "del":
			fmt.Sscan(line, &cmd, &name)
			del(name)
		case "list":
			list()
		case "save":
			var filename string
			fmt.Sscan(line, &cmd, &filename)
			save(filename)
		case "load":
			var filename string
			fmt.Sscan(line, &cmd, &filename)
			load(filename)
		case "exit":
			return
		default:
			usage()
		}
	}
}

func add(id int, name string) {
	if _, ok := stuMap[name]; ok {
		fmt.Printf("Duplicate name: %s\n", name)
		return
	}
	student := Student{Id: id, Name: name}
	stuMap[name] = student
}

func list() {
	fmt.Printf("ID\tNAME\n")
	for _, stu := range stuMap {
		fmt.Printf("%d\t%s\n", stu.Id, stu.Name)
	}
}

func del(name string) {
	delete(stuMap, name)
}

func save(filename string) {
	f, _ := os.Create(filename)
	defer f.Close()

	buf, _ := json.Marshal(stuMap)
	f.Write(buf)
}

func load(filename string) {
	buf, _ := ioutil.ReadFile(filename)
	json.Unmarshal(buf, &stuMap)
}

func usage() {}
