package car

type Director struct{}

func (d Director) MakeAutomaticCar(builder Builder) Car {
	car := builder.
		SetEngine("automatic").
		SetSeats(5).
		Build()

	return *car
}
