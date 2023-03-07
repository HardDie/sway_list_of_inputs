## sway_list_of_inputs
Allows you to get a list of inputs in a compact form with name and vendor attributes.

The usual way to get the name and vendor is to run the command **swaymsg -t get_inputs**.
But in such a large json response, it's hard to find the right fields.
This utility allows you to get a list of inputs in this view:
```
[
        {
                "vendor": 1,
                "names": [
                        "AT Translated Set 2 keyboard"
                ]
        },
        {
                "vendor": 1242,
                "names": [
                        "USB-HID Keyboard",
                        "USB-HID Keyboard Consumer Control",
                        "USB-HID Keyboard System Control"
                ]
        }
]
```

## How to install
```
go install github.com/HardDie/sway_list_of_inputs
```

## How to run
```
export PATH=$PATH:$HOME/go/bin
sway_list_of_inputs
```
