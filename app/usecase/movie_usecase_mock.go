package usecase

import (
	"github.com/stretchr/testify/mock"
	"xsis-code-test/models/request"
	"xsis-code-test/models/response"
)

type MockAppUsecase struct {
	Mock mock.Mock
}

func (mau *MockAppUsecase) CreateMovie(req request.CreateMovie) error {
	args := mau.Mock.Called(req)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (mau *MockAppUsecase) ListMovie() (*[]response.ListMovie, error) {
	args := mau.Mock.Called()
	if args.Get(1) == nil {
		return args.Get(0).(*[]response.ListMovie), nil
	}
	return args.Get(0).(*[]response.ListMovie), args.Get(1).(error)
}

func (mau *MockAppUsecase) GetMovie(id int64) (*response.GetMovie, error) {
	args := mau.Mock.Called(id)
	if args.Get(1) == nil {
		return args.Get(0).(*response.GetMovie), nil
	}
	return args.Get(0).(*response.GetMovie), args.Get(1).(error)
}

func (mau *MockAppUsecase) UpdateMovie(id int64, req request.UpdateMovie) error {
	args := mau.Mock.Called(id, req)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (mau *MockAppUsecase) DeleteMovie(id int64) error {
	args := mau.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}
