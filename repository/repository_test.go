package repository

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

var (
	client *redis.Client
)

var (
	key = "name"
	val = "aayush"
)

func TestMain(m *testing.M) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	client = redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	os.Exit(m.Run())
}

func TestSet(t *testing.T) {
	ctx := context.TODO()
	exp := time.Duration(0)

	db, mock := redismock.NewClientMock()
	mock.ExpectSet(key, val, exp).SetVal(val)

	r := NewRedisRepository(db)
	err := r.Set(ctx, key, val, exp)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGet(t *testing.T) {
	ctx := context.TODO()
	db, mock := redismock.NewClientMock()
	mock.ExpectGet(key).SetVal(val)

	r := NewRedisRepository(db)
	_, err := r.Get(ctx, key)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
