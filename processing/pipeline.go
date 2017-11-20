package processing

// A pipeline is a list of processes
type Pipeline struct {
	processes []Process
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (pl *Pipeline) AddProcess(proc Process) {
	pl.processes = append(pl.processes, proc)
}

func (pl *Pipeline) AddProcesses(procs ...Process) {
	for _, proc := range procs {
		pl.AddProcess(proc)
	}
}

func (pl *Pipeline) Run() {
	for i, proc := range pl.processes {
		if i < len(pl.processes)-1 {
			go proc.Run()
		} else {
			proc.Run()
		}
	}
}
