package runner

/*import (
	"os"
)*/

type Runner struct {
	
}

func New() (runner *Runner) {
	return new(Runner)
}

func (r *Runner) Start() {
	println("Started")	
}

func (r *Runner) Stop() {

}
