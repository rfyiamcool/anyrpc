package anyrpc

import (
	"bytes"
	"encoding/gob"
	"github.com/vmihailenco/msgpack"
)

type RpcError struct {
	Message string
}

func (r RpcError) Error() string {
	return r.Message
}

func RegisterType(value interface{}) (err error) {
	gob.Register(value)
	return
}

type rpcData struct {
	Name string
	Args []interface{}
}

func encodeData(name string, args []interface{}) ([]byte, error) {
	d := rpcData{}
	d.Name = name
	d.Args = args

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(d); err != nil {
		return nil, err
	} else {
		return buf.Bytes(), nil
	}
}

func encodeDataMsgp(name string, args []interface{}) ([]byte, error) {
	d := rpcData{}
	d.Name = name
	d.Args = args

	buf, err := msgpack.Marshal(&d)
	return buf, err
}

func decodeData(data []byte) (name string, args []interface{}, err error) {
	var (
		d   rpcData
		buf = bytes.NewBuffer(data)
	)
	dec := gob.NewDecoder(buf)
	if err = dec.Decode(&d); err != nil {
		return
	}

	name = d.Name
	args = d.Args

	return
}

func decodeDataMsgp(data []byte) (name string, args []interface{}, err error) {
	var d rpcData
	if err = msgpack.Unmarshal(data, &d); err != nil {
		return
	}

	name = d.Name
	args = d.Args

	return
}
