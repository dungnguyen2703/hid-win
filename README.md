# Windows HID Tool

A lightweight utility to remap Keyboard and Mouse inputs on Windows. This tool allows you to define custom profiles to trigger specific actions (like moving windows, pressing keys) based on input triggers.

## Installation & Usage

1.  **Get the executable**: Go to the `build/` folder in this repository and copy the `hid.exe` file.
2.  **Setup**:
    - Create a new folder anywhere on your computer (e.g., `HIDTool`).
    - Place the `hid.exe` file in that folder.
    - Place your profile configuration files (`profile*.json`) in the **same folder** as the `.exe`.
3.  **Run**: Double-click `hid.exe` to start the application. It runs quietly in the system tray.

## Configuration

The application scans for any file matching the pattern `profile*.json` in its directory (e.g., `profile.json`, `profile_game.json`, `profile_office.json`).

Each file must contain a **single** JSON object defining a profile.

### JSON Structure

```json
{
  "name": "My Custom Profile",
  "description": "Short description of what this profile does",
  "bindings": [
    {
      "disabled_latest_input": true, // true: block the original input; false: let it pass through
      "triggers": [
        // All triggers in this list must be satisfied (AND logic)
        { "type": "key", "key": "CTRL" },
        { "type": "key", "key": "F1" }
      ],
      "actions": [
        // List of actions to execute sequentially
        { "type": "window_left" }
      ]
    }
  ]
}
```

### Triggers

Triggers define **when** an action happens. A binding can have multiple triggers, effectively creating a **combo**.
For a binding to activate, **ALL** triggers in the list must be satisfied (AND logic).

Example: Trigger when pressing `CTRL` + `F1`.

```json
"triggers": [
  { "type": "key", "key": "CTRL" },
  { "type": "key", "key": "F1" }
]
```

#### 1. Keyboard Trigger

Triggers when a specific key is pressed.

```json
{
  "type": "key",
  "key": "F1" // See "Supported Keys" below
}
```

#### 2. Mouse Trigger

Triggers when a specific mouse button is clicked.

```json
{
  "type": "mouse",
  "button": "BACK" // See "Supported Mouse Buttons" below
}
```

### Actions

Actions define **what** happens when triggered. Actions in the list are executed sequentially.

#### 1. Move Window

Moves the currently active window to the left or right half of the screen (equivalent to `Win + Arrow`).

```json
{ "type": "window_left" }
{ "type": "window_right" }
```

#### 2. Press Key

Simulates a key press (down and up).

```json
{
  "type": "press",
  "key": "A"
}
```

#### 3. Tap Down / Tap Up

Simulates a key down or key up event.

```json
{
  "type": "tap_down",
  "key": "A"
}
{
  "type": "tap_up",
  "key": "A"
}
```

#### 4. Delay

Waits for a specified duration before extracting the next action.

```json
{
  "type": "delay",
  "duration": 100 // Duration in milliseconds
}
```

---

## Reference

### Supported Mouse Buttons

Use these exact string values for `"button"`:

- `LEFT`
- `RIGHT`
- `MIDDLE`
- `BACK` (Side button 1)
- `FORWARD` (Side button 2)
- `V_WHEEL` (Vertical Wheel)
- `H_WHEEL` (Horizontal Wheel)

### Supported Keys

Use these exact string values for `"key"` (**case-sensitive** matching the values below):

**Letters & Numbers:**
`A` - `Z`, `0` - `9`

**Function Keys:**
`F1` - `F12`

**Special Keys:**

- `BACKSPACE`, `TAB`, `ENTER`, `SPACE`
- `ESCAPE`, `CAPS_LOCK`, `PAUSE`, `PRTSC`
- `INSERT`, `DELETE`, `HOME`, `END`, `PAGE_UP`, `PAGE_DOWN`
- `UP`, `DOWN`, `LEFT`, `RIGHT`

**Modifiers & Locks:**

- `SHIFT`, `CTRL`, `ALT`, `WIN`, `RWIN`
- `L_SHIFT`, `R_SHIFT`, `L_CTRL`, `R_CTRL`, `L_ALT`, `R_ALT`
- `NUM_LOCK`, `SCROLL_LOCK`

**Numpad:**

- `NUM_0` to `NUM_9`
- `NUM_MUL`, `NUM_ADD`, `NUM_SUB`, `NUM_DEC`, `NUM_DIV`

**Media:**

- `VOL_MUTE`, `VOL_DOWN`, `VOL_UP`
- `MEDIA_NEXT`, `MEDIA_PREV`, `MEDIA_STOP`, `MEDIA_PLAY`

**Browser & Launch:**

- `BROWSER_BACK`, `BROWSER_FORWARD`, `BROWSER_REFRESH`, `BROWSER_HOME`
- `BROWSER_STOP`, `BROWSER_SEARCH`, `BROWSER_FAVORITES`
- `LAUNCH_MAIL`, `LAUNCH_MEDIA`, `LAUNCH_APP1`, `LAUNCH_APP2`
- `BRIGHT_UP`, `BRIGHT_DOWN`

**Punctuation & Symbols:**

- `;`, `=`, `,`, `-`, `.`, `/`, `` ` ``
- `[`, `\`, `]`, `'`

**System:**

- `CONTEXT_MENU`, `SLEEP`, `SELECT`, `EXECUTE`, `HELP`

## Troubleshooting

### Profile Errors

If a profile file has invalid syntax or uses unknown keys/buttons, the application **will not load it**.
Instead, it will generate an error log file in the same folder with a name like:
`error_profile_20260111_153000.log`

Open this file to see exactly what went wrong (e.g., `unknown key name: F13`, `invalid mouse button: SIDE3`). The log is only created if errors exist.
