// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AdjustmentsVerticalIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M7.25 13.25L7.25 7.5H8.75L8.75 13.25C8.75 13.6642 8.41421 14 8 14C7.58579 14 7.25 13.6642 7.25 13.25Z\" fill=\"#0F172A\"></path> <path d=\"M8.75 2.75V5L9.5 5C9.91421 5 10.25 5.33579 10.25 5.75C10.25 6.16421 9.91421 6.5 9.5 6.5H6.5C6.08579 6.5 5.75 6.16421 5.75 5.75C5.75 5.33579 6.08579 5 6.5 5H7.25V2.75C7.25 2.33579 7.58579 2 8 2C8.41421 2 8.75 2.33579 8.75 2.75Z\" fill=\"#0F172A\"></path> <path d=\"M2.25 9.5C1.83579 9.5 1.5 9.83579 1.5 10.25C1.5 10.6642 1.83579 11 2.25 11H5.25C5.66421 11 6 10.6642 6 10.25C6 9.83579 5.66421 9.5 5.25 9.5H4.5L4.5 2.75C4.5 2.33579 4.16421 2 3.75 2C3.33579 2 3 2.33579 3 2.75L3 9.5H2.25Z\" fill=\"#0F172A\"></path> <path d=\"M10 10.25C10 9.83579 10.3358 9.5 10.75 9.5H11.5L11.5 2.75C11.5 2.33579 11.8358 2 12.25 2C12.6642 2 13 2.33579 13 2.75L13 9.5H13.75C14.1642 9.5 14.5 9.83579 14.5 10.25C14.5 10.6642 14.1642 11 13.75 11H10.75C10.3358 11 10 10.6642 10 10.25Z\" fill=\"#0F172A\"></path> <path d=\"M3 12L3 13.25C3 13.6642 3.33579 14 3.75 14C4.16421 14 4.5 13.6642 4.5 13.25V12H3Z\" fill=\"#0F172A\"></path> <path d=\"M11.5 13.25V12H13V13.25C13 13.6642 12.6642 14 12.25 14C11.8358 14 11.5 13.6642 11.5 13.25Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
