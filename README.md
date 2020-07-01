# Elektron Sysex Processor

Elektron sysex reader and writer.

This is my first attempt at reading and writing sysex.

Not sure if Go is the right language or I should switch to C++.

SysEx or System Exclusive is a part of the MIDI standard which can be used by manufacturers to define their own specific message format.

SysEx starts with `F0` byte and ends with `F7`.
In between you can find the manufacturers identification and the actual data. 

The idea is that I would be able to read and write files in SysEx format for various devices I own. E.g.
- Elektron MachineDrum
- Elektron MonoMachine
- Akai S3000XL

For testing purposes I use the SysEx files of the Autechre 2008 Quaristice tour.
I haven't includes there in the repo.
You probably can find them on the internet, or just use another SysEx file.

## Links

- A simple [video](https://www.youtube.com/watch?v=Tj9uJ8i0ocg) on SysEx.
- A very nice [article](http://www.muzines.co.uk/articles/everything-you-ever-wanted-to-know-about-system-exclusive/4558) os SysEx.
- [Midi SysEx](https://en.wikipedia.org/wiki/MIDI_Machine_Control) on wiki
- [SysEx Librarian](https://www.snoize.com/SysExLibrarian/) [source](https://github.com/krevis/MIDIApps) for sending and receiving SysEx
- Autechre on [youtube](https://www.youtube.com/watch?v=UiFWYtgRBHk)
- Autechre on [Elektronauts](https://www.elektronauts.com/t/autechre-md-mnm-sysex-files-mpc-nord/67208)

https://electronicmusic.fandom.com/wiki/System_exclusive
https://electronicmusic.fandom.com/wiki/List_of_MIDI_Manufacturer_IDs

## Analysis of Elektron MonoMachine and MachineDrum SysEx dumps

It appears that a SysEx dump itself is not wrapped with SysEx message start and end bytes but that a full dump contains
a number of messages that all have the following format:

```
- F0 byte
- header 5 bytes (synth specific)
- message type 1 byte
- data
- F7 byte
```

Notes:
- Not sure if the header is always 5 bytes long. Maybe older SysEx files (when the synth manufacturer code was just one
 byte long) have a shorter header
- The message types for MonoMachine and MachineDrum are identical. The content most probably not.