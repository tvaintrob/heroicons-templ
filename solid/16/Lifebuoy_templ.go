// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func LifebuoyIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M12.9497 3.05025C10.2161 0.316582 5.78392 0.316582 3.05025 3.05025C0.316582 5.78392 0.316583 10.2161 3.05025 12.9497C5.78392 15.6834 10.2161 15.6834 12.9497 12.9497C15.6834 10.2161 15.6834 5.78392 12.9497 3.05025ZM5.68843 3.00769C7.15077 2.33077 8.84922 2.33077 10.3116 3.00769L8.69825 5.08196C8.23966 4.97269 7.76034 4.97269 7.30175 5.08196L5.68843 3.00769ZM3.00769 5.68845C2.33077 7.15079 2.33077 8.84923 3.0077 10.3116L5.08196 8.69826C4.97268 8.23967 4.97268 7.76035 5.08195 7.30177L3.00769 5.68845ZM5.68845 12.9923L7.30176 10.9181C7.76034 11.0273 8.23966 11.0273 8.69824 10.9181L10.3116 12.9923C8.84922 13.6692 7.15078 13.6692 5.68845 12.9923ZM12.9923 10.3116C13.6692 8.84923 13.6692 7.15078 12.9923 5.68845L10.918 7.30176C11.0273 7.76035 11.0273 8.23967 10.918 8.69825L12.9923 10.3116ZM6.93934 6.93935C7.52513 6.35356 8.47487 6.35356 9.06066 6.93935C9.64645 7.52513 9.64645 8.47488 9.06066 9.06067C8.47487 9.64645 7.52513 9.64645 6.93934 9.06067C6.35355 8.47488 6.35355 7.52513 6.93934 6.93935Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
