// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func TicketIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1.5 6.375C1.5 5.33947 2.33947 4.5 3.375 4.5H20.625C21.6605 4.5 22.5 5.33947 22.5 6.375V9.40135C22.5 9.66907 22.3573 9.91649 22.1255 10.0506C21.4511 10.4407 21 11.1681 21 12C21 12.8319 21.4511 13.5593 22.1255 13.9494C22.3573 14.0835 22.5 14.3309 22.5 14.5987V17.625C22.5 18.6605 21.6605 19.5 20.625 19.5H3.375C2.33947 19.5 1.5 18.6605 1.5 17.625V14.5987C1.5 14.3309 1.64271 14.0835 1.87446 13.9494C2.54894 13.5593 3 12.8319 3 12C3 11.1681 2.54894 10.4407 1.87446 10.0506C1.64271 9.91649 1.5 9.66907 1.5 9.40135V6.375ZM16.5 5.25C16.9142 5.25 17.25 5.58579 17.25 6V6.75C17.25 7.16421 16.9142 7.5 16.5 7.5C16.0858 7.5 15.75 7.16421 15.75 6.75V6C15.75 5.58579 16.0858 5.25 16.5 5.25ZM17.25 9.75C17.25 9.33579 16.9142 9 16.5 9C16.0858 9 15.75 9.33579 15.75 9.75V10.5C15.75 10.9142 16.0858 11.25 16.5 11.25C16.9142 11.25 17.25 10.9142 17.25 10.5V9.75ZM16.5 12.75C16.9142 12.75 17.25 13.0858 17.25 13.5V14.25C17.25 14.6642 16.9142 15 16.5 15C16.0858 15 15.75 14.6642 15.75 14.25V13.5C15.75 13.0858 16.0858 12.75 16.5 12.75ZM17.25 17.25C17.25 16.8358 16.9142 16.5 16.5 16.5C16.0858 16.5 15.75 16.8358 15.75 17.25V18C15.75 18.4142 16.0858 18.75 16.5 18.75C16.9142 18.75 17.25 18.4142 17.25 18V17.25ZM6 12C6 11.5858 6.33579 11.25 6.75 11.25H12C12.4142 11.25 12.75 11.5858 12.75 12C12.75 12.4142 12.4142 12.75 12 12.75H6.75C6.33579 12.75 6 12.4142 6 12ZM6.75 14.25C6.33579 14.25 6 14.5858 6 15C6 15.4142 6.33579 15.75 6.75 15.75H9.75C10.1642 15.75 10.5 15.4142 10.5 15C10.5 14.5858 10.1642 14.25 9.75 14.25H6.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}