// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowRightStartOnRectangleIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2 4.75C2 3.23122 3.23122 2 4.75 2H7.75C9.26878 2 10.5 3.23122 10.5 4.75V5.25C10.5 5.66421 10.1642 6 9.75 6C9.33579 6 9 5.66421 9 5.25V4.75C9 4.05964 8.44036 3.5 7.75 3.5H4.75C4.05964 3.5 3.5 4.05964 3.5 4.75V11.25C3.5 11.9404 4.05964 12.5 4.75 12.5H7.75C8.44036 12.5 9 11.9404 9 11.25V10.75C9 10.3358 9.33579 10 9.75 10C10.1642 10 10.5 10.3358 10.5 10.75V11.25C10.5 12.7688 9.26878 14 7.75 14H4.75C3.23122 14 2 12.7688 2 11.25V4.75ZM11.4697 5.21967C11.7626 4.92678 12.2374 4.92678 12.5303 5.21967L14.7803 7.46967C15.0732 7.76256 15.0732 8.23744 14.7803 8.53033L12.5303 10.7803C12.2374 11.0732 11.7626 11.0732 11.4697 10.7803C11.1768 10.4874 11.1768 10.0126 11.4697 9.71967L12.4393 8.75L5.25 8.75C4.83579 8.75 4.5 8.41421 4.5 8C4.5 7.58579 4.83579 7.25 5.25 7.25L12.4393 7.25L11.4697 6.28033C11.1768 5.98744 11.1768 5.51256 11.4697 5.21967Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
