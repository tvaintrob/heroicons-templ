// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowTrendingDownIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1.21967 5.22211C1.51256 4.92922 1.98744 4.92922 2.28033 5.22211L7 9.94178L10.7685 6.17329C10.9187 6.02306 11.1256 5.94359 11.3378 5.95463C11.55 5.96568 11.7475 6.06619 11.8813 6.23121C13.5732 8.31739 14.888 10.7612 15.6939 13.4849L17.2685 10.7576C17.4756 10.3989 17.9343 10.276 18.293 10.4831C18.6517 10.6902 18.7747 11.1489 18.5675 11.5076L16.0927 15.7942C15.8856 16.153 15.4269 16.2759 15.0682 16.0688L10.7815 13.5939C10.4228 13.3868 10.2999 12.9281 10.507 12.5694C10.7141 12.2106 11.1728 12.0877 11.5315 12.2949L14.2401 13.8586C13.5741 11.6301 12.5419 9.60646 11.2278 7.83529L7.53033 11.5328C7.38968 11.6734 7.19891 11.7524 7 11.7524C6.80109 11.7524 6.61032 11.6734 6.46967 11.5328L1.21967 6.28277C0.926777 5.98988 0.926777 5.515 1.21967 5.22211Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
