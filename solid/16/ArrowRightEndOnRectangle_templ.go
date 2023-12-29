// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowRightEndOnRectangleIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M11.25 2C12.7688 2 14 3.23122 14 4.75V11.25C14 12.7688 12.7688 14 11.25 14H8.25C6.73122 14 5.5 12.7688 5.5 11.25V10.75C5.5 10.3358 5.83579 10 6.25 10C6.66421 10 7 10.3358 7 10.75V11.25C7 11.9404 7.55964 12.5 8.25 12.5H11.25C11.9404 12.5 12.5 11.9404 12.5 11.25V4.75C12.5 4.05964 11.9404 3.5 11.25 3.5H8.25C7.55964 3.5 7 4.05964 7 4.75V5.25C7 5.66421 6.66421 6 6.25 6C5.83579 6 5.5 5.66421 5.5 5.25V4.75C5.5 3.23122 6.73122 2 8.25 2H11.25Z\"></path> <path d=\"M7.96967 6.28033C7.67678 5.98744 7.67678 5.51256 7.96967 5.21967C8.26256 4.92678 8.73744 4.92678 9.03033 5.21967L11.2803 7.46967C11.5732 7.76256 11.5732 8.23744 11.2803 8.53033L9.03033 10.7803C8.73744 11.0732 8.26256 11.0732 7.96967 10.7803C7.67678 10.4874 7.67678 10.0126 7.96967 9.71967L8.93934 8.75H1.75C1.33579 8.75 1 8.41421 1 8C1 7.58579 1.33579 7.25 1.75 7.25H8.93934L7.96967 6.28033Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
