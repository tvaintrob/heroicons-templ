// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func DocumentPlusIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4 2C3.17157 2 2.5 2.67157 2.5 3.5V12.5C2.5 13.3284 3.17157 14 4 14H12C12.8284 14 13.5 13.3284 13.5 12.5V6.62132C13.5 6.2235 13.342 5.84197 13.0607 5.56066L9.93934 2.43934C9.65804 2.15804 9.2765 2 8.87868 2H4ZM8.75 6.75C8.75 6.33579 8.41421 6 8 6C7.58579 6 7.25 6.33579 7.25 6.75V8.25H5.75C5.33579 8.25 5 8.58579 5 9C5 9.41421 5.33579 9.75 5.75 9.75H7.25V11.25C7.25 11.6642 7.58579 12 8 12C8.41421 12 8.75 11.6642 8.75 11.25V9.75H10.25C10.6642 9.75 11 9.41421 11 9C11 8.58579 10.6642 8.25 10.25 8.25H8.75V6.75Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
