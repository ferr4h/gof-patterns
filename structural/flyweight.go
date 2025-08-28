package structural

type TreeType struct {
	color   string
	texture string
}

type TreeTypeFactory struct {
	treeTypes map[string]*TreeType
}

func NewTreeTypeFactory() *TreeTypeFactory {
	return &TreeTypeFactory{make(map[string]*TreeType)}
}

func (factory *TreeTypeFactory) GetTreeType(color, texture string) *TreeType {
	key := color + "-" + texture

	if treeType, exists := factory.treeTypes[key]; exists {
		return treeType
	}

	newTreeType := &TreeType{color, texture}
	factory.treeTypes[key] = newTreeType
	return newTreeType
}

type Tree struct {
	name     string
	treeType *TreeType
}

type Forest struct {
	trees           []*Tree
	treeTypeFactory *TreeTypeFactory
}

func NewForest() *Forest {
	return &Forest{make([]*Tree, 0), NewTreeTypeFactory()}
}

func (f *Forest) AddTree(name, color, texture string) {
	treeType := f.treeTypeFactory.GetTreeType(color, texture)
	f.trees = append(f.trees, &Tree{name, treeType})
}
