package main

import (
	"github.com/gin-gonic/gin"
	goproto "github.com/gogo/protobuf/proto"
	"github.com/sansna/sanser.go/proto"
)

func main() {
	router := gin.Default()
	// return data as json
	router.GET("/config/get", getting)
	// return data using param, as json
	router.POST("/config/post", posting)
	// return data as pb
	router.POST("/config/proto", protoing)
	router.Run(":8888")
}

type A struct {
	AA int64                  `json:"aa,omitempty"`
	CC string                 `json:"cc,omitempty"`
	DD map[string]interface{} `json:"dd"`
}

func getting(c *gin.Context) {
	a := &A{
		AA: 10,
		CC: "axzcvkl",
		DD: map[string]interface{}{
			"xcv":  100,
			"xxcv": "zxcv",
		},
	}
	c.JSON(200, a)
	//c.ProtoBuf(200, a)
}

type Param struct {
	Id int64
}

func posting(c *gin.Context) {
	p := &Param{}
	c.BindJSON(p)
	id := p.Id
	c.JSON(200, &A{
		AA: id,
		DD: make(map[string]interface{}),
	})
}

func protoing(c *gin.Context) {
	p := &Param{}
	c.BindJSON(p)
	a := &proto.C{
		A: goproto.Int64(1000),
		B: goproto.Int64(p.Id),
	}
	// How to decode from curl?
	// protoc -Ipath/to/proto --decode=proto.C a.proto < <(curl -sd '{}' localhost:8888/config/proto)
	c.ProtoBuf(200, a)
}
