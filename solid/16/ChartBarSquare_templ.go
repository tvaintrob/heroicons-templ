// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChartBarSquareIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4 2C2.89543 2 2 2.89543 2 4V12C2 13.1046 2.89543 14 4 14H12C13.1046 14 14 13.1046 14 12V4C14 2.89543 13.1046 2 12 2H4ZM4.75 9C4.33579 9 4 9.33579 4 9.75V11.25C4 11.6642 4.33579 12 4.75 12C5.16421 12 5.5 11.6642 5.5 11.25V9.75C5.5 9.33579 5.16421 9 4.75 9ZM7.25 7.25C7.25 6.83579 7.58579 6.5 8 6.5C8.41421 6.5 8.75 6.83579 8.75 7.25V11.25C8.75 11.6642 8.41421 12 8 12C7.58579 12 7.25 11.6642 7.25 11.25V7.25ZM11.25 4C10.8358 4 10.5 4.33579 10.5 4.75V11.25C10.5 11.6642 10.8358 12 11.25 12C11.6642 12 12 11.6642 12 11.25V4.75C12 4.33579 11.6642 4 11.25 4Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
