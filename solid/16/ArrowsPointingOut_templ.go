// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowsPointingOutIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2.75 9C3.16421 9 3.5 9.33579 3.5 9.75V11.4393L5.71967 9.21967C6.01256 8.92678 6.48744 8.92678 6.78033 9.21967C7.07322 9.51256 7.07322 9.98744 6.78033 10.2803L4.56066 12.5H6.25C6.66421 12.5 7 12.8358 7 13.25C7 13.6642 6.66421 14 6.25 14H2.75C2.33579 14 2 13.6642 2 13.25V9.75C2 9.33579 2.33579 9 2.75 9Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2.75 7C3.16421 7 3.5 6.66421 3.5 6.25V4.56066L5.71967 6.78033C6.01256 7.07322 6.48744 7.07322 6.78033 6.78033C7.07322 6.48744 7.07322 6.01256 6.78033 5.71967L4.56066 3.5H6.25C6.66421 3.5 7 3.16421 7 2.75C7 2.33579 6.66421 2 6.25 2H2.75C2.33579 2 2 2.33579 2 2.75V6.25C2 6.66421 2.33579 7 2.75 7Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M13.25 9C12.8358 9 12.5 9.33579 12.5 9.75V11.4393L10.2803 9.21967C9.98744 8.92678 9.51256 8.92678 9.21967 9.21967C8.92678 9.51256 8.92678 9.98744 9.21967 10.2803L11.4393 12.5H9.75C9.33579 12.5 9 12.8358 9 13.25C9 13.6642 9.33579 14 9.75 14H13.25C13.6642 14 14 13.6642 14 13.25V9.75C14 9.33579 13.6642 9 13.25 9Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M13.25 7C12.8358 7 12.5 6.66421 12.5 6.25V4.56066L10.2803 6.78033C9.98744 7.07322 9.51256 7.07322 9.21967 6.78033C8.92678 6.48744 8.92678 6.01256 9.21967 5.71967L11.4393 3.5H9.75C9.33579 3.5 9 3.16421 9 2.75C9 2.33579 9.33579 2 9.75 2H13.25C13.6642 2 14 2.33579 14 2.75V6.25C14 6.66421 13.6642 7 13.25 7Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}