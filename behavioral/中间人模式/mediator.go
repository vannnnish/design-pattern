package main

type Mediator struct {
	cpu  *CPU
	gpu  *GPU
	disk *Disk
	mem  *Mem
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CPU:
		m.cpu.Process(inst.data)
	case *GPU:
		m.gpu.Display(inst.data)
	case *Disk:
		m.disk.Store(inst.data)
	case *Mem:
		m.mem.Dump(inst.data)
	}
}
