// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowLeftStartOnRectangleIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M14 4.75C14 3.23122 12.7688 2 11.25 2H8.25C6.73122 2 5.5 3.23122 5.5 4.75V5.25C5.5 5.66421 5.83579 6 6.25 6C6.66421 6 7 5.66421 7 5.25V4.75C7 4.05964 7.55964 3.5 8.25 3.5H11.25C11.9404 3.5 12.5 4.05964 12.5 4.75V11.25C12.5 11.9404 11.9404 12.5 11.25 12.5H8.25C7.55964 12.5 7 11.9404 7 11.25V10.75C7 10.3358 6.66421 10 6.25 10C5.83579 10 5.5 10.3358 5.5 10.75V11.25C5.5 12.7688 6.73122 14 8.25 14H11.25C12.7688 14 14 12.7688 14 11.25V4.75ZM4.53033 5.21967C4.23744 4.92678 3.76256 4.92678 3.46967 5.21967L1.21967 7.46967C0.926777 7.76256 0.926777 8.23744 1.21967 8.53033L3.46967 10.7803C3.76256 11.0732 4.23744 11.0732 4.53033 10.7803C4.82322 10.4874 4.82322 10.0126 4.53033 9.71967L3.56066 8.75L10.75 8.75C11.1642 8.75 11.5 8.41421 11.5 8C11.5 7.58579 11.1642 7.25 10.75 7.25L3.56066 7.25L4.53033 6.28033C4.82322 5.98744 4.82322 5.51256 4.53033 5.21967Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
