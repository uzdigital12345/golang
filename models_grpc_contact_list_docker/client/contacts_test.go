package client

import (
	"context"
	// "fmt"
	"net"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/jmoiron/sqlx"
	ast "github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/uzdigital12345/golang/models_grpc/proto"
)

const (
	port = 4040
)


type Sqlx struct {
	fb *sqlx.DB
}

var (
	err         error
	client      pb.ContactManagerInterfaceClient
	jspbMarshal jsonpb.Marshaler
	contact     pb.ContactManagerInterfaceServer
)

func TestSqlx_Add(t *testing.T) {

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		t.Error("Connection error : ", err)
	}

	client = pb.NewContactManagerInterfaceClient(conn)
	
	assert := ast.New(t)

	testCases := []struct {
		Name string
		Req  *pb.Contact
		Resp *pb.Contact
		Err  error
	}{
		{
			Name: "Contact 1",
			Req: &pb.Contact{
				Age:    "19",
				Name:   "Enver",
				Gender: "Male",
				Number: "95218125",
			},
			Resp: &pb.Contact{
				Age:    "19",
				Name:   "Enver",
				Gender: "Male",
				Number: "95218125",
			},
			Err: nil,
		},
		{
			Name: "Contact 2",
			Req: &pb.Contact{
				Id:     2,
				Age:    "15",
				Name:   "Abdurahmon",
				Gender: "Male",
				Number: "941563245",
			},
			Resp: &pb.Contact{
				Id:     2,
				Age:    "15",
				Name:   "Abdurahmon",
				Gender: "Male",
				Number: "941563245",
			},
			Err: nil,
		},
		{
			Name: "Contact 3",
			Req: &pb.Contact{
				Id:     3,
				Age:    "85",
				Name:   "Boboy",
				Gender: "Male",
				Number: "9415625315",
			},
			Resp: &pb.Contact{
				Id:     3,
				Age:    "85",
				Name:   "Boboy",
				Gender: "Male",
				Number: "9415625315",
			},
			Err: nil,
		},
		{
			Name: "Contact 4",
			Req: &pb.Contact{
				Id:     3,
				Age:    "85",
				Name:   "Boboy",
				Gender: "Male",
				Number: "9415625315",
			},
			Resp: &pb.Contact{
				Id:     3,
				Age:    "85",
				Name:   "Boboy",
				Gender: "Male",
				Number: "9415625315",
			},
			Err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			addedContact, err := client.Add(context.Background(), testCase.Req)
			assert.Equal(testCase.Err, err)
			if err == nil {
				cl, err := jspbMarshal.MarshalToString(addedContact)
				assert.Nil(err)

				cl2, err := jspbMarshal.MarshalToString(testCase.Resp)
				assert.Nil(err)

				assert.Equal(cl, cl2)

			} else {
				assert.Equal(testCase.Resp, addedContact )
			}
		})
	}

}

// func TestSqlx_Update(t *testing.T) {

// 	c := &pb.Contact{
// 		Age:    "15",
// 		Name:   "Akbar",
// 		Gender: "Male",
// 		Number: "1581818",
// 	}

// 	contact, err := client.Update(context.Background(), c)
// 	if err != nil {
// 		t.Error("Can't update in testing...", err)
// 	}
// 	fmt.Printf("Updated item : %+v\n", contact)
// }

// func TestSqlx_Delete(t *testing.T) {
// 	//var b  Contact
// 	//c := Contact{ 5, 15, "Akbar", "male", 9898915}
// 	//
// 	//err = client.Call("Sqlx.Delete",c,&b)
// 	c := &pb.Contact{
// 		Age:    "15",
// 		Name:   "Akbar",
// 		Gender: "Male",
// 		Number: "1581818",
// 	}

// 	b, err := client.Delete(context.Background(), c)
// 	if err != nil {
// 		t.Error("Can't delete in testing..", err)
// 	}
// 	fmt.Println(b)
// }

// func TestSqlx_GetAll(t *testing.T) {
// 	b, err := client.GetAll(context.Background(), &empty.Empty{})
// 	if err != nil {
// 		t.Error("Can't get all", err)
// 	}
// 	fmt.Println(b)
// }

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}

func getGrpcConnection() (conn *grpc.ClientConn, err error) {
	ctx := context.Background()
	return grpc.DialContext(ctx, "", grpc.WithContextDialer(getBufDialer(lis)), grpc.WithInsecure())
}
