# Memory management in GO

- Memory allocaiton and deallocation does by go only in GoLang
- new() Allocate memory but no init, returns memory address and zeroed storage, can be referenced using pointers
- make() Allocate memory but init, returns memory address and non-zeroed storage, can be referenced using pointers

## GC happens automatically

- Out of scope or nil
- Rewrote entire GC once go became opensrouce results were skyrocketed - [Ref](https://pkg.go.dev/runtime#hdr-Environment_Variables)
