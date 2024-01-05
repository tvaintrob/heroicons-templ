// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowUpOnSquareIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path d=\"M8.75 6.00012H7.25024V3.56078L6.03057 4.78045C5.73768 5.07334 5.26281 5.07334 4.96991 4.78045C4.67702 4.48756 4.67702 4.01268 4.96991 3.71979L7.46991 1.21979C7.76281 0.926899 8.23768 0.926899 8.53057 1.21979L11.0306 3.71979C11.3235 4.01268 11.3235 4.48756 11.0306 4.78045C10.7377 5.07335 10.2628 5.07335 9.96991 4.78045L8.75024 3.56078L8.75 6.00012L11 6C12.1046 6 13 6.89543 13 8V13C13 14.1046 12.1046 15 11 15H5C3.89543 15 3 14.1046 3 13V8C3 6.89543 3.89543 6 5 6L7.25024 6.00012L7.25 11.25C7.25 11.6642 7.58579 12 8 12C8.41421 12 8.75 11.6642 8.75 11.25L8.75 6.00012Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
