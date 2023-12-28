// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowsRightLeftIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M10.4697 2.21967C10.7626 1.92678 11.2374 1.92678 11.5303 2.21967L13.7803 4.46967C14.0732 4.76256 14.0732 5.23744 13.7803 5.53033L11.5303 7.78033C11.2374 8.07322 10.7626 8.07322 10.4697 7.78033C10.1768 7.48744 10.1768 7.01256 10.4697 6.71967L11.4393 5.75H5.75C5.33579 5.75 5 5.41421 5 5C5 4.58579 5.33579 4.25 5.75 4.25H11.4393L10.4697 3.28033C10.1768 2.98744 10.1768 2.51256 10.4697 2.21967ZM5.53033 8.21967C5.82322 8.51256 5.82322 8.98744 5.53033 9.28033L4.56066 10.25L10.25 10.25C10.6642 10.25 11 10.5858 11 11C11 11.4142 10.6642 11.75 10.25 11.75L4.56066 11.75L5.53033 12.7197C5.82322 13.0126 5.82322 13.4874 5.53033 13.7803C5.23744 14.0732 4.76256 14.0732 4.46967 13.7803L2.21967 11.5303C1.92678 11.2374 1.92678 10.7626 2.21967 10.4697L4.46967 8.21967C4.76256 7.92678 5.23744 7.92678 5.53033 8.21967Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
