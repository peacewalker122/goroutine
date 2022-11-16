package syncpackage

import (
	"log"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	val := []string{"HELLO WORLD","HALO DUNIA","HAI DUNIA"}
	new := Newg(&sync.WaitGroup{},val...)
	
	log.Println(new.Helloworld())
}
