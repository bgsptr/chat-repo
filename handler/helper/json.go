// helper package
package helper

import (
	"encoding/json"
	"io"
	"log"
)

type Helper interface {
	EncodeJson(rw io.Writer, v interface{}) error
	DecodeJson(r io.Reader, v interface{}) error
}

type HelperHandler struct{}

func (h *HelperHandler) EncodeJson(rw io.Writer, v interface{}) error {
	log.Println("error encode")
	d := json.NewEncoder(rw)
	return d.Encode(v)
}

func (h *HelperHandler) DecodeJson(r io.Reader, v interface{}) error {
	log.Println("error decode")
	d := json.NewDecoder(r)
	return d.Decode(v)
}