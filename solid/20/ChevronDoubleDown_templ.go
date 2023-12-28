// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M9.46967 15.2803C9.76256 15.5732 10.2374 15.5732 10.5303 15.2803L14.7803 11.0303C15.0732 10.7374 15.0732 10.2626 14.7803 9.96967C14.4874 9.67678 14.0126 9.67678 13.7197 9.96967L10 13.6893L6.28033 9.96967C5.98744 9.67678 5.51256 9.67678 5.21967 9.96967C4.92678 10.2626 4.92678 10.7374 5.21967 11.0303L9.46967 15.2803ZM5.21967 6.03033L9.46967 10.2803C9.76256 10.5732 10.2374 10.5732 10.5303 10.2803L14.7803 6.03033C15.0732 5.73744 15.0732 5.26256 14.7803 4.96967C14.4874 4.67678 14.0126 4.67678 13.7197 4.96967L10 8.68934L6.28033 4.96967C5.98744 4.67678 5.51256 4.67678 5.21967 4.96967C4.92678 5.26256 4.92678 5.73744 5.21967 6.03033Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
