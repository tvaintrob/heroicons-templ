// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BoltSlashIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M20.7983 11.0118L17.6099 14.4279L9.4624 6.28042L13.7017 1.73829C13.937 1.48615 14.314 1.42701 14.6152 1.59495C14.9165 1.76289 15.0643 2.11463 14.9736 2.44736L12.982 9.75003H20.25C20.5487 9.75003 20.8189 9.92721 20.9379 10.2011C21.0569 10.475 21.0021 10.7934 20.7983 11.0118Z\" fill=\"#0F172A\"></path> <path d=\"M3.20173 12.9883L6.39014 9.57212L14.5376 17.7196L10.2983 22.2618C10.063 22.5139 9.68604 22.573 9.38481 22.4051C9.08357 22.2372 8.9357 21.8854 9.02644 21.5527L11.0181 14.25H3.75002C3.45137 14.25 3.18118 14.0728 3.06216 13.7989C2.94313 13.525 2.99795 13.2066 3.20173 12.9883Z\" fill=\"#0F172A\"></path> <path d=\"M3.53033 2.46967C3.23744 2.17678 2.76256 2.17678 2.46967 2.46967C2.17678 2.76256 2.17678 3.23744 2.46967 3.53033L20.4697 21.5303C20.7626 21.8232 21.2374 21.8232 21.5303 21.5303C21.8232 21.2374 21.8232 20.7626 21.5303 20.4697L3.53033 2.46967Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
