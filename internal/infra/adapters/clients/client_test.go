package clients

import (
	"fmt"
	"github.com/SafetyLink/commons/config"
	log "github.com/SafetyLink/commons/logger"
	"github.com/SafetyLink/webService/internal"
	"testing"
)

func TestConnectToAuthentication(t *testing.T) {
	logger := log.InitLogger()
	cfg, err := config.ReadConfigInTest[internal.Config]()
	if err != nil {
		t.Error(err)
	}

	authClient := GrpcAuthenticationClient(logger, cfg)
	fmt.Println(authClient)

}
