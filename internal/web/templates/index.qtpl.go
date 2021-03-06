// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import (
	"github.com/dustin/go-humanize"
)

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
	ChannelCount int64
	BotCount     int64
}

func (p *IndexPage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamPageBrand(qw422016)
	qw422016.N().S(`
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
        <div class="tile is-ancestor is-vertical content">
            <div class="tile">
                <div class="tile is-parent">
                    <div class="tile is-child box">
                        <p class="title">`)
	p.StreamPageBrand(qw422016)
	qw422016.N().S(`</p>
                        <p>
                            `)
	p.StreamPageBrand(qw422016)
	qw422016.N().S(` is a Twitch chat bot, with:
                        </p>
                        <ul>
                            <li>Moderation</li>
                            <li>Custom commands</li>
                            <li>Repeated commands</li>
                            <li>Quotes</li>
                            <li>Variables</li>
                            <li>LastFM, Steam integration, and more!</li>
                        </ul>
                    </div>
                </div>
                <div class="tile is-4 is-vertical is-parent">
                    <div class="tile is-child box">
                        <p class="title">Join</p>
                        <p>
                            To have the bot join your channel, simply type <code>!join</code> in the bot's Twitch chat.
                            You may need to <a href="/login">log in</a> to enable some features.
                        </p>
                    </div>
                    <div class="tile is-child box">
                        <p class="title">Help</p>
                        <p>
                            Check out the <a href="/docs">documentation</a> for information about the builtin commands,
                            custom commands, and more. For questions, check out our <a href="https://discord.gg/V9Eza32">Discord server</a>.
                        </p>
                    </div>
                </div>
            </div>
            <div class="tile is-parent">
                <div class="tile is-child box">
                    <div class="level">
                        <div class="level-item has-text-centered">
                            <div>
                                <p class="heading">Active channels</p>
                                <p class="title">`)
	qw422016.E().S(humanize.Comma(p.ChannelCount))
	qw422016.N().S(`</p>
                            </div>
                        </div>
                        <div class="level-item has-text-centered">
                            <div>
                                <p class="heading">Active bots</p>
                                <p class="title">`)
	qw422016.E().S(humanize.Comma(p.BotCount))
	qw422016.N().S(`</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
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
