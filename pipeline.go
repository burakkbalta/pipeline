package pipeline

type PipeLine[K any] struct {
	head *pipelineStep[K]
	tail *pipelineStep[K]
}

func (p *PipeLine[K]) Execute(context *K) {
	if p.head != nil {
		p.head.Execute(context)
	}
}

func CreatePipeline[K any]() PipeLine[K] {
	return PipeLine[K]{
		head: nil,
		tail: nil,
	}
}

func (p *PipeLine[K]) Next(step Step[K]) *PipeLine[K] {
	newPipelineStep := createPipelineStep(step)

	if p.head == nil {
		p.head = &newPipelineStep
		p.tail = &newPipelineStep
	} else {
		p.tail.next = &newPipelineStep
		p.tail = p.tail.next
	}

	return p
}
