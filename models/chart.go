package models

type Chart struct {
	ID string `json:"id"`
	Title string `json:"title"`
	XAxis string `json:"x_axis"`
	YAxis string `json:"y_axis"`
	Data []Point `json:"data"`
	Description string `json:"description"`
}

func (c *Chart) GetID() string {
	return c.ID
}

func (c *Chart) GetType() string {
	return "chart"
}

func (c *Chart) GetDescription() string {
	return c.Description
}

func (c *Chart) SetDescription(desc string) {
	c.Description = desc
}