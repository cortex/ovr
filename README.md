Oculus SDK for Go
=================
This package provides Go bindings to the Oculus SDK 0.4.x via its C-API. Some functions have been given an Go idiomatic spin, but for the most part the functions can be used in the same manner the [Oculus Developer Guide v0.4](http://static.oculusvr.com/sdk-downloads/documents/Oculus_Developer_Guide_0.4.1.pdf) describes. At this point it's all **EXPERIMENTAL** and not well tested.

Setup
-----
Only MacOS X and Windows are supported at the moment, as there is no SDK 0.4 release for Linux yet.

#### MacOS X
Make sure you have [XCode](https://itunes.apple.com/en/app/xcode/id497799835) installed and perform the following steps to install the SDK on your system:

* Download the latest 0.4.x SDK [from the Oculus VR website](https://developer.oculusvr.com/?action=dl).
* Copy **OVR_CAPI.h** and **OVR_CAPI_GL.h** to /usr/local/include/.
* Copy **libovr.a** to /usr/local/lib/.

This should be enough for cgo to find the files it needs to compile the ovr package.

#### Windows
Make sure you have [a C compiler installed](https://gist.github.com/prep/e19d7d9e2a1e77316a7f) that Go can use for its cgo feature.

The official SDK package only comes with static libraries for Microsoft's Visual Studio, which are useless to us, because cgo uses a GNU C compiler and the libraries aren't interchangeable. Luckily for us, [jspenguin](https://developer.oculusvr.com/forums/memberlist.php?mode=viewprofile&u=28837) on the Oculus Developer Forums has provided DLL's for the latest SDK which Go can use.

* Download the [Oculus SDK 0.4.1 DLL's](https://www.jspenguin.org/software/glbumper/files/libovr_dll_0.4.1.zip) ([mirror](download.codeninja.nl/ovr/libovr_dll_0.4.1.zip)).
* Put the contents in **C:\libovr_0.4\**, which is where this ovr package expects them to be.

That is all there is to it.

There are three points specific to Windows that are worth mentioning:

* As far as I know, there is no DirectX package for Go, so your program will only have OpenGL available as its API.
* Because a DLL is used as an interface to the Oculus C-API, your compiled Go programs will require it. Chances are you'll want to have a copy of **libovr.dll** in the same directory that your compiled Go binary is in.
* If you get an error about **msvcr120.dll** not being found, you need to install the [Visual C++ Redistributable Packages for Visual Studio 2013](http://www.microsoft.com/en-us/download/details.aspx?id=40784).

Installation
------------
    $ go get github.com/prep/ovr

Todo
----
* Integrate more tightly with [Go-GL](https://github.com/go-gl/gl), in particular the *ovr.Texture* and *ovr.GLTextureData* structs.
* For the most part, the tests just checks if the functions are callable without crashing. It would be nice if they test more than that.
* Create some demos.

License
-------
This Go ovr package is licensed under [The BSD 2-Clause License](http://opensource.org/licenses/BSD-2-Clause).

The Oculus SDK is licensed under the [Oculus VR, Inc. Software Development Kit License Agreement](https://developer.oculusvr.com/license).
