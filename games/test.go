package games

//tableName := "rockPaperAction"
//// Create the DB schema
//schema := &memdb.DBSchema{
//	Tables: map[string]*memdb.TableSchema{
//		tableName: &memdb.TableSchema{
//			Name: tableName,
//			Indexes: map[string]*memdb.IndexSchema{
//				"id": {
//					Name:    "id",
//					Unique:  true,
//					Indexer: &memdb.StringFieldIndex{Field: "Id"},
//				},
//				"gameId": {
//					Name:    "gameId",
//					Indexer: &memdb.StringFieldIndex{Field: "GameId"},
//				},
//				"playerAddress": {
//					Name:    "playerAddress",
//					Indexer: &memdb.StringFieldIndex{Field: "PlayerAddress"},
//				},
//				"action": {
//					Name:    "action",
//					Indexer: &memdb.StringFieldIndex{Field: "Action"},
//				},
//			},
//		},
//	},
//}
//
//// Create a new database
//db, err := memdb.NewMemDB(schema)
//if err != nil {
//	panic(err)
//}
//
//// Create a White transaction
//txn := db.Txn(true)
//
//// Insert some people
//testData := []*entity.RockPaperAction{
//	{time.Now().String(), "1", "p1", "rock"},
//	{time.Now().String(), "1", "p2", "paper"},
//
//	{time.Now().String(), "1", "p2", "paper"},
//	{time.Now().String(), "1", "p1", "rock"},
//
//	{time.Now().String(), "1", "p2", "paper"},
//	{time.Now().String(), "1", "p1", "paper"},
//
//	{time.Now().String(), "1", "p2", "paper"},
//	{time.Now().String(), "1", "p1", "rock"},
//}
//for _, p := range testData {
//	if err := txn.Insert(tableName, p); err != nil {
//		panic(err)
//	}
//}
//
//// Commit the transaction
//txn.Commit()
//
//// Create read-only transaction
//txn = db.Txn(false)
//defer txn.Abort()
//
//// List all the people
//it, err := txn.Get(tableName, "gameId", "1")
//if err != nil {
//	panic(err)
//}
//
//count := 0
//var allRounds []*entity.RockPaperAction
//fmt.Println("all record:")
//for obj := it.Next(); obj != nil; obj = it.Next() {
//	p := obj.(*entity.RockPaperAction)
//	allRounds = append(allRounds, p)
//	fmt.Printf("  Action: %s PlayerAddress: %s \n", p.Action, p.PlayerAddress)
//	count++
//}
//println()
//println()
//println()
////this round is last round of this game, and game should end and send to all result
//player1Address := allRounds[0].PlayerAddress
//player1Wins := 0
//player2Address := allRounds[1].PlayerAddress
//player2Wins := 0
//count = 0
//for range allRounds {
//	if count+2 <= len(allRounds) {
//		winner := rockPaperScissors.CalculaterRockPaperRoundWinner(
//			allRounds[count], allRounds[count+1],
//		)
//		if winner == player1Address {
//			player1Wins++
//		} else {
//			player2Wins++
//		}
//	} else {
//		if player1Wins > player2Wins {
//			println("victory for: " + player1Address)
//		} else {
//			println("victory for: " + player2Address)
//		}
//	}
//	count++
//}
