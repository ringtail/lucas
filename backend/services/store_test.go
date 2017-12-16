package services

import (
	"testing"
)

var store *Store

func init() {
	store, _ = NewWithOutTLS("http://127.0.0.1:2379")
}

func Test_Store_List(t *testing.T) {
}

func Test_Store_Put(t *testing.T) {

}

func Test_Store_Update(t *testing.T) {

}

func Test_Store_Delete(t *testing.T) {

}
