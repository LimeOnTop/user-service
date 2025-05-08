package grpcuser

import (
	"context"

	pb "user-service/gen/user"
	usecase "user-service/internal/usecase/user"
)

var _ pb.UserServiceServer = (*UserServer)(nil)

// UserServer - структура для обработки RPC-методов, реализующая интерфейс pb.UserServiceServer
type UserServer struct {
	pb.UnimplementedUserServiceServer
	user usecase.UserUseCase
}

// New - конструктор для UserServer
func New(user usecase.UserUseCase) *UserServer {
	return &UserServer{user: user}
}

// GetUserProducts - метод для получения продуктов пользователя
func (s *UserServer) GetUserProducts(ctx context.Context, req *pb.UserRequest) (*pb.GetProductsResponse, error) {
	products, err := s.user.GetUserProducts(req.AccessToken)
	if err != nil {
		return nil, err
	}

	response := &pb.GetProductsResponse{
		ProductNames: products,
	}

	return response, nil
}

// GetUserPreference - метод для получения предпочтений пользователя
func (s *UserServer) GetUserPreference(ctx context.Context, req *pb.UserRequest) (*pb.GetPreferenceResponse, error) {
	preference, err := s.user.GetUserPreference(req.AccessToken)
	if err != nil {
		return nil, err
	}

	response := &pb.GetPreferenceResponse{
		PreferenceName: preference,
	}

	return response, nil
}

// UpdateUserPreference - метод для обновления предпочтений пользователя
func (s *UserServer) UpdateUserPreference(ctx context.Context, req *pb.UpdatePreferenceRequest) (*pb.UpdatePreferenceResponse, error) {
	err := s.user.UpdateUserPreference(req.AccessToken, req.PreferenceName)
	if err != nil {
		return nil, err
	}

	response := &pb.UpdatePreferenceResponse{
		Success: true,
	}

	return response, nil
}

// RemoveUserPreference - метод для удаления предпочтений пользователя
func (s *UserServer) RemoveUserPreference(ctx context.Context, req *pb.RemovePreferenceRequest) (*pb.RemovePreferenceResponse, error) {
	err := s.user.RemoveUserPreference(req.AccessToken)
	if err != nil {
		return nil, err
	}
	response := &pb.RemovePreferenceResponse{
		Success: true,
	}
	return response, nil
}

// AddUserProduct - метод для добавления продукта пользователю
func (s *UserServer) AddUserProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	err := s.user.AddUserProduct(req.AccessToken, req.ProductName)
	if err != nil {
		return nil, err
	}

	response := &pb.AddProductResponse{
		Success: true,
	}

	return response, nil
}

// RemoveUserProduct - метод для удаления продукта у пользователя
func (s *UserServer) RemoveUserProduct(ctx context.Context, req *pb.RemoveProductRequest) (*pb.RemoveProductResponse, error) {
	err := s.user.RemoveUserProduct(req.AccessToken, req.ProductName)
	if err != nil {
		return nil, err
	}

	response := &pb.RemoveProductResponse{
		Success: true,
	}

	return response, nil
}
