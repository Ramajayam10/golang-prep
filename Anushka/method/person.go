package method

type Person struct{
	Name string
	Age int
}

func (p Person) IsAdult() bool {
	return p.Age >= 18

}

func (p *Person) HaveBirthday(){
	p.Age++
}