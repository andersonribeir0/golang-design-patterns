package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
}

type ManufactoringDirector struct {
	builder BuildProcess
}

type CarBuilder struct {
	v VehicleProduct
}

type BikeBuilder struct {
	v VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Structure string
	Seats     int
}

func (m *ManufactoringDirector) Construct(b BuildProcess) {
	m.builder.SetSeats().SetStructure().SetWheels()
}

func (m *ManufactoringDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "car"
	return c
}
func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "bike"
	return b
}
func (b *BikeBuilder) Build() VehicleProduct {
	return b.v
}
