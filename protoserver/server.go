package protoserver

import (
	"fmt"
	"log"
	"net"

	md "gopkg-spider/models"
	pb "gopkg-spider/protoserver/echospec"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type EchoPkg struct{}

func (e *EchoPkg) EchoPkg(ctx context.Context, req *pb.GOPkgRequest) (resp *pb.GOPkgReply, err error) {

	log.Printf("receive client request, client send:%s\n", req.PkgName)
	gopkg, ov := md.FindPkgNameAndOV(req.PkgName)
	return &pb.GOPkgReply{
		PkgName:  gopkg.PkgName,
		Url:      gopkg.Link,
		Overview: ov,
	}, nil

}

func Start_gRPC() {
	apiListener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Println(err)
		return
	}

	// 註冊 grpc
	es := &EchoPkg{}

	grpc := grpc.NewServer()
	pb.RegisterEchoPkgServer(grpc, es)
	reflection.Register(grpc)
	fmt.Println("gRPC Listening on 9999...")
	grpc.Serve(apiListener)

}
