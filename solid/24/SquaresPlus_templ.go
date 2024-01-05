// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SquaresPlusIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 24 24\" class=\"w-6 h-6\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path d=\"M6 3C4.34315 3 3 4.34315 3 6V8.25C3 9.90685 4.34315 11.25 6 11.25H8.25C9.90685 11.25 11.25 9.90685 11.25 8.25V6C11.25 4.34315 9.90685 3 8.25 3H6Z\"></path><path d=\"M15.75 3C14.0931 3 12.75 4.34315 12.75 6V8.25C12.75 9.90685 14.0931 11.25 15.75 11.25H18C19.6569 11.25 21 9.90685 21 8.25V6C21 4.34315 19.6569 3 18 3H15.75Z\"></path><path d=\"M6 12.75C4.34315 12.75 3 14.0931 3 15.75V18C3 19.6569 4.34315 21 6 21H8.25C9.90685 21 11.25 19.6569 11.25 18V15.75C11.25 14.0931 9.90685 12.75 8.25 12.75H6Z\"></path><path d=\"M17.625 13.5C17.625 13.0858 17.2892 12.75 16.875 12.75C16.4608 12.75 16.125 13.0858 16.125 13.5V16.125H13.5C13.0858 16.125 12.75 16.4608 12.75 16.875C12.75 17.2892 13.0858 17.625 13.5 17.625H16.125V20.25C16.125 20.6642 16.4608 21 16.875 21C17.2892 21 17.625 20.6642 17.625 20.25V17.625H20.25C20.6642 17.625 21 17.2892 21 16.875C21 16.4608 20.6642 16.125 20.25 16.125H17.625V13.5Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
