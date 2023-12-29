// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CommandLineIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3.25 3C2.00736 3 1 4.00736 1 5.25V14.75C1 15.9926 2.00736 17 3.25 17H16.75C17.9926 17 19 15.9926 19 14.75V5.25C19 4.00736 17.9926 3 16.75 3H3.25ZM4.19253 11.7517C3.91544 11.4438 3.94039 10.9696 4.24828 10.6925L6.12886 9L4.24828 7.30747C3.94039 7.03038 3.91543 6.55616 4.19253 6.24828C4.46962 5.94039 4.94384 5.91544 5.25172 6.19253L7.75172 8.44253C7.90976 8.58476 8 8.78738 8 9C8 9.21261 7.90976 9.41524 7.75172 9.55747L5.25172 11.8075C4.94384 12.0846 4.46962 12.0596 4.19253 11.7517ZM9.75 10.25C9.33579 10.25 9 10.5858 9 11C9 11.4142 9.33579 11.75 9.75 11.75H12.25C12.6642 11.75 13 11.4142 13 11C13 10.5858 12.6642 10.25 12.25 10.25H9.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
