// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

type IndexPage struct {
	BasePage
}

func (p *IndexPage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
HortBot
`)
}

func (p *IndexPage) WritePageTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *IndexPage) PageTitle() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *IndexPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<section class="section">
    <div class="container">
        <h1 class="title">
            Welcome
        </h1>
        <h2 class="subtitle">
            To the index
        </h2>
    </div>
</section>
`)
}

func (p *IndexPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *IndexPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
