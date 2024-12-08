package domain_test

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"testing"

	"app/internal/data"
	"app/internal/domain"
	mockDomain "app/internal/domain/mock"
)

func TestArticleCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// arrange
	mockStorage := mockDomain.NewMockArticleRepository(ctrl)
	service := domain.NewArticle(mockStorage)

	data := data.NewArticle(uuid.New())
	beforeTest := time.Now()

	mockStorage.EXPECT().Create(gomock.Any(), gomock.Eq(data)).
		Return(nil)

	// act
	assert.NoError(t, service.Create(context.Background(), data))

	// assert
	assert.True(t, data.CreatedAt.After(beforeTest))
}

func TestArticleUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// arrange
	mockStorage := mockDomain.NewMockArticleRepository(ctrl)
	service := domain.NewArticle(mockStorage)

	data := data.NewArticle(uuid.New())
	beforeTest := time.Now()

	mockStorage.EXPECT().Update(gomock.Any(), gomock.Eq(data)).
		Return(nil)

	// act
	assert.NoError(t, service.Update(context.Background(), data))

	// assert
	assert.True(t, data.UpdatedAt.After(beforeTest))
}
