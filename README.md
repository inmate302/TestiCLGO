<img width="1175" height="241" alt="image" src="https://github.com/user-attachments/assets/6d375719-8c5a-4105-8715-b3e527204206" />




Rework of "TestiCL" a CLI controller test app (hence the anagram).
It is now written in Go and it uses bubbletea. So it should be faster, lighter
and with a bunch less of dependencies and libraries installed on your computer.

What's new?
Thumbsticks are rendered dinamically now and with the correct tolerances, 
so they shouldn't jitter and flick erratically (unless they're pretty mangled).

What is left to be done?

- Bring back quotes! And possibly dial down on the Dave Chapelle ones!
- Integrate SDL_GameControllerDB in order to parse
  different controller mappings (Xbox, Playstation, Nintendo Switch, etc)

  OR
  
- Allow user to set controller mappings.
- Parse a string with Controller name, GUID, mappings so they can use it
  with their projects and games.

Installation
- 
