package m

/*
#cgo LDFLAGS: -framework ApplicationServices -framework CoreFoundation
#include <ApplicationServices/ApplicationServices.h>

static CGEventRef eventCallback(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *refcon) {
    if (type == kCGEventOtherMouseDown || type == kCGEventOtherMouseUp) {
        int64_t button = CGEventGetIntegerValueField(event, kCGMouseEventButtonNumber);
        // Block Back (3) and Forward (4) buttons
        if (button == 3 || button == 4) {
            return NULL;
        }
    }
    return event;
}

static void stopSystemButtonAction() {
    CGEventMask mask = CGEventMaskBit(kCGEventOtherMouseDown) | CGEventMaskBit(kCGEventOtherMouseUp);
    CFMachPortRef tap = CGEventTapCreate(kCGSessionEventTap, kCGHeadInsertEventTap, 0, mask, eventCallback, NULL);
    if (!tap) return;
    CFRunLoopSourceRef src = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, tap, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(), src, kCFRunLoopCommonModes);
    CGEventTapEnable(tap, true);
    CFRunLoopRun();
}
*/
import "C"
import (
	"fmt"
	"os/exec"
	"time"

	hook "github.com/robotn/gohook"
)

func Run() {
	fmt.Println("🚀 Please accept permission requests if any...")

	// Start C function to block system actions for Back/Forward buttons
	go C.stopSystemButtonAction()

	// Listen mouse events
	evChan := hook.Start()
	defer hook.End()

	var lastActionTime time.Time
	debounceDuration := 350 * time.Millisecond

	for ev := range evChan {
		if ev.Kind == hook.MouseDown {
			if time.Since(lastActionTime) < debounceDuration {
				continue
			}

			switch ev.Button {
			case 4: // Back button
				lastActionTime = time.Now()
				go switchDesktop("left")
			case 5: // Forward button
				lastActionTime = time.Now()
				go switchDesktop("right")
			}
		}
	}
}

func switchDesktop(direction string) {
	var script string
	if direction == "left" {
		script = `tell application "System Events" to key code 123 using control down`
	} else {
		script = `tell application "System Events" to key code 124 using control down`
	}
	exec.Command("osascript", "-e", script).Run()
}
