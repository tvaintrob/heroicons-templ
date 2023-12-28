// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CurrencyEuroIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M10 18C14.4183 18 18 14.4183 18 10C18 5.58172 14.4183 2 10 2C5.58172 2 2 5.58172 2 10C2 14.4183 5.58172 18 10 18ZM8.79767 7.45038C9.31032 6.78 9.93284 6.5 10.5 6.5C11.0672 6.5 11.6897 6.78 12.2023 7.45038C12.4539 7.77942 12.9247 7.84218 13.2537 7.59056C13.5827 7.33895 13.6455 6.86824 13.3939 6.53921C12.6373 5.54986 11.5963 5 10.5 5C9.40366 5 8.36269 5.54986 7.60613 6.53921C7.27168 6.97657 7.01308 7.47233 6.83031 8H5.75C5.33579 8 5 8.33579 5 8.75C5 9.16421 5.33579 9.5 5.75 9.5H6.51977C6.49341 9.83243 6.49341 10.1676 6.51977 10.5H5.75C5.33579 10.5 5 10.8358 5 11.25C5 11.6642 5.33579 12 5.75 12H6.83031C7.01308 12.5277 7.27168 13.0234 7.60613 13.4608C8.36269 14.4501 9.40366 15 10.5 15C11.5963 15 12.6373 14.4501 13.3939 13.4608C13.6455 13.1318 13.5827 12.661 13.2537 12.4094C12.9247 12.1578 12.4539 12.2206 12.2023 12.5496C11.6897 13.22 11.0672 13.5 10.5 13.5C9.93284 13.5 9.31032 13.22 8.79767 12.5496C8.66735 12.3792 8.55299 12.1948 8.4546 12H10.25C10.6642 12 11 11.6642 11 11.25C11 10.8358 10.6642 10.5 10.25 10.5H8.02592C7.99136 10.1685 7.99136 9.83152 8.02592 9.5H10.25C10.6642 9.5 11 9.16421 11 8.75C11 8.33579 10.6642 8 10.25 8H8.4546C8.55299 7.80515 8.66735 7.62081 8.79767 7.45038Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}