package grob

// Code generated by go-plotly/generator. DO NOT EDIT.// Config Plot config options
type Config struct {

	// Autosizable
	// arrayOK: false
	// type: boolean
	// Determines whether the graphs are plotted with respect to layout.autosize:true and infer its container size.
	Autosizable Bool `json:"autosizable,omitempty"`

	// Displaymodebar
	// default: hover
	// type: enumerated
	// Determines the mode bar display mode. If *true*, the mode bar is always visible. If *false*, the mode bar is always hidden. If *hover*, the mode bar is visible while the mouse cursor is on the graph container.
	Displaymodebar ConfigDisplaymodebar `json:"displayModeBar,omitempty"`

	// Displaylogo
	// arrayOK: false
	// type: boolean
	// Determines whether or not the plotly logo is displayed on the end of the mode bar.
	Displaylogo Bool `json:"displaylogo,omitempty"`

	// Doubleclick
	// default: reset+autosize
	// type: enumerated
	// Sets the double click interaction mode. Has an effect only in cartesian plots. If *false*, double click is disable. If *reset*, double click resets the axis ranges to their initial values. If *autosize*, double click set the axis ranges to their autorange values. If *reset+autosize*, the odd double clicks resets the axis ranges to their initial values and even double clicks set the axis ranges to their autorange values.
	Doubleclick ConfigDoubleclick `json:"doubleClick,omitempty"`

	// Doubleclickdelay
	// arrayOK: false
	// type: number
	// Sets the delay for registering a double-click in ms. This is the time interval (in ms) between first mousedown and 2nd mouseup to constitute a double-click. This setting propagates to all on-subplot double clicks (except for geo and mapbox) and on-legend double clicks.
	Doubleclickdelay float64 `json:"doubleClickDelay,omitempty"`

	// Editable
	// arrayOK: false
	// type: boolean
	// Determines whether the graph is editable or not. Sets all pieces of `edits` unless a separate `edits` config item overrides individual parts.
	Editable Bool `json:"editable,omitempty"`

	// Edits
	// role: Object
	Edits *ConfigEdits `json:"edits,omitempty"`

	// Fillframe
	// arrayOK: false
	// type: boolean
	// When `layout.autosize` is turned on, determines whether the graph fills the container (the default) or the screen (if set to *true*).
	Fillframe Bool `json:"fillFrame,omitempty"`

	// Framemargins
	// arrayOK: false
	// type: number
	// When `layout.autosize` is turned on, set the frame margins in fraction of the graph size.
	Framemargins float64 `json:"frameMargins,omitempty"`

	// Globaltransforms
	// arrayOK: false
	// type: any
	// Set global transform to be applied to all traces with no specification needed
	Globaltransforms interface{} `json:"globalTransforms,omitempty"`

	// Linktext
	// arrayOK: false
	// type: string
	// Sets the text appearing in the `showLink` link.
	Linktext String `json:"linkText,omitempty"`

	// Locale
	// arrayOK: false
	// type: string
	// Which localization should we use? Should be a string like 'en' or 'en-US'.
	Locale String `json:"locale,omitempty"`

	// Locales
	// arrayOK: false
	// type: any
	// Localization definitions Locales can be provided either here (specific to one chart) or globally by registering them as modules. Should be an object of objects {locale: {dictionary: {...}, format: {...}}} {   da: {       dictionary: {'Reset axes': 'Nulstil aksler', ...},       format: {months: [...], shortMonths: [...]}   },   ... } All parts are optional. When looking for translation or format fields, we look first for an exact match in a config locale, then in a registered module. If those fail, we strip off any regionalization ('en-US' -> 'en') and try each (config, registry) again. The final fallback for translation is untranslated (which is US English) and for formats is the base English (the only consequence being the last fallback date format %x is DD/MM/YYYY instead of MM/DD/YYYY). Currently `grouping` and `currency` are ignored for our automatic number formatting, but can be used in custom formats.
	Locales interface{} `json:"locales,omitempty"`

	// Logging
	// arrayOK: false
	// type: integer
	// Turn all console logging on or off (errors will be thrown) This should ONLY be set via Plotly.setPlotConfig Available levels: 0: no logs 1: warnings and errors, but not informational messages 2: verbose logs
	Logging int64 `json:"logging,omitempty"`

	// Mapboxaccesstoken
	// arrayOK: false
	// type: string
	// Mapbox access token (required to plot mapbox trace types) If using an Mapbox Atlas server, set this option to '' so that plotly.js won't attempt to authenticate to the public Mapbox server.
	Mapboxaccesstoken String `json:"mapboxAccessToken,omitempty"`

	// Modebarbuttons
	// arrayOK: false
	// type: any
	// Define fully custom mode bar buttons as nested array, where the outer arrays represents button groups, and the inner arrays have buttons config objects or names of default buttons See ./components/modebar/buttons.js for more info.
	Modebarbuttons interface{} `json:"modeBarButtons,omitempty"`

	// Modebarbuttonstoadd
	// arrayOK: false
	// type: any
	// Add mode bar button using config objects See ./components/modebar/buttons.js for list of arguments.
	Modebarbuttonstoadd interface{} `json:"modeBarButtonsToAdd,omitempty"`

	// Modebarbuttonstoremove
	// arrayOK: false
	// type: any
	// Remove mode bar buttons by name. See ./components/modebar/buttons.js for the list of names.
	Modebarbuttonstoremove interface{} `json:"modeBarButtonsToRemove,omitempty"`

	// Notifyonlogging
	// arrayOK: false
	// type: integer
	// Set on-graph logging (notifier) level This should ONLY be set via Plotly.setPlotConfig Available levels: 0: no on-graph logs 1: warnings and errors, but not informational messages 2: verbose logs
	Notifyonlogging int64 `json:"notifyOnLogging,omitempty"`

	// Plotglpixelratio
	// arrayOK: false
	// type: number
	// Set the pixel ratio during WebGL image export. This config option was formerly named `plot3dPixelRatio` which is now deprecated.
	Plotglpixelratio float64 `json:"plotGlPixelRatio,omitempty"`

	// Plotlyserverurl
	// arrayOK: false
	// type: string
	// When set it determines base URL for the 'Edit in Chart Studio' `showEditInChartStudio`/`showSendToCloud` mode bar button and the showLink/sendData on-graph link. To enable sending your data to Chart Studio Cloud, you need to set both `plotlyServerURL` to 'https://chart-studio.plotly.com' and also set `showSendToCloud` to true.
	Plotlyserverurl String `json:"plotlyServerURL,omitempty"`

	// Queuelength
	// arrayOK: false
	// type: integer
	// Sets the length of the undo/redo queue.
	Queuelength int64 `json:"queueLength,omitempty"`

	// Responsive
	// arrayOK: false
	// type: boolean
	// Determines whether to change the layout size when window is resized. In v2, this option will be removed and will always be true.
	Responsive Bool `json:"responsive,omitempty"`

	// Scrollzoom
	// default: gl3d+geo+mapbox
	// type: flaglist
	// Determines whether mouse wheel or two-finger scroll zooms is enable. Turned on by default for gl3d, geo and mapbox subplots (as these subplot types do not have zoombox via pan), but turned off by default for cartesian subplots. Set `scrollZoom` to *false* to disable scrolling for all subplots.
	Scrollzoom ConfigScrollzoom `json:"scrollZoom,omitempty"`

	// Senddata
	// arrayOK: false
	// type: boolean
	// If *showLink* is true, does it contain data just link to a Chart Studio Cloud file?
	Senddata Bool `json:"sendData,omitempty"`

	// Setbackground
	// arrayOK: false
	// type: any
	// Set function to add the background color (i.e. `layout.paper_color`) to a different container. This function take the graph div as first argument and the current background color as second argument. Alternatively, set to string *opaque* to ensure there is white behind it.
	Setbackground interface{} `json:"setBackground,omitempty"`

	// Showaxisdraghandles
	// arrayOK: false
	// type: boolean
	// Set to *false* to omit cartesian axis pan/zoom drag handles.
	Showaxisdraghandles Bool `json:"showAxisDragHandles,omitempty"`

	// Showaxisrangeentryboxes
	// arrayOK: false
	// type: boolean
	// Set to *false* to omit direct range entry at the pan/zoom drag points, note that `showAxisDragHandles` must be enabled to have an effect.
	Showaxisrangeentryboxes Bool `json:"showAxisRangeEntryBoxes,omitempty"`

	// Showeditinchartstudio
	// arrayOK: false
	// type: boolean
	// Same as `showSendToCloud`, but use a pencil icon instead of a floppy-disk. Note that if both `showSendToCloud` and `showEditInChartStudio` are turned, only `showEditInChartStudio` will be honored.
	Showeditinchartstudio Bool `json:"showEditInChartStudio,omitempty"`

	// Showlink
	// arrayOK: false
	// type: boolean
	// Determines whether a link to Chart Studio Cloud is displayed at the bottom right corner of resulting graphs. Use with `sendData` and `linkText`.
	Showlink Bool `json:"showLink,omitempty"`

	// Showsendtocloud
	// arrayOK: false
	// type: boolean
	// Should we include a ModeBar button, labeled "Edit in Chart Studio", that sends this chart to chart-studio.plotly.com (formerly plot.ly) or another plotly server as specified by `plotlyServerURL` for editing, export, etc? Prior to version 1.43.0 this button was included by default, now it is opt-in using this flag. Note that this button can (depending on `plotlyServerURL` being set) send your data to an external server. However that server does not persist your data until you arrive at the Chart Studio and explicitly click "Save".
	Showsendtocloud Bool `json:"showSendToCloud,omitempty"`

	// Showsources
	// arrayOK: false
	// type: any
	// Adds a source-displaying function to show sources on the resulting graphs.
	Showsources interface{} `json:"showSources,omitempty"`

	// Showtips
	// arrayOK: false
	// type: boolean
	// Determines whether or not tips are shown while interacting with the resulting graphs.
	Showtips Bool `json:"showTips,omitempty"`

	// Staticplot
	// arrayOK: false
	// type: boolean
	// Determines whether the graphs are interactive or not. If *false*, no interactivity, for export or image generation.
	Staticplot Bool `json:"staticPlot,omitempty"`

	// Toimagebuttonoptions
	// arrayOK: false
	// type: any
	// Statically override options for toImage modebar button allowed keys are format, filename, width, height, scale see ../components/modebar/buttons.js
	Toimagebuttonoptions interface{} `json:"toImageButtonOptions,omitempty"`

	// Topojsonurl
	// arrayOK: false
	// type: string
	// Set the URL to topojson used in geo charts. By default, the topojson files are fetched from cdn.plot.ly. For example, set this option to: <path-to-plotly.js>/dist/topojson/ to render geographical feature using the topojson files that ship with the plotly.js module.
	Topojsonurl String `json:"topojsonURL,omitempty"`

	// Watermark
	// arrayOK: false
	// type: boolean
	// watermark the images with the company's logo
	Watermark Bool `json:"watermark,omitempty"`
}

// ConfigEdits
type ConfigEdits struct {

	// Annotationposition
	// arrayOK: false
	// type: boolean
	// Determines if the main anchor of the annotation is editable. The main anchor corresponds to the text (if no arrow) or the arrow (which drags the whole thing leaving the arrow length & direction unchanged).
	Annotationposition Bool `json:"annotationPosition,omitempty"`

	// Annotationtail
	// arrayOK: false
	// type: boolean
	// Has only an effect for annotations with arrows. Enables changing the length and direction of the arrow.
	Annotationtail Bool `json:"annotationTail,omitempty"`

	// Annotationtext
	// arrayOK: false
	// type: boolean
	// Enables editing annotation text.
	Annotationtext Bool `json:"annotationText,omitempty"`

	// Axistitletext
	// arrayOK: false
	// type: boolean
	// Enables editing axis title text.
	Axistitletext Bool `json:"axisTitleText,omitempty"`

	// Colorbarposition
	// arrayOK: false
	// type: boolean
	// Enables moving colorbars.
	Colorbarposition Bool `json:"colorbarPosition,omitempty"`

	// Colorbartitletext
	// arrayOK: false
	// type: boolean
	// Enables editing colorbar title text.
	Colorbartitletext Bool `json:"colorbarTitleText,omitempty"`

	// Legendposition
	// arrayOK: false
	// type: boolean
	// Enables moving the legend.
	Legendposition Bool `json:"legendPosition,omitempty"`

	// Legendtext
	// arrayOK: false
	// type: boolean
	// Enables editing the trace name fields from the legend
	Legendtext Bool `json:"legendText,omitempty"`

	// Shapeposition
	// arrayOK: false
	// type: boolean
	// Enables moving shapes.
	Shapeposition Bool `json:"shapePosition,omitempty"`

	// Titletext
	// arrayOK: false
	// type: boolean
	// Enables editing the global layout title.
	Titletext Bool `json:"titleText,omitempty"`
}

// ConfigDisplaymodebar Determines the mode bar display mode. If *true*, the mode bar is always visible. If *false*, the mode bar is always hidden. If *hover*, the mode bar is visible while the mouse cursor is on the graph container.
type ConfigDisplaymodebar interface{}

var (
	ConfigDisplaymodebarHover ConfigDisplaymodebar = "hover"
	ConfigDisplaymodebarTrue  ConfigDisplaymodebar = true
	ConfigDisplaymodebarFalse ConfigDisplaymodebar = false
)

// ConfigDoubleclick Sets the double click interaction mode. Has an effect only in cartesian plots. If *false*, double click is disable. If *reset*, double click resets the axis ranges to their initial values. If *autosize*, double click set the axis ranges to their autorange values. If *reset+autosize*, the odd double clicks resets the axis ranges to their initial values and even double clicks set the axis ranges to their autorange values.
type ConfigDoubleclick interface{}

var (
	ConfigDoubleclickFalse             ConfigDoubleclick = false
	ConfigDoubleclickReset             ConfigDoubleclick = "reset"
	ConfigDoubleclickAutosize          ConfigDoubleclick = "autosize"
	ConfigDoubleclickResetPlusautosize ConfigDoubleclick = "reset+autosize"
)

// ConfigScrollzoom Determines whether mouse wheel or two-finger scroll zooms is enable. Turned on by default for gl3d, geo and mapbox subplots (as these subplot types do not have zoombox via pan), but turned off by default for cartesian subplots. Set `scrollZoom` to *false* to disable scrolling for all subplots.
type ConfigScrollzoom interface{}

var (
	// Flags
	ConfigScrollzoomCartesian ConfigScrollzoom = "cartesian"
	ConfigScrollzoomGl3d      ConfigScrollzoom = "gl3d"
	ConfigScrollzoomGeo       ConfigScrollzoom = "geo"
	ConfigScrollzoomMapbox    ConfigScrollzoom = "mapbox"

	// Extra
	ConfigScrollzoomTrue  ConfigScrollzoom = true
	ConfigScrollzoomFalse ConfigScrollzoom = false
)
