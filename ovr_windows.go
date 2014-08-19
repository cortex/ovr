package ovr

// Used to configure slave GL rendering (i.e. for devices created externally).
type GLConfigData struct {
	Header RenderAPIConfigHeader
	Window C.HWND
	DC     C.HDC
}

func (hmd *Hmd) AttachToWindow(hwnd syscall.Handle) bool {
	return C.char(C.ovrHmd_AttachToWindow(hmd.hmdRef, uintptr(hwnd), nil, nil)) == 1
}
