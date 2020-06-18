package graph

import (
	logf "github.com/theSuess/keypub/pkg/log"
	"github.com/theSuess/keypub/pkg/service"
)

var log = logf.Log.WithName("resolver")

type Resolver struct {
	UserService  *service.UserService
	GroupService *service.GroupService
	KeyService   *service.KeyService
}
