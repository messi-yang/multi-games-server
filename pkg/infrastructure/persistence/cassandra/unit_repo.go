package cassandra

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/gocql/gocql"
	"github.com/samber/lo"
)

type unitRepo struct {
	session *gocql.Session
}

var unitRepoSingleton *unitRepo

func NewUnitRepo() (unitmodel.Repo, error) {
	if unitRepoSingleton == nil {
		newSession, err := newSession()
		if err != nil {
			return nil, err
		}
		unitRepoSingleton = &unitRepo{
			session: newSession,
		}
		return unitRepoSingleton, nil
	}
	return unitRepoSingleton, nil
}

func (repo *unitRepo) Add(unit unitmodel.UnitAgg) error {
	if err := repo.session.Query(
		"INSERT INTO units (game_id, pos_x, pos_z, item_id) VALUES (?, ?, ?, ?)",
		unit.GetWorldId().String(),
		unit.GetPosition().GetX(),
		unit.GetPosition().GetZ(),
		unit.GetItemId().ToInt16(),
	).Exec(); err != nil {
		return err
	}
	return nil
}

func (repo *unitRepo) GetUnitAt(worldId worldmodel.WorldIdVo, position commonmodel.PositionVo) (unitmodel.UnitAgg, bool, error) {
	iter := repo.session.Query(
		"SELECT * FROM units WHERE game_id = ? AND pos_x = ? AND pos_z = ? LIMIT 1",
		worldId.String(),
		position.GetX(),
		position.GetZ(),
	).Iter()
	var unit *unitmodel.UnitAgg = nil
	var rawWorldId gocql.UUID
	var rawPosX int
	var rawPosZ int
	var rawItemId int16
	for iter.Scan(&rawWorldId, &rawPosX, &rawPosZ, &rawItemId) {
		parsedWorldId, _ := worldmodel.ParseWorldIdVo(rawWorldId.String())
		position := commonmodel.NewPositionVo(rawPosX, rawPosZ)
		itemId := itemmodel.NewItemIdVo(rawItemId)
		unitFound := unitmodel.NewUnitAgg(parsedWorldId, position, itemId)
		unit = &unitFound
	}
	if err := iter.Close(); err != nil {
		return unitmodel.UnitAgg{}, false, err
	}
	if unit == nil {
		return unitmodel.UnitAgg{}, false, nil
	}
	return *unit, true, nil
}

func (repo *unitRepo) GetUnitsInBound(worldId worldmodel.WorldIdVo, bound commonmodel.BoundVo) ([]unitmodel.UnitAgg, error) {
	fromX := bound.GetFrom().GetX()
	toX := bound.GetTo().GetX()
	fromZ := bound.GetFrom().GetZ()
	toZ := bound.GetTo().GetZ()
	xPositions := lo.RangeFrom(fromX, toX-fromX+1)
	iter := repo.session.Query(
		"SELECT game_id, pos_x, pos_z, item_id FROM units WHERE game_id = ? AND pos_x IN ? AND pos_z >= ? AND pos_z <= ?",
		worldId.String(),
		xPositions,
		fromZ,
		toZ,
	).Iter()
	var units []unitmodel.UnitAgg = make([]unitmodel.UnitAgg, 0)
	var rawWorldId gocql.UUID
	var rawPosX int
	var rawPosZ int
	var rawItemId int16
	for iter.Scan(&rawWorldId, &rawPosX, &rawPosZ, &rawItemId) {
		parsedWorldId, _ := worldmodel.ParseWorldIdVo(rawWorldId.String())
		position := commonmodel.NewPositionVo(rawPosX, rawPosZ)
		itemId := itemmodel.NewItemIdVo(rawItemId)
		units = append(units, unitmodel.NewUnitAgg(parsedWorldId, position, itemId))
	}
	if err := iter.Close(); err != nil {
		return units, err
	}
	return units, nil
}

func (repo *unitRepo) Delete(worldId worldmodel.WorldIdVo, position commonmodel.PositionVo) error {
	if err := repo.session.Query(
		"DELETE FROM units WHERE game_id = ? AND pos_x = ? AND pos_z = ?",
		worldId.String(),
		position.GetX(),
		position.GetZ(),
	).Exec(); err != nil {
		return err
	}
	return nil
}
