// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func MagnifyingGlassMinusIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path d=\"M6.75 8.25C6.33579 8.25 6 8.58579 6 9C6 9.41421 6.33579 9.75 6.75 9.75L11.25 9.75C11.6642 9.75 12 9.41421 12 9C12 8.58579 11.6642 8.25 11.25 8.25L6.75 8.25Z\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M9 2C5.13401 2 2 5.13401 2 9C2 12.866 5.13401 16 9 16C10.6625 16 12.1906 15.4197 13.3911 14.4517L16.7197 17.7803C17.0126 18.0732 17.4874 18.0732 17.7803 17.7803C18.0732 17.4874 18.0732 17.0126 17.7803 16.7197L14.4517 13.3911C15.4197 12.1906 16 10.6625 16 9C16 5.13401 12.866 2 9 2ZM3.5 9C3.5 5.96243 5.96243 3.5 9 3.5C12.0376 3.5 14.5 5.96243 14.5 9C14.5 10.519 13.8852 11.893 12.8891 12.8891C11.893 13.8852 10.519 14.5 9 14.5C5.96243 14.5 3.5 12.0376 3.5 9Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
