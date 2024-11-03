package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// memasukkan data ke dalam map
//** store(key, value) ******
// mengambil data menggunkan key
//******* load(key) *******************//
// menghapus data menggunkan key//
//******* Delete(key) *******************//
// mengiterasi data yang ada //
// ******* range(fung(key,valu interface{}) bool) *******************//

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	// memasukkan data ke dalam map
	data.Store(value,value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i <= 10; i++ {
		group.Add(1)
		go AddToMap(data, i, group)
	}

	group.Wait()
	
	// men iterasi data yang ada di map
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, " : ", value)
		return true
	})

	dataMap, _:= data.Load(10)
	fmt.Println(dataMap)
}