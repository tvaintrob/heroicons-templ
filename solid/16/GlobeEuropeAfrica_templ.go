// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func GlobeEuropeAfricaIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M8 1C4.13401 1 1 4.13401 1 8C1 11.866 4.13401 15 8 15C11.866 15 15 11.866 15 8C15 4.13401 11.866 1 8 1ZM5.65653 3.02281C3.79099 3.90272 2.5 5.8006 2.5 8C2.5 11.0376 4.96243 13.5 8 13.5C11.0376 13.5 13.5 11.0376 13.5 8C13.5 7.4168 13.4092 6.8548 13.241 6.32736L12.2942 5.69612C12.1231 5.58208 11.8954 5.60464 11.75 5.75C11.6046 5.89536 11.3769 5.91792 11.2058 5.80388L10.7387 5.49248C10.423 5.28197 10 5.50833 10 5.88783C10 5.9616 10.0172 6.03435 10.0502 6.10033L10.2764 6.55279C10.4234 6.84689 10.5 7.17119 10.5 7.5C10.5 7.82881 10.4234 8.15311 10.2764 8.44721L10.1 8.8C10.0342 8.93153 10 9.07656 10 9.22361V9.33333C10 9.76607 9.85964 10.1871 9.6 10.5333L8.8 11.6C8.61115 11.8518 8.31476 12 8 12C7.44772 12 7 11.5523 7 11V10.618C7 10.2393 6.786 9.893 6.44721 9.72361L6.02492 9.51246C5.39678 9.19839 5 8.55638 5 7.8541C5 6.83011 5.83011 6 6.8541 6H7.56155C7.8037 6 8 5.8037 8 5.56155C8 5.27631 7.73194 5.06702 7.45521 5.1362L6.61996 5.34501C6.23534 5.44116 5.82847 5.32847 5.54813 5.04813C5.2179 4.7179 5.12516 4.2184 5.31483 3.79164L5.65653 3.02281Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
