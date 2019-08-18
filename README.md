# go-rpc-exmaple

Goで試しにRPCサーバを作ってみた記録

## Requirements for exported methods by `net/rpc`
- The method's type is exported.
- The method itself is exported.
- The method has **2** arguments, both exported (or builtin) types.
    > The method’s first argument represents the arguments provided by the caller; the second argument represents the result parameters to be returned to the caller. The method’s return value, if non-nil, is passed back as a string that the client sees as if created by errors.New. If an error is returned, the reply parameter will not be sent back to the client.
- The method's second argument is a POINTER.
- The method has return type error.

