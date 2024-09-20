package test

import (
	"log"
	"testing"

	"github.com/nipun2003/url-shortner/utils"
)

func TestBase62(t *testing.T) {
	var encodedStr = utils.Base62Encode(100)
	log.Printf("Encoded String: %s", encodedStr)
	// if encodedStr != "1C" {
	// 	t.Errorf("Expected 1C but got %s", encodedStr)
	// }
}
