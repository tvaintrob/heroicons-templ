// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func MapPinIcon() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M7.53854 14.8413L7.54225 14.8441L7.54402 14.8455C7.81 15.05 8.19 15.05 8.45647 14.8451L8.45775 14.8441L8.46146 14.8413L8.47342 14.832C8.48058 14.8264 8.48976 14.8192 8.50087 14.8104C8.50524 14.8069 8.5099 14.8032 8.51485 14.7993C8.54997 14.7713 8.59982 14.7311 8.66228 14.6791C8.78713 14.5752 8.96279 14.4243 9.17232 14.2312C9.59046 13.8459 10.1484 13.2881 10.7079 12.5969C11.8105 11.2349 13 9.25535 13 7C13 4.23858 10.7614 2 8 2C5.23858 2 3 4.23858 3 7C3 9.25535 4.18946 11.2349 5.29207 12.5969C5.8516 13.2881 6.40954 13.8459 6.82768 14.2312C7.03721 14.4243 7.21286 14.5752 7.33772 14.6791C7.40018 14.7311 7.45003 14.7713 7.48515 14.7993C7.50272 14.8133 7.51661 14.8242 7.52658 14.832L7.53854 14.8413ZM8 8.5C8.82843 8.5 9.5 7.82843 9.5 7C9.5 6.17157 8.82843 5.5 8 5.5C7.17157 5.5 6.5 6.17157 6.5 7C6.5 7.82843 7.17157 8.5 8 8.5Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
