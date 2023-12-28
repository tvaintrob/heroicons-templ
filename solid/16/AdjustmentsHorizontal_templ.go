// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AdjustmentsHorizontalIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M6.5 2.25C6.5 1.83579 6.16421 1.5 5.75 1.5C5.33579 1.5 5 1.83579 5 2.25V5.25C5 5.66421 5.33579 6 5.75 6C6.16421 6 6.5 5.66421 6.5 5.25V4.5H13.25C13.6642 4.5 14 4.16421 14 3.75C14 3.33579 13.6642 3 13.25 3H6.5V2.25Z\" fill=\"#0F172A\"></path> <path d=\"M11 6.5C11 6.08579 10.6642 5.75 10.25 5.75C9.83579 5.75 9.5 6.08579 9.5 6.5V9.5C9.5 9.91421 9.83579 10.25 10.25 10.25C10.6642 10.25 11 9.91421 11 9.5V8.75H13.25C13.6642 8.75 14 8.41421 14 8C14 7.58579 13.6642 7.25 13.25 7.25H11V6.5Z\" fill=\"#0F172A\"></path> <path d=\"M5.75 10C6.16421 10 6.5 10.3358 6.5 10.75V11.5H13.25C13.6642 11.5 14 11.8358 14 12.25C14 12.6642 13.6642 13 13.25 13H6.5V13.75C6.5 14.1642 6.16421 14.5 5.75 14.5C5.33579 14.5 5 14.1642 5 13.75V10.75C5 10.3358 5.33579 10 5.75 10Z\" fill=\"#0F172A\"></path> <path d=\"M2.75 7.25H8.5V8.75H2.75C2.33579 8.75 2 8.41421 2 8C2 7.58579 2.33579 7.25 2.75 7.25Z\" fill=\"#0F172A\"></path> <path d=\"M4 3H2.75C2.33579 3 2 3.33579 2 3.75C2 4.16421 2.33579 4.5 2.75 4.5H4V3Z\" fill=\"#0F172A\"></path> <path d=\"M2.75 11.5H4V13H2.75C2.33579 13 2 12.6642 2 12.25C2 11.8358 2.33579 11.5 2.75 11.5Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}