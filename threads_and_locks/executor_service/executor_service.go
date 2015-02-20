package executor_service

type Worker interface {
	Run()
}

type ExecutorService struct {
	MaxPoolSize int
	Jobs chan Worker
}

func NewExecutorService(maxPoolSize int) *ExecutorService {
	e := &ExecutorService{
		MaxPoolSize: maxPoolSize,
		Jobs: make(chan Worker),
	}
	for i:=0; i<e.MaxPoolSize; i++ {
		go func(){
			for job := range e.Jobs {
				job.Run()
			}
		}()
	}
	return e
}



func (e *ExecutorService) Execute(toDo Worker) {
	e.Jobs <- toDo
}
