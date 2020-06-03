package text

import "testing"

func TestAntiXssLite(t *testing.T) {
	in := []string{
		`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
		`<a href="javascript:alert('XSS1')" onmouseover="alert('XSS2')">XSS<a>`,
		`<a href="javascript:alert('XSS1')" onmouseover="alert('XSS2')">XSS`,
		`<script>alert(xss1)</script>`,
		`<script>alert(xss1)`,
		`alert(xss1)</script>`,
		`>alert(xss1)</script>`,
		`Hello <STYLE>.XSS{background-image:url("javascript:alert('XSS')");}</STYLE><A CLASS=XSS></A>World`,
	}

	for i, item := range in {
		t.Logf("%d, \nitem=%v, \nresp=%v\n", i, item, AntiXssLite(item))
	}
}

func TestAntiXssStrict(t *testing.T) {
	in := []string{
		`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
		`<a href="javascript:alert('XSS1')" onmouseover="alert('XSS2')">XSS<a>`,
		`<a href="javascript:alert('XSS1')" onmouseover="alert('XSS2')">XSS`,
		`<script>alert(xss1)</script>`,
		`<script>alert(xss1)`,
		`alert(xss1)</script>`,
		`>alert(xss1)</script>`,
		`Hello <STYLE>.XSS{background-image:url("javascript:alert('XSS')");}</STYLE><A CLASS=XSS></A>World`,
		`hello jim.`,
		`你好, 小明.`,
		`你好, 小明.<script>alert(xss1)</script>你看到我了吗?`,
	}

	for i, item := range in {
		t.Logf("%d, \nitem=%v, \nresp=%v\n", i, item, AntiXssStrict(item))
	}
}
