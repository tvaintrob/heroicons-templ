// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChevronDoubleLeftIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M10.7197 11.4697C10.4268 11.7626 10.4268 12.2374 10.7197 12.5303L18.2197 20.0303C18.5126 20.3232 18.9874 20.3232 19.2803 20.0303C19.5732 19.7374 19.5732 19.2626 19.2803 18.9697L12.3107 12L19.2803 5.03033C19.5732 4.73744 19.5732 4.26256 19.2803 3.96967C18.9874 3.67678 18.5126 3.67678 18.2197 3.96967L10.7197 11.4697Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4.71967 11.4697C4.42678 11.7626 4.42678 12.2374 4.71967 12.5303L12.2197 20.0303C12.5126 20.3232 12.9874 20.3232 13.2803 20.0303C13.5732 19.7374 13.5732 19.2626 13.2803 18.9697L6.31066 12L13.2803 5.03033C13.5732 4.73744 13.5732 4.26256 13.2803 3.96967C12.9874 3.67678 12.5126 3.67678 12.2197 3.96967L4.71967 11.4697Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
