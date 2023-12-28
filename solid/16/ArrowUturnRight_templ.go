// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowUturnRightIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3.5 9.75C3.5 8.23122 4.73122 7 6.25 7H11.4393L9.21967 9.21967C8.92678 9.51256 8.92678 9.98744 9.21967 10.2803C9.51256 10.5732 9.98744 10.5732 10.2803 10.2803L13.7803 6.78033C14.0732 6.48744 14.0732 6.01256 13.7803 5.71967L10.2803 2.21967C9.98744 1.92678 9.51256 1.92678 9.21967 2.21967C8.92678 2.51256 8.92678 2.98744 9.21967 3.28033L11.4393 5.5L6.25 5.5C3.90279 5.5 2 7.40279 2 9.75C2 12.0972 3.90279 14 6.25 14H7.25C7.66421 14 8 13.6642 8 13.25C8 12.8358 7.66421 12.5 7.25 12.5H6.25C4.73122 12.5 3.5 11.2688 3.5 9.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
