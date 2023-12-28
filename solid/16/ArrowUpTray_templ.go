// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowUpTrayIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M7.25 10.25C7.25 10.6642 7.58579 11 8 11C8.41421 11 8.75 10.6642 8.75 10.25L8.75 4.56066L10.9697 6.78033C11.2626 7.07322 11.7374 7.07322 12.0303 6.78033C12.3232 6.48744 12.3232 6.01256 12.0303 5.71967L8.53033 2.21967C8.23744 1.92678 7.76256 1.92678 7.46967 2.21967L3.96967 5.71967C3.67678 6.01256 3.67678 6.48744 3.96967 6.78033C4.26256 7.07322 4.73744 7.07322 5.03033 6.78033L7.25 4.56066L7.25 10.25Z\" fill=\"#0F172A\"></path> <path d=\"M3.5 9.75C3.5 9.33579 3.16421 9 2.75 9C2.33579 9 2 9.33579 2 9.75V11.25C2 12.7688 3.23122 14 4.75 14H11.25C12.7688 14 14 12.7688 14 11.25V9.75C14 9.33579 13.6642 9 13.25 9C12.8358 9 12.5 9.33579 12.5 9.75V11.25C12.5 11.9404 11.9404 12.5 11.25 12.5H4.75C4.05964 12.5 3.5 11.9404 3.5 11.25V9.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
