package jwtlocal

import (
	"github.com/go-ozzo/ozzo-routing"
	"strings"
	"errors"
	"github.com/go-redis/redis"
	"fmt"
	"github.com/Zhanat87/go/db"
)

func GetJWTToken(c *routing.Context) string {
	header := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(header, "Bearer ") {
		return header[7:]
	} else {
		panic(errors.New("no jwt token"))
	}
}

func CheckJWTTokenIsValid(c *routing.Context) (bool, error) {
	client := db.NewRedis()
	token := GetJWTToken(c)
	_, err := client.Get(token).Result()
	if err == redis.Nil {
		return true, nil
	} else if err != nil {
		return false, err
	} else {
		return false, errors.New(fmt.Sprintf("token <<%s>> was invalidated early", token))
	}
}
