package main

import (
	"db/pkg/typetask"

	"db/pkg/storage"
	"db/pkg/storage/postgres"
	"fmt"
	"log"
	//	"os"
)

var db storage.DBInterface

func main() {
	var err error
	/* pwd := os.Getenv("dbpass")
	if pwd == "" {
		os.Exit(1)
	} */
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	connectstr := "postgres://postgres:1C_Db_post@172.16.0.44:5432/skill31"
	dbc, err := postgres.New(connectstr)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}
	defer dbc.Close()
	fmt.Println("\nAll task\n")
	ts, err := dbc.Tasks(0, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range ts {
		fmt.Println(v)
	}

	fmt.Println("\ntasks label 1\n")
	ts, err = dbc.Tasks(0, 0, 1)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range ts {
		fmt.Println(v)
	}

	fmt.Println("\ntasks author  1\n")
	ts, err = dbc.Tasks(0, 1, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range ts {
		fmt.Println(v)
	}

	fmt.Println("\ntasks id  2\n")
	ts, err = dbc.Tasks(2, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range ts {
		fmt.Println(v)
	}

	fmt.Println("\nadd new task\n")
	tasknew := typetask.New(0, 0, 0, 1, 3, "Важная Задача", "Изучить pgx")

	ret, err := dbc.NewTask(*tasknew)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nadded new task ", ret)
	fmt.Println("\nPrint new task ")
	ts, err = dbc.Tasks(ret, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range ts {
		tasknew = &v
		fmt.Println(v)
	}
	fmt.Println("\nchange  task id ", ret)
	str := "Совсем  очень-очень Важная Задача"
	tasknew.Title =  &str
ret,err = dbc.UpdateTask(*tasknew)
if err != nil{
	log.Fatal(err)
}
fmt.Println("\nchanged  task id ", ret)
ts, err = dbc.Tasks(ret, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range ts {
		tasknew = &v
		fmt.Println(v)
	}

	fmt.Println("\ndelete  task id ", ret)	
	err = dbc.DeleteTask(ret)
	if err != nil {
		log.Fatal(err)
	}

	ts, err = dbc.Tasks(0, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range ts {
		tasknew = &v
		fmt.Println(v)
	}

}
