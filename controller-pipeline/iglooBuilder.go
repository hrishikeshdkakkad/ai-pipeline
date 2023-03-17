package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = ""
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 0
}

func (b *IglooBuilder) getHouse() Response {
	return Response{}
}
