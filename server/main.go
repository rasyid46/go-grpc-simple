package main

//buat struck yang implement data student ini
import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	pb "sule.id/learn/grpc-simple/student" //2
)

type dataStudentServer struct { //1
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*pb.Student //dao
}

//membuat receiver 4
func (d *dataStudentServer) FindStundentByemail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	//ctx dan error harus ada ketika req data request student

	fmt.Println(d.students, "student  ===")

	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataStudentServer) loadData() {
	data, err := ioutil.ReadFile("data/datas.json")
	if err != nil {
		log.Fatalln("error in read fie", err.Error())
	}
	//
	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("error in un marshall data json", err.Error())
	}

}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	//grpc serve
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("error in lister ", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error when serve grpc", err.Error())
	}
}
