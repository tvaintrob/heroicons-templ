// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChartBarSquareIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4.25 2C3.00736 2 2 3.00736 2 4.25V15.75C2 16.9926 3.00736 18 4.25 18H15.75C16.9926 18 18 16.9926 18 15.75V4.25C18 3.00736 16.9926 2 15.75 2H4.25ZM15 5.75C15 5.33579 14.6642 5 14.25 5C13.8358 5 13.5 5.33579 13.5 5.75V14.25C13.5 14.6642 13.8358 15 14.25 15C14.6642 15 15 14.6642 15 14.25V5.75ZM6.5 11.75C6.5 11.3358 6.16421 11 5.75 11C5.33579 11 5 11.3358 5 11.75V14.25C5 14.6642 5.33579 15 5.75 15C6.16421 15 6.5 14.6642 6.5 14.25V11.75ZM8.5835 9C8.99771 9 9.3335 9.33579 9.3335 9.75V14.25C9.3335 14.6642 8.99771 15 8.5835 15C8.16928 15 7.8335 14.6642 7.8335 14.25V9.75C7.8335 9.33579 8.16928 9 8.5835 9ZM12.1636 7.75C12.1636 7.33579 11.8278 7 11.4136 7C10.9994 7 10.6636 7.33579 10.6636 7.75V14.25C10.6636 14.6642 10.9994 15 11.4136 15C11.8278 15 12.1636 14.6642 12.1636 14.25V7.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
