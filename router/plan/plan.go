package plan

import (
	"github.com/pg-sharding/spqr/pkg/models/kr"
	"github.com/pg-sharding/spqr/pkg/spqrlog"
)

type Plan interface {
	iPlan()
	ExecutionTargets() []*kr.ShardKey
}

type ScatterPlan struct {
	Plan
	SubPlan Plan
	/* Empty means execute everywhere */
	ExecTargets []*kr.ShardKey
}

func (sp ScatterPlan) ExecutionTargets() []*kr.ShardKey {
	return sp.ExecTargets
}

type ModifyTable struct {
	Plan
}

func (mt ModifyTable) ExecutionTargets() []*kr.ShardKey {
	return nil
}

type ShardDispatchPlan struct {
	Plan

	ExecTarget         *kr.ShardKey
	TargetSessionAttrs string
}

func (sms ShardDispatchPlan) ExecutionTargets() []*kr.ShardKey {
	return []*kr.ShardKey{sms.ExecTarget}
}

type DDLState struct {
	Plan
}

func (ddl DDLState) ExecutionTargets() []*kr.ShardKey {
	return nil
}

type RandomDispatchPlan struct {
	Plan
}

func (rdp RandomDispatchPlan) ExecutionTargets() []*kr.ShardKey {
	return nil
}

type VirtualPlan struct {
	Plan
}

func (vp VirtualPlan) ExecutionTargets() []*kr.ShardKey {
	return nil
}

type CopyState struct {
	Plan
}

func (cs CopyState) ExecutionTargets() []*kr.ShardKey {
	return nil
}

type ReferenceRelationState struct {
	Plan
}

func (rrs ReferenceRelationState) ExecutionTargets() []*kr.ShardKey {
	return nil
}

const NOSHARD = ""

// TODO : unit tests
func Combine(p1, p2 Plan) Plan {
	if p1 == nil && p2 == nil {
		return nil
	}
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}

	spqrlog.Zero.Debug().
		Interface("plan1", p1).
		Interface("plan2", p2).
		Msg("combine two plans")

	switch shq1 := p1.(type) {
	case ScatterPlan:
		return p1
	case RandomDispatchPlan:
		return p2
	case ReferenceRelationState:
		return p2
	case ShardDispatchPlan:
		switch shq2 := p2.(type) {
		case ScatterPlan:
			return p2
		case ReferenceRelationState:
			return p1
		case ShardDispatchPlan:
			if shq2.ExecTarget.Name == shq1.ExecTarget.Name {
				return p1
			}
		}
	}

	/* execute on all shards */
	return ScatterPlan{}
}
