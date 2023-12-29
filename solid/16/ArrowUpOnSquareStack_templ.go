// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M5.26756 14C5.61337 14.5978 6.25972 15 7 15H11C12.1046 15 13 14.1046 13 13V10C13 9.25972 12.5978 8.61337 12 8.26756V11C12 12.6569 10.6569 14 9 14H5.26756Z\" fill=\"black\"></path> <path d=\"M6.25 6H7.75L7.75 3.56066L8.96967 4.78033C9.26256 5.07322 9.73744 5.07322 10.0303 4.78033C10.3232 4.48744 10.3232 4.01256 10.0303 3.71967L7.53033 1.21967C7.23744 0.926777 6.76256 0.926777 6.46967 1.21967L3.96967 3.71967C3.67678 4.01256 3.67678 4.48744 3.96967 4.78033C4.26256 5.07322 4.73744 5.07322 5.03033 4.78033L6.25 3.56066L6.25 6Z\" fill=\"black\"></path> <path d=\"M6.25 8.75C6.25 9.16421 6.58579 9.5 7 9.5C7.41421 9.5 7.75 9.16421 7.75 8.75L7.75 6H9C10.1046 6 11 6.89543 11 8V11C11 12.1046 10.1046 13 9 13H5C3.89543 13 3 12.1046 3 11V8C3 6.89543 3.89543 6 5 6H6.25L6.25 8.75Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
