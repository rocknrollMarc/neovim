package neovim

import (
	"net"
	"sync"

	"github.com/vmihailenco/msgpack"
)

type Encoder func() error
type Decoder func() (interface{}, error)

type Client struct {
	conn     net.Conn
	dec      *msgpack.Decoder
	enc      *msgpack.Encoder
	next_req uint32
	resp_map *sync_map
	lock     *sync.Mutex
}

// neovim types

type Buffer struct {
	Id     uint32
	client *Client
}

type API struct {
	Classes   []APIClass
	Functions []APIFunction
}

type APIClass struct {
	Name string
}

type APIFunction struct {
	Name              string
	ReturnType        string
	Id                uint32
	CanFail           bool
	ReceivesChannelId bool
	Parameters        []APIFunctionParameter
}

type APIFunctionParameter struct {
	Type, Name string
}
