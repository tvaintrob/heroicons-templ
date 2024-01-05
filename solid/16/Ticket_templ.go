// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func TicketIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1 4.5C1 3.67157 1.67157 3 2.5 3H13.5C14.3284 3 15 3.67157 15 4.5V5.5C15 5.77614 14.7727 5.99359 14.5051 6.06171C13.6399 6.28197 13 7.06626 13 8C13 8.93374 13.6399 9.71803 14.5051 9.93829C14.7727 10.0064 15 10.2239 15 10.5V11.5C15 12.3284 14.3284 13 13.5 13H2.5C1.67157 13 1 12.3284 1 11.5V10.5C1 10.2239 1.22733 10.0064 1.49494 9.93829C2.36012 9.71803 3 8.93374 3 8C3 7.06626 2.36012 6.28197 1.49494 6.06171C1.22733 5.99359 1 5.77614 1 5.5V4.5ZM10 5.75C10 5.33579 10.3358 5 10.75 5C11.1642 5 11.5 5.33579 11.5 5.75V6.75C11.5 7.16421 11.1642 7.5 10.75 7.5C10.3358 7.5 10 7.16421 10 6.75V5.75ZM10.75 8.5C10.3358 8.5 10 8.83579 10 9.25V10.25C10 10.6642 10.3358 11 10.75 11C11.1642 11 11.5 10.6642 11.5 10.25V9.25C11.5 8.83579 11.1642 8.5 10.75 8.5Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
