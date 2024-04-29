# xkb-i3

**xkb-i3** is a daemon that automate switching keyboard layout(keyboard language).
On changing focus from some window or workspace to another, **xkb-i3** will automatically switch the keyboard layout to the last layout used in the newly focused window or workspace.

## Requirements
- [Xlib](https://gitlab.freedesktop.org/xorg/lib/libx11)
- Go 1.22+

## Installation

```sh
go install github.com/IslamWalid/xkb-i3/cmd/xkbi3@latest
```

## Usage

|   Flags    |                             Value                              |                                                                                             Description                                                                                              |
| :--------: | :------------------------------------------------------------: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |
|   `mode`   | Mode can be `workspace` or `window`. Default value is `window` | Set which i3 element should the program track. If its value is `workspace`, all windows in the workspace share the same keyboard layout. Or if it is `window`, each window get tracked individually. |
| `i3blocks` |                           <signal\>                            |                                          Enable notifying i3blocks by setting the signal value which will be sent after any change in the keyboard layout.                                           |

### Examples:

- window mode:

```sh
xkbi3
```

- workspace mode:

```sh
xkbi3 -mode workspace
```

- window mode with sending notifications to i3blocks

```sh
xkbi3 -i3blocks SIGRTMIN+12
```

## TODO

- [ ] restore the last used layout after system shutdown or restart (works with workspace mode only).
- [ ] support notifying i3status.
- [ ] support notifying polybar.
