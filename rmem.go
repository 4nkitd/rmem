package main

import (
	"rmem/entity"
)

type RMEM struct {
	Nodes []entity.Node
}

func (r *RMEM) CreateNode(nodeName string) (bool, string) {

	if r.doesNodeByThisNameExists(nodeName) {
		return false, "Node By This Name Already Exists"
	}

	tmpNode := entity.Node{
		Name: nodeName,
	}

	r.Nodes = append(r.Nodes, tmpNode)

	return true, "Node Created"

}

func (r *RMEM) doesNodeByThisNameExists(nodeName string) bool {
	for _, v := range r.Nodes {
		if v.Name == nodeName {
			return true
		}
	}
	return false
}
