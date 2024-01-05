// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func MegaphoneIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\" class=\"w-5 h-5\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path d=\"M13.9202 3.84494C11.9859 4.84176 9.86368 5.52434 7.62054 5.82553C6.76376 5.94057 5.88902 6.00001 5 6.00001C2.79086 6.00001 1 7.79087 1 10C1 12.0384 2.52477 13.7207 4.49597 13.9686C4.78782 15.1063 5.20979 16.2269 5.76704 17.3098C6.1636 18.0805 7.10902 18.3086 7.81763 17.8995L8.68366 17.3995C9.41014 16.9801 9.62418 16.0784 9.27228 15.3782C9.10619 15.0477 8.95684 14.7129 8.82394 14.3747C10.6243 14.7325 12.3353 15.3383 13.9201 16.1551C14.6189 14.2348 15 12.1619 15 10C15 7.83812 14.6189 5.76526 13.9202 3.84494Z\"></path><path d=\"M15.2428 3.09656C16.0553 5.24255 16.5 7.56934 16.5 10C16.5 12.4307 16.0553 14.7575 15.2428 16.9035C15.2427 16.9034 15.2428 16.9035 15.2428 16.9035L15.2135 16.9807C15.0652 17.3674 15.2585 17.8012 15.6452 17.9495C16.032 18.0979 16.4657 17.9046 16.6141 17.5178C16.7002 17.2933 16.7825 17.067 16.8611 16.8389C17.4152 15.2293 17.7791 13.5316 17.9262 11.7729C18.5645 11.4388 19 10.7707 19 10C19 9.22929 18.5645 8.5612 17.9262 8.22715C17.7791 6.46842 17.4152 4.77069 16.8611 3.16114C16.7825 2.93301 16.7002 2.70667 16.6141 2.48219C16.4657 2.09544 16.032 1.90216 15.6452 2.05048C15.2585 2.19881 15.0652 2.63257 15.2135 3.01932L15.2428 3.09656C15.2428 3.09652 15.2427 3.09659 15.2428 3.09656Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
