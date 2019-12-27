package main
const port=4040
import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)
func main() {


	c := sPrintf("localhost:%d",port)
	conn, err := grpc.Dial(a,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context){
        a,err := strconv.ParseUint(ctx.Param("a", 10 ,64)
        if err!=nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid paramet A"})
            return
        }

        b,err := strconv.ParseUint(ctx.Param("a", 10 ,64)
        if err!=nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid paramet B"})
            return
        }


        req := &proto.Request{A:int64(a),B:int64(b)}
        if response, err := client.Add(ctx, req); err==nil {
            ctx.JSON(http.StatusOK, gin.H {
                "result": fmt.Sprint(response.Result),
            }
         else {
            ctx.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error()
            })
	g.GET("/mult/:a/:b", func(ctx *gin.Context){

	})
}