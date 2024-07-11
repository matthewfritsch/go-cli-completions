More than anything, this is a personal project. I have no idea if this will become something usable, but I hope it does.

## go-cli-completions

### Simple tooling for creating shell completions for your Go program
The hope is to create a library that allows developers to generate CLI command autocompletions/tab completions for their program.
The library should allow a developer to use their own list of strings, use completions from the result of other bash commands, and should be highly configurable.


### TODO
- [ ] Basic testing of individual functions
- [ ] Toy CLI program that uses the library to generate internal completions ((used for testing))
- [ ] Toy CLI program that uses external completions (e.g `ls` completions in the cwd) ((used for testing))
- [ ] A binary that allows for completions using a command and regex (E.g if you want to use `gco` to replace `git checkout`, we may write this as `alias gco='completions "git branch -a" | git checkout` or similar) ((is this literally just impossible without writing my own shell? lol))
- [ ] Interaction with bash completions
- [ ] Other shell completions
- [ ] Optional caching (?) so commands that may take longer can be fairly short... maybe with a timeout or number of re-hits before generating new completions?
