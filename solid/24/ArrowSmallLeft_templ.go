// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowSmallLeftIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M20.25 12C20.25 12.4142 19.9142 12.75 19.5 12.75L6.31066 12.75L11.7803 18.2197C12.0732 18.5126 12.0732 18.9874 11.7803 19.2803C11.4874 19.5732 11.0126 19.5732 10.7197 19.2803L3.96967 12.5303C3.67678 12.2374 3.67678 11.7626 3.96967 11.4697L10.7197 4.71967C11.0126 4.42678 11.4874 4.42678 11.7803 4.71967C12.0732 5.01256 12.0732 5.48744 11.7803 5.78033L6.31066 11.25L19.5 11.25C19.9142 11.25 20.25 11.5858 20.25 12Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
