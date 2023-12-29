// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func FaceSmileIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15 8C15 11.866 11.866 15 8 15C4.13401 15 1 11.866 1 8C1 4.13401 4.13401 1 8 1C11.866 1 15 4.13401 15 8ZM6 8C6.55228 8 7 7.32843 7 6.5C7 5.67157 6.55228 5 6 5C5.44772 5 5 5.67157 5 6.5C5 7.32843 5.44772 8 6 8ZM11 6.5C11 7.32843 10.5523 8 10 8C9.44771 8 9 7.32843 9 6.5C9 5.67157 9.44771 5 10 5C10.5523 5 11 5.67157 11 6.5ZM11.0052 10.7447C10.7123 10.4518 10.2374 10.4518 9.94454 10.7447C8.8706 11.8186 7.1294 11.8186 6.05546 10.7447C5.76256 10.4518 5.28769 10.4518 4.9948 10.7447C4.7019 11.0376 4.7019 11.5125 4.9948 11.8053C6.65452 13.4651 9.34548 13.4651 11.0052 11.8053C11.2981 11.5124 11.2981 11.0376 11.0052 10.7447Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
