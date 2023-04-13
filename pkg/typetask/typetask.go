package typetask
import (
	"fmt"
	"time"
)

type Task struct{
ID  *int
Opened *int64
Closed *int64
AuthtorId *int
AssignedId *int
Title *string	
Content *string
}

func ( t Task ) String() string{

s :=  fmt.Sprintf( "{%s %s %s %s %s %s %s}", isNull(t.ID), isNull(t.Opened),
isNull(t.Closed),isNull(t.AuthtorId),isNull(t.AssignedId),isNull(t.Title),isNull(t.Content))

return s
}

func New (id int,o,c int64, aut,ass int, tit,cont string ) * Task{

	if o==0 {
		o =time.Now().Unix()		
	}	
	task := Task{ID:&id,Opened:&o,
				Closed:&c, AuthtorId: &aut,
				AssignedId: &ass,
				Title: &tit,
				Content: &cont	}  
				
	return &task

}

func isNull(a interface{}) string{

	switch a.(type){
	case *int :
		if(a.(*int) == nil){
		 return 	fmt.Sprint("<nil>") 
		}
		 return  fmt.Sprintf("%d"  , *a.(*int))
	case *int64 :
		if(a.(*int64) == nil){
			return 	fmt.Sprint("<nil>") 
		   }
			return  fmt.Sprintf("%d"  , *a.(*int64))
	case *string  :	
	if(a.(*string) == nil){
		return 	fmt.Sprint("<nil>") 
	   }
		return  fmt.Sprintf("'%s'"  , *a.(*string))
	}
	
return ""
}
