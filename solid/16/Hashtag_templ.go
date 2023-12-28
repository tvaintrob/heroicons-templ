// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func HashtagIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M7.48677 2.89033C7.56427 2.48344 7.29725 2.09075 6.89035 2.01325C6.48345 1.93574 6.09077 2.20277 6.01326 2.60967L5.55827 4.99835H3.60963C3.19542 4.99835 2.85963 5.33414 2.85963 5.74835C2.85963 6.16257 3.19542 6.49835 3.60963 6.49835H5.27256L4.7016 9.49589H2.74963C2.33542 9.49589 1.99963 9.83168 1.99963 10.2459C1.99963 10.6601 2.33542 10.9959 2.74963 10.9959H4.41588L4.01326 13.1097C3.93576 13.5166 4.20278 13.9092 4.60968 13.9868C5.01658 14.0643 5.40926 13.7972 5.48677 13.3903L5.94285 10.9959H8.91589L8.51326 13.1097C8.43576 13.5166 8.70278 13.9092 9.10968 13.9868C9.51658 14.0643 9.90927 13.7972 9.98677 13.3903L10.4429 10.9959H12.3896C12.8038 10.9959 13.1396 10.6601 13.1396 10.2459C13.1396 9.83168 12.8038 9.49589 12.3896 9.49589H10.7286L11.2995 6.49835H13.2496C13.6638 6.49835 13.9996 6.16257 13.9996 5.74835C13.9996 5.33414 13.6638 4.99835 13.2496 4.99835H11.5852L11.9868 2.89033C12.0643 2.48344 11.7972 2.09075 11.3903 2.01325C10.9835 1.93574 10.5908 2.20277 10.5133 2.60967L10.0583 4.99835H7.08524L7.48677 2.89033ZM6.79953 6.49835L6.22857 9.49589H9.2016L9.77256 6.49835H6.79953Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
