// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func FireIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M13.4997 4.93762C16.8478 6.87062 17.9949 11.1518 16.0619 14.4998C14.1289 17.8479 9.84775 18.995 6.4997 17.062C3.15166 15.129 2.00453 10.8479 3.93753 7.4998C4.10592 7.20813 4.29214 6.93316 4.49401 6.67548C4.69562 6.41812 5.08463 6.45704 5.28714 6.71368C5.56487 7.06565 5.88119 7.38577 6.22971 7.66764C6.56235 7.93667 7.01647 7.61943 7.00304 7.19183C7.00103 7.12812 7.00003 7.06416 7.00003 6.99997C7.00003 6.08143 7.20643 5.2111 7.57539 4.43282C8.10854 3.30822 8.98111 2.37583 10.0608 1.76798C10.3078 1.62893 10.6112 1.7522 10.7378 2.00584C11.3297 3.1927 12.2651 4.2248 13.4997 4.93762ZM14 12C14 14.2091 12.2092 16 10 16C8.08674 16 6.4791 14.6016 6.09017 12.8183C5.9966 12.3894 6.52967 12.1749 6.90396 12.4045C7.38998 12.7025 7.93731 12.8964 8.50538 12.9685C8.80801 13.0068 9.03609 12.7289 9.01482 12.4246C9.00501 12.2844 9.00002 12.1428 9.00002 12C9.00002 10.5731 9.49812 9.26254 10.3299 8.23269C10.4337 8.10417 10.599 8.04108 10.7612 8.07233C12.6063 8.4278 14 10.0511 14 12Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
