package inRamDb

import (
	"github.com/hashicorp/go-memdb"
	"sync"
	"web3game/models/entity"
)

var dbLockInRamDb = &sync.Mutex{}

var InRamDb *memdb.MemDB
var Txnn *memdb.Txn

const RockPaperActionTable = "rockPaperAction"

func GetInRamDb() *memdb.MemDB {

	var err error
	if InRamDb == nil {

		dbLockInRamDb.Lock()
		defer dbLockInRamDb.Unlock()

		if InRamDb == nil {

			schema := &memdb.DBSchema{
				Tables: map[string]*memdb.TableSchema{
					RockPaperActionTable: &memdb.TableSchema{
						Name: RockPaperActionTable,
						Indexes: map[string]*memdb.IndexSchema{
							"id": {
								Name:    "id",
								Unique:  true,
								Indexer: &memdb.StringFieldIndex{Field: "Id"},
							},
							"gameId": {
								Name:    "gameId",
								Indexer: &memdb.StringFieldIndex{Field: "GameId"},
							},
							"playerAddress": {
								Name:    "playerAddress",
								Indexer: &memdb.StringFieldIndex{Field: "PlayerAddress"},
							},
							"action": {
								Name:    "action",
								Indexer: &memdb.StringFieldIndex{Field: "Action"},
							},
						},
					},
				},
			}
			InRamDb, err = memdb.NewMemDB(schema)
			if err != nil {
				panic(1)
			} else {
				Txnn = InRamDb.Txn(true)
				defer Txnn.Abort()
			}
		}
	}

	// Create a new database
	return InRamDb
}

func ReadAllRockPaperRounds() ([]*entity.RockPaperAction, error) {
	Txnn = GetInRamDb().Txn(false)
	defer Txnn.Abort()
	it, err := Txnn.Get(RockPaperActionTable, "gameId", "1")
	count := 0
	var allRounds []*entity.RockPaperAction
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*entity.RockPaperAction)
		allRounds = append(allRounds, p)
		count++
	}
	return allRounds, err
}

func InsertRockPaperAction(action *entity.RockPaperAction) error {
	Txnn = GetInRamDb().Txn(true)
	defer Txnn.Abort()
	err := Txnn.Insert(RockPaperActionTable, action)
	Txnn.Commit()
	return err
}

func DeleteRockPaper(action *entity.RockPaperAction) error {
	Txnn = GetInRamDb().Txn(true)
	defer Txnn.Abort()
	err := Txnn.Delete(RockPaperActionTable, action)
	Txnn.Commit()
	return err
}
