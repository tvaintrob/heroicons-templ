// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowUpOnSquareStackIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path d=\"M9.96967 0.96967C10.2626 0.676777 10.7374 0.676777 11.0303 0.96967L14.0303 3.96967C14.3232 4.26256 14.3232 4.73744 14.0303 5.03033C13.7374 5.32322 13.2626 5.32322 12.9697 5.03033L11.25 3.31066V6.75H9.75V3.31066L8.03033 5.03033C7.73744 5.32322 7.26256 5.32322 6.96967 5.03033C6.67678 4.73744 6.67678 4.26256 6.96967 3.96967L9.96967 0.96967Z\"></path> <path d=\"M9.75 6.75V12.75C9.75 13.1642 10.0858 13.5 10.5 13.5C10.9142 13.5 11.25 13.1642 11.25 12.75V6.75H14.25C15.9069 6.75 17.25 8.09315 17.25 9.75V17.25C17.25 18.9069 15.9069 20.25 14.25 20.25H6.75C5.09315 20.25 3.75 18.9069 3.75 17.25V9.75C3.75 8.09315 5.09315 6.75 6.75 6.75H9.75Z\"></path> <path d=\"M7.15137 21.75C7.67008 22.6467 8.6396 23.25 9.75002 23.25H17.25C18.9069 23.25 20.25 21.9069 20.25 20.25V12.75C20.25 11.6396 19.6467 10.6701 18.75 10.1514V17.25C18.75 19.7353 16.7353 21.75 14.25 21.75H7.15137Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
