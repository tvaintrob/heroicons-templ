// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func DocumentArrowUpIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4.5 2C3.67157 2 3 2.67157 3 3.5V16.5C3 17.3284 3.67157 18 4.5 18H15.5C16.3284 18 17 17.3284 17 16.5V7.62132C17 7.2235 16.842 6.84197 16.5607 6.56066L12.4393 2.43934C12.158 2.15804 11.7765 2 11.3787 2H4.5ZM9.25 13.25C9.25 13.6642 9.58579 14 10 14C10.4142 14 10.75 13.6642 10.75 13.25V10.7045L11.6925 11.7517C11.9696 12.0596 12.4438 12.0846 12.7517 11.8075C13.0596 11.5304 13.0846 11.0562 12.8075 10.7483L10.5575 8.24828C10.4152 8.09024 10.2126 8 10 8C9.78739 8 9.58476 8.09024 9.44253 8.24828L7.19253 10.7483C6.91544 11.0562 6.94039 11.5304 7.24828 11.8075C7.55616 12.0846 8.03038 12.0596 8.30747 11.7517L9.25 10.7045V13.25Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
