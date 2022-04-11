package proxy

import (
    "testing"
)

func TestProxy(t *testing.T) {
    var human Human
    human = &Proxy{}
    response := human.Say()
    if response != "pre:I'm SuperMan:post" {
        t.Fail()
    }
}
