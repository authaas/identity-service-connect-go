//revive:disable:package-comments
package service

import "testing"

func TestNewServer(_ *testing.T) {
	New(&dataStub{})
}
