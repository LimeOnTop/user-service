package user

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	user usecase.UserUsecase
}