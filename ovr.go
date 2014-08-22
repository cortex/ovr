package ovr

/*
#cgo darwin CFLAGS: -DOVR_OS_MAC
#cgo darwin LDFLAGS: -framework CoreFoundation -framework IOKit -framework CoreGraphics -framework OpenGL -lc++ -lovr
#cgo windows CFLAGS: -DOVR_OS_WIN32 -I C:/libovr_0.4/dynamic
#cgo windows,386 LDFLAGS: -L C:/libovr_0.4/x86 -lovr
#cgo windows,amd64 LDFLAGS: -L C:/libovr_0.4/x64 -lovr
#include <stdlib.h>
#include <string.h>
#include <OVR_CAPI_GL.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

func ovrBool(b bool) C.ovrBool {
	if b {
		return 1
	}
	return 0
}

func newBool(b C.ovrBool) bool {
	if b == 1 {
		return true
	}
	return false
}

// ****************************************************************************
// ************************ [ Simple Math Structures ] ************************
// ****************************************************************************

// A 2D vector with integer components.
type Vector2i struct {
	X int
	Y int
}

func (vector Vector2i) toC() C.ovrVector2i {
	return C.ovrVector2i{x: C.int(vector.X), y: C.int(vector.Y)}
}

func newVector2i(vector C.ovrVector2i) Vector2i {
	return Vector2i{X: int(vector.x), Y: int(vector.y)}
}

// A 2D size with integer components.
type Sizei struct {
	W int
	H int
}

func (size Sizei) toC() C.ovrSizei {
	return C.ovrSizei{w: C.int(size.W), h: C.int(size.H)}
}

func newSizei(size C.ovrSizei) Sizei {
	return Sizei{W: int(size.w), H: int(size.h)}
}

// A 2D rectangle with a position and size.
type Recti struct {
	Pos  Vector2i
	Size Sizei
}

func (rect Recti) toC() C.ovrRecti {
	return C.ovrRecti{Pos: rect.Pos.toC(), Size: rect.Size.toC()}
}

func newRecti(rect C.ovrRecti) Recti {
	return Recti{Pos: newVector2i(rect.Pos), Size: newSizei(rect.Size)}
}

// A quaternion rotation.
type Quatf struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (quat Quatf) toC() C.ovrQuatf {
	return C.ovrQuatf{x: C.float(quat.X), y: C.float(quat.Y), z: C.float(quat.Z), w: C.float(quat.W)}
}

func newQuatf(quat C.ovrQuatf) Quatf {
	return Quatf{X: float32(quat.x), Y: float32(quat.y), Z: float32(quat.z), W: float32(quat.w)}
}

// A 2D vector with float components.
type Vector2f struct {
	X float32
	Y float32
}

func (vector Vector2f) toC() C.ovrVector2f {
	return C.ovrVector2f{x: C.float(vector.X), y: C.float(vector.Y)}
}

func newVector2f(vector C.ovrVector2f) Vector2f {
	return Vector2f{X: float32(vector.x), Y: float32(vector.y)}
}

// A 3D vector with float components.
type Vector3f struct {
	X float32
	Y float32
	Z float32
}

func (vector Vector3f) toC() C.ovrVector3f {
	return C.ovrVector3f{x: C.float(vector.X), y: C.float(vector.Y), z: C.float(vector.Z)}
}

func newVector3f(vector C.ovrVector3f) Vector3f {
	return Vector3f{X: float32(vector.x), Y: float32(vector.y), Z: float32(vector.z)}
}

// A 4x4 matrix with float elements.
type Matrix4f struct {
	M [4][4]float32
}

func (matrix Matrix4f) toC() C.ovrMatrix4f {
	M := [4][4]C.float{}

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			M[x][y] = C.float(matrix.M[x][y])
		}
	}

	return C.ovrMatrix4f{M: M}
}

func newMatrix4f(matrix C.ovrMatrix4f) Matrix4f {
	m := Matrix4f{}

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			m.M[x][y] = float32(matrix.M[x][y])
		}
	}

	return m
}

// Position and orientation together.
type Posef struct {
	Orientation Quatf
	Position    Vector3f
}

func (posef Posef) toC() C.ovrPosef {
	return C.ovrPosef{
		Orientation: posef.Orientation.toC(),
		Position:    posef.Position.toC(),
	}
}

func newPosef(posef C.ovrPosef) Posef {
	return Posef{
		Orientation: newQuatf(posef.Orientation),
		Position:    newVector3f(posef.Position),
	}
}

// A full pose (rigid body) configuration with first and second derivatives.
type PoseStatef struct {
	ThePose             Posef
	AngularVelocity     Vector3f
	LinearVelocity      Vector3f
	AngularAcceleration Vector3f
	LinearAcceleration  Vector3f
	TimeInSeconds       float64
}

func (poseState PoseStatef) toC() C.ovrPoseStatef {
	return C.ovrPoseStatef{
		ThePose:             poseState.ThePose.toC(),
		AngularVelocity:     poseState.AngularVelocity.toC(),
		LinearVelocity:      poseState.LinearVelocity.toC(),
		AngularAcceleration: poseState.AngularAcceleration.toC(),
		LinearAcceleration:  poseState.LinearAcceleration.toC(),
		TimeInSeconds:       C.double(poseState.TimeInSeconds),
	}
}

func newPoseStatef(poseStatef C.ovrPoseStatef) PoseStatef {
	return PoseStatef{
		ThePose:             newPosef(poseStatef.ThePose),
		AngularVelocity:     newVector3f(poseStatef.AngularVelocity),
		LinearVelocity:      newVector3f(poseStatef.LinearVelocity),
		AngularAcceleration: newVector3f(poseStatef.AngularAcceleration),
		LinearAcceleration:  newVector3f(poseStatef.LinearAcceleration),
		TimeInSeconds:       float64(poseStatef.TimeInSeconds),
	}
}

// Field Of View (FOV) in tangent of the angle units.
type FovPort struct {
	UpTan    float32
	DownTan  float32
	LeftTan  float32
	RightTan float32
}

func (fov FovPort) toC() C.ovrFovPort {
	return C.ovrFovPort{
		UpTan:    C.float(fov.UpTan),
		DownTan:  C.float(fov.DownTan),
		LeftTan:  C.float(fov.LeftTan),
		RightTan: C.float(fov.RightTan),
	}
}

func newFovPort(fov C.ovrFovPort) FovPort {
	return FovPort{
		UpTan:    float32(fov.UpTan),
		DownTan:  float32(fov.DownTan),
		LeftTan:  float32(fov.LeftTan),
		RightTan: float32(fov.RightTan),
	}
}

// ****************************************************************************
// ******************************* [ HMD types ] ******************************
// ****************************************************************************

// Enumerates all HMD types that we support.
const (
	Hmd_None  = C.ovrHmd_None
	Hmd_DKHD  = C.ovrHmd_DKHD
	Hmd_DK1   = C.ovrHmd_DK1
	Hmd_DK2   = C.ovrHmd_DK2
	Hmd_Other = C.ovrHmd_Other
)

type HmdType C.ovrHmdType

// HMD capability bits reported by device.
const (
	// Read-only flags.
	HmdCap_Present       = C.ovrHmdCap_Present
	HmdCap_Available     = C.ovrHmdCap_Available
	HmdCap_Captured      = C.ovrHmdCap_Captured
	HmdCap_ExtendDesktop = C.ovrHmdCap_ExtendDesktop

	// Modifiable flags.
	HmdCap_NoMirrorToWindow  = C.ovrHmdCap_NoMirrorToWindow
	HmdCap_DisplayOff        = C.ovrHmdCap_DisplayOff
	HmdCap_LowPersistence    = C.ovrHmdCap_LowPersistence
	HmdCap_DynamicPrediction = C.ovrHmdCap_DynamicPrediction
	HmdCap_NoVSync           = C.ovrHmdCap_NoVSync

	HmdCap_Writable_Mask = C.ovrHmdCap_Writable_Mask
	HmdCap_Service_Mask  = C.ovrHmdCap_Service_Mask
)

type HmdCaps C.ovrHmdCaps

// Tracking capability bits reported by the device.
const (
	TrackingCap_Orientation      = C.ovrTrackingCap_Orientation
	TrackingCap_MagYawCorrection = C.ovrTrackingCap_MagYawCorrection
	TrackingCap_Position         = C.ovrTrackingCap_Position
	TrackingCap_Idle             = C.ovrTrackingCap_Idle
)

type TrackingCaps C.ovrTrackingCaps

// Distortion capability bits reported by device.
const (
	DistortionCap_Chromatic                  = C.ovrDistortionCap_Chromatic
	DistortionCap_TimeWarp                   = C.ovrDistortionCap_TimeWarp
	DistortionCap_Vignette                   = C.ovrDistortionCap_Vignette
	DistortionCap_NoRestore                  = C.ovrDistortionCap_NoRestore
	DistortionCap_FlipInput                  = C.ovrDistortionCap_FlipInput
	DistortionCap_SRGB                       = C.ovrDistortionCap_SRGB
	DistortionCap_Overdrive                  = C.ovrDistortionCap_Overdrive
	DistortionCap_ProfileNoTimewarpSpinWaits = C.ovrDistortionCap_ProfileNoTimewarpSpinWaits
)

type DistortionCaps C.ovrDistortionCaps

// Specifies which eye is being used for rendering.
const (
	Eye_Left  = 0
	Eye_Right = 1
	Eye_Count = 2
)

type EyeType C.ovrEyeType

// This is a complete descriptor of the HMD.
type Hmd struct {
	hmdRef                     C.ovrHmd
	Type                       HmdType
	ProductName                string
	Manufacturer               string
	VendorId                   int
	ProductId                  int
	SerialNumber               string
	FirmwareMajor              int
	FirmwareMinor              int
	CameraFrustumHFovInRadians float32
	CameraFrustumVFovInRadians float32
	CameraFrustumNearZInMeters float32
	CameraFrustumFarZInMeters  float32
	HmdCaps                    uint
	TrackingCaps               uint
	DistortionCaps             uint
	DefaultEyeFov              [Eye_Count]FovPort
	MaxEyeFov                  [Eye_Count]FovPort
	EyeRenderOrder             [Eye_Count]EyeType
	Resolution                 Sizei
	WindowsPos                 Vector2i
	DisplayDeviceName          string
	DisplayId                  int
}

func newHmd(hmd C.ovrHmd) *Hmd {
	if hmd == nil {
		return nil
	}

	return &Hmd{
		hmdRef:                     hmd,
		Type:                       HmdType(hmd.Type),
		ProductName:                C.GoString(hmd.ProductName),
		Manufacturer:               C.GoString(hmd.Manufacturer),
		VendorId:                   int(hmd.VendorId),
		ProductId:                  int(hmd.ProductId),
		SerialNumber:               C.GoString(&hmd.SerialNumber[0]),
		FirmwareMajor:              int(hmd.FirmwareMajor),
		FirmwareMinor:              int(hmd.FirmwareMinor),
		CameraFrustumHFovInRadians: float32(hmd.CameraFrustumHFovInRadians),
		CameraFrustumVFovInRadians: float32(hmd.CameraFrustumVFovInRadians),
		CameraFrustumNearZInMeters: float32(hmd.CameraFrustumNearZInMeters),
		CameraFrustumFarZInMeters:  float32(hmd.CameraFrustumFarZInMeters),
		HmdCaps:                    uint(hmd.HmdCaps),
		TrackingCaps:               uint(hmd.TrackingCaps),
		DistortionCaps:             uint(hmd.DistortionCaps),
		DefaultEyeFov: [Eye_Count]FovPort{
			newFovPort(hmd.DefaultEyeFov[0]),
			newFovPort(hmd.DefaultEyeFov[1]),
		},
		MaxEyeFov: [Eye_Count]FovPort{
			newFovPort(hmd.MaxEyeFov[0]),
			newFovPort(hmd.MaxEyeFov[1]),
		},
		EyeRenderOrder: [Eye_Count]EyeType{
			EyeType(hmd.EyeRenderOrder[0]),
			EyeType(hmd.EyeRenderOrder[1]),
		},
		Resolution:        newSizei(hmd.Resolution),
		WindowsPos:        newVector2i(hmd.WindowsPos),
		DisplayDeviceName: C.GoString(hmd.DisplayDeviceName),
		DisplayId:         int(hmd.DisplayId),
	}
}

const (
	Status_OrientationTracked = C.ovrStatus_OrientationTracked
	Status_PositionTracked    = C.ovrStatus_PositionTracked
	Status_CameraPoseTracked  = C.ovrStatus_CameraPoseTracked
	Status_PositionConnected  = C.ovrStatus_PositionConnected
	Status_HmdConnected       = C.ovrStatus_HmdConnected
)

type StatusBits C.ovrStatusBits

type SensorData struct {
	Accelerometer Vector3f
	Gyro          Vector3f
	Magnetometer  Vector3f
	Temperature   float32
	TimeInSeconds float32
}

func (data SensorData) toC() C.ovrSensorData {
	return C.ovrSensorData{
		Accelerometer: data.Accelerometer.toC(),
		Gyro:          data.Gyro.toC(),
		Magnetometer:  data.Magnetometer.toC(),
		Temperature:   C.float(data.Temperature),
		TimeInSeconds: C.float(data.TimeInSeconds),
	}
}

func newSensorData(data C.ovrSensorData) SensorData {
	return SensorData{
		Accelerometer: newVector3f(data.Accelerometer),
		Gyro:          newVector3f(data.Gyro),
		Magnetometer:  newVector3f(data.Magnetometer),
		Temperature:   float32(data.Temperature),
		TimeInSeconds: float32(data.TimeInSeconds),
	}
}

type TrackingState struct {
	HeadPose          PoseStatef
	CameraPose        Posef
	LeveledCameraPose Posef
	RawSensorData     SensorData
	StatusFlags       uint
}

func (state TrackingState) toC() C.ovrTrackingState {
	return C.ovrTrackingState{
		HeadPose:          state.HeadPose.toC(),
		CameraPose:        state.CameraPose.toC(),
		LeveledCameraPose: state.LeveledCameraPose.toC(),
		RawSensorData:     state.RawSensorData.toC(),
		StatusFlags:       C.uint(state.StatusFlags),
	}
}

func newTrackingState(trackingState C.ovrTrackingState) TrackingState {
	return TrackingState{
		HeadPose:          newPoseStatef(trackingState.HeadPose),
		CameraPose:        newPosef(trackingState.CameraPose),
		LeveledCameraPose: newPosef(trackingState.LeveledCameraPose),
		RawSensorData:     newSensorData(trackingState.RawSensorData),
		StatusFlags:       uint(trackingState.StatusFlags),
	}
}

type FrameTiming C.ovrFrameTiming

type EyeRenderDesc struct {
	Eye                       EyeType
	Fov                       FovPort
	DistortedViewport         Recti
	PixelsPerTanAngleAtCenter Vector2f
	ViewAdjust                Vector3f
}

func (desc EyeRenderDesc) toC() C.ovrEyeRenderDesc {
	return C.ovrEyeRenderDesc{
		Eye:                       C.ovrEyeType(desc.Eye),
		Fov:                       desc.Fov.toC(),
		DistortedViewport:         desc.DistortedViewport.toC(),
		PixelsPerTanAngleAtCenter: desc.PixelsPerTanAngleAtCenter.toC(),
		ViewAdjust:                desc.ViewAdjust.toC(),
	}
}

func newEyeRenderDesc(eyeRenderDesc C.ovrEyeRenderDesc) EyeRenderDesc {
	return EyeRenderDesc{
		Eye:                       EyeType(eyeRenderDesc.Eye),
		Fov:                       newFovPort(eyeRenderDesc.Fov),
		DistortedViewport:         newRecti(eyeRenderDesc.DistortedViewport),
		PixelsPerTanAngleAtCenter: newVector2f(eyeRenderDesc.PixelsPerTanAngleAtCenter),
		ViewAdjust:                newVector3f(eyeRenderDesc.ViewAdjust),
	}
}

const (
	RenderAPI_None         = C.ovrRenderAPI_None
	RenderAPI_OpenGL       = C.ovrRenderAPI_OpenGL
	RenderAPI_Android_GLES = C.ovrRenderAPI_Android_GLES
	RenderAPI_D39          = C.ovrRenderAPI_D3D9
	RenderAPI_D310         = C.ovrRenderAPI_D3D10
	RenderAPI_D311         = C.ovrRenderAPI_D3D11
	RenderAPI_Count        = C.ovrRenderAPI_Count
)

type RenderAPIType C.ovrRenderAPIType

type RenderAPIConfigHeader struct {
	API         RenderAPIType
	RTSize      Sizei
	Multisample int
}

func (configHeader RenderAPIConfigHeader) toC() C.ovrRenderAPIConfigHeader {
	return C.ovrRenderAPIConfigHeader{
		API:         C.ovrRenderAPIType(configHeader.API),
		RTSize:      configHeader.RTSize.toC(),
		Multisample: C.int(configHeader.Multisample),
	}
}

type RenderAPIConfig C.ovrRenderAPIConfig

type TextureHeader struct {
	API            RenderAPIType
	TextureSize    Sizei
	RenderViewport Recti
}

func (header TextureHeader) toC() C.ovrTextureHeader {
	return C.ovrTextureHeader{
		API:            C.ovrRenderAPIType(header.API),
		TextureSize:    header.TextureSize.toC(),
		RenderViewport: header.RenderViewport.toC(),
	}
}

func newTextureHeader(header C.ovrTextureHeader) TextureHeader {
	return TextureHeader{
		API:            RenderAPIType(header.API),
		TextureSize:    newSizei(header.TextureSize),
		RenderViewport: newRecti(header.RenderViewport),
	}
}

type Texture struct {
	Header       TextureHeader
	PlatformData [8]C.uintptr_t
}

func (texture Texture) toC() C.ovrTexture {
	return C.ovrTexture{
		Header:       texture.Header.toC(),
		PlatformData: texture.PlatformData,
	}
}

func newTexture(texture C.ovrTexture) Texture {
	return Texture{
		Header:       newTextureHeader(texture.Header),
		PlatformData: texture.PlatformData,
	}
}

// ****************************************************************************
// ******************************** [ OpenGL ] ********************************
// ****************************************************************************

// Contains OpenGL-specific rendering information. In C this is a union with
// another structure, but here we just do the conversion with a function.
type GLConfig struct {
	OGL GLConfigData
}

func (config GLConfig) Config() *RenderAPIConfig {
	configData := config.OGL.toC()
	return (*RenderAPIConfig)(unsafe.Pointer(&configData))
}

// Used to pass GL eye texture data to ovrHmd_EndFrame.
type GLTextureData struct {
	Header TextureHeader
	TexId  C.GLuint
}

// Contains platform-specific information about a texture.
type GLTexture struct {
	Texture Texture
	OGL     GLTextureData
}

// ****************************************************************************
// ***************************** [ API interface ] ****************************
// ****************************************************************************

func InitializeRenderingShim() {
	C.ovr_InitializeRenderingShim()
}

func Initialize() bool {
	return C.ovr_Initialize() == 1
}

func Shutdown() {
	C.ovr_Shutdown()
}

func GetVersionString() string {
	return C.GoString(C.ovr_GetVersionString())
}

func HmdDetect() int {
	return int(C.ovrHmd_Detect())
}

func HmdCreate(index int) *Hmd {
	return newHmd(C.ovrHmd_Create(C.int(index)))
}

func HmdCreateDebug(hmdType HmdType) *Hmd {
	return newHmd(C.ovrHmd_CreateDebug(C.ovrHmdType(hmdType)))
}

func (hmd *Hmd) Destroy() {
	C.ovrHmd_Destroy(hmd.hmdRef)
}

// The ovrHmd_GetLastError function has a bug, where it sends back an empty
// string when there is no error, instead of NULL. Work around that.
func (hmd *Hmd) GetLastError() *string {
	if str := C.ovrHmd_GetLastError(hmd.hmdRef); str != nil && C.strlen(str) != 0 {
		goStr := C.GoString(str)
		return &goStr
	}

	return nil
}

func (hmd *Hmd) GetEnabledCaps() uint {
	return uint(C.ovrHmd_GetEnabledCaps(hmd.hmdRef))
}

func (hmd *Hmd) SetEnabledCaps(hmdCaps uint) {
	C.ovrHmd_SetEnabledCaps(hmd.hmdRef, C.uint(hmdCaps))
}

// ****************************************************************************
// ************************** [ Tracking interface ] **************************
// ****************************************************************************

func (hmd *Hmd) ConfigureTracking(supportedTrackingCaps uint, requiredTrackingCaps uint) bool {
	return C.ovrHmd_ConfigureTracking(hmd.hmdRef, C.uint(supportedTrackingCaps), C.uint(requiredTrackingCaps)) == 1
}

func (hmd *Hmd) RecenterPose() {
	C.ovrHmd_RecenterPose(hmd.hmdRef)
}

func (hmd *Hmd) GetTrackingState(absTime float64) TrackingState {
	return newTrackingState(C.ovrHmd_GetTrackingState(hmd.hmdRef, C.double(absTime)))
}

// ****************************************************************************
// **************************** [ Graphics setup ] ****************************
// ****************************************************************************

func (hmd *Hmd) GetFovTextureSize(eye EyeType, fov FovPort, pixelsPerDisplayPixel float32) Sizei {
	return newSizei(C.ovrHmd_GetFovTextureSize(hmd.hmdRef, C.ovrEyeType(eye), fov.toC(), C.float(pixelsPerDisplayPixel)))
}

// ****************************************************************************
// *********************** [ SDK Distortion rendering ] ***********************
// ****************************************************************************

func (hmd *Hmd) ConfigureRendering(apiConfig *RenderAPIConfig, distortionCaps uint, eyeFovIn [2]FovPort) (*[2]EyeRenderDesc, error) {
	_apiConfig := C.ovrRenderAPIConfig(*apiConfig)
	_eyeFovIn := [2]C.ovrFovPort{eyeFovIn[0].toC(), eyeFovIn[1].toC()}
	eyeRenderDescOut := [2]C.ovrEyeRenderDesc{}

	if C.ovrHmd_ConfigureRendering(hmd.hmdRef, &_apiConfig, C.uint(distortionCaps), &_eyeFovIn[0], &eyeRenderDescOut[0]) == 0 {
		if lastError := hmd.GetLastError(); lastError != nil {
			return nil, errors.New(*lastError)
		}

		return nil, errors.New("An unknown error occured")
	}

	_eyeRenderDescOut := [2]EyeRenderDesc{
		newEyeRenderDesc(eyeRenderDescOut[0]),
		newEyeRenderDesc(eyeRenderDescOut[1]),
	}

	return &_eyeRenderDescOut, nil
}

func (hmd *Hmd) BeginFrame(frameIndex uint) FrameTiming {
	return FrameTiming(C.ovrHmd_BeginFrame(hmd.hmdRef, C.uint(frameIndex)))
}

func (hmd *Hmd) EndFrame(renderPose [2]Posef, eyeTexture [2]Texture) {
	_renderPose := [2]C.ovrPosef{renderPose[0].toC(), renderPose[1].toC()}
	_eyeTexture := [2]C.ovrTexture{eyeTexture[0].toC(), eyeTexture[1].toC()}

	C.ovrHmd_EndFrame(hmd.hmdRef, &_renderPose[0], &_eyeTexture[0])
}

func (hmd *Hmd) GetEyePose(eye EyeType) Posef {
	return newPosef(C.ovrHmd_GetEyePose(hmd.hmdRef, C.ovrEyeType(eye)))
}

// ****************************************************************************
// ********************** [ Client Distortion rendering ] *********************
// ****************************************************************************

func (hmd *Hmd) GetRenderDesc(eye EyeType, fov FovPort) EyeRenderDesc {
	return newEyeRenderDesc(C.ovrHmd_GetRenderDesc(hmd.hmdRef, C.ovrEyeType(eye), fov.toC()))
}

// type DistortionVertex C.ovrDistortionVertex
type DistortionMesh C.ovrDistortionMesh

func (mesh *DistortionMesh) Destroy() {
	C.ovrHmd_DestroyDistortionMesh((*C.ovrDistortionMesh)(unsafe.Pointer(mesh)))
}

func (hmd *Hmd) CreateDistortionMesh(eye EyeType, fov FovPort, distortionCaps uint) (*DistortionMesh, error) {
	meshData := DistortionMesh{}

	if C.ovrHmd_CreateDistortionMesh(hmd.hmdRef, C.ovrEyeType(eye), fov.toC(), C.uint(distortionCaps), (*C.ovrDistortionMesh)(unsafe.Pointer(&meshData))) == 0 {
		if lastError := hmd.GetLastError(); lastError != nil {
			return nil, errors.New(*lastError)
		}

		return nil, errors.New("An unknown error occured")
	}

	return &meshData, nil
}

func (hmd *Hmd) GetRenderScaleAndOffset(fov FovPort, textureSize Sizei, renderViewport Recti) [2]Vector2f {
	uvScaleOffsetOut := [2]C.ovrVector2f{}
	C.ovrHmd_GetRenderScaleAndOffset(fov.toC(), textureSize.toC(), renderViewport.toC(), &uvScaleOffsetOut[0])

	return [2]Vector2f{newVector2f(uvScaleOffsetOut[0]), newVector2f(uvScaleOffsetOut[1])}
}

func (hmd *Hmd) GetFrameTiming(frameIndex uint) FrameTiming {
	return FrameTiming(C.ovrHmd_GetFrameTiming(hmd.hmdRef, C.uint(frameIndex)))
}

func (hmd *Hmd) BeginFrameTiming(frameIndex uint) FrameTiming {
	return FrameTiming(C.ovrHmd_BeginFrameTiming(hmd.hmdRef, C.uint(frameIndex)))
}

func (hmd *Hmd) EndFrameTiming() {
	C.ovrHmd_EndFrameTiming(hmd.hmdRef)
}

func (hmd *Hmd) ResetFrameTiming(frameIndex uint) {
	C.ovrHmd_ResetFrameTiming(hmd.hmdRef, C.uint(frameIndex))
}

func (hmd *Hmd) GetEyeTimewarpMatrices(eye EyeType, renderPose Posef) [2]Matrix4f {
	twmOut := [2]C.ovrMatrix4f{}
	C.ovrHmd_GetEyeTimewarpMatrices(hmd.hmdRef, C.ovrEyeType(eye), renderPose.toC(), &twmOut[0])

	return [2]Matrix4f{newMatrix4f(twmOut[0]), newMatrix4f(twmOut[1])}
}

// ****************************************************************************
// *********************** [ Stateless math functions ] ***********************
// ****************************************************************************

func Matrix4f_Projection(fov FovPort, znear float32, zfar float32, rightHanded bool) Matrix4f {
	return newMatrix4f(C.ovrMatrix4f_Projection(fov.toC(), C.float(znear), C.float(zfar), ovrBool(rightHanded)))
}

func Matrix4f_OrthoSubProjection(projection Matrix4f, orthoScale Vector2f, orthoDistance float32, eyeViewAdjustX float32) Matrix4f {
	return newMatrix4f(C.ovrMatrix4f_OrthoSubProjection(projection.toC(), orthoScale.toC(), C.float(orthoDistance), C.float(eyeViewAdjustX)))
}

func GetTimeInSeconds() float64 {
	return float64(C.ovr_GetTimeInSeconds())
}

func WaitTillTime(absTime float64) float64 {
	return float64(C.ovr_WaitTillTime(C.double(absTime)))
}

// ****************************************************************************
// ************************ [ Latency Test interface ] ************************
// ****************************************************************************

func (hmd *Hmd) ProcessLatencyTest() (*[3]uint, bool) {
	rgbColorOut := [3]C.uchar{}
	if C.ovrHmd_ProcessLatencyTest(hmd.hmdRef, &rgbColorOut[0]) == 0 {
		return nil, false
	}

	return &[3]uint{uint(rgbColorOut[0]), uint(rgbColorOut[1]), uint(rgbColorOut[2])}, true
}

// The ovrHmd_GetLatencyTestResult function has a bug, where it sends back an
// empty string when there is no error, instead of a NULL. Work around that.
func (hmd *Hmd) GetLatencyTestResult() *string {
	if str := C.ovrHmd_GetLatencyTestResult(hmd.hmdRef); str != nil && C.strlen(str) != 0 {
		goStr := C.GoString(str)
		return &goStr
	}

	return nil
}

// ****************************************************************************
// ************** [ Health and Safety Warning Display interface ] *************
// ****************************************************************************

type HSWDisplayState struct {
	Displayed       bool
	StartTime       float64
	DismissibleTime float64
}

func (state HSWDisplayState) toC() C.ovrHSWDisplayState {
	return C.ovrHSWDisplayState{
		Displayed:       ovrBool(state.Displayed),
		StartTime:       C.double(state.StartTime),
		DismissibleTime: C.double(state.DismissibleTime),
	}
}

func (hmd *Hmd) GetHSWDisplayState() *HSWDisplayState {
	hasWarningState := HSWDisplayState{}
	C.ovrHmd_GetHSWDisplayState(hmd.hmdRef, (*C.ovrHSWDisplayState)(unsafe.Pointer(&hasWarningState)))

	return &hasWarningState
}

func (hmd *Hmd) DismissHSWDisplay() bool {
	return C.ovrHmd_DismissHSWDisplay(hmd.hmdRef) == 1
}

// ****************************************************************************
// **************************** [ Property Access ] ***************************
// ****************************************************************************

const (
	KEY_USER                 = C.OVR_KEY_USER
	KEY_NAME                 = C.OVR_KEY_NAME
	KEY_GENDER               = C.OVR_KEY_GENDER
	KEY_PLAYER_HEIGHT        = C.OVR_KEY_PLAYER_HEIGHT
	KEY_EYE_HEIGHT           = C.OVR_KEY_EYE_HEIGHT
	KEY_IPD                  = C.OVR_KEY_IPD
	KEY_NECK_TO_EYE_DISTANCE = C.OVR_KEY_NECK_TO_EYE_DISTANCE

	DEFAULT_GENDER                 = C.OVR_DEFAULT_GENDER
	DEFAULT_PLAYER_HEIGHT          = C.float(1.778)
	DEFAULT_EYE_HEIGHT             = C.float(1.675)
	DEFAULT_IPD                    = C.float(0.064)
	DEFAULT_NECK_TO_EYE_HORIZONTAL = C.float(0.0805)
	DEFAULT_NECK_TO_EYE_VERTICAL   = C.float(0.075)
	DEFAULT_EYE_RELIEF_DIAL        = C.OVR_DEFAULT_EYE_RELIEF_DIAL
)

func (hmd *Hmd) GetBool(propertyName string, defaultVal bool) bool {
	_propertyName := C.CString(propertyName)
	defer C.free(unsafe.Pointer(_propertyName))
	return C.ovrHmd_GetBool(hmd.hmdRef, _propertyName, ovrBool(defaultVal)) == 1
}

func (hmd *Hmd) SetBool(propertyName string, value bool) bool {
	_propertyName := C.CString(propertyName)
	defer C.free(unsafe.Pointer(_propertyName))
	return C.ovrHmd_SetBool(hmd.hmdRef, _propertyName, ovrBool(value)) == 1
}

func (hmd *Hmd) GetInt(propertyName string, defaultVal int) int {
	_propertyName := C.CString(propertyName)
	defer C.free(unsafe.Pointer(_propertyName))
	return int(C.ovrHmd_GetInt(hmd.hmdRef, _propertyName, C.int(defaultVal)))
}

func (hmd *Hmd) SetInt(propertyName string, value int) bool {
	_propertyName := C.CString(propertyName)
	defer C.free(unsafe.Pointer(_propertyName))
	return C.ovrHmd_SetInt(hmd.hmdRef, _propertyName, C.int(value)) == 1
}

func (hmd *Hmd) GetFloat(propertyName string, defaultVal float32) float32 {
	_propertyName := C.CString(propertyName)
	defer C.free(unsafe.Pointer(_propertyName))
	return float32(C.ovrHmd_GetFloat(hmd.hmdRef, _propertyName, C.float(defaultVal)))
}

func (hmd *Hmd) SetFloat(propertyName string, value float32) bool {
	_propertyName := C.CString(propertyName)
	defer C.free(unsafe.Pointer(_propertyName))
	return C.ovrHmd_SetFloat(hmd.hmdRef, _propertyName, C.float(value)) == 1
}

func (hmd *Hmd) GetFloatArray(propertyName string, values []float32, arraySize uint) uint {
	_propertyName := C.CString(propertyName)
	defer C.free(unsafe.Pointer(_propertyName))
	_values := C.float(values[0])
	return uint(C.ovrHmd_GetFloatArray(hmd.hmdRef, _propertyName, &_values, C.uint(arraySize)))
}

func (hmd *Hmd) SetFloatArray(propertyName string, values []float32, arraySize uint) bool {
	_propertyName := C.CString(propertyName)
	defer C.free(unsafe.Pointer(_propertyName))
	_values := C.float(values[0])
	return C.ovrHmd_SetFloatArray(hmd.hmdRef, _propertyName, &_values, C.uint(arraySize)) == 1
}

func (hmd *Hmd) GetString(propertyName, defaultVal string) string {
	_propertyName := C.CString(propertyName)
	defer C.free(unsafe.Pointer(_propertyName))
	_defaultVal := C.CString(defaultVal)
	defer C.free(unsafe.Pointer(_defaultVal))
	return C.GoString(C.ovrHmd_GetString(hmd.hmdRef, _propertyName, _defaultVal))
}

func (hmd *Hmd) SetString(propertyName, value string) bool {
	_propertyName := C.CString(propertyName)
	defer C.free(unsafe.Pointer(_propertyName))
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_value))
	return C.ovrHmd_SetString(hmd.hmdRef, _propertyName, _value) == 1
}
