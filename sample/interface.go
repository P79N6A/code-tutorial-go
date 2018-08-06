package main

type IEngine struct {
    engine interface{}
}

type MotoEngine struct {
    num int
}
func (m *MotoEngine)GetNum() int {
    return m.num
}
func (m *MotoEngine)WithInt(n int) *MotoEngine {
    m.num = n
    return m
}
func NewIEngine() *IEngine {
    return &IEngine{engine:&MotoEngine{}}
}
func main() {
    ins := NewIEngine()
    ins.engine.GetNum()
}
