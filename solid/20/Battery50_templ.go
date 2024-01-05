// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Battery50Icon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\" class=\"w-5 h-5\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path d=\"M4.75 8C4.33579 8 4 8.33579 4 8.75V11.25C4 11.6642 4.33579 12 4.75 12H9.5C9.91421 12 10.25 11.6642 10.25 11.25V8.75C10.25 8.33579 9.91421 8 9.5 8H4.75Z\"></path><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3.25 5C2.00736 5 1 6.00736 1 7.25V12.75C1 13.9926 2.00736 15 3.25 15H15.75C16.9926 15 18 13.9926 18 12.75V11.6646C18.5826 11.4587 19 10.9031 19 10.25V9.75C19 9.09689 18.5826 8.54127 18 8.33535V7.25C18 6.00736 16.9926 5 15.75 5H3.25ZM2.5 7.25C2.5 6.83579 2.83579 6.5 3.25 6.5H15.75C16.1642 6.5 16.5 6.83579 16.5 7.25V12.75C16.5 13.1642 16.1642 13.5 15.75 13.5H3.25C2.83579 13.5 2.5 13.1642 2.5 12.75V7.25Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
