package entity

type Branch struct {
	ID int32

	Name     string
	KeyStore []KeyStore

	COUNT int64
}

func (r *Branch) Create(key string, value string) (bool, string) {

	exists, _ := r.KeyStoreExists(key)
	if exists {
		return false, "Branch By This Name Already Exists"
	}

	tmpKeyStore := KeyStore{
		Key:   key,
		Value: value,
	}

	r.KeyStore = append(r.KeyStore, tmpKeyStore)

	return true, "Branch Created"

}

func (r *Branch) UpdateBranch(keyName string, ID int32) (updated bool, err string) {

	exists, branchIndex := r.KeyStoreExists(keyName)

	if !exists {
		updated = false
	}

	r.KeyStore[branchIndex].Key = keyName

	updated = false
	return

}

func (r *Branch) DeleteBranch(branchName string) (deleted bool, err string) {

	exists, branchIndex := r.KeyStoreExists(branchName)

	if !exists {
		deleted = false
		err = "Branch Doesn't Exists"
		return
	}

	r.KeyStore = append(r.KeyStore[:branchIndex], r.KeyStore[branchIndex+1:]...)

	deleted = true
	err = "Done"
	return

}

func (r *Branch) FindBranch(KeyName string) (branch KeyStore, err string) {

	for _, v := range r.KeyStore {
		if v.Key == KeyName {
			branch = v
			return
		}
	}
	err = "Branch Not Found"
	return
}

func (r *Branch) GetAllBranch() (branches []KeyStore) {

	branches = r.KeyStore
	return

}

func (r *Branch) CountAllBranch() int {
	return len(r.KeyStore)
}

func (r *Branch) KeyStoreExists(Key string) (exists bool, keyIndex int) {

	for index, v := range r.KeyStore {
		if v.Key == Key {

			exists = true
			keyIndex = index
			return
		}
	}

	exists = false
	return

}
