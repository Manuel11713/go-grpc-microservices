package main

import (
	"log"

	pb "github.com/Manuel11713/sum/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.PrimesService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v \n", in)

	var k int32 = 2
	n := in.Number

	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimesResponse{
				Prime: k,
			})

			n /= k
		} else {
			k += 1
		}

	}

	return nil
}
