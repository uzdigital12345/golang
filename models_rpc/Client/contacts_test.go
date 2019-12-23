package Client

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/rpc"
	"testing"
)

const (
	port = 4040
	
)

type Sqlx struct {
	fb *sqlx.DB
}

var (
	err error
	reply Contact
	client *rpc.Client
	contacts []Contact
	b Contact
)

type Contact struct {
	Id int `db:"id"`
	Age int `db:"age"`
	Name string `db:"name"`
	Gender string `db:"gender"`
	PhoneNumber int `db:"number"`
}

func TestSqlx_Add(t *testing.T) {
	a:=fmt.Sprintf("localhost:%d",port)
	client, err = rpc.DialHTTP("tcp",a)

	if err != nil {
		t.Error("Connection error : ",err)
	}

	t1 := Contact{1,19,"Enver","Male",991562536}
	t2 := Contact{2,15,"Abdurahmon","Male",941563245}
	t3 := Contact{3,85,"Boboy","Male",9415625315}
	t4 := Contact{4,25,"Sardor","Male",941563515}

	err = client.Call("Sqlx.Add", t1, &reply)
	if err != nil {
		t.Error("Can't Add", err)
	}
	fmt.Printf("Reply item : %+v\n",reply)

	err = client.Call("Sqlx.Add", t2, &reply)
	if err != nil {
		t.Error("Can't Add", err)
	}
	fmt.Printf("Reply item : %+v\n",reply)

	err = client.Call("Sqlx.Add", t3, &reply)
	if err != nil {
		t.Error("Can't Add", err)
	}
	fmt.Printf("Reply item : %+v\n", reply)

	err = client.Call("Sqlx.Add", t4, &reply)
	if err != nil {
		t.Error("Can't Add", err)
	}
	fmt.Printf("Reply item : %+v\n", reply)

}

func TestSqlx_Update(t *testing.T) {

	c := Contact{ 5, 15, "Akbar", "male", 9898915}

	err = client.Call("Sqlx.Update",c,&c)
	if err != nil {
		t.Error("Can't update",err)
	}
	fmt.Println(c)
}

func TestSqlx_Delete(t *testing.T) {
	//var b  Contact
	//c := Contact{ 5, 15, "Akbar", "male", 9898915}
	//
	//err = client.Call("Sqlx.Delete",c,&b)


	c := Contact{5,15,"Akbar","Male",1561}

	err = client.Call("Sqlx.Delete",c,&b)
	if err != nil {
		t.Error("Can't delete",err)
	}
	fmt.Println(b)
}

func TestSqlx_ListAll(t *testing.T)  {

	err = client.Call("Sqlx.ListAll","",&contacts)
	if err != nil {
		t.Error("Can't list all",err)
	}
	fmt.Println(contacts)
}

func TestSqlx_GetAll(t *testing.T) {

	err = client.Call("Sqlx.GetAll","",&contacts)
	if err != nil {
		t.Error("Can't get all",err)
	}
	fmt.Println(contacts)
}
