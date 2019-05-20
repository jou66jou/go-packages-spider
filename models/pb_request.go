package models

import (
	"context"
	"encoding/json"
	"fmt"
	pb "gopkg-spider/protoserver/echospec"

	"google.golang.org/grpc"
)

func RequestToPB(name string) string {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("連線失敗：%v", err)
	}
	defer conn.Close()

	c := pb.NewEchoPkgClient(conn)

	r, err := c.EchoPkg(context.Background(), &pb.GOPkgRequest{PkgName: name})
	if err != nil {
		fmt.Printf("無法執行 Plus 函式：%v \n", err)
	}
	jFmt, _ := json.Marshal(r)
	return string(jFmt)
}
