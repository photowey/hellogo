package prec

import (
	"fmt"
	"net/http"
	"sync"

	zero "github.com/zeromicro/go-zero/core/service"
)

func morning(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintln(w, "morning!")
}

func evening(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintln(w, "evening!")
}

type Morning struct{}

func (m Morning) Start() {
	http.HandleFunc("/morning", morning)
	_ = http.ListenAndServe("localhost:8080", nil)
}

func (m Morning) Stop() {
	fmt.Println("Stop morning service...")
}

type Evening struct{}

func (e Evening) Start() {
	http.HandleFunc("/evening", evening)
	_ = http.ListenAndServe("localhost:8081", nil)
}

func (e Evening) Stop() {
	fmt.Println("Stop evening service...")
}

func Run() {
	// v1()
	// v2()
	v3()
}

func v1() {
	fmt.Println("Start morning service...")
	var morning Morning
	morning.Start()
	defer morning.Stop()

	fmt.Println("Start evening service...")
	var evening Evening
	evening.Start()
	defer evening.Stop()
}

func v2() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Start morning service...")
		var morning Morning
		defer morning.Stop()
		morning.Start()
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Start evening service...")
		var evening Evening
		defer evening.Stop()
		evening.Start()
	}()

	wg.Wait()
}

func v3() {
	group := zero.NewServiceGroup()
	defer group.Stop()
	group.Add(Morning{})
	group.Add(Evening{})
	group.Start()
}
