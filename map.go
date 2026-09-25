package main

func mapCommand(c *config, param string) error {
	return location(c, true)
}

func mapBCommand(c *config, param string) error {
	return location(c, false)
}
