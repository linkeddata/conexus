package conexus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBDropTables(t *testing.T) {
	err := db.DropTables(new(User), new(Subscription))
	assert.NoError(t, err)
}

func TestDBSync(t *testing.T) {
	err := db.Sync(
		new(User),
		new(Subscription),
	)
	assert.NoError(t, err)

	results, err := db.Query("SHOW TABLES")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
}

var (
	testUser = &User{
		Uri: "https://webid.mit.edu/presbrey#",
	}
	testSynd = &Subscription{
		Source:      "https://test.data.fm/foo",
		Destination: "https://test.data.fm/bar",
	}
)

func TestDBSave(t *testing.T) {
	n, err := db.InsertOne(testUser)
	assert.Equal(t, 1, n)
	assert.NoError(t, err)

	testSynd.User = testUser.Id

	n, err = db.Insert(testSynd)
	assert.Equal(t, 1, n)
	assert.NoError(t, err)
}

func TestDBSearch(t *testing.T) {
	res1 := make([]User, 0)
	err := db.Cols("id").Where("uri LIKE ?", `%`+testUser.Uri+`%`).Find(&res1)
	assert.NoError(t, err)
	assert.Equal(t, testUser.Id, res1[0].Id)

	res2 := make([]Subscription, 0)
	err = db.Cols("user").Where("source LIKE ?", `%foo%`).Find(&res2)
	assert.NoError(t, err)
	err = db.Cols("user").Where("destination LIKE ?", `%bar%`).Find(&res2)
	assert.NoError(t, err)
	assert.Equal(t, testUser.Id, res2[0].User)
}

func TestDBDelete(t *testing.T) {
	var (
		n   int64
		err error
	)

	n, err = db.Delete(testUser)
	assert.Equal(t, 1, n)
	assert.NoError(t, err)

	n, err = db.Delete(testSynd)
	assert.Equal(t, 1, n)
	assert.NoError(t, err)
}
