// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChevronDoubleDownIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.4697 13.2803C11.7626 13.5732 12.2374 13.5732 12.5303 13.2803L20.0303 5.78033C20.3232 5.48744 20.3232 5.01256 20.0303 4.71967C19.7374 4.42678 19.2626 4.42678 18.9697 4.71967L12 11.6893L5.03033 4.71967C4.73744 4.42678 4.26256 4.42678 3.96967 4.71967C3.67678 5.01256 3.67678 5.48744 3.96967 5.78033L11.4697 13.2803Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.4697 19.2803C11.7626 19.5732 12.2374 19.5732 12.5303 19.2803L20.0303 11.7803C20.3232 11.4874 20.3232 11.0126 20.0303 10.7197C19.7374 10.4268 19.2626 10.4268 18.9697 10.7197L12 17.6893L5.03033 10.7197C4.73744 10.4268 4.26256 10.4268 3.96967 10.7197C3.67678 11.0126 3.67678 11.4874 3.96967 11.7803L11.4697 19.2803Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
