// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowLeftEndOnRectangleIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M4.75 2C3.23122 2 2 3.23122 2 4.75V11.25C2 12.7688 3.23122 14 4.75 14H7.75C9.26878 14 10.5 12.7688 10.5 11.25V10.75C10.5 10.3358 10.1642 10 9.75 10C9.33579 10 9 10.3358 9 10.75V11.25C9 11.9404 8.44036 12.5 7.75 12.5H4.75C4.05964 12.5 3.5 11.9404 3.5 11.25V4.75C3.5 4.05964 4.05964 3.5 4.75 3.5H7.75C8.44036 3.5 9 4.05964 9 4.75V5.25C9 5.66421 9.33579 6 9.75 6C10.1642 6 10.5 5.66421 10.5 5.25V4.75C10.5 3.23122 9.26878 2 7.75 2H4.75Z\" fill=\"#0F172A\"></path> <path d=\"M8.03033 6.28033C8.32322 5.98744 8.32322 5.51256 8.03033 5.21967C7.73744 4.92678 7.26256 4.92678 6.96967 5.21967L4.71967 7.46967C4.42678 7.76256 4.42678 8.23744 4.71967 8.53033L6.96967 10.7803C7.26256 11.0732 7.73744 11.0732 8.03033 10.7803C8.32322 10.4874 8.32322 10.0126 8.03033 9.71967L7.06066 8.75H14.25C14.6642 8.75 15 8.41421 15 8C15 7.58579 14.6642 7.25 14.25 7.25H7.06066L8.03033 6.28033Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}