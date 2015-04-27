package store

type store struct {
}

func GetStorage(data_dir string) *store {
	return new(store)
}
