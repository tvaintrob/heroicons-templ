// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChevronUpDownIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M10.5303 3.46967C10.2374 3.17678 9.76256 3.17678 9.46967 3.46967L6.21967 6.71967C5.92678 7.01256 5.92678 7.48744 6.21967 7.78033C6.51256 8.07322 6.98744 8.07322 7.28033 7.78033L10 5.06066L12.7197 7.78033C13.0126 8.07322 13.4874 8.07322 13.7803 7.78033C14.0732 7.48744 14.0732 7.01256 13.7803 6.71967L10.5303 3.46967ZM6.21967 13.2803L9.46967 16.5303C9.76256 16.8232 10.2374 16.8232 10.5303 16.5303L13.7803 13.2803C14.0732 12.9874 14.0732 12.5126 13.7803 12.2197C13.4874 11.9268 13.0126 11.9268 12.7197 12.2197L10 14.9393L7.28033 12.2197C6.98744 11.9268 6.51256 11.9268 6.21967 12.2197C5.92678 12.5126 5.92678 12.9874 6.21967 13.2803Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}