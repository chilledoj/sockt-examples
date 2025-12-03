# SOCKT Examples

This repo shows some examples of how to use the [sockt](https://github.com/chilledoj/sockt) library.
They don't do much, but show how it all joins together.


## How to Run

<div style="display: flex; flex-direction: row; gap: 1rem;">
<div style="flex: 0.75;">
Type
</div>
<div style="flex: 1;">
Server
</div>
<div style="flex: 1;">
Client
</div>
</div>

<div style="display: flex; flex-direction: row; gap: 1rem;">
<div style="flex: 0.5;">

WS using [coder/websocket](github.com/coder/websocket)

</div>
<div style="flex: 1;">

```shell
go run main.go coderws
```

</div>
<div style="flex: 1;">

```shell
go run cmd/client.go coderws
```

</div>
</div>

<div style="display: flex; flex-direction: row; gap: 1rem;">
<div style="flex: 0.5;">

[gorilla/websocket](https://github.com/gorilla/websocket)

</div>
<div style="flex: 1;">

```shell
go run main.go gorillaws
```

</div>
<div style="flex: 1;">

```shell
go run cmd/client.go gorillaws
```

</div>
</div>

<div style="display: flex; flex-direction: row; gap: 1rem;">
<div style="flex: 0.5;">

TCP

</div>
<div style="flex: 1;">

```shell
go run main.go tcp
```

</div>
<div style="flex: 1;">

```shell
go run cmd/client.go tcp
```

</div>
</div>
