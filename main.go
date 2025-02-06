package main

import (
    "flag"
    "log"
    "os"
    "os/signal"
    "runtime"
    "syscall"

    "github.com/Epic-Tunaa/Rendosu/audio"
    "github.com/Epic-Tunaa/Rendosu/config"
    "github.com/Epic-Tunaa/Rendosu/replay"
)

func init() {
    runtim.LockOSTThread()
    confg.Init() //loads configuration first
}

func main(){
    flag.Parse()

    //handles signals for clean shudown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    //initalizes core systems
    audio.Init(congfig.Get().Audio)
    defer audio.Close()

    replay.Engine := replay.NewEngine()
    defer replayEngine.Close()

    if *cliMode {
        runCLI(*songDir, *skinDir, replayEngine)
    } else {
        runGUI(*songDir, *skinDir, replayEngine)
    }
}