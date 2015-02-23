all: install

install:
	go install ./
	go install ./threads_locks/wiki
	go install ./threads_locks/word_count_batch_sync_map
	go install ./threads_locks/word_count



