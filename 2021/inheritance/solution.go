package main

type ThroneInheritance struct {
	king   string
	childs map[string][]string
	dead   map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		king:   kingName,
		childs: map[string][]string{},
		dead:   make(map[string]bool),
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.childs[parentName] = append(this.childs[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	result := make([]string, 0)

	var preOrder func(name string)
	preOrder = func(name string) {
		if !this.dead[name] {
			result = append(result, name)
		}

		for _, ch := range this.childs[name] {
			preOrder(ch)
		}
	}
	preOrder(this.king)
	return result
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
