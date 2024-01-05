// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func FingerPrintIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 16 16\" class=\"w-4 h-4\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M8 3C7.01162 3 6.09234 3.28601 5.3177 3.77951C4.96836 4.00206 4.50474 3.89928 4.28218 3.54994C4.05963 3.20059 4.16241 2.73698 4.51176 2.51442C5.52005 1.87207 6.71772 1.5 8 1.5C11.5898 1.5 14.5 4.41015 14.5 8C14.5 9.66508 14.1674 11.2544 13.5642 12.7036C13.4051 13.086 12.9661 13.267 12.5836 13.1079C12.2012 12.9487 12.0202 12.5097 12.1794 12.1273C12.7078 10.8574 13 9.46379 13 8C13 5.23858 10.7614 3 8 3ZM3.54994 4.28219C3.89928 4.50474 4.00206 4.96836 3.77951 5.3177C3.28601 6.09234 3 7.01162 3 8C3 8.41421 2.66421 8.75 2.25 8.75C1.83578 8.75 1.5 8.41421 1.5 8C1.5 6.71772 1.87207 5.52005 2.51442 4.51176C2.73697 4.16241 3.20059 4.05963 3.54994 4.28219ZM8 5.875C6.82639 5.875 5.875 6.8264 5.875 8C5.875 10.002 4.25203 11.625 2.25 11.625H2.24577L2.21276 11.6248C1.79855 11.6225 1.46466 11.2848 1.46699 10.8706C1.46933 10.4564 1.807 10.1225 2.2212 10.1248L2.25195 10.125C3.42466 10.1239 4.375 9.17295 4.375 8C4.375 5.99797 5.99796 4.375 8 4.375C10.002 4.375 11.625 5.99797 11.625 8C11.625 8.07795 11.624 8.15568 11.6221 8.2332C11.612 8.64729 11.2681 8.97477 10.854 8.96465C10.4399 8.95452 10.1125 8.61063 10.1226 8.19654C10.1242 8.13123 10.125 8.06572 10.125 8C10.125 6.8264 9.1736 5.875 8 5.875ZM7.995 7.25C8.40921 7.25 8.745 7.58579 8.745 8C8.745 10.8351 6.93038 13.2444 4.40155 14.1335C4.01079 14.2709 3.58263 14.0655 3.44524 13.6747C3.30785 13.2839 3.51325 12.8558 3.90402 12.7184C5.85103 12.0338 7.245 10.1787 7.245 8C7.245 7.58579 7.58078 7.25 7.995 7.25ZM10.6462 10.1205C11.0375 10.2563 11.2447 10.6837 11.1088 11.075C10.5227 12.764 9.46801 14.2322 8.10143 15.3251C7.77795 15.5838 7.30599 15.5313 7.04727 15.2078C6.78856 14.8844 6.84106 14.4124 7.16455 14.1537C8.31427 13.2342 9.20006 11.9999 9.69177 10.5832C9.82758 10.1918 10.2549 9.98472 10.6462 10.1205Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
