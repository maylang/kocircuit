package sys

import (
	"github.com/kocircuit/kocircuit/lang/circuit/compile"
	"github.com/kocircuit/kocircuit/lang/circuit/model"
	"github.com/kocircuit/kocircuit/lang/go/runtime"
)

type Compile struct {
	RepoDirs []string `ko:"name=repoDir"`
	PkgPath  string   `ko:"name=pkgPath"`
	Show     bool     `ko:"name=show"`
}

type CompileResult struct {
	Compile *Compile         `ko:"name=compile"`
	Repo    model.Repo       `ko:"name=repo"`
	Stats   *model.RepoStats `ko:"name=stats"`
	Error   error            `ko:"name=error"`
}

func (c *Compile) Play(ctx *runtime.Context) *CompileResult {
	r := &CompileResult{Compile: c}
	if r.Repo, r.Error = compile.CompileRepo(
		c.RepoDirs,
		c.PkgPath,
	); r.Error != nil {
		return r
	}
	r.Stats = r.Repo.Stats()
	ctx.Printf(
		"compiled functions=%d steps=%d steps-per-function=%0.2f",
		r.Stats.TotalFunc, r.Stats.TotalStep, r.Stats.StepPerFunc,
	)
	if c.Show {
		ctx.Printf("%s\n", r.Repo.BodyString())
	}
	return r
}
