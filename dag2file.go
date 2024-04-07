package merkledag

import (
	"encoding/json"
	"fmt"
)

// Hash to file
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	var data []byte 
	var tree Object
	h := hp.Get()
	h.Write(hash)
	value, err := store.Get(h.Sum(nil))
	if err != nil {
		_ = fmt.Errorf("get value from store failed, err: %s", err)
		return nil
	}


	err = json.Unmarshal(value, &tree)
	if err != nil {
		_ = fmt.Errorf("unmarshal value to tree failed, err: %s", err)
		return nil
	}

	
	var start uint64 
	var link Link
	
	for link = range tree.Links {
		if link.Name == path {
			data = tree.Data[start : start+link.Size]
		}
		start += link.Size
	}
	return data
}