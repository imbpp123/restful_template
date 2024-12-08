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

	data := &data.CreateArticle{
		UUID:   uuid.New(),
		Title:  "title",
		Author: "author",
		Text:   "text",
	}
	beforeTest := time.Now()

	mockStorage.EXPECT().Create(gomock.Any(), gomock.Eq(data)).
		Return(nil)

	// act
	result, err := service.Create(context.Background(), data)
	assert.NoError(t, err)

	// assert
	assert.True(t, result.CreatedAt.After(beforeTest))
}
