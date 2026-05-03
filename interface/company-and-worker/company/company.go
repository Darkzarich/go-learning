package company

type Worker interface {
	// All a worker should do is work
	Work(tasks []string) string
}

type Company struct {
	// personal of the company, we use here Worker interface, 
	// so everyone in this array should be able to Work
	personal []Worker
}

// We're hiring a person, a robot, whatever, but who can Work
func (c *Company) Hire(newbie Worker) {
	c.personal = append(c.personal, newbie)
}

func (c Company) Process(id int, tasks []string) (res string) {
	return c.personal[id].Work(tasks)
}
