// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Bars4Icon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2 2.75C2 2.33579 2.33579 2 2.75 2H13.25C13.6642 2 14 2.33579 14 2.75C14 3.16421 13.6642 3.5 13.25 3.5H2.75C2.33579 3.5 2 3.16421 2 2.75ZM2 13.25C2 12.8358 2.33579 12.5 2.75 12.5H13.25C13.6642 12.5 14 12.8358 14 13.25C14 13.6642 13.6642 14 13.25 14H2.75C2.33579 14 2 13.6642 2 13.25Z\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2 6.25C2 5.83579 2.33579 5.5 2.75 5.5H13.25C13.6642 5.5 14 5.83579 14 6.25C14 6.66421 13.6642 7 13.25 7H2.75C2.33579 7 2 6.66421 2 6.25ZM2 9.75C2 9.33579 2.33579 9 2.75 9H13.25C13.6642 9 14 9.33579 14 9.75C14 10.1642 13.6642 10.5 13.25 10.5H2.75C2.33579 10.5 2 10.1642 2 9.75Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
