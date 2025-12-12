# SOCKT Examples

This repo shows some examples of how to use the [sockt](https://github.com/chilledoj/sockt) library.
They don't do much, but show how it all joins together.


## How to Run

<table>
<thead>
    <tr>
        <th>Type</th>
        <th>Server</th>
        <th>Client</th>
    </tr>
</thead>
<tbody>
<tr>
<td>

WS using [coder/websocket](github.com/coder/websocket)

</td>
<td>

```shell
go run main.go coderws
```

</td>
<td>

```shell
go run cmd/client.go coderws
```

</td>
</tr>

<tr>
<td>

WS using [gorilla/websocket](https://github.com/gorilla/websocket)

</td>
<td>

```shell
go run main.go gorillaws
```

</td>
<td>

```shell
go run cmd/client.go gorillaws
```

</td>
</tr>

<tr>
<td>
TCP
</td>
<td>

```shell
go run main.go tcp
```

</td>
<td>

```shell
go run cmd/client.go tcp
```
</td>
</tr>

<tr>
<td>

WS using [gobwas/ws](https://github.com/gobwas/ws)

</td>
<td>

```shell
go run main.go gobwasws
```

</td>
<td>

```shell
go run cmd/client.go gobwasws
```

</td>
</tr>

</tbody>
</table>


