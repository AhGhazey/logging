package format

import (
	"github.com/ahghazey/logging/pkg/constant"
	"net/http"
	"strings"
)

func TokenRefactor(request *http.Request) (map[string]string, bool) {
	auth := request.Header.Get(constant.AuthorizationHeader)
	parts := strings.Split(auth, " ")

	if auth == "" || len(parts) != 2 {
		return nil, false
	}
	if !strings.EqualFold(parts[0], constant.BearerPrefix) {
		return nil, false
	}

	// TODO: parse token second part

	return nil, false
}
