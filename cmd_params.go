package main

type GitStateComponent interface {

}

type ParamAction func() (GitStateComponent, error)
type CmdParams map[string]ParamAction

func (params CmdParams) AddAction(name string, action ParamAction) {
	params[name] = action
}
