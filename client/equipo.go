package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	pb "github.com/dmoyan0/TareasSD/tree/main/proto"
	"google.golang.org/grpc"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	//Conexión con servidor Tierra
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	//Error conexión
	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	c := pb.NewWishListc(conn)
	equipoID := 1

	//Esperar 10 seg
	time.Sleep(10 * time.Second)
	for {
		//Cantidad aleatoria munición
		municionAT := randInt(20, 30)
		municionMP := randInt(10, 15)

		//Solicitud
		req := &pb.Request{ID: (equipoID), AT: (municionAT), MP: (municionMP)}
		fmt.Printf("Equipo %d: Solicitando %d AT y %d MP\n", equipoID, municionAT, municionMP)
		res, err := c.SolicitarM(context.Background(), req)
		if err != nil {
			panic("Equipo %d: Error al realizar la solicitud: %v\n", equipoID, err)
		} else {
			if res.Success {
				fmt.Printf("Equipo %d: Aprobada la solicitud! Conquista Exitosa!\n", equipoID)
				break
			} else {
				fmt.Printf(" Equipo %d: Denegada la solicitud, reintentando en 3 segundos...\n", equipoID)
				time.Sleep(3 * time.Second)
			}
		}

	}

}
