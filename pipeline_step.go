package pipeline

type stepFunction[K any] func(context *K)

type Step[K any] interface {
	Execute(context *K)
}

type pipelineStep[K any] struct {
	stepFunction stepFunction[K]
	next         *pipelineStep[K]
}

func (p *pipelineStep[K]) Execute(context *K) {
	p.stepFunction(context)
	if p.next != nil {
		p.next.Execute(context)
	}
}

func (p *pipelineStep[K]) Next(next *pipelineStep[K]) {
	p.next = next
}

func createPipelineStep[K any](step Step[K]) pipelineStep[K] {
	return pipelineStep[K]{
		stepFunction: step.Execute,
		next:         nil,
	}
}
