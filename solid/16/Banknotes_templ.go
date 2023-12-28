// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BanknotesIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1 3C1 2.44772 1.44772 2 2 2H14C14.5523 2 15 2.44772 15 3V9C15 9.55228 14.5523 10 14 10H2C1.44772 10 1 9.55228 1 9V3ZM10 6C10 7.10457 9.10457 8 8 8C6.89543 8 6 7.10457 6 6C6 4.89543 6.89543 4 8 4C9.10457 4 10 4.89543 10 6ZM3.75 5.25C3.33579 5.25 3 5.58579 3 6C3 6.41421 3.33579 6.75 3.75 6.75C4.16421 6.75 4.5 6.41421 4.5 6C4.5 5.58579 4.16421 5.25 3.75 5.25ZM11.5 6C11.5 5.58579 11.8358 5.25 12.25 5.25C12.6642 5.25 13 5.58579 13 6C13 6.41421 12.6642 6.75 12.25 6.75C11.8358 6.75 11.5 6.41421 11.5 6Z\" fill=\"#0F172A\"></path> <path d=\"M13 11.75C13 11.3358 12.6642 11 12.25 11C11.8358 11 11.5 11.3358 11.5 11.75V11.9286C11.5 12.079 11.3616 12.2084 11.1939 12.1839C8.11075 11.7333 4.95739 11.5 1.75 11.5C1.33579 11.5 1 11.8358 1 12.25C1 12.6642 1.33579 13 1.75 13C4.88456 13 7.96543 13.228 10.977 13.6681C12.031 13.8221 13 13.0128 13 11.9286V11.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
