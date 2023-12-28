// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CurrencyPoundIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M12 2.25C6.61522 2.25 2.25 6.61522 2.25 12C2.25 17.3848 6.61522 21.75 12 21.75C17.3848 21.75 21.75 17.3848 21.75 12C21.75 6.61522 17.3848 2.25 12 2.25ZM9.76273 9.51002C9.81529 9.01621 10.0302 8.53785 10.409 8.15901C11.2877 7.28033 12.7123 7.28033 13.591 8.15901C13.8839 8.4519 14.3588 8.4519 14.6517 8.15901C14.9445 7.86612 14.9445 7.39124 14.6517 7.09835C13.1872 5.63388 10.8128 5.63388 9.34835 7.09835C8.7183 7.7284 8.3587 8.52881 8.27116 9.35126C8.23849 9.6582 8.27044 9.94628 8.30277 10.1726L8.45668 11.25H8.25C7.83579 11.25 7.5 11.5858 7.5 12C7.5 12.4142 7.83579 12.75 8.25 12.75H8.67097L8.80874 13.7144C8.91661 14.4695 8.79165 15.2396 8.45053 15.9218L8.32918 16.1645C8.19352 16.4358 8.23466 16.7621 8.43341 16.9913C8.63216 17.2205 8.94939 17.3074 9.23717 17.2114L10.7757 16.6986C11.0836 16.596 11.4164 16.596 11.7243 16.6986L12.3787 16.9167C13.1335 17.1683 13.9573 17.1098 14.669 16.7539L15.3354 16.4207C15.7059 16.2355 15.8561 15.785 15.6708 15.4145C15.4856 15.044 15.0351 14.8939 14.6646 15.0791L13.9982 15.4123C13.6424 15.5902 13.2304 15.6195 12.853 15.4937L12.1987 15.2756C11.5829 15.0703 10.9171 15.0703 10.3013 15.2756L10.2401 15.296C10.3601 14.7089 10.3794 14.1022 10.2937 13.5023L10.1862 12.75H12C12.4142 12.75 12.75 12.4142 12.75 12C12.75 11.5858 12.4142 11.25 12 11.25H9.97191L9.7877 9.9605C9.7598 9.76518 9.75032 9.62664 9.76273 9.51002Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
