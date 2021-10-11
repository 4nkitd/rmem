package entity

type Node struct {
	ID int32

	Name   string
	Branch []Branch

	COUNT int64
}

func (r *Node) CreateBranch(branchName string) (bool, string) {

	exists, _ := r.BranchExists(branchName)
	if exists {
		return false, "Branch By This Name Already Exists"
	}

	tmpBranch := Branch{
		Name: branchName,
	}

	r.Branch = append(r.Branch, tmpBranch)

	return true, "Branch Created"

}

func (r *Node) UpdateBranch(branchName string, ID int32) (updated bool, err string) {

	exists, branchIndex := r.BranchExists(branchName)

	if !exists {
		updated = false
	}

	r.Branch[branchIndex].Name = branchName

	updated = false
	return

}

func (r *Node) DeleteBranch(branchName string) (deleted bool, err string) {

	exists, branchIndex := r.BranchExists(branchName)

	if !exists {
		deleted = false
		err = "Branch Doesn't Exists"
		return
	}

	r.Branch = append(r.Branch[:branchIndex], r.Branch[branchIndex+1:]...)

	deleted = true
	err = "Done"
	return

}

func (r *Node) FindBranch(branchName string) (branch Branch, err string) {

	for _, v := range r.Branch {
		if v.Name == branchName {
			branch = v
			return
		}
	}
	err = "Branch Not Found"
	return
}

func (r *Node) GetAllBranch() (branches []Branch, err string) {

	branches = r.Branch
	return

}

func (r *Node) CountAllBranch() int {
	return len(r.Branch)
}

func (r *Node) BranchExists(branchName string) (exists bool, branchIndex int) {

	for index, v := range r.Branch {
		if v.Name == branchName {

			exists = true
			branchIndex = index
			return
		}
	}

	exists = true
	return

}
