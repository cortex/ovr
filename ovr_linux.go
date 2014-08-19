package ovr

// Used to configure slave GL rendering (i.e. for devices created externally).
type GLConfigData struct {
	Header RenderAPIConfigHeader
	Disp   *C._XDisplay
	Win    C.Window
}
