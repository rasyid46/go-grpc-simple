package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "sule.id/learn/grpc-simple/student" //2
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	//memanggil layaknya sebagai method

	s := pb.Student{Email: email}

	student, err := client.FindStundentByemail(ctx, &s)

	if err != nil {
		log.Fatalln("error when get student by email", err.Error())
	}

	fmt.Println(student)

}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)

	if err != nil {
		log.Fatalln("error in dial")

	}

	defer conn.Close()

	client := pb.NewDataStudentClient(conn)
	//bikin method manggil endpoin/grpc

	getDataStudentByEmail(client, "info@shekkrean.id")
	getDataStudentByEmail(client, "wandinak17@gmail.com")
}
