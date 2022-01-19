package test

import (
	"testing"
	"time"

	"github.com/sandeepkumardev/simplebank/token"
	"github.com/sandeepkumardev/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := token.NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := token.NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	ps_token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, ps_token)

	payload, err := maker.VerifyToken(ps_token)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrExpiredToken.Error())
	require.Nil(t, payload)
}
