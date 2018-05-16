package embedded

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/chrislusf/seaweedfs/weed/filer2"
	"github.com/chrislusf/seaweedfs/weed/pb/filer_pb"
)

type EmbeddedStore struct {
	db *leveldb.DB
}

func NewEmbeddedStore(dir string) (filer *EmbeddedStore, err error) {
	filer = &EmbeddedStore{}
	if filer.db, err = leveldb.OpenFile(dir, nil); err != nil {
		return
	}
	return
}

func (filer *EmbeddedStore) InsertEntry(entry *filer2.Entry) (err error) {
	return nil
}

func (filer *EmbeddedStore) AddDirectoryLink(directory *filer2.Entry, delta int32) (err error) {
	return nil
}

func (filer *EmbeddedStore) AppendFileChunk(fullpath filer2.FullPath, fileChunks []*filer_pb.FileChunk) (err error) {
	return nil
}

func (filer *EmbeddedStore) FindEntry(fullpath filer2.FullPath) (found bool, entry *filer2.Entry, err error) {
	return false, nil, nil
}

func (filer *EmbeddedStore) DeleteEntry(fullpath filer2.FullPath) (entry *filer2.Entry, err error) {
	return nil, nil
}

func (filer *EmbeddedStore) ListDirectoryEntries(fullpath filer2.FullPath, startFileName string, inclusive bool, limit int) (entries []*filer2.Entry, err error) {
	return nil, nil
}
