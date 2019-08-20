# Notes

The aim of this tool is to ease the development & packaging flows.
It is a fairly opinionated tool, but I will aim to make it somewhat configurable if necessary!


Let's see what's possible.


# Features

- If no snapshot is specified, then set one and write back to disk
  - maybe exit with non-zero status...
- Automatic packaging of Application and Dependencies
  - Specify language & it 'just builds'
  - Allow for specific build instructions (a custom mkDerivation)
- Automatic build of Docker images
  - Automatic push of docker images? (would also be nice)
- Drop into shell environmment where all dependencies are imported
