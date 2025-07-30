package creational

type Prototype interface {
	Clone() Prototype
}

type Document struct {
	Title    string
	Metadata map[string]string
}

func (d *Document) Clone() Prototype {
	metaData := make(map[string]string, len(d.Metadata))
	for k, v := range d.Metadata {
		metaData[k] = v
	}
	return &Document{
		Title:    d.Title,
		Metadata: metaData,
	}
}
