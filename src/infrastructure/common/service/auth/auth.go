package auth_service

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	auth_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/service/auth"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/valueobject/access"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/valueobject/token"
	"strconv"
	"time"
)

type ClientData struct {
	client *redis.Client
}

func NewAuth(client *redis.Client) auth_service.AuthInterface {
	return &ClientData{client: client}
}

//Save token metadata to Redis
func (tk *ClientData) CreateAuth(userid string, td *token_detail.TokenDetail) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(td.TokenUuid, userid, at.Sub(now)).Result()

	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(td.RefreshUuid, userid, rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

//Check the metadata saved
func (tk *ClientData) FetchAuth(tokenUuid string) (uint64, error) {
	userid, err := tk.client.Get(tokenUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

//Once a user row in the token table
func (tk *ClientData) DeleteTokens(authD *access_detail.AccessDetail) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%d", authD.TokenUuid, authD.UserId)
	//delete access token
	deletedAt, err := tk.client.Del(authD.TokenUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := tk.client.Del(refreshUuid).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

func (tk *ClientData) DeleteRefresh(refreshUuid string) error {
	//delete refresh token
	deleted, err := tk.client.Del(refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}
