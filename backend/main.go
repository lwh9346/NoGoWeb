package main

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	threads int
	lock    sync.Mutex
)

func main() {
	r := gin.Default()
	r.POST("/api/nogo", handleNoGoRequest)
	r.Static("/", "./ui")
	r.Run("127.0.0.1:8888")
}

type nogoResponse struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Allow  []bool `json:"allow"`
	Winner string `json:"winner"`
}

type nogoRequest struct {
	Board      []int  `json:"board" binding:"required"`
	Difficulty string `json:"difficulty" binding:"required"`
}

func handleNoGoRequest(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var req nogoRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"msg": "数据绑定失败"})
		return
	}
	if len(req.Board) < 81 {
		c.JSON(400, gin.H{"msg": "棋子数过少"})
		return
	}
	for i := 0; i < 81; i++ {
		if !(req.Board[i] == 1 || req.Board[i] == 0 || req.Board[i] == -1) {
			c.JSON(400, gin.H{"msg": "棋盘数据不合规"})
			return
		}
	}
	var maxStep int
	switch req.Difficulty {
	case "easy":
		maxStep = 20000
	case "normal":
		maxStep = 100000
	case "hard":
		maxStep = 200000
	default:
		c.JSON(400, gin.H{"msg": "不存在的难度"})
	}
	if err := waitForFree(); err != nil {
		c.JSON(503, gin.H{"msg": "服务器过载"})
		return
	}
	//读取棋盘，这里需要翻转一下
	board := make([]int, 81)
	for i := 0; i < 81; i++ {
		board[i] = -req.Board[i]
	}
	var resp nogoResponse
	res := GoGetValidPosition(board)
	if res.numS == 0 {
		resp.Winner = "player"
		c.JSON(200, resp)
		freeThread()
		return
	}
	x, y := GoGetBestAction(board, maxStep)
	board[x*9+y] = 1
	res = GoGetValidPosition(board)
	resp.X = x
	resp.Y = y
	resp.Allow = res.resR //因为电脑是玩家的反方
	if res.numR == 0 {
		resp.Winner = "computer"
	} else {
		resp.Winner = "none"
	}
	c.JSON(200, resp)
	freeThread()
}

func waitForFree() error {
	count := 0
	for true {
		lock.Lock()
		if threads > 2 {
			lock.Unlock()
			count++
			time.Sleep(100 * time.Millisecond)
		} else {
			threads++
			lock.Unlock()
			break
		}
		if count > 30 {
			return errors.New("TimeOut")
		}
	}
	return nil
}

func freeThread() {
	lock.Lock()
	threads--
	lock.Unlock()
}
