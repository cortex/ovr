package ovr

/*
#include <OVR_CAPI_GL.h>
*/
import "C"

// Used to configure slave GL rendering (i.e. for devices created externally).
type GLConfigData struct {
	Header RenderAPIConfigHeader
}

func (configData GLConfigData) toC() C.ovrGLConfigData {
	return C.ovrGLConfigData{Header: configData.Header.toC()}
}
