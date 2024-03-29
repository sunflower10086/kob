package matchutil

import (
	"context"
	"log"
	"matching/internal/grpc/client/result"
	pb "matching/internal/grpc/client/result/pb"
	"matching/pkg/mw"
	"time"
)

type Player struct {
	UserId   int32
	BotId    int32
	Rating   int32
	WaitTime int32
}

var (
	Players = make([]*Player, 0)
)

func MatchingPool() {
	for {
		time.Sleep(1 * time.Second)
		increaseWaitingTime()
		matchPlayers()
	}
}

// 给所有玩家的匹配时间增加一秒
func increaseWaitingTime() {
	for _, player := range Players {
		player.WaitTime += 1
	}
}

// 尝试匹配所有玩家
func matchPlayers() {
	used := make([]bool, len(Players))

	for i := 0; i < len(Players); i++ {
		if Players[i].UserId == 0 || used[i] {
			continue
		}
		for j := i + 1; j < len(Players); j++ {
			if Players[j].UserId == 0 || used[j] {
				continue
			}

			if Players[i].UserId == Players[j].UserId {
				continue
			}
			a, b := Players[i], Players[j]

			if checkMatched(a, b) {
				used[i] = true
				used[j] = true
				mw.SugarLogger.Infof("matching success playerA:%v, playerB:%v", a, b)
				go sendResult(a, b)
				break
			}
		}
	}

	newPlayer := make([]*Player, 0)
	for i := 0; i < len(Players); i++ {
		if !used[i] {
			newPlayer = append(newPlayer, Players[i])
		}
	}
	Players = newPlayer
}

func checkMatched(a, b *Player) bool {
	ratingDelta := Abs(a.Rating - b.Rating)
	waitTime := Min(a.WaitTime, b.WaitTime)
	return ratingDelta <= waitTime*10
}

func sendResult(a, b *Player) {
	// TODO : 把结果返回给backend层
	// EventType 0表示匹配的结果，1表示游戏的结果
	req := pb.ResultReq{
		EventType: 0,
		MatchResult: &pb.MatchResult{
			AId:    a.UserId,
			ABotId: a.BotId,
			BId:    b.UserId,
			BBotId: b.BotId,
		},
	}
	ctx := context.Background()
	result, err := result.Result(ctx, &req)
	if err != nil {
		log.Println(err)
	}
	log.Println("sendResult", result)
}
