// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BeakerIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11 3.5V5.75736C11 6.3541 11.2371 6.92639 11.659 7.34835L14.3916 10.0809C14.7812 10.4705 15 10.9988 15 11.5498C15 12.5352 14.3015 13.4103 13.2976 13.5736C11.5728 13.8542 9.80313 14 8 14C6.19687 14 4.42725 13.8542 2.70242 13.5736C1.69854 13.4103 1 12.5352 1 11.5498C1 10.9988 1.21885 10.4705 1.6084 10.0809L4.34099 7.34835C4.76295 6.92639 5 6.3541 5 5.75736V3.5H4.75C4.33579 3.5 4 3.16421 4 2.75C4 2.33579 4.33579 2 4.75 2H11.25C11.6642 2 12 2.33579 12 2.75C12 3.16421 11.6642 3.5 11.25 3.5H11ZM6.5 5.75736V3.5H9.5V5.75736C9.5 6.75192 9.89509 7.70575 10.5983 8.40901L10.7562 8.56685C10.7312 8.57785 10.7062 8.58922 10.6812 8.60096C10.2573 8.80044 9.76537 8.79458 9.34634 8.58507L8.15576 7.98978C7.50358 7.66369 6.79678 7.49016 6.08618 7.4698C6.35566 6.94476 6.5 6.3585 6.5 5.75736Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
