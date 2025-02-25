package grob

// Code generated by go-plotly/generator. DO NOT EDIT.

var TraceTypePie TraceType = "pie"

func (trace *Pie) GetType() TraceType {
	return TraceTypePie
}

// Pie A data visualized by the sectors of the pie is set in `values`. The sector labels are set in `labels`. The sector colors are set in `marker.colors`
type Pie struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Automargin
	// arrayOK: false
	// type: boolean
	// Determines whether outside text labels can push the margins.
	Automargin Bool `json:"automargin,omitempty"`

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Direction
	// default: counterclockwise
	// type: enumerated
	// Specifies the direction at which succeeding sectors follow one another.
	Direction PieDirection `json:"direction,omitempty"`

	// Dlabel
	// arrayOK: false
	// type: number
	// Sets the label step. See `label0` for more info.
	Dlabel float64 `json:"dlabel,omitempty"`

	// Domain
	// role: Object
	Domain *PieDomain `json:"domain,omitempty"`

	// Hole
	// arrayOK: false
	// type: number
	// Sets the fraction of the radius to cut out of the pie. Use this to make a donut chart.
	Hole float64 `json:"hole,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo PieHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *PieHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `label`, `color`, `value`, `percent` and `text`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate String `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext
	// arrayOK: true
	// type: string
	// Sets hover text elements associated with each sector. If a single string, the same string appears for all data points. If an array of string, the items are mapped in order of this trace's sectors. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext String `json:"hovertext,omitempty"`

	// Hovertextsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids
	// arrayOK: false
	// type: data_array
	// Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Insidetextfont
	// role: Object
	Insidetextfont *PieInsidetextfont `json:"insidetextfont,omitempty"`

	// Insidetextorientation
	// default: auto
	// type: enumerated
	// Controls the orientation of the text inside chart sectors. When set to *auto*, text may be oriented in any direction in order to be as big as possible in the middle of a sector. The *horizontal* option orients text to be parallel with the bottom of the chart, and may make text smaller in order to achieve that goal. The *radial* option orients text along the radius of the sector. The *tangential* option orients text perpendicular to the radius of the sector.
	Insidetextorientation PieInsidetextorientation `json:"insidetextorientation,omitempty"`

	// Label0
	// arrayOK: false
	// type: number
	// Alternate to `labels`. Builds a numeric set of labels. Use with `dlabel` where `label0` is the starting label and `dlabel` the step.
	Label0 float64 `json:"label0,omitempty"`

	// Labels
	// arrayOK: false
	// type: data_array
	// Sets the sector labels. If `labels` entries are duplicated, we sum associated `values` or simply count occurrences if `values` is not provided. For other array attributes (including color) we use the first non-empty entry among all occurrences of the label.
	Labels interface{} `json:"labels,omitempty"`

	// Labelssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  labels .
	Labelssrc String `json:"labelssrc,omitempty"`

	// Legendgroup
	// arrayOK: false
	// type: string
	// Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker
	// role: Object
	Marker *PieMarker `json:"marker,omitempty"`

	// Meta
	// arrayOK: true
	// type: any
	// Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name
	// arrayOK: false
	// type: string
	// Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Outsidetextfont
	// role: Object
	Outsidetextfont *PieOutsidetextfont `json:"outsidetextfont,omitempty"`

	// Pull
	// arrayOK: true
	// type: number
	// Sets the fraction of larger radius to pull the sectors out from the center. This can be a constant to pull all slices apart from each other equally or an array to highlight one or more slices.
	Pull float64 `json:"pull,omitempty"`

	// Pullsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  pull .
	Pullsrc String `json:"pullsrc,omitempty"`

	// Rotation
	// arrayOK: false
	// type: number
	// Instead of the first slice starting at 12 o'clock, rotate to some other angle.
	Rotation float64 `json:"rotation,omitempty"`

	// Scalegroup
	// arrayOK: false
	// type: string
	// If there are multiple pie charts that should be sized according to their totals, link them by providing a non-empty group id here shared by every trace in the same group.
	Scalegroup String `json:"scalegroup,omitempty"`

	// Showlegend
	// arrayOK: false
	// type: boolean
	// Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Sort
	// arrayOK: false
	// type: boolean
	// Determines whether or not the sectors are reordered from largest to smallest.
	Sort Bool `json:"sort,omitempty"`

	// Stream
	// role: Object
	Stream *PieStream `json:"stream,omitempty"`

	// Text
	// arrayOK: false
	// type: data_array
	// Sets text elements associated with each sector. If trace `textinfo` contains a *text* flag, these elements will be seen on the chart. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont
	// role: Object
	Textfont *PieTextfont `json:"textfont,omitempty"`

	// Textinfo
	// default: %!s(<nil>)
	// type: flaglist
	// Determines which trace information appear on the graph.
	Textinfo PieTextinfo `json:"textinfo,omitempty"`

	// Textposition
	// default: auto
	// type: enumerated
	// Specifies the location of the `textinfo`.
	Textposition PieTextposition `json:"textposition,omitempty"`

	// Textpositionsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `label`, `color`, `value`, `percent` and `text`.
	Texttemplate String `json:"texttemplate,omitempty"`

	// Texttemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Title
	// role: Object
	Title *PieTitle `json:"title,omitempty"`

	// Transforms
	// It's an items array and what goes inside it's... messy... check the docs
	// I will be happy if you want to contribute by implementing this
	// just raise an issue before you start so we do not overlap
	Transforms interface{} `json:"transforms,omitempty"`

	// Uid
	// arrayOK: false
	// type: string
	// Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision
	// arrayOK: false
	// type: any
	// Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Values
	// arrayOK: false
	// type: data_array
	// Sets the values of the sectors. If omitted, we count occurrences of each label.
	Values interface{} `json:"values,omitempty"`

	// Valuessrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  values .
	Valuessrc String `json:"valuessrc,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible PieVisible `json:"visible,omitempty"`
}

// PieDomain
type PieDomain struct {

	// Column
	// arrayOK: false
	// type: integer
	// If there is a layout grid, use the domain for this column in the grid for this pie trace .
	Column int64 `json:"column,omitempty"`

	// Row
	// arrayOK: false
	// type: integer
	// If there is a layout grid, use the domain for this row in the grid for this pie trace .
	Row int64 `json:"row,omitempty"`

	// X
	// arrayOK: false
	// type: info_array
	// Sets the horizontal domain of this pie trace (in plot fraction).
	X interface{} `json:"x,omitempty"`

	// Y
	// arrayOK: false
	// type: info_array
	// Sets the vertical domain of this pie trace (in plot fraction).
	Y interface{} `json:"y,omitempty"`
}

// PieHoverlabelFont Sets the font used in hover labels.
type PieHoverlabelFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family String `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  family .
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size float64 `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  size .
	Sizesrc String `json:"sizesrc,omitempty"`
}

// PieHoverlabel
type PieHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align PieHoverlabelAlign `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  align .
	Alignsrc String `json:"alignsrc,omitempty"`

	// Bgcolor
	// arrayOK: true
	// type: color
	// Sets the background color of the hover labels for this trace
	Bgcolor Color `json:"bgcolor,omitempty"`

	// Bgcolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  bgcolor .
	Bgcolorsrc String `json:"bgcolorsrc,omitempty"`

	// Bordercolor
	// arrayOK: true
	// type: color
	// Sets the border color of the hover labels for this trace.
	Bordercolor Color `json:"bordercolor,omitempty"`

	// Bordercolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  bordercolor .
	Bordercolorsrc String `json:"bordercolorsrc,omitempty"`

	// Font
	// role: Object
	Font *PieHoverlabelFont `json:"font,omitempty"`

	// Namelength
	// arrayOK: true
	// type: integer
	// Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis.
	Namelength int64 `json:"namelength,omitempty"`

	// Namelengthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  namelength .
	Namelengthsrc String `json:"namelengthsrc,omitempty"`
}

// PieInsidetextfont Sets the font used for `textinfo` lying inside the sector.
type PieInsidetextfont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family String `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  family .
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size float64 `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  size .
	Sizesrc String `json:"sizesrc,omitempty"`
}

// PieMarkerLine
type PieMarkerLine struct {

	// Color
	// arrayOK: true
	// type: color
	// Sets the color of the line enclosing each sector.
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Width
	// arrayOK: true
	// type: number
	// Sets the width (in px) of the line enclosing each sector.
	Width float64 `json:"width,omitempty"`

	// Widthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  width .
	Widthsrc String `json:"widthsrc,omitempty"`
}

// PieMarker
type PieMarker struct {

	// Colors
	// arrayOK: false
	// type: data_array
	// Sets the color of each sector. If not specified, the default trace color set is used to pick the sector colors.
	Colors interface{} `json:"colors,omitempty"`

	// Colorssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  colors .
	Colorssrc String `json:"colorssrc,omitempty"`

	// Line
	// role: Object
	Line *PieMarkerLine `json:"line,omitempty"`
}

// PieOutsidetextfont Sets the font used for `textinfo` lying outside the sector.
type PieOutsidetextfont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family String `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  family .
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size float64 `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  size .
	Sizesrc String `json:"sizesrc,omitempty"`
}

// PieStream
type PieStream struct {

	// Maxpoints
	// arrayOK: false
	// type: number
	// Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot.
	Maxpoints float64 `json:"maxpoints,omitempty"`

	// Token
	// arrayOK: false
	// type: string
	// The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details.
	Token String `json:"token,omitempty"`
}

// PieTextfont Sets the font used for `textinfo`.
type PieTextfont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family String `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  family .
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size float64 `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  size .
	Sizesrc String `json:"sizesrc,omitempty"`
}

// PieTitleFont Sets the font used for `title`. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type PieTitleFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family String `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  family .
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size float64 `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  size .
	Sizesrc String `json:"sizesrc,omitempty"`
}

// PieTitle
type PieTitle struct {

	// Font
	// role: Object
	Font *PieTitleFont `json:"font,omitempty"`

	// Position
	// default: %!s(<nil>)
	// type: enumerated
	// Specifies the location of the `title`. Note that the title's position used to be set by the now deprecated `titleposition` attribute.
	Position PieTitlePosition `json:"position,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the chart. If it is empty, no title is displayed. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// PieDirection Specifies the direction at which succeeding sectors follow one another.
type PieDirection string

const (
	PieDirectionClockwise        PieDirection = "clockwise"
	PieDirectionCounterclockwise PieDirection = "counterclockwise"
)

// PieHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type PieHoverlabelAlign string

const (
	PieHoverlabelAlignLeft  PieHoverlabelAlign = "left"
	PieHoverlabelAlignRight PieHoverlabelAlign = "right"
	PieHoverlabelAlignAuto  PieHoverlabelAlign = "auto"
)

// PieInsidetextorientation Controls the orientation of the text inside chart sectors. When set to *auto*, text may be oriented in any direction in order to be as big as possible in the middle of a sector. The *horizontal* option orients text to be parallel with the bottom of the chart, and may make text smaller in order to achieve that goal. The *radial* option orients text along the radius of the sector. The *tangential* option orients text perpendicular to the radius of the sector.
type PieInsidetextorientation string

const (
	PieInsidetextorientationHorizontal PieInsidetextorientation = "horizontal"
	PieInsidetextorientationRadial     PieInsidetextorientation = "radial"
	PieInsidetextorientationTangential PieInsidetextorientation = "tangential"
	PieInsidetextorientationAuto       PieInsidetextorientation = "auto"
)

// PieTextposition Specifies the location of the `textinfo`.
type PieTextposition string

const (
	PieTextpositionInside  PieTextposition = "inside"
	PieTextpositionOutside PieTextposition = "outside"
	PieTextpositionAuto    PieTextposition = "auto"
	PieTextpositionNone    PieTextposition = "none"
)

// PieTitlePosition Specifies the location of the `title`. Note that the title's position used to be set by the now deprecated `titleposition` attribute.
type PieTitlePosition string

const (
	PieTitlePositionTopLeft      PieTitlePosition = "top left"
	PieTitlePositionTopCenter    PieTitlePosition = "top center"
	PieTitlePositionTopRight     PieTitlePosition = "top right"
	PieTitlePositionMiddleCenter PieTitlePosition = "middle center"
	PieTitlePositionBottomLeft   PieTitlePosition = "bottom left"
	PieTitlePositionBottomCenter PieTitlePosition = "bottom center"
	PieTitlePositionBottomRight  PieTitlePosition = "bottom right"
)

// PieVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type PieVisible interface{}

var (
	PieVisibleTrue       PieVisible = true
	PieVisibleFalse      PieVisible = false
	PieVisibleLegendonly PieVisible = "legendonly"
)

// PieHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type PieHoverinfo string

const (
	// Flags
	PieHoverinfoLabel   PieHoverinfo = "label"
	PieHoverinfoText    PieHoverinfo = "text"
	PieHoverinfoValue   PieHoverinfo = "value"
	PieHoverinfoPercent PieHoverinfo = "percent"
	PieHoverinfoName    PieHoverinfo = "name"

	// Extra
	PieHoverinfoAll  PieHoverinfo = "all"
	PieHoverinfoNone PieHoverinfo = "none"
	PieHoverinfoSkip PieHoverinfo = "skip"
)

// PieTextinfo Determines which trace information appear on the graph.
type PieTextinfo string

const (
	// Flags
	PieTextinfoLabel   PieTextinfo = "label"
	PieTextinfoText    PieTextinfo = "text"
	PieTextinfoValue   PieTextinfo = "value"
	PieTextinfoPercent PieTextinfo = "percent"

	// Extra
	PieTextinfoNone PieTextinfo = "none"
)
