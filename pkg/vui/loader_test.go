package vui

import "testing"

func TestLoader(t *testing.T) {
	xml := `
<app>
  <tabs>
		<window title="personal" visibled="true">
		  Hello personal
		</window>
		<window title="zbeaver">
		  <text>Hello</text>
		</window>
	</tabs>
	<status-bar>
	</status-bar>
	<command>
	</command>
</app>
`
	loader := NewLoader([]byte(xml))
	loader.Load()
	// loader.Registry()
	// loader.Build()
}
