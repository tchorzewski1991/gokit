package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tchorzewski1991/gokit/slice"
)

func TestFind(t *testing.T) {
	type user struct {
		email string
	}
	users := []user{
		{
			email: "user1@email.com",
		},
		{
			email: "user2@email.com",
		},
	}

	byEmail := func(u user) bool {
		return u.email == "user2@email.com"
	}
	res, ok := slice.Find(users, byEmail)
	assert.True(t, ok)
	assert.Equal(t, res, users[1])

	byUser := func(u user) bool {
		return u == users[0]
	}
	res, ok = slice.Find(users, byUser)
	assert.True(t, ok)
	assert.Equal(t, res, users[0])
}
