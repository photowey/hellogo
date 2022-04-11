package proxy

type Human interface {
    Say() string
}

type Man struct {
}

func (Man) Say() string {
    return "I'm SuperMan"
}

type Proxy struct {
    subject Man
}

func (p Proxy) Say() string {
    var response string

    // before
    response = p.before(response)
    response += p.subject.Say()
    // after
    response = p.after(response)

    return response
}

func (p Proxy) after(response string) string {
    response += ":post"
    return response
}

func (p Proxy) before(response string) string {
    response += "pre:"
    return response
}
