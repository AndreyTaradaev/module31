package storage

import (
"db/pkg/typetask"	
)

type DBInterface interface{
	Tasks(int,int,int) ([]typetask.Task,error)
	NewTask(typetask.Task)(int,error)
	DeleteTask(int) error
	UpdateTask(typetask.Task) (int,error)
	Close()
}
