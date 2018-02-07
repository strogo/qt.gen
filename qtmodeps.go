package main

// auto generated by genmodeps.sh, do not edit.

// for version: 5.10.0

var modDepsAll = map[string][]string{

	"3danimation":            []string{"network", "concurrent", "core", "gui", "3dcore", "3drender"},
	"3DAnimation":            []string{"Network", "Concurrent", "Core", "Gui", "3DCore", "3DRender"},
	"3dcore":                 []string{"core", "network", "gui"},
	"3DCore":                 []string{"Core", "Network", "Gui"},
	"3dextras":               []string{"network", "concurrent", "core", "gui", "3dcore", "3dlogic", "3dinput", "3drender"},
	"3DExtras":               []string{"Network", "Concurrent", "Core", "Gui", "3DCore", "3DLogic", "3DInput", "3DRender"},
	"3dinput":                []string{"network", "core", "gui", "3dcore"},
	"3DInput":                []string{"Network", "Core", "Gui", "3DCore"},
	"3dlogic":                []string{"network", "gui", "core", "3dcore"},
	"3DLogic":                []string{"Network", "Gui", "Core", "3DCore"},
	"3dquickanimation":       []string{"network", "concurrent", "gui", "core", "qml", "3dcore", "3drender", "3danimation"},
	"3DQuickAnimation":       []string{"Network", "Concurrent", "Gui", "Core", "Qml", "3DCore", "3DRender", "3DAnimation"},
	"3dquickextras":          []string{"network", "concurrent", "quick", "core", "qml", "gui", "3dcore", "3dlogic", "3drender", "3dquick", "3dinput", "3dextras"},
	"3DQuickExtras":          []string{"Network", "Concurrent", "Quick", "Core", "Qml", "Gui", "3DCore", "3DLogic", "3DRender", "3DQuick", "3DInput", "3DExtras"},
	"3dquickinput":           []string{"network", "gui", "core", "qml", "3dcore", "3dinput"},
	"3DQuickInput":           []string{"Network", "Gui", "Core", "Qml", "3DCore", "3DInput"},
	"3dquickrender":          []string{"network", "concurrent", "gui", "core", "qml", "3dcore", "3drender"},
	"3DQuickRender":          []string{"Network", "Concurrent", "Gui", "Core", "Qml", "3DCore", "3DRender"},
	"3dquickscene2d":         []string{"network", "concurrent", "core", "qml", "gui", "quick", "3dcore", "3drender"},
	"3DQuickScene2D":         []string{"Network", "Concurrent", "Core", "Qml", "Gui", "Quick", "3DCore", "3DRender"},
	"3dquick":                []string{"network", "core", "qml", "gui", "quick", "3dcore"},
	"3DQuick":                []string{"Network", "Core", "Qml", "Gui", "Quick", "3DCore"},
	"3drender":               []string{"network", "concurrent", "core", "gui", "3dcore"},
	"3DRender":               []string{"Network", "Concurrent", "Core", "Gui", "3DCore"},
	"bluetooth":              []string{"dbus", "core"},
	"Bluetooth":              []string{"DBus", "Core"},
	"charts":                 []string{"core", "gui", "widgets"},
	"Charts":                 []string{"Core", "Gui", "Widgets"},
	"concurrent":             []string{"core"},
	"Concurrent":             []string{"Core"},
	"core":                   []string{},
	"Core":                   []string{},
	"datavisualization":      []string{"core", "gui"},
	"DataVisualization":      []string{"Core", "Gui"},
	"dbus":                   []string{"core"},
	"DBus":                   []string{"Core"},
	"declarative":            []string{"core", "network", "xmlpatterns", "sql", "script", "gui", "widgets"},
	"Declarative":            []string{"Core", "Network", "XmlPatterns", "Sql", "Script", "Gui", "Widgets"},
	"designercomponents":     []string{"core", "xml", "gui", "widgets", "designer"},
	"DesignerComponents":     []string{"Core", "Xml", "Gui", "Widgets", "Designer"},
	"designer":               []string{"core", "xml", "gui", "widgets"},
	"Designer":               []string{"Core", "Xml", "Gui", "Widgets"},
	"eglfsdeviceintegration": []string{"core", "gui", "dbus"},
	"EglFSDeviceIntegration": []string{"Core", "Gui", "DBus"},
	"eglfskmssupport":        []string{"dbus", "core", "gui", "eglfsdeviceintegration"},
	"EglFsKmsSupport":        []string{"DBus", "Core", "Gui", "EglFSDeviceIntegration"},
	"gamepad":                []string{"core", "gui"},
	"Gamepad":                []string{"Core", "Gui"},
	"glib-2":                 []string{"core"},
	"GLib-2":                 []string{"Core"},
	"gstreamer-1":            []string{"core", "glib-2"},
	"GStreamer-1":            []string{"Core", "GLib-2"},
	"gstreamerquick-1":       []string{"network", "qml", "core", "gui", "glib-2", "quick", "gstreamer-1"},
	"GStreamerQuick-1":       []string{"Network", "Qml", "Core", "Gui", "GLib-2", "Quick", "GStreamer-1"},
	"gstreamerui-1":          []string{"core", "gui", "widgets", "glib-2", "opengl", "gstreamer-1"},
	"GStreamerUi-1":          []string{"Core", "Gui", "Widgets", "GLib-2", "OpenGL", "GStreamer-1"},
	"gstreamerutils-1":       []string{"core", "glib-2", "gstreamer-1"},
	"GStreamerUtils-1":       []string{"Core", "GLib-2", "GStreamer-1"},
	"gui":                    []string{"core"},
	"Gui":                    []string{"Core"},
	"help":                   []string{"core", "sql", "gui", "widgets"},
	"Help":                   []string{"Core", "Sql", "Gui", "Widgets"},
	"inline":                 []string{"quicktemplates2", "core", "network", "qml", "gui", "quick", "widgets", "quickwidgets", "quickcontrols2"},
	"Inline":                 []string{"QuickTemplates2", "Core", "Network", "Qml", "Gui", "Quick", "Widgets", "QuickWidgets", "QuickControls2"},
	"location":               []string{"network", "core", "qml", "gui", "quick", "positioning"},
	"Location":               []string{"Network", "Core", "Qml", "Gui", "Quick", "Positioning"},
	"multimediagsttools":     []string{"network", "opengl", "core", "gui", "widgets", "multimedia", "multimediawidgets"},
	"MultimediaGstTools":     []string{"Network", "OpenGL", "Core", "Gui", "Widgets", "Multimedia", "MultimediaWidgets"},
	"multimediaquick":        []string{"qml", "network", "core", "gui", "quick", "multimedia"},
	"MultimediaQuick":        []string{"Qml", "Network", "Core", "Gui", "Quick", "Multimedia"},
	"multimedia":             []string{"core", "gui", "network"},
	"Multimedia":             []string{"Core", "Gui", "Network"},
	"multimediawidgets":      []string{"network", "opengl", "core", "gui", "widgets", "multimedia"},
	"MultimediaWidgets":      []string{"Network", "OpenGL", "Core", "Gui", "Widgets", "Multimedia"},
	"networkauth":            []string{"core", "network"},
	"NetworkAuth":            []string{"Core", "Network"},
	"network":                []string{"core"},
	"Network":                []string{"Core"},
	"nfc":                    []string{"dbus", "core"},
	"Nfc":                    []string{"DBus", "Core"},
	"opengl":                 []string{"core", "gui", "widgets"},
	"OpenGL":                 []string{"Core", "Gui", "Widgets"},
	"positioning":            []string{"core"},
	"Positioning":            []string{"Core"},
	"printsupport":           []string{"core", "gui", "widgets"},
	"PrintSupport":           []string{"Core", "Gui", "Widgets"},
	"qml":                    []string{"core", "network"},
	"Qml":                    []string{"Core", "Network"},
	"quickcontrols2":         []string{"network", "quicktemplates2", "core", "qml", "gui", "quick"},
	"QuickControls2":         []string{"Network", "QuickTemplates2", "Core", "Qml", "Gui", "Quick"},
	"quickparticles":         []string{"network", "core", "gui", "qml", "quick"},
	"QuickParticles":         []string{"Network", "Core", "Gui", "Qml", "Quick"},
	"quick":                  []string{"core", "network", "gui", "qml"},
	"Quick":                  []string{"Core", "Network", "Gui", "Qml"},
	"quicktemplates2":        []string{"network", "core", "qml", "gui", "quick"},
	"QuickTemplates2":        []string{"Network", "Core", "Qml", "Gui", "Quick"},
	"quicktest":              []string{"network", "qml", "quick", "core", "gui", "widgets", "test"},
	"QuickTest":              []string{"Network", "Qml", "Quick", "Core", "Gui", "Widgets", "Test"},
	"quickwidgets":           []string{"network", "core", "gui", "widgets", "qml", "quick"},
	"QuickWidgets":           []string{"Network", "Core", "Gui", "Widgets", "Qml", "Quick"},
	"remoteobjects":          []string{"core", "network"},
	"RemoteObjects":          []string{"Core", "Network"},
	"script":                 []string{"core"},
	"Script":                 []string{"Core"},
	"scripttools":            []string{"gui", "widgets", "script", "core"},
	"ScriptTools":            []string{"Gui", "Widgets", "Script", "Core"},
	"scxml":                  []string{"network", "core", "qml"},
	"Scxml":                  []string{"Network", "Core", "Qml"},
	"sensors":                []string{"core"},
	"Sensors":                []string{"Core"},
	"serialbus":              []string{"serialport", "network", "core"},
	"SerialBus":              []string{"SerialPort", "Network", "Core"},
	"serialport":             []string{"core"},
	"SerialPort":             []string{"Core"},
	"sql":                    []string{"core"},
	"Sql":                    []string{"Core"},
	"svg":                    []string{"core", "gui", "widgets"},
	"Svg":                    []string{"Core", "Gui", "Widgets"},
	"test":                   []string{"core"},
	"Test":                   []string{"Core"},
	"texttospeech":           []string{"core"},
	"TextToSpeech":           []string{"Core"},
	"waylandclient":          []string{"dbus", "core", "gui"},
	"WaylandClient":          []string{"DBus", "Core", "Gui"},
	"waylandcompositor":      []string{"network", "core", "qml", "gui", "quick"},
	"WaylandCompositor":      []string{"Network", "Core", "Qml", "Gui", "Quick"},
	"webchannel":             []string{"network", "core", "qml"},
	"WebChannel":             []string{"Network", "Core", "Qml"},
	"webenginecore":          []string{"qml", "positioning", "webchannel", "core", "network", "gui", "quick"},
	"WebEngineCore":          []string{"Qml", "Positioning", "WebChannel", "Core", "Network", "Gui", "Quick"},
	"webengine":              []string{"positioning", "network", "core", "qml", "webchannel", "gui", "quick", "webenginecore"},
	"WebEngine":              []string{"Positioning", "Network", "Core", "Qml", "WebChannel", "Gui", "Quick", "WebEngineCore"},
	"webenginewidgets":       []string{"qml", "positioning", "webchannel", "quickwidgets", "core", "network", "gui", "widgets", "printsupport", "quick", "webenginecore"},
	"WebEngineWidgets":       []string{"Qml", "Positioning", "WebChannel", "QuickWidgets", "Core", "Network", "Gui", "Widgets", "PrintSupport", "Quick", "WebEngineCore"},
	"webkit":                 []string{"core", "network", "qml", "gui", "sensors", "positioning", "webchannel", "quick"},
	"WebKit":                 []string{"Core", "Network", "Qml", "Gui", "Sensors", "Positioning", "WebChannel", "Quick"},
	"webkitwidgets":          []string{"qml", "sensors", "positioning", "webchannel", "quick", "core", "network", "gui", "webkit", "widgets", "printsupport"},
	"WebKitWidgets":          []string{"Qml", "Sensors", "Positioning", "WebChannel", "Quick", "Core", "Network", "Gui", "WebKit", "Widgets", "PrintSupport"},
	"websockets":             []string{"core", "network"},
	"WebSockets":             []string{"Core", "Network"},
	"webview":                []string{"positioning", "network", "webchannel", "webenginecore", "core", "qml", "gui", "quick", "webengine"},
	"WebView":                []string{"Positioning", "Network", "WebChannel", "WebEngineCore", "Core", "Qml", "Gui", "Quick", "WebEngine"},
	"widgets":                []string{"core", "gui"},
	"Widgets":                []string{"Core", "Gui"},
	"x11extras":              []string{"core", "gui"},
	"X11Extras":              []string{"Core", "Gui"},
	"xcbqpa":                 []string{"core", "dbus", "gui"},
	"XcbQpa":                 []string{"Core", "DBus", "Gui"},
	"xdgiconloader":          []string{"core", "gui"},
	"XdgIconLoader":          []string{"Core", "Gui"},
	"xdg":                    []string{"core", "gui", "widgets", "xdgiconloader", "dbus", "xml"},
	"Xdg":                    []string{"Core", "Gui", "Widgets", "XdgIconLoader", "DBus", "Xml"},
	"xmlpatterns":            []string{"core", "network"},
	"XmlPatterns":            []string{"Core", "Network"},
	"xml":                    []string{"core"},
	"Xml":                    []string{"Core"},
	"AndroidExtras":          []string{"Core"},
	"WinExtras":              []string{"Core", "Gui"},
	"MacExtras":              []string{"Core", "Gui"},
	"androidextras":          []string{"core"},
	"winextras":              []string{"core", "gui"},
	"macextras":              []string{"core", "gui"},
}