package word_count_batch_sync_map

type Counter struct {
	LocalCounts map[string]int64
}

func NewCounter() *Counter {
	return &Counter{LocalCounts: make(map[string]int64)}
}

func (c* Counter) CountWord(word string) {
	if val, ok := c.LocalCounts[word]; ok {
		c.LocalCounts[word] = val + 1
	} else {
		c.LocalCounts[word] = 1
	}
}

func (c* Counter) MergeMap(source map[string]int64) map[string]int64 {
	for key, value := range c.LocalCounts {
		if val, ok := source[key]; ok {
			source[key] = val + value
		} else {
			source[key] = value
		}
	}
	return source
}
