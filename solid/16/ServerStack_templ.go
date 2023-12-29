// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ServerStackIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M5.35401 2C4.5362 2 3.80078 2.4979 3.49706 3.25722L3.1593 4.10161C3.42859 4.03522 3.71016 4 3.99995 4H11.9999C12.2897 4 12.5713 4.03522 12.8406 4.10161L12.5028 3.25722C12.1991 2.4979 11.4637 2 10.6459 2H5.35401Z\" fill=\"black\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2 13C2 11.8954 2.89543 11 4 11H12C13.1046 11 14 11.8954 14 13C14 14.1046 13.1046 15 12 15H4C2.89543 15 2 14.1046 2 13ZM12.75 13C12.75 13.4142 12.4142 13.75 12 13.75C11.5858 13.75 11.25 13.4142 11.25 13C11.25 12.5858 11.5858 12.25 12 12.25C12.4142 12.25 12.75 12.5858 12.75 13ZM9 13.75C9.41421 13.75 9.75 13.4142 9.75 13C9.75 12.5858 9.41421 12.25 9 12.25C8.58579 12.25 8.25 12.5858 8.25 13C8.25 13.4142 8.58579 13.75 9 13.75Z\" fill=\"black\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4 5.5C2.89543 5.5 2 6.39543 2 7.5C2 8.60457 2.89543 9.5 4 9.5H12C13.1046 9.5 14 8.60457 14 7.5C14 6.39543 13.1046 5.5 12 5.5H4ZM12 8.25C12.4142 8.25 12.75 7.91421 12.75 7.5C12.75 7.08579 12.4142 6.75 12 6.75C11.5858 6.75 11.25 7.08579 11.25 7.5C11.25 7.91421 11.5858 8.25 12 8.25ZM9.75 7.5C9.75 7.91421 9.41421 8.25 9 8.25C8.58579 8.25 8.25 7.91421 8.25 7.5C8.25 7.08579 8.58579 6.75 9 6.75C9.41421 6.75 9.75 7.08579 9.75 7.5Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
