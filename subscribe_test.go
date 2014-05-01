package conexus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserExists(t *testing.T) {
	n, err := db.InsertOne(testUser)
	assert.Equal(t, 1, n)
	assert.NoError(t, err)
	assert.True(t, userExists(testUser.Uri))

	n, err = db.Delete(testUser)
	assert.Equal(t, 1, n)
	assert.NoError(t, err)
}

func TestSubscribe(t *testing.T) {
	status := testUser.Subscribe("http://test.rww.io/feeds")
	assert.True(t, status)
}
