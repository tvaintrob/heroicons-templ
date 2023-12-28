// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowDownOnSquareIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M8 1C8.41421 1 8.75 1.33579 8.75 1.75V5H7.25V1.75C7.25 1.33579 7.58579 1 8 1Z\" fill=\"#0F172A\"></path> <path d=\"M7.25 5V9.43934L6.03033 8.21967C5.73744 7.92678 5.26256 7.92678 4.96967 8.21967C4.67678 8.51256 4.67678 8.98744 4.96967 9.28033L7.46967 11.7803C7.76256 12.0732 8.23744 12.0732 8.53033 11.7803L11.0303 9.28033C11.3232 8.98744 11.3232 8.51256 11.0303 8.21967C10.7374 7.92678 10.2626 7.92678 9.96967 8.21967L8.75 9.43934V5H11C12.1046 5 13 5.89543 13 7V13C13 14.1046 12.1046 15 11 15H5C3.89543 15 3 14.1046 3 13V7C3 5.89543 3.89543 5 5 5H7.25Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}