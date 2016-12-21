package serviceImplement

import (
	"log"
)

type KeepAliveImpl struct {
}

func (srv *KeepAliveImpl) DefaultKeepAlive(clientId string) (r bool, err error) {

	log.Printf("Client: %v keep alived.", clientId)
	r = true
	return
}
