// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ComputerDesktopIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2 4.25C2 3.00736 3.00736 2 4.25 2H15.75C16.9926 2 18 3.00736 18 4.25V12.75C18 13.9926 16.9926 15 15.75 15H12.6448C12.8417 15.6619 13.2292 16.2418 13.7449 16.6767C13.9856 16.8798 14.0738 17.2116 13.9657 17.5074C13.8576 17.8032 13.5762 18 13.2613 18H6.73881C6.42387 18 6.14248 17.8032 6.03437 17.5074C5.92627 17.2116 6.01449 16.8798 6.25522 16.6767C6.77086 16.2418 7.15838 15.6619 7.35525 15H4.25C3.00736 15 2 13.9926 2 12.75V4.25ZM3.5 4.25C3.5 3.83579 3.83579 3.5 4.25 3.5H15.75C16.1642 3.5 16.5 3.83579 16.5 4.25V11.75C16.5 12.1642 16.1642 12.5 15.75 12.5H4.25C3.83579 12.5 3.5 12.1642 3.5 11.75V4.25Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
