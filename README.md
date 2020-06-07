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
- [SysEx Librarian](https://www.snoize.com/SysExLibrarian/) [source](https://github.com/krevis/MIDIApps) for sending and receiving SysEx
- Autechre on [youtube](https://www.youtube.com/watch?v=UiFWYtgRBHk)
- Autechre on [Elektronauts](https://www.elektronauts.com/t/autechre-md-mnm-sysex-files-mpc-nord/67208)
