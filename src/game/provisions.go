package game

type Provisions struct {
	Fuel        int
	Passengers  int
	Food        int
	Water       int
	Oxygen      int
	RepairParts int
	Materials   int
	Supplies    int
}

func (p *Provisions) LosePassengers(amount int) {
	if amount > p.Passengers {
		p.Passengers = 0
		return
	}
	p.Passengers -= amount
}

func (p *Provisions) LoseOxygen(amount int) {
	if amount > p.Oxygen {
		p.Oxygen = 0
		return
	}
	p.Oxygen -= amount
}

func (p *Provisions) LoseWater(amount int) {
	if amount > p.Water {
		p.Water = 0
		return
	}
	p.Water -= amount
}

func (p *Provisions) LoseFood(amount int) {
	if amount > p.Food {
		p.Food = 0
		return
	}
	p.Food -= amount
}

func (p *Provisions) LoseRepairParts(amount int) {
	if amount > p.RepairParts {
		p.RepairParts = 0
		return
	}
	p.RepairParts -= amount
}

func (p *Provisions) LoseSupplies(amount int) {
	if amount > p.Supplies {
		p.Supplies = 0
		return
	}
	p.Supplies -= amount
}

func (p *Provisions) LoseMaterials(amount int) {
	if amount > p.Materials {
		p.Materials = 0
		return
	}
	p.Materials -= amount
}
