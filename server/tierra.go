package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/dmoyan0/TareasSD/tree/main/proto"
	"google.golang.org/grpc"
)

// Definición servidor Tierra
type server struct {
	pb.UnimplementedWishListServiceServer     //<-Interfaz servicio gRPC
	municionAT                            int //Contador Munición Anti-Terrestre
	municionMP                            int //Contador Munición Pesada
}

func (s *server) solicitarM(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	fmt.Printf("Recepción de solicitud desde equipo %d, %d AT y %d MP\n", request.ID, request.AT, request.MP)

	//Verifiamos munición
	if s.municionAT >= int(request.AT) {
		if s.municionMP >= int(request.MP) {
			//Actializamos contadores
			s.municionAT -= int(request.AT)
			s.municionMP -= int(request.MP)
			//Exito
			fmt.Printf("AT EN SISTEMA: %d; MP EN SISTEMA: %d\n", s.municionAT, s.municionMP)
			return &pb.response{Success: true}, nil
		}
		//No hay munición suficiente
		fmt.Printf("AT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", s.municionAT, s.municionMP)
		return &pb.response{Success: false}, nil
	}
	//No hay munición suficiente
	fmt.Printf("AT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", s.municionAT, s.municionMP)
	return &pb.response{Success: false}, nil
}

func main() {
	//Creamos conexion TCP escuchando en el puerto 50051
	listner, err := net.Listen("tcp", ":50051")
	//Manejo error
	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	server := &server{municionAT: 0, municionMP: 0} //Municiones parten en 0
	pb.RegisterWishListServiceServer(serv, &server)
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
}
