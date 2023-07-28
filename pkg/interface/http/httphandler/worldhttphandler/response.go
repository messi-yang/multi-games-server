package worldhttphandler

import "github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"

type getWorldResponse viewmodel.WorldViewModel

type queryWorldsResponse []viewmodel.WorldViewModel

type getMyWorldsResponse []viewmodel.WorldViewModel

type createWorldResponse viewmodel.WorldViewModel

type updateWorldResponse viewmodel.WorldViewModel
