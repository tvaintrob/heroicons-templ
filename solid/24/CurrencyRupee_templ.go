// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CurrencyRupeeIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M12 2.25C6.61522 2.25 2.25 6.61522 2.25 12C2.25 17.3848 6.61522 21.75 12 21.75C17.3848 21.75 21.75 17.3848 21.75 12C21.75 6.61522 17.3848 2.25 12 2.25ZM9 7.5C8.58579 7.5 8.25 7.83579 8.25 8.25C8.25 8.66421 8.58579 9 9 9H10.5C11.4797 9 12.3131 9.62611 12.622 10.5H9C8.58579 10.5 8.25 10.8358 8.25 11.25C8.25 11.6642 8.58579 12 9 12H12.622C12.3131 12.8739 11.4797 13.5 10.5 13.5H9C8.69665 13.5 8.42318 13.6827 8.30709 13.963C8.191 14.2432 8.25517 14.5658 8.46967 14.7803L11.4697 17.7803C11.7626 18.0732 12.2374 18.0732 12.5303 17.7803C12.8232 17.4874 12.8232 17.0126 12.5303 16.7197L10.7989 14.9883C12.4785 14.8558 13.8468 13.6168 14.175 12H15C15.4142 12 15.75 11.6642 15.75 11.25C15.75 10.8358 15.4142 10.5 15 10.5H14.175C14.0625 9.94584 13.8278 9.43606 13.5003 9H15C15.4142 9 15.75 8.66421 15.75 8.25C15.75 7.83579 15.4142 7.5 15 7.5H9Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
