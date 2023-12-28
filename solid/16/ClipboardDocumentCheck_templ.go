// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ClipboardDocumentCheckIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.9862 3H12C13.1045 3 14 3.89543 14 5V11C14 11.9319 13.3626 12.715 12.5 12.937V7C12.5 5.61929 11.3807 4.5 9.99998 4.5H4.06299C4.28501 3.63739 5.06806 3 5.99998 3H6.01371C6.13807 1.87501 7.09184 1 8.24998 1H9.74998C10.9081 1 11.8619 1.87501 11.9862 3ZM10.5 4V3.25C10.5 2.83579 10.1642 2.5 9.74998 2.5H8.24998C7.83577 2.5 7.49998 2.83579 7.49998 3.25V4H10.5Z\" fill=\"black\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2 7C2 6.44772 2.44772 6 3 6H10C10.5523 6 11 6.44772 11 7V14C11 14.5523 10.5523 15 10 15H3C2.44772 15 2 14.5523 2 14V7ZM8.58542 8.0791C8.95591 8.26434 9.10608 8.71485 8.92083 9.08533L7.17083 12.5853C7.06755 12.7919 6.87485 12.9394 6.64847 12.9851C6.42209 13.0308 6.18727 12.9697 6.01192 12.8194L4.26192 11.3194C3.94743 11.0498 3.911 10.5763 4.18057 10.2618C4.45014 9.94733 4.92361 9.91091 5.23811 10.1805L6.25878 11.0553L7.57919 8.41451C7.76443 8.04403 8.21494 7.89386 8.58542 8.0791Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
