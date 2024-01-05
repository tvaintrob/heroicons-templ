// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CpuChipIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path d=\"M14 6H6V14H14V6Z\"></path><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M9.25 3V1.75C9.25 1.33579 9.58579 1 10 1C10.4142 1 10.75 1.33579 10.75 1.75V3H12.25V1.75C12.25 1.33579 12.5858 1 13 1C13.4142 1 13.75 1.33579 13.75 1.75V3H14.25C15.7688 3 17 4.23122 17 5.75V6.25H18.25C18.6642 6.25 19 6.58579 19 7C19 7.41421 18.6642 7.75 18.25 7.75H17V9.25H18.25C18.6642 9.25 19 9.58579 19 10C19 10.4142 18.6642 10.75 18.25 10.75H17V12.25H18.25C18.6642 12.25 19 12.5858 19 13C19 13.4142 18.6642 13.75 18.25 13.75H17V14.25C17 15.7688 15.7688 17 14.25 17H13.75V18.25C13.75 18.6642 13.4142 19 13 19C12.5858 19 12.25 18.6642 12.25 18.25V17H10.75V18.25C10.75 18.6642 10.4142 19 10 19C9.58579 19 9.25 18.6642 9.25 18.25V17H7.75V18.25C7.75 18.6642 7.41421 19 7 19C6.58579 19 6.25 18.6642 6.25 18.25V17H5.75C4.23122 17 3 15.7688 3 14.25V13.75H1.75C1.33579 13.75 1 13.4142 1 13C1 12.5858 1.33579 12.25 1.75 12.25H3V10.75H1.75C1.33579 10.75 1 10.4142 1 10C1 9.58579 1.33579 9.25 1.75 9.25H3V7.75H1.75C1.33579 7.75 1 7.41421 1 7C1 6.58579 1.33579 6.25 1.75 6.25H3V5.75C3 4.23122 4.23122 3 5.75 3H6.25V1.75C6.25 1.33579 6.58579 1 7 1C7.41421 1 7.75 1.33579 7.75 1.75V3H9.25ZM4.5 5.75C4.5 5.05964 5.05964 4.5 5.75 4.5H14.25C14.9404 4.5 15.5 5.05964 15.5 5.75V14.25C15.5 14.9404 14.9404 15.5 14.25 15.5H5.75C5.05964 15.5 4.5 14.9404 4.5 14.25V5.75Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
