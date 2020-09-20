# AnimatedLEDStrip Client Library for Go

[![Build Status](https://travis-ci.com/AnimatedLEDStrip/client-go.svg?branch=master)](https://travis-ci.com/AnimatedLEDStrip/client-go)
[![godoc](https://godoc.org/github.com/AnimatedLEDStrip/client-go?status.svg)](http://godoc.org/github.com/AnimatedLEDStrip/client-go)
[![codecov](https://codecov.io/gh/AnimatedLEDStrip/client-go/branch/master/graph/badge.svg)](https://codecov.io/gh/AnimatedLEDStrip/client-go)

This library allows a Go client to connect to an AnimatedLEDStrip server, allowing the client to send animations to the server and receive currently running animations from the server, among other information.

## Using the Library in a Project
The library can be downloaded with:

```bash
go get github.com/AnimatedLEDStrip/client-go
```

## Creating an `AnimationSender`
An `AnimationSender` struct contains an `Ip` field (type `string`) and a `Port` field (type `int`).

```go
import als "github.com/AnimatedLEDStrip/client-go"

sender := als.AnimationSender{}
sender.Ip = "10.0.0.254"
sender.Port = 5

// or

sender := als.AnimationSender{
	Ip:   "10.0.0.254",
	Port: 5,
}
```

## Starting the `AnimationSender`
An `AnimationSender` is started by calling the `Start()` method on the instance.

```go
sender.Start()
```

## Stopping the `AnimationSender`
To stop the `AnimationSender`, call its `End()` method.

```go
sender.End()
```

## Sending Data
An animation can be sent to the server by creating an instance of the `AnimationData` struct, then calling `SendAnimation` with the struct as the argument.

```go
cc := als.ColorContainer{}
cc.AddColor(0xFF)
cc.AddColor(0xFF00)

data := als.AnimationData()        // Note that this is a function call 
                                   // that returns an animationData struct pointer
data.AddColor(&cc)

sender.SendAnimation(data)
```

#### `AnimationData` type notes
The Go library uses the following values for `continuous` and `direction`:
- `continuous`: `DEFAULT`, `CONTINUOUS`, `NONCONTINUOUS`
- `direction`: `FORWARD`, `BACKWARD`

## Receiving Data
Received animations are saved to `RunningAnimations`, which is a `RunningAnimationMap` (which is a thread-safe map).

To retrieve an animation, use
```go
sender.RunningAnimations.Load(ID)
```
where `ID` is the string ID of the animation.
