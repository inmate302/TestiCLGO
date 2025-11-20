<img width="1175" height="241" alt="image" src="https://github.com/user-attachments/assets/6d375719-8c5a-4105-8715-b3e527204206" />




Rework of "TestiCL" a CLI controller test app (hence the anagram).
It is now written in Go and it uses bubbletea. So it should be faster, lighter
and with a bunch less of dependencies and libraries installed on your computer.

Why TestiCL?
===============
Because GUI apps are **BLOAT**. Just kidding.
Ever since changing fully to linux I have spent a fair amount of time in the terminal
and enjoyed many TUI apps. So if you live in the terminal and need a controller tested
this is the app for you.

What's new?
===============


Thumbsticks are rendered dinamically now and with the correct tolerances, 
so they shouldn't jitter and flick erratically (unless they're pretty mangled).

What is left to be done?
===============

- Bring back quotes! And possibly dial down on the Dave Chapelle ones!
- Integrate SDL_GameControllerDB in order to parse
  different controller mappings (Xbox, Playstation, Nintendo Switch, etc)

  OR
  
- Allow user to set controller mappings.
- Parse a string with Controller name, GUID, mappings so they can use it
  with their projects and games.

Installation
===============
- You may download the binary from the releases section [here](https://github.com/inmate302/TestiCL-go/releases) on github
- Or if you want to build it from source make sure you have the latest version of Go installed in your system and
  from your terminal run:

      go install github.com/inmate302/TestiCLGO@latest

  After that you may run the program from your terminal:

        TestiCLGO



