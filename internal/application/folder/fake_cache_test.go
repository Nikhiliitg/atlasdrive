package folder

import "context"

type fakeCache struct {
	store map[string]interface{}
}

func newFakeCache() *fakeCache {
	return &fakeCache{
		store: make(map[string]interface{}),
	}
}

func (f *fakeCache) Get(ctx context.Context, key string, dest interface{}) (bool, error) {
	val, ok := f.store[key]
	if !ok {
		return false, nil
	}
	*(dest.(*FolderContents)) = val.(FolderContents)
	return true, nil
}

func (f *fakeCache) Set(ctx context.Context, key string, value interface{}) error {
	f.store[key] = value
	return nil
}
