// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func FilmIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1 4.75C1 3.7835 1.7835 3 2.75 3H17.25C18.2165 3 19 3.7835 19 4.75V15.265C19 16.2315 18.2165 17.015 17.25 17.015H15.75C15.6719 17.015 15.5951 17.0099 15.5197 17H4.48027C4.40492 17.0099 4.32806 17.015 4.25 17.015H2.75C1.7835 17.015 1 16.2315 1 15.265V4.75ZM17.5 12.135V11.01C17.5 10.8719 17.3881 10.76 17.25 10.76H15.75C15.6119 10.76 15.5 10.8719 15.5 11.01V12.135C15.5 12.2731 15.6119 12.385 15.75 12.385H17.25C17.3881 12.385 17.5 12.2731 17.5 12.135ZM17.5 14.14C17.5 14.0019 17.3881 13.89 17.25 13.89H15.75C15.6119 13.89 15.5 14.0019 15.5 14.14V15.265C15.5 15.373 15.5685 15.4651 15.6645 15.5H17.25C17.3881 15.5 17.5 15.3881 17.5 15.25V14.14ZM2.5 15.25V14.14C2.5 14.0019 2.61193 13.89 2.75 13.89H4.25C4.38807 13.89 4.5 14.0019 4.5 14.14V15.265C4.5 15.373 4.43148 15.4651 4.33553 15.5H2.75C2.61193 15.5 2.5 15.3881 2.5 15.25ZM4.5 11.01V12.135C4.5 12.2731 4.38807 12.385 4.25 12.385H2.75C2.61193 12.385 2.5 12.2731 2.5 12.135V11.01C2.5 10.8719 2.61193 10.76 2.75 10.76H4.25C4.38807 10.76 4.5 10.8719 4.5 11.01ZM17.5 9.005V7.88C17.5 7.74193 17.3881 7.63 17.25 7.63H15.75C15.6119 7.63 15.5 7.74193 15.5 7.88V9.005C15.5 9.14308 15.6119 9.255 15.75 9.255H17.25C17.3881 9.255 17.5 9.14308 17.5 9.005ZM4.25 7.63C4.38807 7.63 4.5 7.74193 4.5 7.88V9.005C4.5 9.14308 4.38807 9.255 4.25 9.255H2.75C2.61193 9.255 2.5 9.14308 2.5 9.005V7.88C2.5 7.74193 2.61193 7.63 2.75 7.63H4.25ZM4.25 4.5C4.38807 4.5 4.5 4.61193 4.5 4.75V5.875C4.5 6.01307 4.38807 6.125 4.25 6.125H2.75C2.61193 6.125 2.5 6.01307 2.5 5.875V4.75C2.5 4.61193 2.61193 4.5 2.75 4.5H4.25ZM15.75 6.125C15.6119 6.125 15.5 6.01307 15.5 5.875V4.75C15.5 4.61193 15.6119 4.5 15.75 4.5H17.25C17.3881 4.5 17.5 4.61193 17.5 4.75V5.875C17.5 6.01307 17.3881 6.125 17.25 6.125H15.75ZM6.75 9.25C6.33579 9.25 6 9.58579 6 10C6 10.4142 6.33579 10.75 6.75 10.75H13.25C13.6642 10.75 14 10.4142 14 10C14 9.58579 13.6642 9.25 13.25 9.25H6.75Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
