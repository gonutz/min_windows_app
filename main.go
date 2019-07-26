package main

import (
	"runtime"
	"syscall"

	"github.com/gonutz/w32"
)

func main() {
	runtime.LockOSThread()

	const className = "my_window_class"
	w32.RegisterClassEx(&w32.WNDCLASSEX{
		Cursor: w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW)),
		WndProc: syscall.NewCallback(func(window w32.HWND, msg uint32, w, l uintptr) uintptr {
			if msg == w32.WM_DESTROY {
				w32.PostQuitMessage(0)
				return 0
			}
			return w32.DefWindowProc(window, msg, w, l)
		}),
		ClassName: syscall.StringToUTF16Ptr(className),
	})

	window := w32.CreateWindow(
		syscall.StringToUTF16Ptr(className),
		syscall.StringToUTF16Ptr("title"),
		w32.WS_OVERLAPPEDWINDOW|w32.WS_VISIBLE,
		w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, 640, 480,
		0, 0, 0, nil,
	)

	var msg w32.MSG
	for w32.GetMessage(&msg, 0, 0, 0) != 0 {
		if !w32.TranslateAccelerator(window, 0, &msg) {
			w32.TranslateMessage(&msg)
			w32.DispatchMessage(&msg)
		}
	}
}
