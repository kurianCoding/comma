# COMMA
- this package executes bash commands, shows their output and waits for a
key press to execute the next command or does it automatically with a time
delay. 
- It is aimed at providing instructional content for teachers and
hackers who wish to demo their script in front of an audience
## Websocket 
### TODO: make websocket optional
commands locally executed and their output is written
to a websocket connection. Right now result of command output is visible
only to someone connected to a websocket server. planning to modify this
so that it is also visible locally. In addition websocket writing may be
made optional.

```
mkdir $HOME/bin
export PATH=$PATH:$HOME/bin
make build
make test
```
