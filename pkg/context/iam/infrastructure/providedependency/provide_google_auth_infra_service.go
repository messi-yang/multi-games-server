package providedependency

import "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/service/googleauthinfrasrv"

func ProvideGoogleAuthInfraService() googleauthinfrasrv.Service {
	return googleauthinfrasrv.NewService()
}
