package worldhttpcontroller

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/jsondto"

type getWorldResponseDto jsondto.WorldAggDto

type getWorldsResponseDto []jsondto.WorldAggDto

type createWorldResponseDto jsondto.WorldAggDto
