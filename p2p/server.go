package p2p

import (
	"io"
	"log"
	"encoding/json"
	"golang.org/x/net/websocket"

	"github.com/heeeeeng/naivechain/core"
)

const (
	queryBlock = iota
	queryAll
	respQueryBlock
	respQueryChain
)

type P2PServer struct {
	conns		[]*websocket.Conn
	blockchain	*core.BlockChain
}

func NewP2PServer(bc *core.BlockChain) *P2PServer {
	srv := &P2PServer{
		conns:		[]*websocket.Conn{},
		blockchain: bc,
	}
	return srv
}

func (srv *P2PServer) Start() {
	for _, conn := range srv.conns {
		go srv.readloop(conn)
	}
}

func (srv *P2PServer) AddPeers(peers []string) {
	for _, peer := range peers {
		if peer == "" {
			continue
		}
		ws, err := websocket.Dial(peer, "", peer)
		if err != nil {
			log.Println("dial to peer", err)
			continue
		}
		srv.conns = append(srv.conns, ws)
		go srv.readloop(ws)
	}
}

type ResponseBlockchain struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}

type QueryBlockMsg struct {
	Index int 	`json:"index"`
}

func (srv *P2PServer) readloop(ws *websocket.Conn) {
	var (
		v    = &ResponseBlockchain{}
		peer = ws.LocalAddr().String()
	)

	for {
		var msg []byte
		err := websocket.Message.Receive(ws, &msg)
		if err == io.EOF {
			log.Printf("p2p Peer[%s] shutdown, remove it form peers pool.\n", peer)
			break
		}
		if err != nil {
			log.Println("Can't receive p2p msg from ", peer, err.Error())
			break
		}

		log.Printf("Received[from %s]: %s.\n", peer, msg)
		err = json.Unmarshal(msg, v)
		errFatal("invalid p2p msg", err)

		switch v.Type {
		case queryBlock:
			v.Type = respQueryBlock

			var queryData QueryBlockMsg
			err = json.Unmarshal([]byte(v.Data), &queryData)
			if err != nil {
				log.Println("Can't unmarshal QueryBlockMsg from ", peer, err.Error())
				continue
			}
			block := srv.blockchain.GetBlock(queryData.Index)
			if block == nil {
				log.Println("GetBlock with wrong index")
				bs := []byte("wrong index")
				ws.Write(bs)
				continue
			}
			bs := block.Bytes()
			log.Printf("responseLatestMsg: %s\n", bs)
			ws.Write(bs)

		case queryAll:
			// TODO query all
			v.Type = respQueryChain
			v.Data = srv.blockchain.String()
			bs, _ := json.Marshal(v)
			log.Printf("responseChainMsg: %s\n", bs)
			ws.Write(bs)

		case respQueryBlock:
			// handleBlockchainResponse([]byte(v.Data))

		case respQueryChain:

		}

	}
}

// ----------

func errFatal(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
