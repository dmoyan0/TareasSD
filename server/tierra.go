package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
)

// Definición servidor Tierra
type server struct {
	pb.UnimplementedWishListServiceServer     //<-Interfaz servicio gRPC
	municionAT                            int //Contador Munición Anti-Terrestre
	municionMP                            int //Contador Munición Pesada
	maxAT                                 int
	maxMP                                 int
}

func (s *server) SolicitarM(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	fmt.Printf("Recepción de solicitud desde equipo %d, %d AT y %d MP\n", request.ID, request.AT, request.MP)

	// Verificamos munición
	if s.municionAT >= int(request.AT) && s.municionMP >= int(request.MP) {
		// Actualizamos contadores
		s.municionAT -= int(request.AT)
		s.municionMP -= int(request.MP)
		// Éxito
		fmt.Printf("AT EN SISTEMA: %d; MP EN SISTEMA: %d\n", s.municionAT, s.municionMP)
		return &pb.Response{Success: true}, nil
	}

	// No hay suficiente munición
	fmt.Printf("AT EN SISTEMA: %d; MP EN SISTEMA: %d\n", s.municionAT, s.municionMP)
	return &pb.Response{Success: false}, nil
}

func (s *server) generarMunicion() {
	for {
		// Generamos munición cada 5 seg
		time.Sleep(5 * time.Second)

		// Vemos si hay espacio para la municion AT
		if s.municionAT < s.maxAT {
			if s.municionAT+10 <= s.maxAT {
				s.municionAT += 10
			} else {
				s.municionAT = s.maxAT
			}
		}

		// Vemos si hay espacio para la municion MP
		if s.municionMP < s.maxMP {
			if s.municionMP+5 <= s.maxMP {
				s.municionMP += 5
			} else {
				s.municionMP = s.maxMP
			}
		}
	}
}

func main() {
	//Creamos conexion TCP escuchando en el puerto 50051
	listner, err := net.Listen("tcp", ":50051")
	//Manejo error
	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	server := &server{municionAT: 0, municionMP: 0, maxAT: 50, maxMP: 20} //Municiones parten en 0
	pb.RegisterWishListServiceServer(serv, &server)

	//Generamos munición
	go server.generarMunicion()

	//Error
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
}
