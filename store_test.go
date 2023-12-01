package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathkey := CASPathTransformFunc(key)
	expectedOriginal := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathkey.Pathname != expectedPathName {
		t.Errorf("have %s expected %s", pathkey.Pathname, expectedPathName)
	}
	if pathkey.Original != expectedOriginal {
		t.Errorf("have %s expected %s", pathkey.Original, expectedOriginal)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("momsbestpicture", data); err != nil {
		t.Error(err)
	}

}
