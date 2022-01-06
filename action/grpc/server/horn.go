package server

import (
	"context"
	"hcc/horn/lib/mysqlUtil"
	"innogrid.com/hcloud-classic/pb"
)

type hornServer struct {
	pb.UnimplementedHornServer
}

func (s *hornServer) GetMYSQLDEncryptedPassword(_ context.Context, _ *pb.Empty) (*pb.ResGetMYSQLDEncryptedPassword, error) {
	encryptedPassword, err := mysqlUtil.GetEncryptPassword()
	if err != nil {
		return nil, err
	}

	return &pb.ResGetMYSQLDEncryptedPassword{EncryptedPassword: encryptedPassword}, nil
}
