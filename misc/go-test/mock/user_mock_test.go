package mock

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestMockUserRepository_EXPECT(t *testing.T) {
	c := gomock.NewController(t)
	h := NewMockUserRepository(c)

	h.EXPECT().FindOne(1).Return(User{Name: "1"}, nil)
	h.EXPECT().FindOne(gomock.Any()).DoAndReturn(func(id int) (User, error) {
		if id < 100 {
			return User{Name: "<100"}, nil
		}

		if id < 200 {
			return User{Name: "<200"}, nil
		}

		return User{}, errors.New("invalid id")
	}).AnyTimes()

	// gomock.InOrder()

	t.Log(h.FindOne(1))
	t.Log(h.FindOne(2))

	t.Log(h.FindOne(3))
	t.Log(h.FindOne(999))

	t.Error()

}
