package html

import "github.com/tdewolff/parse/v2/html"

type traits uint16

const (
	rawTag traits = 1 << iota
	nonPhrasingTag
	objectTag
	omitPTag // omit p end tag if it is followed by this start tag
	keepPTag // keep p end tag if it is followed by this end tag
	booleanAttr
	caselessAttr
	urlAttr
	trimAttr
)

var tagMap = map[html.Hash]traits{
	html.A:          keepPTag,
	html.Address:    nonPhrasingTag | omitPTag,
	html.Article:    nonPhrasingTag | omitPTag,
	html.Aside:      nonPhrasingTag | omitPTag,
	html.Audio:      objectTag | keepPTag,
	html.Blockquote: nonPhrasingTag | omitPTag,
	html.Body:       nonPhrasingTag,
	html.Br:         nonPhrasingTag,
	html.Button:     objectTag,
	html.Canvas:     objectTag,
	html.Caption:    nonPhrasingTag,
	html.Col:        nonPhrasingTag,
	html.Colgroup:   nonPhrasingTag,
	html.Dd:         nonPhrasingTag,
	html.Del:        keepPTag,
	html.Details:    omitPTag,
	html.Div:        nonPhrasingTag | omitPTag,
	html.Dl:         nonPhrasingTag | omitPTag,
	html.Dt:         nonPhrasingTag,
	html.Embed:      nonPhrasingTag,
	html.Fieldset:   nonPhrasingTag | omitPTag,
	html.Figcaption: nonPhrasingTag | omitPTag,
	html.Figure:     nonPhrasingTag | omitPTag,
	html.Footer:     nonPhrasingTag | omitPTag,
	html.Form:       nonPhrasingTag | omitPTag,
	html.H1:         nonPhrasingTag | omitPTag,
	html.H2:         nonPhrasingTag | omitPTag,
	html.H3:         nonPhrasingTag | omitPTag,
	html.H4:         nonPhrasingTag | omitPTag,
	html.H5:         nonPhrasingTag | omitPTag,
	html.H6:         nonPhrasingTag | omitPTag,
	html.Head:       nonPhrasingTag,
	html.Header:     nonPhrasingTag | omitPTag,
	html.Hgroup:     nonPhrasingTag,
	html.Hr:         nonPhrasingTag | omitPTag,
	html.Html:       nonPhrasingTag,
	html.Iframe:     rawTag | objectTag,
	html.Img:        objectTag,
	html.Input:      objectTag,
	html.Ins:        keepPTag,
	html.Keygen:     objectTag,
	html.Li:         nonPhrasingTag,
	html.Main:       nonPhrasingTag | omitPTag,
	html.Map:        keepPTag,
	html.Math:       rawTag,
	html.Menu:       omitPTag,
	html.Meta:       nonPhrasingTag,
	html.Meter:      objectTag,
	html.Nav:        nonPhrasingTag | omitPTag,
	html.Noscript:   nonPhrasingTag | keepPTag,
	html.Object:     objectTag,
	html.Ol:         nonPhrasingTag | omitPTag,
	html.Output:     nonPhrasingTag,
	html.P:          nonPhrasingTag | omitPTag,
	html.Picture:    objectTag,
	html.Pre:        nonPhrasingTag | omitPTag,
	html.Progress:   objectTag,
	html.Q:          objectTag,
	html.Script:     rawTag,
	html.Section:    nonPhrasingTag | omitPTag,
	html.Select:     objectTag,
	html.Style:      rawTag | nonPhrasingTag,
	html.Svg:        rawTag | objectTag,
	html.Table:      nonPhrasingTag | omitPTag,
	html.Tbody:      nonPhrasingTag,
	html.Td:         nonPhrasingTag,
	html.Textarea:   rawTag | objectTag,
	html.Tfoot:      nonPhrasingTag,
	html.Th:         nonPhrasingTag,
	html.Thead:      nonPhrasingTag,
	html.Title:      nonPhrasingTag,
	html.Tr:         nonPhrasingTag,
	html.Ul:         nonPhrasingTag | omitPTag,
	html.Video:      objectTag | keepPTag,
}

var attrMap = map[html.Hash]traits{
	html.Accept:          caselessAttr,
	html.Accept_Charset:  caselessAttr,
	html.Action:          urlAttr,
	html.Align:           caselessAttr,
	html.Alink:           caselessAttr,
	html.Allowfullscreen: booleanAttr,
	html.Async:           booleanAttr,
	html.Autofocus:       booleanAttr,
	html.Autoplay:        booleanAttr,
	html.Axis:            caselessAttr,
	html.Background:      urlAttr,
	html.Bgcolor:         caselessAttr,
	html.Charset:         caselessAttr,
	html.Checked:         booleanAttr,
	html.Cite:            urlAttr,
	html.Class:           trimAttr,
	html.Classid:         urlAttr,
	html.Clear:           caselessAttr,
	html.Codebase:        urlAttr,
	html.Codetype:        caselessAttr,
	html.Color:           caselessAttr,
	html.Cols:            trimAttr,
	html.Colspan:         trimAttr,
	html.Compact:         booleanAttr,
	html.Controls:        booleanAttr,
	html.Data:            urlAttr,
	html.Declare:         booleanAttr,
	html.Default:         booleanAttr,
	html.DefaultChecked:  booleanAttr,
	html.DefaultMuted:    booleanAttr,
	html.DefaultSelected: booleanAttr,
	html.Defer:           booleanAttr,
	html.Dir:             caselessAttr,
	html.Disabled:        booleanAttr,
	html.Enabled:         booleanAttr,
	html.Enctype:         caselessAttr,
	html.Face:            caselessAttr,
	html.Formaction:      urlAttr,
	html.Formnovalidate:  booleanAttr,
	html.Frame:           caselessAttr,
	html.Hidden:          booleanAttr,
	html.Href:            urlAttr,
	html.Hreflang:        caselessAttr,
	html.Http_Equiv:      caselessAttr,
	html.Icon:            urlAttr,
	html.Inert:           booleanAttr,
	html.Ismap:           booleanAttr,
	html.Itemscope:       booleanAttr,
	html.Lang:            caselessAttr,
	html.Language:        caselessAttr,
	html.Link:            caselessAttr,
	html.Longdesc:        urlAttr,
	html.Manifest:        urlAttr,
	html.Maxlength:       trimAttr,
	html.Media:           caselessAttr | trimAttr,
	html.Method:          caselessAttr,
	html.Multiple:        booleanAttr,
	html.Muted:           booleanAttr,
	html.Nohref:          booleanAttr,
	html.Noresize:        booleanAttr,
	html.Noshade:         booleanAttr,
	html.Novalidate:      booleanAttr,
	html.Nowrap:          booleanAttr,
	html.Open:            booleanAttr,
	html.Pauseonexit:     booleanAttr,
	html.Poster:          urlAttr,
	html.Profile:         urlAttr,
	html.Readonly:        booleanAttr,
	html.Rel:             caselessAttr,
	html.Required:        booleanAttr,
	html.Rev:             caselessAttr,
	html.Reversed:        booleanAttr,
	html.Rows:            trimAttr,
	html.Rowspan:         trimAttr,
	html.Rules:           caselessAttr,
	html.Scope:           caselessAttr,
	html.Scoped:          booleanAttr,
	html.Scrolling:       caselessAttr,
	html.Seamless:        booleanAttr,
	html.Selected:        booleanAttr,
	html.Shape:           caselessAttr,
	html.Size:            trimAttr,
	html.Sortable:        booleanAttr,
	html.Span:            trimAttr,
	html.Src:             urlAttr,
	html.Srcset:          trimAttr,
	html.Tabindex:        trimAttr,
	html.Target:          caselessAttr,
	html.Text:            caselessAttr,
	html.Translate:       booleanAttr,
	html.Truespeed:       booleanAttr,
	html.Typemustmatch:   booleanAttr,
	html.Undeterminate:   booleanAttr,
	html.Usemap:          urlAttr,
	html.Valign:          caselessAttr,
	html.Valuetype:       caselessAttr,
	html.Vlink:           caselessAttr,
	html.Visible:         booleanAttr,
	html.Xmlns:           urlAttr,
}

var jsMimetypes = map[string]bool{
	"text/javascript":        true,
	"application/javascript": true,
}
