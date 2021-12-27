package repository_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (mock *MockRepo) TestCreateNewUser(t *testing.T) {

}