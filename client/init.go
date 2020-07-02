package client

var projects map[string]int

func (c *ClientConfig) Init() error {
	if p, err := c.GetProjects(); err != nil {
		return err
	} else {
		projects = p
	}
	return nil
}
