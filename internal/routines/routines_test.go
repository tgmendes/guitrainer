package routines_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tgmendes/guitrainer/internal/mocks"
	"github.com/tgmendes/guitrainer/internal/routines"
	"testing"
)

func TestAdd(t *testing.T) {
	mockRow := new(mocks.Row)
	mockRow.On("Scan", mock.Anything).Return(nil)
	mockRepo := new(mocks.Repo)
	mockRepo.On("QueryRow", context.Background(), mock.Anything, mock.Anything).Return(mockRow)

	nr := routines.Routine{
		Name:        "someName",
		Description: "someDescription",
		Level:       "someLevel",
	}

	r, err := routines.Add(context.Background(), mockRepo, nr)

	mockRow.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.IsType(t, &routines.Routine{}, r)
	assert.Equal(t, nr.Name, r.Name)
	assert.Equal(t, nr.Description, r.Description)
	assert.Equal(t, nr.Level, r.Level)

}

func TestList(t *testing.T) {
	mockRows := new(mocks.Rows)
	mockRows.On("Next").Return(true).Once()
	mockRows.On("Next").Return(false)
	mockRows.On("Scan", mock.Anything).Return(nil)
	mockRepo := new(mocks.Repo)
	mockRepo.On("Query", context.Background(), mock.Anything, mock.Anything).Return(mockRows, nil)

	r, err := routines.List(context.Background(), mockRepo)

	mockRows.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.IsType(t, []routines.Routine{}, r)
}

func TestGet(t *testing.T) {
	mockRow := new(mocks.Row)
	mockRow.On("Scan", mock.Anything).Return(nil)
	mockRepo := new(mocks.Repo)
	mockRepo.On("QueryRow", context.Background(), mock.Anything, mock.Anything).Return(mockRow)

	r, err := routines.Get(context.Background(), mockRepo, int64(123))

	mockRow.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.IsType(t, &routines.Routine{}, r)
}

func TestDelete(t *testing.T) {
	mockRepo := new(mocks.Repo)
	mockRepo.On("Exec", context.Background(), mock.Anything, mock.Anything).Return([]byte("DELETE"), nil)

	err := routines.Delete(context.Background(), mockRepo, int64(123))

	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
}
