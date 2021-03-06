// Code generated by qtc from "channel.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import (
	"time"

	"github.com/hako/durafmt"
	"github.com/hortbot/hortbot/internal/cbp"
	"github.com/hortbot/hortbot/internal/db/models"
)

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

type ChannelPage struct {
	BasePage
	Channel *models.Channel
}

func (p *ChannelPage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamPageBrand(qw422016)
	qw422016.N().S(` - `)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`
`)
}

func (p *ChannelPage) WritePageTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelPage) PageTitle() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ChannelPage) StreamPageMeta(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.BasePage.StreamPageMeta(qw422016)
	qw422016.N().S(`
`)
	streamsidebarStyle(qw422016)
	qw422016.N().S(`
<style>
.subtitle {
    padding-left: 1rem;
}
</style>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-table/1.16.0/themes/bulma/bootstrap-table-bulma.min.css" integrity="sha256-u49ra7w4V15McfEDsAvfE6A+W18iGcM7mWlkBmASoAs=" crossorigin="anonymous" />
`)
}

func (p *ChannelPage) WritePageMeta(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageMeta(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelPage) PageMeta() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageMeta(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ChannelPage) StreamPageScripts(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.BasePage.StreamPageScripts(qw422016)
	qw422016.N().S(`
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.4.1/jquery.min.js" integrity="sha256-CSXorXvZcTkaix6Yvo6HppcZGetbYMGWSFlBw8HfCJo=" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-table/1.16.0/bootstrap-table.min.js" integrity="sha256-JFzlEUS2cZGdNFhVNH3GSFuqZFLjzWIjOqG5BY+Yhvw=" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-table/1.16.0/themes/bulma/bootstrap-table-bulma.min.js" integrity="sha256-EiXUuCvnTWCYTHg2nWE3Yd0N0fprGRbVTkPMJ7s4enU=" crossorigin="anonymous"></script>
<script>
function timeFormatter(value) {
    try {
        var d = new Date(value);
        return d.toLocaleString();
    } catch {
        return "";
    }
}

function timeSorter(a, b) {
    try {
        return Date.parse(a) - Date.parse(b);
    } catch {
        return 0;
    }
}
</script>
`)
}

func (p *ChannelPage) WritePageScripts(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageScripts(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelPage) PageScripts() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageScripts(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func streamisActive(qw422016 *qt422016.Writer, a, b string) {
	if a == b {
		qw422016.N().S(`is-active`)
	}
}

func writeisActive(qq422016 qtio422016.Writer, a, b string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamisActive(qw422016, a, b)
	qt422016.ReleaseWriter(qw422016)
}

func isActive(a, b string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeisActive(qb422016, a, b)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ChannelPage) StreamSidebar(qw422016 *qt422016.Writer, item string) {
	qw422016.N().S(`
<div class="is-sidebar-menu is-hidden-mobile">
    <aside class="menu">
        <p class="menu-label">
            General
        </p>
        <ul class="menu-list">
            <li><a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`" class='`)
	streamisActive(qw422016, item, "overview")
	qw422016.N().S(`'>Overview</a></li>
            <li><a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/commands" class='`)
	streamisActive(qw422016, item, "commands")
	qw422016.N().S(`'>Commands</a></li>
            <li><a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/quotes" class='`)
	streamisActive(qw422016, item, "quotes")
	qw422016.N().S(`'>Quotes</a></li>
            <li><a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/lists" class='`)
	streamisActive(qw422016, item, "lists")
	qw422016.N().S(`'>Lists</a></li>
            <li><a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/autoreplies" class='`)
	streamisActive(qw422016, item, "autoreplies")
	qw422016.N().S(`'>Autoreplies</a></li>
            <li><a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/scheduled" class='`)
	streamisActive(qw422016, item, "scheduled")
	qw422016.N().S(`'>Repeated / scheduled</a></li>
            <li><a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/variables" class='`)
	streamisActive(qw422016, item, "variables")
	qw422016.N().S(`'>Variables</a></li>
            <li><a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/highlights" class='`)
	streamisActive(qw422016, item, "highlights")
	qw422016.N().S(`'>Highlights</a></li>
        </ul>
        <p class="menu-label">
            Settings
        </p>
        <ul class="menu-list">
            <li><a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/regulars" class='`)
	streamisActive(qw422016, item, "regulars")
	qw422016.N().S(`'>Regulars</a></li>
            <li><a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/chatrules" class='`)
	streamisActive(qw422016, item, "chatrules")
	qw422016.N().S(`'>Chat rules</a></li>
        </ul>
    </aside>
</div>
`)
}

func (p *ChannelPage) WriteSidebar(qq422016 qtio422016.Writer, item string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamSidebar(qw422016, item)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelPage) Sidebar(item string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WriteSidebar(qb422016, item)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ChannelPage) StreamSidebarMobile(qw422016 *qt422016.Writer, item string) {
	qw422016.N().S(`
<nav class="navbar is-hidden-tablet is-info" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
        <p class="navbar-item">Menu</p>
        <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="sidebarNav">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>

    <div id="sidebarNav" class="navbar-menu">
        <div class="navbar-start">
            <div class="navbar-item has-dropdown is-hoverable">
                <p class="navbar-link">General</p>

                <div class="navbar-dropdown">
                    <a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`" class='navbar-item `)
	streamisActive(qw422016, item, "overview")
	qw422016.N().S(`'>Overview</a>
                    <a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/commands" class='navbar-item `)
	streamisActive(qw422016, item, "commands")
	qw422016.N().S(`'>Commands</a>
                    <a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/quotes" class='navbar-item `)
	streamisActive(qw422016, item, "quotes")
	qw422016.N().S(`'>Quotes</a>
                    <a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/lists" class='navbar-item `)
	streamisActive(qw422016, item, "lists")
	qw422016.N().S(`'>Lists</a>
                    <a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/autoreplies" class='navbar-item `)
	streamisActive(qw422016, item, "autoreplies")
	qw422016.N().S(`'>Autoreplies</a>
                    <a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/scheduled" class='navbar-item `)
	streamisActive(qw422016, item, "scheduled")
	qw422016.N().S(`'>Repeated / scheduled</a>
                    <a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/variables" class='navbar-item `)
	streamisActive(qw422016, item, "variables")
	qw422016.N().S(`'>Variables</a>
                    <a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/highlights" class='navbar-item `)
	streamisActive(qw422016, item, "highlights")
	qw422016.N().S(`'>Highlights</a>
                </div>
            </div>

            <div class="navbar-item has-dropdown is-hoverable">
                <p class="navbar-link">Settings</a>

                <div class="navbar-dropdown">
                    <a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/regulars" class='navbar-item `)
	streamisActive(qw422016, item, "regulars")
	qw422016.N().S(`'>Regulars</a>
                    <a href="/c/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`/chatrules" class='navbar-item `)
	streamisActive(qw422016, item, "chatrules")
	qw422016.N().S(`'>Chat rules</a>
                </div>
            </div>
        </div>
    </div>
</nav>
`)
}

func (p *ChannelPage) WriteSidebarMobile(qq422016 qtio422016.Writer, item string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamSidebarMobile(qw422016, item)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelPage) SidebarMobile(item string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WriteSidebarMobile(qb422016, item)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ChannelPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamSidebarMobile(qw422016, "overview")
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    `)
	p.StreamSidebar(qw422016, "overview")
	qw422016.N().S(`

    <div class="column is-main-content content">
        <span class="title is-1">`)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`</span>
        <hr>

        <a class="button is-outlined" href="https://www.twitch.tv/`)
	qw422016.N().U(p.Channel.Name)
	qw422016.N().S(`" target="_blank">
            <span class="icon"><i class="fab fa-twitch"></i></span>
            <span>Twitch</span>
        </a>

        `)
	if p.Channel.LastFM != "" {
		qw422016.N().S(`
            <a class="button is-outlined" href="https://www.last.fm/user/`)
		qw422016.N().U(p.Channel.LastFM)
		qw422016.N().S(`" target="_blank">
                <span class="icon"><i class="fab fa-lastfm"></i></span>
                <span>LastFM</span>
            </a>
        `)
	}
	qw422016.N().S(`

        `)
	if p.Channel.SteamID != "" {
		qw422016.N().S(`
            <a class="button is-outlined" href="https://steamcommunity.com/profiles/`)
		qw422016.N().U(p.Channel.SteamID)
		qw422016.N().S(`" target="_blank">
                <span class="icon"><i class="fab fa-steam"></i></span>
                <span>Steam</span>
            </a>
        `)
	}
	qw422016.N().S(`

        `)
	if p.Channel.ExtraLifeID != 0 {
		qw422016.N().S(`
            <a class="button is-outlined" href="https://www.extra-life.org/index.cfm?fuseaction=donorDrive.participant&participantID=`)
		qw422016.N().D(p.Channel.ExtraLifeID)
		qw422016.N().S(`" target="_blank">
                <span class="icon"><i class="fas fa-gamepad"></i></span>
                <span>Extra-Life</span>
            </a>
        `)
	}
	qw422016.N().S(`

        <br>
        <br>

        <p><b>Bot name:</b> <code>`)
	qw422016.E().S(p.Channel.BotName)
	qw422016.N().S(`</code></p>
        <p><b>Command prefix:</b> <code>`)
	qw422016.E().S(p.Channel.Prefix)
	qw422016.N().S(`</code></p>
    </div>
</div>

<script>
var hash = window.location.hash;
var realPage = window.location.origin + window.location.pathname;
if (hash) {
    if (hash == "#overview") {
        history.replaceState(null, null, ' ');
    } else {
        var redirect = {
            "#commands": "commands",
            "#quotes": "quotes",
            "#autoreplies": "autoreplies",
            "#scheduled": "scheduled",
            "#regulars": "regulars",
            "#chatrules": "chatrules",
        }[hash];

        window.location.href = realPage + "/" + redirect;
    }
}
</script>
`)
}

func (p *ChannelPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type ChannelCommandsPage struct {
	ChannelPage
	Commands models.CustomCommandSlice
}

func streamaccessLevel(qw422016 *qt422016.Writer, level string) {
	qw422016.N().S(`
`)
	switch level {
	case models.AccessLevelEveryone:
		qw422016.N().S(`
All
`)
	case models.AccessLevelSubscriber:
		qw422016.N().S(`
<span class="has-text-info">Subs</span>
`)
	case models.AccessLevelModerator:
		qw422016.N().S(`
<span class="has-text-success">Mods</span>
`)
	case models.AccessLevelBroadcaster:
		qw422016.N().S(`
<span class="has-text-danger">Broadcaster</span>
`)
	case models.AccessLevelAdmin:
		qw422016.N().S(`
<span>Admins</span>
`)
	default:
		qw422016.N().S(`
Unknown
`)
	}
	qw422016.N().S(`
`)
}

func writeaccessLevel(qq422016 qtio422016.Writer, level string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamaccessLevel(qw422016, level)
	qt422016.ReleaseWriter(qw422016)
}

func accessLevel(level string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeaccessLevel(qb422016, level)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func streamconvertCommand(qw422016 *qt422016.Writer, s string) {
	c, _ := cbp.Parse(s)

	for _, node := range c {
		if node.Text != "" {
			qw422016.E().S(node.Text)
		} else {
			qw422016.N().S(`<code>(_`)
			qw422016.E().S(cbp.NodesString(node.Children))
			qw422016.N().S(`_)</code>`)
		}
	}
}

func writeconvertCommand(qq422016 qtio422016.Writer, s string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamconvertCommand(qw422016, s)
	qt422016.ReleaseWriter(qw422016)
}

func convertCommand(s string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeconvertCommand(qb422016, s)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ChannelCommandsPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamSidebarMobile(qw422016, "commands")
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    `)
	p.StreamSidebar(qw422016, "commands")
	qw422016.N().S(`

    <div class="column is-main-content content">
        <span class="title is-1">`)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`</span><span class="subtitle is-3">Commands</span>
        <hr>

        `)
	if len(p.Commands) == 0 {
		qw422016.N().S(`
        <p>No commands.</p>
        `)
	} else {
		qw422016.N().S(`
        <table
            class="table is-striped is-hoverable is-fullwidth"
            data-toggle="table"
            data-sort-class="table-active"
            data-sort-name="command"
            data-search="true"
            data-sortable="true"
        >
            <thead>
                <tr>
                    <th data-sortable="true" data-field="command">Command</th>
                    <th data-sortable="true">Access</th>
                    <th data-sortable="true">Response</th>
                    <th data-sortable="true">Count</th>
                    <th data-sortable="true">Editor</th>
                    <th data-sortable="true" data-formatter="timeFormatter" data-sorter="timeSorter">Edited</th>
                </tr>
            </thead>
            <tbody>
                `)
		for _, c := range p.Commands {
			qw422016.N().S(`
                <tr>
                    <td><code>`)
			qw422016.E().S(p.Channel.Prefix)
			qw422016.E().S(c.R.CommandInfo.Name)
			qw422016.N().S(`</code></td>
                    <td>`)
			streamaccessLevel(qw422016, c.R.CommandInfo.AccessLevel)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.N().S(convertCommand(c.Message))
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().V(c.R.CommandInfo.Count)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(c.R.CommandInfo.Editor)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(c.UpdatedAt.Format(time.RFC3339))
			qw422016.N().S(`</td>
                </tr>
                `)
		}
		qw422016.N().S(`
            </tbody>
        </table>
        `)
	}
	qw422016.N().S(`
    </div>
</div>
`)
}

func (p *ChannelCommandsPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelCommandsPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type ChannelQuotesPage struct {
	ChannelPage
	Quotes models.QuoteSlice
}

func (p *ChannelQuotesPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamSidebarMobile(qw422016, "quotes")
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    `)
	p.StreamSidebar(qw422016, "quotes")
	qw422016.N().S(`

    <div class="column is-main-content content">
        <span class="title is-1">`)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`</span><span class="subtitle is-3">Quotes</span>
        <hr>

        `)
	if len(p.Quotes) == 0 {
		qw422016.N().S(`
        <p>No quotes.</p>
        `)
	} else {
		qw422016.N().S(`
        <table
            class="table is-striped is-hoverable is-fullwidth"
            data-toggle="table"
            data-sort-class="table-active"
            data-sort-name="number"
            data-search="true"
            data-sortable="true"
        >
            <thead>
                <tr>
                    <th data-sortable="true" data-field="number">#</th>
                    <th data-sortable="true">Quote</th>
                    <th data-sortable="true">Editor</th>
                    <th data-sortable="true" data-formatter="timeFormatter" data-sorter="timeSorter">Edited</th>
                </tr>
            </thead>
            <tbody>
                `)
		for _, q := range p.Quotes {
			qw422016.N().S(`
                <tr>
                    <td>`)
			qw422016.N().D(q.Num)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(q.Quote)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(q.Editor)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(q.UpdatedAt.Format(time.RFC3339))
			qw422016.N().S(`</td>
                </tr>
                `)
		}
		qw422016.N().S(`
            </tbody>
        </table>
        `)
	}
	qw422016.N().S(`
    </div>
</div>
`)
}

func (p *ChannelQuotesPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelQuotesPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type ChannelAutorepliesPage struct {
	ChannelPage
	Autoreplies models.AutoreplySlice
}

func (p *ChannelAutorepliesPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamSidebarMobile(qw422016, "autoreplies")
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    `)
	p.StreamSidebar(qw422016, "autoreplies")
	qw422016.N().S(`

    <div class="column is-main-content content">
        <span class="title is-1">`)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`</span><span class="subtitle is-3">Autoreplies</span>
        <hr>

        `)
	if len(p.Autoreplies) == 0 {
		qw422016.N().S(`
        <p>No autoreplies.</p>
        `)
	} else {
		qw422016.N().S(`
        <table
            class="table is-striped is-hoverable is-fullwidth"
            data-toggle="table"
            data-sort-class="table-active"
            data-sort-name="number"
            data-search="true"
            data-sortable="true"
        >
            <thead>
                <tr>
                    <th data-sortable="true" data-field="number">#</th>
                    <th data-sortable="true">Trigger</th>
                    <th data-sortable="true">Regex</th>
                    <th data-sortable="true">Response</th>
                    <th data-sortable="true">Count</th>
                    <th data-sortable="true">Editor</th>
                    <th data-sortable="true" data-formatter="timeFormatter" data-sorter="timeSorter">Edited</th>
                </tr>
            </thead>
            <tbody>
                `)
		for _, a := range p.Autoreplies {
			qw422016.N().S(`
                <tr>
                    <td>`)
			qw422016.N().D(a.Num)
			qw422016.N().S(`</td>
                    <td><code>`)
			qw422016.E().S(a.OrigPattern.String)
			qw422016.N().S(`</code></td>
                    <td><code>`)
			qw422016.E().S(a.Trigger)
			qw422016.N().S(`</code></td>
                    <td>`)
			qw422016.N().S(convertCommand(a.Response))
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.N().D(a.Count)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(a.Editor)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(a.UpdatedAt.Format(time.RFC3339))
			qw422016.N().S(`</td>
                </tr>
                `)
		}
		qw422016.N().S(`
            </tbody>
        </table>
        `)
	}
	qw422016.N().S(`

    </div>
</div>
`)
}

func (p *ChannelAutorepliesPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelAutorepliesPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type ChannelListsPage struct {
	ChannelPage
	Lists models.CommandListSlice
}

func (p *ChannelListsPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamSidebarMobile(qw422016, "lists")
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    `)
	p.StreamSidebar(qw422016, "lists")
	qw422016.N().S(`

    <div class="column is-main-content content">
        <span class="title is-1">`)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`</span><span class="subtitle is-3">Lists</span>
        <hr>

        `)
	if len(p.Lists) == 0 {
		qw422016.N().S(`
        <p>No lists.</p>
        `)
	} else {
		qw422016.N().S(`
        <table
            class="table is-striped is-hoverable is-fullwidth"
            data-toggle="table"
            data-sort-class="table-active"
            data-sort-name="command"
            data-search="true"
            data-sortable="true"
            data-detail-view="true"
            data-detail-formatter="detailFormatter"
        >
            <thead>
                <tr>
                    <th data-sortable="true" data-field="command">Command</th>
                    <th data-sortable="true">Access</th>
                    <th data-sortable="true">Count</th>
                    <th data-sortable="true">Editor</th>
                    <th data-sortable="true" data-formatter="timeFormatter" data-sorter="timeSorter">Edited</th>
                </tr>
            </thead>
            <tbody>
                `)
		for _, l := range p.Lists {
			qw422016.N().S(`
                <tr data-has-detail-view="true">
                    <td><code>`)
			qw422016.E().S(p.Channel.Prefix)
			qw422016.E().S(l.R.CommandInfo.Name)
			qw422016.N().S(`</code></td>
                    <td>`)
			streamaccessLevel(qw422016, l.R.CommandInfo.AccessLevel)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().V(l.R.CommandInfo.Count)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(l.R.CommandInfo.Editor)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(l.UpdatedAt.Format(time.RFC3339))
			qw422016.N().S(`</td>
                </tr>
                `)
		}
		qw422016.N().S(`
            </tbody>
        </table>
        `)
	}
	qw422016.N().S(`
    </div>
</div>
`)
}

func (p *ChannelListsPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelListsPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ChannelListsPage) StreamPageScripts(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.ChannelPage.StreamPageScripts(qw422016)
	qw422016.N().S(`
<script>
`)
	qw422016.N().S(`var all = [`)
	for i, l := range p.Lists {
		if i != 0 {
			qw422016.N().S(`,`)
		}
		qw422016.N().S(`[`)
		for j, item := range l.Items {
			if j != 0 {
				qw422016.N().S(`,`)
			}
			qw422016.N().Q(item)
		}
		qw422016.N().S(`]`)
	}
	qw422016.N().S(`];`)
	qw422016.N().S(`

function detailFormatter(index, row) {
    var html = ["<p><ol>"];

    var items = all[index];

    if (items.length == 0) {
        return "<p>No items.</p>"
    }

    for (var i = 0; i < items.length; i++) {
        html.push("<li>");
        var div = document.createElement("div");
        div.innerText = items[i];
        html.push(div.innerHTML);
        html.push("</li>");
    }

    html.push("</ol></p>")

    return html.join("");
}
</script>
`)
}

func (p *ChannelListsPage) WritePageScripts(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageScripts(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelListsPage) PageScripts() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageScripts(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type ChannelRegularsPage struct {
	ChannelPage
}

func (p *ChannelRegularsPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamSidebarMobile(qw422016, "regulars")
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    `)
	p.StreamSidebar(qw422016, "regulars")
	qw422016.N().S(`

    <div class="column is-main-content content">
        <span class="title is-1">`)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`</span><span class="subtitle is-3">Regulars</span>
        <hr>

        `)
	if len(p.Channel.CustomRegulars) == 0 {
		qw422016.N().S(`
        <p>No regulars.</p>
        `)
	} else {
		qw422016.N().S(`
        <p>The following users are considered regulars in chat:</p>
        <ul>
            `)
		for _, reg := range p.Channel.CustomRegulars {
			qw422016.N().S(`
            <li>`)
			qw422016.E().S(reg)
			qw422016.N().S(`</li>
            `)
		}
		qw422016.N().S(`
        </ul>
        `)
	}
	qw422016.N().S(`
    </div>
</div>
`)
}

func (p *ChannelRegularsPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelRegularsPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type ChannelRulesPage struct {
	ChannelPage
}

func (p *ChannelRulesPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamSidebarMobile(qw422016, "chatrules")
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    `)
	p.StreamSidebar(qw422016, "chatrules")
	qw422016.N().S(`

    <div class="column is-main-content content">
        <span class="title is-1">`)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`</span><span class="subtitle is-3">Chat rules</span>
        <hr>

        <h2>Allowed URLs</h2>

        `)
	if len(p.Channel.PermittedLinks) == 0 {
		qw422016.N().S(`
        <p>No allowed URLs.</p>
        `)
	} else {
		qw422016.N().S(`
        <p>The following URL patterns are allowed:</p>
        <ul>
            `)
		for _, link := range p.Channel.PermittedLinks {
			qw422016.N().S(`
            <li><code>`)
			qw422016.E().S(link)
			qw422016.N().S(`</code></li>
            `)
		}
		qw422016.N().S(`
        </ul>
        `)
	}
	qw422016.N().S(`

        <h2>Banned phrases</h2>
        `)
	if len(p.Channel.FilterBannedPhrasesPatterns) == 0 {
		qw422016.N().S(`
        <p>No banned phrases.</p>
        `)
	} else {
		qw422016.N().S(`
        <p>The following phrases are banned:</p>
        <ul>
            `)
		for _, p := range p.Channel.FilterBannedPhrasesPatterns {
			qw422016.N().S(`
            <li><code>`)
			qw422016.E().S(p)
			qw422016.N().S(`</code></li>
            `)
		}
		qw422016.N().S(`
        </ul>
        `)
	}
	qw422016.N().S(`
    </div>
</div>
`)
}

func (p *ChannelRulesPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelRulesPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type ChannelScheduledPage struct {
	ChannelPage
	Repeated  models.RepeatedCommandSlice
	Scheduled models.ScheduledCommandSlice
}

func streamyesNo(qw422016 *qt422016.Writer, b bool) {
	qw422016.N().S(`
`)
	if b {
		qw422016.N().S(`
<span class="has-text-success">Yes</span>
`)
	} else {
		qw422016.N().S(`
<span class="has-text-danger">No</span>
`)
	}
	qw422016.N().S(`
`)
}

func writeyesNo(qq422016 qtio422016.Writer, b bool) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamyesNo(qw422016, b)
	qt422016.ReleaseWriter(qw422016)
}

func yesNo(b bool) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeyesNo(qb422016, b)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ChannelScheduledPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamSidebarMobile(qw422016, "scheduled")
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    `)
	p.StreamSidebar(qw422016, "scheduled")
	qw422016.N().S(`

    <div class="column is-main-content content">
        <span class="title is-1">`)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`</span><span class="subtitle is-3">Repeated / scheduled commands</span>
        <hr>

        <h2>Repeated commands</h2>

        `)
	if len(p.Repeated) == 0 {
		qw422016.N().S(`
        <p>No repeated commands.</p>
        `)
	} else {
		qw422016.N().S(`
        <table
            class="table is-striped is-hoverable is-fullwidth"
            data-toggle="table"
            data-sort-class="table-active"
            data-sort-name="enabled"
            data-sort-order="desc"
            data-search="true"
            data-sortable="true"
        >
            <thead>
                <tr>
                    <th data-sortable="true">Command</th>
                    <th data-sortable="true" data-field="enabled">Enabled</th>
                    <th data-sortable="true">Interval</th>
                    <th data-sortable="true">Message diff</th>
                </tr>
            </thead>
            <tbody>
                `)
		for _, c := range p.Repeated {
			qw422016.N().S(`
                <tr>
                    <td><code>`)
			qw422016.E().S(p.Channel.Prefix)
			qw422016.E().S(c.R.CommandInfo.Name)
			qw422016.N().S(`</code></td>
                    <td>`)
			streamyesNo(qw422016, c.Enabled)
			qw422016.N().S(`</td>
                    <td>Every `)
			qw422016.E().S(durafmt.Parse(time.Duration(c.Delay) * time.Second).String())
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().V(c.MessageDiff)
			qw422016.N().S(`</td>
                </tr>
                `)
		}
		qw422016.N().S(`
            </tbody>
        </table>
        `)
	}
	qw422016.N().S(`

        <h2>Scheduled commands</h2>

        `)
	if len(p.Scheduled) == 0 {
		qw422016.N().S(`
        <p>No scheduled commands.</p>
        `)
	} else {
		qw422016.N().S(`
        <table
            class="table is-striped is-hoverable is-fullwidth"
            data-toggle="table"
            data-sort-class="table-active"
            data-sort-name="enabled"
            data-sort-order="desc"
            data-search="true"
            data-sortable="true"
        >
            <thead>
                <tr>
                    <th data-sortable="true">Command</th>
                    <th data-sortable="true" data-field="enabled">Enabled</th>
                    <th data-sortable="true">Cron expression</th>
                    <th data-sortable="true">Message diff</th>
                </tr>
            </thead>
            <tbody>
                `)
		for _, c := range p.Scheduled {
			qw422016.N().S(`
                <tr>
                    <td><code>`)
			qw422016.E().S(p.Channel.Prefix)
			qw422016.E().S(c.R.CommandInfo.Name)
			qw422016.N().S(`</code></td>
                    <td>`)
			streamyesNo(qw422016, c.Enabled)
			qw422016.N().S(`</td>
                    <td><code>`)
			qw422016.E().S(c.CronExpression)
			qw422016.N().S(`</code></td>
                    <td>`)
			qw422016.E().V(c.MessageDiff)
			qw422016.N().S(`</td>
                </tr>
                `)
		}
		qw422016.N().S(`
            </tbody>
        </table>
        `)
	}
	qw422016.N().S(`
    </div>
</div>
`)
}

func (p *ChannelScheduledPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelScheduledPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type ChannelVariablesPage struct {
	ChannelPage
	Variables models.VariableSlice
}

func (p *ChannelVariablesPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamSidebarMobile(qw422016, "variables")
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    `)
	p.StreamSidebar(qw422016, "variables")
	qw422016.N().S(`

    <div class="column is-main-content content">
        <span class="title is-1">`)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`</span><span class="subtitle is-3">Variables</span>
        <hr>

        `)
	if len(p.Variables) == 0 {
		qw422016.N().S(`
        <p>No variables.</p>
        `)
	} else {
		qw422016.N().S(`
        <table
            class="table is-striped is-hoverable is-fullwidth"
            data-toggle="table"
            data-sort-class="table-active"
            data-sort-name="name"
            data-sort-order="asc"
            data-search="true"
            data-sortable="true"
        >
            <thead>
                <tr>
                    <th data-sortable="true" data-field="name">Name</th>
                    <th data-sortable="true">Value</th>
                </tr>
            </thead>
            <tbody>
                `)
		for _, v := range p.Variables {
			qw422016.N().S(`
                <tr>
                    <td>`)
			qw422016.E().S(v.Name)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(v.Value)
			qw422016.N().S(`</td>
                </tr>
                `)
		}
		qw422016.N().S(`
            </tbody>
        </table>
        `)
	}
	qw422016.N().S(`
    </div>
</div>
`)
}

func (p *ChannelVariablesPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelVariablesPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type ChannelHighlightsPage struct {
	ChannelPage
	Highlights models.HighlightSlice
}

func (p *ChannelHighlightsPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamSidebarMobile(qw422016, "highlights")
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    `)
	p.StreamSidebar(qw422016, "highlights")
	qw422016.N().S(`

    <div class="column is-main-content content">
        <span class="title is-1">`)
	qw422016.E().S(displayNameFor(p.Channel))
	qw422016.N().S(`</span><span class="subtitle is-3">Highlights</span>
        <hr>

        `)
	if len(p.Highlights) == 0 {
		qw422016.N().S(`
        <p>No highlights.</p>
        `)
	} else {
		qw422016.N().S(`
        <table
            class="table is-striped is-hoverable is-fullwidth"
            data-toggle="table"
            data-sort-class="table-active"
            data-sort-name="created_at"
            data-sort-order="desc"
            data-search="true"
            data-sortable="true"
        >
            <thead>
                <tr>
                    <th data-sortable="true" data-field="created_at" data-formatter="timeFormatter" data-sorter="timeSorter">Created at</th>
                    <th data-sortable="true">Timestamp</th>
                    <th data-sortable="true">Status</th>
                    <th data-sortable="true">Game</th>
                </tr>
            </thead>
            <tbody>
                `)
		for _, h := range p.Highlights {
			qw422016.N().S(`
                <tr>
                    <td>`)
			qw422016.E().S(h.HighlightedAt.Format(time.RFC3339))
			qw422016.N().S(`</td>
                    <td>`)
			if h.StartedAt.Valid {
				qw422016.E().S(durafmt.Parse(h.HighlightedAt.Sub(h.StartedAt.Time).Truncate(time.Second)).String())
			}
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(h.Status)
			qw422016.N().S(`</td>
                    <td>`)
			qw422016.E().S(h.Game)
			qw422016.N().S(`</td>
                </tr>
                `)
		}
		qw422016.N().S(`
            </tbody>
        </table>
        `)
	}
	qw422016.N().S(`
    </div>
</div>
`)
}

func (p *ChannelHighlightsPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ChannelHighlightsPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
