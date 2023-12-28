// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PlayPauseIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M1 4.80433C1 4.0189 1.86395 3.54006 2.53 3.95633L7.6432 7.15209C8.26987 7.54376 8.26987 8.45642 7.6432 8.84808L2.53 12.0438C1.86395 12.4601 1 11.9813 1 11.1958V4.80433Z\" fill=\"black\"></path> <path d=\"M13.5 4.50009C13.5 4.22394 13.7239 4.00009 14 4.00009H14.5C14.7761 4.00009 15 4.22394 15 4.50009V11.5001C15 11.7762 14.7761 12.0001 14.5 12.0001H14C13.7239 12.0001 13.5 11.7762 13.5 11.5001V4.50009Z\" fill=\"black\"></path> <path d=\"M10.5 4.00009C10.2239 4.00009 10 4.22394 10 4.50009V11.5001C10 11.7762 10.2239 12.0001 10.5 12.0001H11C11.2761 12.0001 11.5 11.7762 11.5 11.5001V4.50009C11.5 4.22394 11.2761 4.00009 11 4.00009H10.5Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
