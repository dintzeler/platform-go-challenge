package models

type Insight struct {
	ID string `json:"id"`
	Text string `json:"text"`
	Description string `json:"description"`
}

func (i *Insight) GetID() string {
	return i.ID
}

func (i *Insight) GetType() string {
	return "insight"
}

func (i *Insight) GetDescription() string {
	return i.Description
}

func (i *Insight) SetDescription(desc string) {
	i.Description = desc
}