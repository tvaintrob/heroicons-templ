// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BackspaceIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 16 16\" class=\"w-4 h-4\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M6.41421 3C5.95008 3 5.50497 3.18437 5.17678 3.51256L1.21967 7.46967C1.07902 7.61032 1 7.80109 1 8C1 8.19891 1.07902 8.38968 1.21967 8.53033L5.17678 12.4874C5.50497 12.8156 5.95009 13 6.41421 13H12.25C13.7688 13 15 11.7688 15 10.25V5.75C15 4.23122 13.7688 3 12.25 3H6.41421ZM8.28033 5.71967C7.98744 5.42678 7.51256 5.42678 7.21967 5.71967C6.92678 6.01256 6.92678 6.48744 7.21967 6.78033L8.43934 8L7.21967 9.21967C6.92678 9.51256 6.92678 9.98744 7.21967 10.2803C7.51256 10.5732 7.98744 10.5732 8.28033 10.2803L9.5 9.06066L10.7197 10.2803C11.0126 10.5732 11.4874 10.5732 11.7803 10.2803C12.0732 9.98744 12.0732 9.51256 11.7803 9.21967L10.5607 8L11.7803 6.78033C12.0732 6.48744 12.0732 6.01256 11.7803 5.71967C11.4874 5.42678 11.0126 5.42678 10.7197 5.71967L9.5 6.93934L8.28033 5.71967Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
