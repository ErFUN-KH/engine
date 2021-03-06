// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package qollenge

import (
	impl2 "elbix.dev/engine/modules/misc/impl"
	"elbix.dev/engine/modules/misc/proto"
	"elbix.dev/engine/modules/user/impl"
	"elbix.dev/engine/modules/user/proto"
	"elbix.dev/engine/pkg/grpcgw"
	"elbix.dev/engine/pkg/sec"
	"elbix.dev/engine/pkg/token/jwt"
)

import (
	_ "elbix.dev/engine/modules/misc"
	_ "elbix.dev/engine/modules/user"
)

// Injectors from wire.go:

func userMod() (grpcgw.Controller, error) {
	string2 := getPrivateKey()
	rsaPrivateKey, err := sec.ParseRSAPrivateKeyFromBase64PEM(string2)
	if err != nil {
		return nil, err
	}
	provider := jwt.NewJWTTokenProvider(rsaPrivateKey)
	userSystemServer := impl.NewUserController(provider)
	wrappedUserSystemController := userpb.NewWrappedUserSystemServer(userSystemServer)
	return wrappedUserSystemController, nil
}

func miscMod() (grpcgw.Controller, error) {
	string2 := getPrivateKey()
	rsaPrivateKey, err := sec.ParseRSAPrivateKeyFromBase64PEM(string2)
	if err != nil {
		return nil, err
	}
	publicKey := sec.ExtractPublicFromPrivate(rsaPrivateKey)
	miscSystemServer := impl2.NewMiscController(publicKey)
	wrappedMiscSystemController := miscpb.NewWrappedMiscSystemServer(miscSystemServer)
	return wrappedMiscSystemController, nil
}
