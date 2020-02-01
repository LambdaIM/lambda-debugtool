package main

import "os"

const (
	FlagPubkeyPrefix  = "pubkey"
	FlagBechPrefix    = "bech"
	FlagLambdaCliHome = "lambdacli-home"
)

var (
	DefaultLambdaCliHome = os.ExpandEnv("$HOME/.lambdacli")
)
