// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CloudArrowDownIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M5.5 17C3.01472 17 1 14.9853 1 12.5C1 10.5184 2.28084 8.83595 4.05979 8.2354C4.02046 7.9961 4 7.75044 4 7.5C4 5.01472 6.01472 3 8.5 3C10.1404 3 11.5758 3.87771 12.3621 5.18913C12.7189 5.06655 13.1017 5 13.5 5C15.433 5 17 6.567 17 8.5C17 8.83334 16.9534 9.15579 16.8664 9.4612C18.1353 10.1318 19 11.4649 19 13C19 15.2091 17.2091 17 15 17H5.5ZM10.75 7.75C10.75 7.33579 10.4142 7 10 7C9.58579 7 9.25 7.33579 9.25 7.75V12.3401L7.29959 10.2397C7.01774 9.93613 6.54319 9.91855 6.23966 10.2004C5.93613 10.4823 5.91855 10.9568 6.20041 11.2603L9.45041 14.7603C9.59231 14.9132 9.79145 15 10 15C10.2086 15 10.4077 14.9132 10.5496 14.7603L13.7996 11.2603C14.0814 10.9568 14.0639 10.4823 13.7603 10.2004C13.4568 9.91855 12.9823 9.93613 12.7004 10.2397L10.75 12.3401V7.75Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
