// +build ignore

// if dev then add the exlamation mark ! in front of the ignore word.
package main

type asset struct {
	bytes []byte
}

// Asset dev
func Asset(name string) (b []byte, err error) { return }

// AssetNames dev
func AssetNames() []string { return nil }

// MustAsset dev
func MustAsset(string) []byte { return nil }
