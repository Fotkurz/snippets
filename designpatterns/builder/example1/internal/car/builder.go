package car

type Builder interface {
	Reset()
	SetSeats(seats int) Builder
	SetEngine(engine string) Builder
	Build() *Car
}

type CarAutomaticBuilder struct {
	car *Car
}

func NewCarAutomaticBuilder() Builder {
	return CarAutomaticBuilder{
		car: &Car{},
	}
}

func (cab CarAutomaticBuilder) Reset() {
	cab.car = &Car{}
}

func (cab CarAutomaticBuilder) SetSeats(num int) Builder {
	cab.car.seats = num
	return cab
}

func (cab CarAutomaticBuilder) SetEngine(engine string) Builder {
	cab.car.engine = engine
	return cab
}

func (cab CarAutomaticBuilder) Build() *Car {
	return cab.car
}

type Car struct {
	seats  int
	engine string
}
