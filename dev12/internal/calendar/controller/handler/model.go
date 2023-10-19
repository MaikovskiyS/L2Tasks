package handler

import (
	"dev11/internal/calendar/domain/entity"
	"encoding/json"
)

type Resp struct {
	Result result `json:"result"`
	Error  string `json:"error"`
}
type result struct {
	Msg  string          `json:"msg"`
	Data []*entity.Event `jsin:"data"`
}

func (r *Resp) SetEvents(e []*entity.Event) {
	r.Result.Data = e
}
func (r *Resp) SetMsg(msg string) {
	r.Result.Msg = msg
}
func (r *Resp) SetErr(err string) {
	r.Error = err
}
func (r *Resp) Bytes() []byte {
	bytes, _ := json.Marshal(&r)
	return bytes
}
