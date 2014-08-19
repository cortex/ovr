package ovr

import (
	"testing"
	"time"
)

func approxFloat(expVal, calcVal, delta float32) bool {
	if calcVal < expVal-delta || calcVal > expVal+delta {
		return false
	}
	return true
}

// Create a debugging Hmd object, so the test suite runs predictably without
// a device attached.
func initializeWithHmd() *Hmd {
	Initialize()
	return HmdCreateDebug(Hmd_DK2)
}

func destroyAndShutdown(hmd *Hmd) {
	hmd.Destroy()
	Shutdown()
}

// ****************************************************************************
// ***************************** [ API interface ] ****************************
// ****************************************************************************

func TestInitializeRenderingShimAndShutdown(t *testing.T) {
	InitializeRenderingShim()
	Shutdown()
}

func TestInitializeAndShutdown(t *testing.T) {
	Initialize()
	Shutdown()
}

func TestGetVersionString(t *testing.T) {
	version := GetVersionString()
	if version[0:10] != "libOVR:0.4" {
		t.Errorf("Expected 'libOVR:0.4', instead of '%s'", version[0:10])
	}
}

func TestHmdDetect(t *testing.T) {
	numDevices := HmdDetect()
	if numDevices != 0 && numDevices != 1 {
		t.Errorf("Expected 0 or 1 device to be attached, instead of %d", numDevices)
	}
}

func TestGetLastError(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	if errStr := hmd.GetLastError(); errStr != nil {
		t.Errorf("Expected GetLastError() to return nil, instead of '%s'", errStr)
	}

	// XXX: Add a case here where an error is triggered and read back.
}

func TestGetEnabledCapsAndSetEnabledCaps(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	if caps := hmd.GetEnabledCaps(); caps != 0 {
		t.Errorf("Expected GetEnabledCaps() to return 0, instead of %d", caps)
	}

	hmd.SetEnabledCaps(HmdCap_LowPersistence | HmdCap_NoVSync)

	if caps := hmd.GetEnabledCaps(); caps != (HmdCap_LowPersistence | HmdCap_NoVSync) {
		t.Errorf("Expected GetEnabledCaps() to return HmdCap_LowPersistence | HmdCap_NoVSync")
	}
}

// ****************************************************************************
// ************************** [ Tracking interface ] **************************
// ****************************************************************************

func TestConfigureTracking(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	if !hmd.ConfigureTracking(TrackingCap_Orientation|TrackingCap_MagYawCorrection|TrackingCap_Position, 0) {
		t.Errorf("Expected ConfigureTracking() to return true")
	}
}

func TestRecenterPose(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	// We can just test that it doesn't cause a panic.
	hmd.RecenterPose()
}

func TestGetTrackingState(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	// XXX: Not sure what to test here.
	if tState := hmd.GetTrackingState(0.0); tState.HeadPose.ThePose.Orientation.W != 1 {
		t.Error("Unexpected tracking state value")
	}
}

// ****************************************************************************
// **************************** [ Graphics setup ] ****************************
// ****************************************************************************

func TestGetFovTextureSize(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	textureSize := hmd.GetFovTextureSize(Eye_Left, hmd.DefaultEyeFov[0], 1.0)

	if textureSize.W != 1182 || textureSize.H != 1461 {
		t.Error("Expected a textureSize of 1182x1461, instead of %dx%d", textureSize.W, textureSize.H)
	}
}

// ****************************************************************************
// *********************** [ SDK Distortion rendering ] ***********************
// ****************************************************************************

func TestConfigureRendering(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	renderConfig := GLConfig{}
	renderConfig.OGL.Header.API = RenderAPI_OpenGL
	renderConfig.OGL.Header.RTSize = Sizei{W: hmd.Resolution.W, H: hmd.Resolution.H}
	renderConfig.OGL.Header.Multisample = 0

	// XXX: The line below yields a panic. This might be because no window has
	//      been created to default to.

	//hmd.ConfigureRendering(&renderConfig.Config, DistortionCap_Chromatic | DistortionCap_TimeWarp | DistortionCap_Overdrive, hmd.DefaultEyeFov)
}

func TestBeginFrame(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	hmd.BeginFrame(0)
}

func TestEndFrame(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	/*
		headPose := [2]Posef{
			Posef{
				Orientation: Quatf{0, 0, 0, 1},
				Position: Vector3f{0, 0, 0},
			},
			Posef{
				Orientation: Quatf{0, 0, 0, 1},
				Position: Vector3f{0, 0, 0},
			},
		}
	*/

	eyeTexture := [2]Texture{}
	for eye := 0; eye < 2; eye++ {
		eyeTexture[eye].Header.API = RenderAPI_OpenGL
		eyeTexture[eye].Header.TextureSize = Sizei{W: hmd.Resolution.W, H: hmd.Resolution.H}
		eyeTexture[eye].Header.RenderViewport.Pos = Vector2i{0, 0}
		eyeTexture[eye].Header.RenderViewport.Size = Sizei{hmd.Resolution.W / 2, hmd.Resolution.H}
		// XXX: Give Texture{} a function to deal with textures.
	}

	// XXX: The line below results in a panic. Probably because the eyeTextures
	//      don't have a texture set.

	//hmd.EndFrame(headPose, eyeTexture)
}

func TestGetEyePose(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	// XXX: Not sure what to test here.
	hmd.GetEyePose(hmd.EyeRenderOrder[0])
}

// ****************************************************************************
// ********************** [ Client Distortion rendering ] *********************
// ****************************************************************************

func TestGetRenderDesc(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	// XXX: Not sure what to test here.
	hmd.GetRenderDesc(Eye_Left, hmd.DefaultEyeFov[0])
}

func TestCreateDistortionMesh_and_DestroyDistortionMesh_and_GetRenderScaleAndOffset(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	eyeRenderDesc := hmd.GetRenderDesc(Eye_Left, hmd.DefaultEyeFov[0])

	meshData, err := hmd.CreateDistortionMesh(eyeRenderDesc.Eye, eyeRenderDesc.Fov, DistortionCap_Chromatic|DistortionCap_TimeWarp|DistortionCap_Overdrive)
	if err != nil {
		t.Error("Expected CreateDistortionMesh() to return mesh data")
	}

	defer meshData.Destroy()

	vp := Recti{
		Pos:  Vector2i{5, 5},
		Size: Sizei{10, 10},
	}

	uvScaleOffsetOut := hmd.GetRenderScaleAndOffset(
		eyeRenderDesc.Fov,
		vp.Size,
		vp)

	checkVal := func(expVal, calcVal float32) {
		if !approxFloat(expVal, calcVal, 0.0001) {
			t.Errorf("Expected an uvScaleOffsetOut value of %f to be within range of %f", calcVal, expVal)
		}
	}

	checkVal(0.46489447, uvScaleOffsetOut[0].X)
	checkVal(0.3761417, uvScaleOffsetOut[0].Y)
	checkVal(0.99216413, uvScaleOffsetOut[1].X)
	checkVal(1.0, uvScaleOffsetOut[1].Y)
}

func TestGetFrameTiming_and_EndFrameTiming(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	// XXX: Not sure what to test here.
	hmd.GetFrameTiming(0)
}

func TestBeginFrameTiming(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	// XXX: Not sure what to test here.
	hmd.BeginFrameTiming(0)
}

func TestEndFrameTiming(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	// XXX: Not sure what to test here.
	hmd.EndFrameTiming()
}

func TestResetFrameTiming(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	// XXX: Not sure what to test here.
	hmd.ResetFrameTiming(0)
}

func TestGetEyeTimewarpMatrices(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	headPose := hmd.GetEyePose(hmd.EyeRenderOrder[0])

	// XXX: Not sure what to test here.
	hmd.GetEyeTimewarpMatrices(Eye_Left, headPose)
}

// ****************************************************************************
// *********************** [ Stateless math functions ] ***********************
// ****************************************************************************

func TestMatrix4f_Projection(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	fov := FovPort{
		UpTan:    1.3292863,
		DownTan:  1.3292863,
		LeftTan:  1.0586576,
		RightTan: 1.092368,
	}

	proj := Matrix4f_Projection(fov, 10, 1000, true)

	if !approxFloat(0.92978895, proj.M[0][0], 0.0001) || !approxFloat(0.7522834, proj.M[1][1], 0.0001) {
		t.Error("Unexpected results from Matrix4f_Projection()")
	}
}

func TestMatrix4f_OrthoSubProjection(t *testing.T) {
	proj := Matrix4f{
		M: [4][4]float32{
			{0.92978895, 0, 0.01567176, 0},
			{0, 0.7522834, 0, 0},
			{0, 0, -1.010101, -10.10101},
			{0, 0, -1, 0},
		},
	}

	ortho := Vector2f{0.0, 0.0}
	subProj := Matrix4f_OrthoSubProjection(proj, ortho, 0.8, 0.5)

	if !approxFloat(0.5654464, subProj.M[0][3], 0.0001) || !approxFloat(1, subProj.M[3][3], 0.0001) {
		t.Error("Unexpected results from Matrix4f_OrthoSubProjection()")
	}
}

func TestGetTimeInSeconds(t *testing.T) {
	if val := GetTimeInSeconds(); val < 0 {
		t.Errorf("Unexpected results %f from GetTimeInSeconds()", val)
	}
}

func TestWaitTillTime(t *testing.T) {
	startTime := time.Now()
	WaitTillTime(GetTimeInSeconds() + 0.15)
	runTime := float32(time.Since(startTime).Seconds())

	if !approxFloat(0.15, runTime, 0.0001) {
		t.Error("Expected to WaitTillTime() to pause for approximately 0.2 seconds")
	}
}

// ****************************************************************************
// ************************ [ Latency Test interface ] ************************
// ****************************************************************************

func TestProcessLatencyTest(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	_, useRgb := hmd.ProcessLatencyTest()
	if useRgb {
		t.Error("Expected ProcessLatencyTest() to tell us NOT to use the returned RGB values")
	}
}

func TestGetLatencyTestResult(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	if hmd.GetLatencyTestResult() != nil {
		t.Error("Didn't expect GetLatencyTestResult() to return testresults")
	}
}

// ****************************************************************************
// ************** [ Health and Safety Warning Display interface ] *************
// ****************************************************************************

func TestGetHSWDisplayState(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	displayState := hmd.GetHSWDisplayState()
	if !displayState.Displayed {
		t.Error("Expected the Health and Safety Warning Display to be shown")
	}
}

func TestDismissHSWDisplay(t *testing.T) {
	hmd := initializeWithHmd()
	defer destroyAndShutdown(hmd)

	if hmd.DismissHSWDisplay() {
		t.Error("Shouldn't be able to dismiss the Health and Safety Warning Display so soon")
	}
}
