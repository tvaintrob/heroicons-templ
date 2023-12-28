// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15.75 3C16.9926 3 18 4.00736 18 5.25V6.46353C18 6.88735 17.7226 7.25186 17.3672 7.48284C16.5442 8.01776 16 8.94532 16 10C16 11.0547 16.5442 11.9822 17.3672 12.5172C17.7226 12.7481 18 13.1127 18 13.5365V14.75C18 15.9926 16.9926 17 15.75 17H4.25C3.00736 17 2 15.9926 2 14.75V13.5365C2 13.1127 2.27744 12.7481 2.63279 12.5172C3.45575 11.9822 4 11.0547 4 10C4 8.94532 3.45575 8.01776 2.63279 7.48284C2.27744 7.25186 2 6.88735 2 6.46353V5.25C2 4.00736 3.00736 3 4.25 3H15.75ZM13.5 7.39583C13.5 6.98162 13.1642 6.64583 12.75 6.64583C12.3358 6.64583 12 6.98162 12 7.39583V8.4375C12 8.85171 12.3358 9.1875 12.75 9.1875C13.1642 9.1875 13.5 8.85171 13.5 8.4375V7.39583ZM13.5 11.5625C13.5 11.1483 13.1642 10.8125 12.75 10.8125C12.3358 10.8125 12 11.1483 12 11.5625V12.6042C12 13.0184 12.3358 13.3542 12.75 13.3542C13.1642 13.3542 13.5 13.0184 13.5 12.6042V11.5625Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}