// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func TrashIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M5 3.25V4H2.75C2.33579 4 2 4.33579 2 4.75C2 5.16421 2.33579 5.5 2.75 5.5H3.05L3.86493 13.6493C3.94161 14.4161 4.58685 15 5.35748 15H10.6425C11.4131 15 12.0584 14.4161 12.1351 13.6493L12.95 5.5H13.25C13.6642 5.5 14 5.16421 14 4.75C14 4.33579 13.6642 4 13.25 4H11V3.25C11 2.00736 9.99264 1 8.75 1H7.25C6.00736 1 5 2.00736 5 3.25ZM7.25 2.5C6.83579 2.5 6.5 2.83579 6.5 3.25V4H9.5V3.25C9.5 2.83579 9.16421 2.5 8.75 2.5H7.25ZM6.05044 6.00094C6.46413 5.98025 6.81627 6.29885 6.83696 6.71255L7.11195 12.2125C7.13264 12.6262 6.81404 12.9784 6.40034 12.9991C5.98665 13.0197 5.63451 12.7011 5.61383 12.2875L5.33883 6.78745C5.31814 6.37376 5.63674 6.02162 6.05044 6.00094ZM9.95034 6.00094C10.364 6.02162 10.6826 6.37376 10.662 6.78745L10.387 12.2875C10.3663 12.7011 10.0141 13.0197 9.60044 12.9991C9.18674 12.9784 8.86814 12.6262 8.88883 12.2125L9.16383 6.71255C9.18451 6.29885 9.53665 5.98025 9.95034 6.00094Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
