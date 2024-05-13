package providedependency

import (
	"os"

	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/authappsrv"
)

func ProvideAuthAppService() authappsrv.Service {
	return authappsrv.NewService(os.Getenv("AUTH_SECRET"))
}
