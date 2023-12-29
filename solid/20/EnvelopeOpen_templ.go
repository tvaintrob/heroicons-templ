// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func EnvelopeOpenIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2.10557 6.44723C1.428 6.78601 1 7.47854 1 8.23608V16C1 17.1046 1.89543 18 3 18H17C18.1046 18 19 17.1046 19 16V8.23608C19 7.47854 18.572 6.78601 17.8944 6.44723L10.8944 2.94723C10.3314 2.6657 9.66863 2.6657 9.10557 2.94723L2.10557 6.44723ZM3.58541 10.4542C3.21493 10.269 2.76442 10.4191 2.57918 10.7896C2.39394 11.1601 2.54411 11.6106 2.91459 11.7958L8.77016 14.7236C9.54436 15.1107 10.4556 15.1107 11.2298 14.7236L17.0823 11.7974C17.4528 11.6122 17.6029 11.1617 17.4177 10.7912C17.2325 10.4207 16.7819 10.2705 16.4115 10.4558L10.559 13.382C10.2071 13.5579 9.79289 13.5579 9.44098 13.382L3.58541 10.4542Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
