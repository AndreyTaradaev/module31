package memdb

import (
	"db/pkg/typetask"	
	)

type MemoDb [] typetask.Task

func (m *MemoDb )Tasks(TaskId, autohor ,LabelId int) ([]typetask.Task,error){
	return []typetask.Task {},nil
}
func (m *MemoDb )	NewTask( t typetask.Task) (int,error) {
	return 1,nil
}
func (m *MemoDb )	DeleteTask(TaskId int) error{
	return nil
}
func (m *MemoDb ) 	UpdateTask( Id typetask.Task) (int,error){
	return 1,nil
}

func (m *MemoDb ) 	Close( ) {	
}
