package executor_service

// Interface for wrapping work
type Job interface {
	Run()
}

type ExecutorService struct {
	MaxPoolSize int  // Max parallel jobs executing
	Jobs chan Job    // Channel queuing all jobs
}

func NewExecutorService(maxPoolSize int) *ExecutorService {
	e := &ExecutorService{
		MaxPoolSize: maxPoolSize,
		Jobs: make(chan Job),
	}
	// Dispatch N workers to distrubte the jobs
	for i:=0; i<e.MaxPoolSize; i++ {
		go Worker(e.Jobs)
	}
	return e
}

func Worker(jobs chan Job){
	for job := range jobs { // Grab next job
		// Execute work
		job.Run()
	}
}

func (e *ExecutorService) Execute(toDo Job) {
	// Enqueue incoming job
	e.Jobs <- toDo
}
